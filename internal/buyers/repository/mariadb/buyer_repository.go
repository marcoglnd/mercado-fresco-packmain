package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.BuyerRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	var buyers []domain.Buyer = []domain.Buyer{}

	rows, err := m.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &buyers, err
	}

	defer rows.Close()

	for rows.Next() {
		var buyer domain.Buyer

		if err := rows.Scan(
			&buyer.ID,
			&buyer.CardNumberID,
			&buyer.FirstName,
			&buyer.LastName,
		); err != nil {
			return &buyers, err
		}

		buyers = append(buyers, buyer)
	}

	return &buyers, nil
}

func (m mariadbRepository) GetById(ctx context.Context, id int64) (*domain.Buyer, error) {
	row := m.db.QueryRowContext(ctx, sqlGetById, id)

	var buyer domain.Buyer

	err := row.Scan(
		&buyer.ID,
		&buyer.CardNumberID,
		&buyer.FirstName,
		&buyer.LastName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return &buyer, domain.ErrIDNotFound
	}

	if err != nil {
		return &buyer, err
	}

	return &buyer, nil
}

func (m mariadbRepository) GetByCardNumberId(
	ctx context.Context,
	cardNumberId string,
) (*domain.Buyer, error) {
	row := m.db.QueryRowContext(
		ctx, sqlGetByCardNumberId, cardNumberId,
	)

	foundBuyer := &domain.Buyer{}
	err := row.Scan(
		&foundBuyer.ID,
		&foundBuyer.CardNumberID,
		&foundBuyer.FirstName,
		&foundBuyer.LastName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return foundBuyer, nil
}

func (m mariadbRepository) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	var newBuyer = domain.Buyer{
		CardNumberID: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}

	query := sqlInsert

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newBuyer.CardNumberID,
		&newBuyer.FirstName,
		&newBuyer.LastName,
	)
	if err != nil {
		return &newBuyer, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return &newBuyer, err
	}

	newBuyer.ID = lastID

	return &newBuyer, nil
}

func (m mariadbRepository) Update(ctx context.Context, id int64, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	var newBuyer = domain.Buyer{
		ID:           id,
		CardNumberID: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}

	query := sqlUpdate

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newBuyer.CardNumberID,
		&newBuyer.FirstName,
		&newBuyer.LastName,
		&newBuyer.ID,
	)
	if err != nil {
		return &newBuyer, err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return &newBuyer, domain.ErrIDNotFound
	}

	if err != nil {
		return &newBuyer, err
	}

	return &newBuyer, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (m mariadbRepository) ReportAllPurchaseOrders(ctx context.Context) (*[]domain.PurchaseOrdersResponse, error) {
	var report = []domain.PurchaseOrdersResponse{}

	rows, err := m.db.QueryContext(ctx, sqlFindAllPurchaseOrders)

	if err != nil {
		return &report, err
	}

	for rows.Next() {
		var response = domain.PurchaseOrdersResponse{}

		if err := rows.Scan(
			&response.ID,
			&response.CardNumberID,
			&response.FirstName,
			&response.LastName,
			&response.PurchaseOrdersCount,
		); err != nil {
			return &report, err
		}

		report = append(report, response)
	}

	return &report, nil
}

func (m mariadbRepository) ReportPurchaseOrders(ctx context.Context, buyerId int64) (*domain.PurchaseOrdersResponse, error) {
	var response = domain.PurchaseOrdersResponse{}
	row := m.db.QueryRowContext(ctx, sqlFindPurchaseOrderByBuyerId, buyerId)

	err := row.Scan(
		&response.ID,
		&response.CardNumberID,
		&response.FirstName,
		&response.LastName,
		&response.PurchaseOrdersCount,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &response, err
}
