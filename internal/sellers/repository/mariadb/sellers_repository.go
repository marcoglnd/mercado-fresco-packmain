package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.SellerRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	sellers := []domain.Seller{}

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM sellers")
	if err != nil {
		return &sellers, err
	}

	defer rows.Close()

	for rows.Next() {
		var seller domain.Seller

		err := rows.Scan(
			&seller.ID,
			&seller.Cid,
			&seller.Company_name,
			&seller.Address,
			&seller.Telephone,
		)
		if err != nil {
			return &sellers, err
		}

		sellers = append(sellers, seller)
	}

	return &sellers, nil
}

func (m mariadbRepository) GetByID(ctx context.Context, id int64) (*domain.Seller, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM sellers WHERE ID = ?", id)

	seller := domain.Seller{}

	err := row.Scan(
		&seller.ID,
		&seller.Cid,
		&seller.Company_name,
		&seller.Address,
		&seller.Telephone,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &seller, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &seller, err
	}

	return &seller, nil
}

func (m mariadbRepository) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	newSeller := domain.Seller{}

	query := `INSERT INTO sellers 
	(cid, company_name, address, telephone) VALUES (?, ?, ?, ?)`

	result, err := m.db.ExecContext(
		ctx,
		query,
		newSeller.Cid,
		newSeller.Company_name,
		newSeller.Address,
		newSeller.Telephone,
	)
	if err != nil {
		return &newSeller, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return &newSeller, err
	}

	newSeller.ID = lastID

	return &newSeller, nil
}

func (m mariadbRepository) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	newSeller := domain.Seller{}

	query := `UPDATE sellers SET
	cid=?, company_name=?, address=?, telephone=? WHERE id=?`

	result, err := m.db.ExecContext(
		ctx,
		query,
		newSeller.Cid,
		newSeller.Company_name,
		newSeller.Address,
		newSeller.Telephone,
		newSeller.ID,
	)
	if err != nil {
		return &newSeller, err
	}

	affectedRows, err := result.RowsAffected()
	// ID not found
	if affectedRows == 0 {
		return &newSeller, domain.ErrIDNotFound
	}

	// other errors
	if err != nil {
		return &newSeller, err
	}

	return &newSeller, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, "DELETE FROM sellers WHERE id=?", id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return err
	}
	return nil
}
