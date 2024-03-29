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
	var employees []domain.Employee

	rows, err := m.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &employees, err
	}

	defer rows.Close()

	for rows.Next() {
		var employee domain.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.CardNumberId,
			&employee.FirstName,
			&employee.LastName,
			&employee.WarehouseId,
		); err != nil {
			return &employees, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (m mariadbRepository) GetById(ctx context.Context, id int64) (*domain.Employee, error) {

	row := m.db.QueryRowContext(ctx, sqlGetById, id)

	employee := domain.Employee{}

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
	newEmployee := domain.Employee{}

	result, err := m.db.ExecContext(
		ctx,
		sqlInsert,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
	)
	if err != nil {
		return &newEmployee, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return &newEmployee, err
	}

	employee.ID = lastId

	return employee, nil
}

func (m mariadbRepository) Update(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	newEmployee := domain.Employee{}

	result, err := m.db.ExecContext(
		ctx,
		sqlUpdate,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
		&employee.ID,
	)

	if err != nil {
		return &newEmployee, err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return &newEmployee, domain.ErrIdNotFound
	}

	// Other errors
	if err != nil {
		return &newEmployee, err
	}

	return employee, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, sqlDelete, id)
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

func (m mariadbRepository) ReportAllInboundOrders(ctx context.Context) (*[]domain.InboundOrderResponse, error) {
	var inboundOrders = []domain.InboundOrderResponse{}

	rows, err := m.db.QueryContext(ctx, sqlAllInboundOrdersCount)

	if err != nil {
		return &inboundOrders, err
	}

	for rows.Next() {
		var inboundOrder = domain.InboundOrderResponse{}

		if err := rows.Scan(
			&inboundOrder.ID,
			&inboundOrder.CardNumberId,
			&inboundOrder.FirstName,
			&inboundOrder.LastName,
			&inboundOrder.WarehouseId,
			&inboundOrder.InboundOrdersCount,
		); err != nil {
			return &inboundOrders, err
		}

		inboundOrders = append(inboundOrders, inboundOrder)
	}

	return &inboundOrders, nil
}

func (m mariadbRepository) ReportInboundOrders(ctx context.Context, employeeId int64) (*domain.InboundOrderResponse, error) {
	var inboundOrder = domain.InboundOrderResponse{}

	row := m.db.QueryRowContext(ctx, sqlInboundOrdersCountByEmployeeId, employeeId)

	err := row.Scan(
		&inboundOrder.ID,
		&inboundOrder.CardNumberId,
		&inboundOrder.FirstName,
		&inboundOrder.LastName,
		&inboundOrder.WarehouseId,
		&inboundOrder.InboundOrdersCount,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &inboundOrder, err
}
