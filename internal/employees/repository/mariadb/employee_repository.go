package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.EmployeeRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Employee, error) {
	rows, err := m.db.QueryContext(ctx, queryGetAll)

	if err != nil {
		return &[]domain.Employee{}, err
	}

	defer rows.Close()

	var employees []domain.Employee

	for rows.Next() {
		var employee domain.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.CardNumberId,
			&employee.FirstName,
			&employee.LastName,
			&employee.WarehouseId,
		); err != nil {
			return &[]domain.Employee{}, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (m mariadbRepository) GetById(ctx context.Context, id int) (*domain.Employee, error) {

	row := m.db.QueryRowContext(ctx, queryGetById, id)

	var employee domain.Employee

	err := row.Scan(
		&employee.ID,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &employee, domain.ErrIdNotFound
	}

	// Other errors
	if err != nil {
		return &employee, err
	}

	return &employee, nil
}

func (m mariadbRepository) Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	var newEmployee domain.Employee

	result, err := m.db.ExecContext(ctx, queryCreate,
		&newEmployee.CardNumberId,
		&newEmployee.FirstName,
		&newEmployee.LastName,
		&newEmployee.WarehouseId,
	)

	if err != nil {
		return &newEmployee, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return &newEmployee, err
	}

	newEmployee.ID = int(lastId)

	return &newEmployee, nil
}

func (m mariadbRepository) Update(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	var newEmployee domain.Employee

	result, err := m.db.ExecContext(ctx, queryCreate,
		&newEmployee.CardNumberId,
		&newEmployee.FirstName,
		&newEmployee.LastName,
		&newEmployee.WarehouseId,
		&newEmployee.ID,
	)

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return &newEmployee, domain.ErrIdNotFound
	}

	// Other errors
	if err != nil {
		return &newEmployee, err
	}

	return &newEmployee, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int) error {
	result, err := m.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return domain.ErrIdNotFound
	}

	// Other errors
	if err != nil {
		return err
	}

	return nil
}
