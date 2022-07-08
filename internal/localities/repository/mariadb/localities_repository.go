package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.LocalityRepository {
	return &mariadbRepository{db: db}
}

func (r *mariadbRepository) CreateLocality(ctx context.Context, local *domain.Locality) (int64, error) {
	newLocal := domain.Locality{
		LocalityName: local.LocalityName,
		ProvinceID:   local.ProvinceID,
	}
	result, err := r.db.ExecContext(
		ctx,
		sqlCreateLocality,
		&newLocal.LocalityName,
		&newLocal.ProvinceID,
	)
	if err != nil {
		return 0, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}

func (m mariadbRepository) GetLocalityByID(ctx context.Context, id int64) (*domain.GetLocality, error) {
	row := m.db.QueryRowContext(ctx, sqlGetLocalityById, id)

	getLocality := domain.GetLocality{}

	err := row.Scan(
		&getLocality.ID,
		&getLocality.LocalityName,
		&getLocality.ProvinceName,
		&getLocality.CountryName,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &getLocality, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &getLocality, err
	}

	return &getLocality, nil
}

func (m mariadbRepository) GetQtyOfSellers(ctx context.Context) (*[]domain.QtyOfSellers, error) {
	listOfSellers := []domain.QtyOfSellers{}

	rows, err := m.db.QueryContext(ctx, sqlGetQtyOfSellersLocalityId)
	if err != nil {
		return &listOfSellers, err
	}

	defer rows.Close()

	for rows.Next() {
		var seller domain.QtyOfSellers

		err := rows.Scan(
			&seller.LocalityID,
			&seller.LocalityName,
			&seller.SellersCount,
		)
		if err != nil {
			return &listOfSellers, err
		}

		listOfSellers = append(listOfSellers, seller)
	}

	return &listOfSellers, nil
}

func (m mariadbRepository) GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*domain.QtyOfSellers, error) {
	row := m.db.QueryRowContext(ctx, sqlGetQtyOfSellersByLocalityId, id)

	sellers := domain.QtyOfSellers{}

	err := row.Scan(
		&sellers.LocalityID,
		&sellers.LocalityName,
		&sellers.SellersCount,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &sellers, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &sellers, err
	}

	return &sellers, nil
}
