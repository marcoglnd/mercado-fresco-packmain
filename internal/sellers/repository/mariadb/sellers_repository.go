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

	rows, err := m.db.QueryContext(ctx, sqlGetAllSellers)
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
			&seller.LocalityID,
		)
		if err != nil {
			return &sellers, err
		}

		sellers = append(sellers, seller)
	}

	return &sellers, nil
}

func (m mariadbRepository) GetByID(ctx context.Context, id int64) (*domain.Seller, error) {
	row := m.db.QueryRowContext(ctx, sqlGetSellerById, id)

	seller := domain.Seller{}

	err := row.Scan(
		&seller.ID,
		&seller.Cid,
		&seller.Company_name,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityID,
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
	newSeller := domain.Seller{
		Cid:          seller.Cid,
		Company_name: seller.Company_name,
		Address:      seller.Address,
		Telephone:    seller.Telephone,
		LocalityID:   seller.LocalityID,
	}

	query := sqlInsertSeller

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newSeller.Cid,
		&newSeller.Company_name,
		&newSeller.Address,
		&newSeller.Telephone,
		&newSeller.LocalityID,
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
	newSeller := domain.Seller{
		ID:           seller.ID,
		Cid:          seller.Cid,
		Company_name: seller.Company_name,
		Address:      seller.Address,
		Telephone:    seller.Telephone,
		LocalityID:   seller.LocalityID,
	}

	query := sqlUpdateSeller

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newSeller.Cid,
		&newSeller.Company_name,
		&newSeller.Address,
		&newSeller.Telephone,
		&newSeller.LocalityID,
		&newSeller.ID,
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
	result, err := m.db.ExecContext(ctx, sqlDeleteSeller, id)
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
