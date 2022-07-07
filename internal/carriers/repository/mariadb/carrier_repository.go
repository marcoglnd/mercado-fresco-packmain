package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"
)

type carrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) domain.CarrierRepository {
	return &carrierRepository{db: db}
}

func (r *carrierRepository) Create(
	ctx context.Context,
	carrier *domain.Carrier,
) (*domain.Carrier, error) {
	result, err := r.db.ExecContext(
		ctx,
		sqlStore,
		&carrier.Cid,
		&carrier.CompanyName,
		&carrier.Address,
		&carrier.Telephone,
		&carrier.LocalityId,
	)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	carrier.ID = lastID

	return carrier, nil
}

func (r *carrierRepository) FindById(
	ctx context.Context,
	id int64,
) (*domain.Carrier, error) {
	row := r.db.QueryRowContext(
		ctx, sqlGetById, id,
	)

	foundCarrier := &domain.Carrier{}
	err := row.Scan(
		&foundCarrier.ID,
		&foundCarrier.Cid,
		&foundCarrier.CompanyName,
		&foundCarrier.Address,
		&foundCarrier.Telephone,
		&foundCarrier.LocalityId,
	)

	if err != nil {
		return nil, err
	}

	return foundCarrier, nil
}

func (r *carrierRepository) FindByCid(
	ctx context.Context,
	cid string,
) (*domain.Carrier, error) {
	row := r.db.QueryRowContext(
		ctx, sqlGetByCid, cid,
	)

	foundCarrier := &domain.Carrier{}
	err := row.Scan(
		&foundCarrier.ID,
		&foundCarrier.Cid,
		&foundCarrier.CompanyName,
		&foundCarrier.Address,
		&foundCarrier.Telephone,
		&foundCarrier.LocalityId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return foundCarrier, nil
}

func (r *carrierRepository) GetAll(
	ctx context.Context,
) (*[]domain.Carrier, error) {
	carriers := []domain.Carrier{}

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &carriers, err
	}

	defer rows.Close()

	for rows.Next() {
		var carrier domain.Carrier

		if err := rows.Scan(
			&carrier.ID,
			&carrier.Cid,
			&carrier.CompanyName,
			&carrier.Address,
			&carrier.Telephone,
			&carrier.LocalityId,
		); err != nil {
			return &carriers, err
		}

		carriers = append(carriers, carrier)
	}

	return &carriers, nil
}
