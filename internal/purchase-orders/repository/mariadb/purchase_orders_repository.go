package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.PurchaseOrderRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetByOrderNumber(
	ctx context.Context,
	orderNumber string,
) (*domain.PurchaseOrder, error) {
	row := m.db.QueryRowContext(
		ctx, sqlGetByOrderNumber, orderNumber,
	)

	foundPurchaseOrder := &domain.PurchaseOrder{}
	err := row.Scan(
		&foundPurchaseOrder.ID,
		&foundPurchaseOrder.OrderNumber,
		&foundPurchaseOrder.OrderDate,
		&foundPurchaseOrder.TrackingCode,
		&foundPurchaseOrder.BuyerId,
		&foundPurchaseOrder.CarrierId,
		&foundPurchaseOrder.OrderStatusId,
		&foundPurchaseOrder.WarehouseId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return foundPurchaseOrder, nil
}

func (m mariadbRepository) Create(
	ctx context.Context,
	orderNumber,
	orderDate,
	trackingCode string,
	buyerId,
	carrierId,
	orderStatusId,
	warehouseId int64,
) (*domain.PurchaseOrder, error) {
	var newPurchaseOrder = domain.PurchaseOrder{
		OrderNumber:   orderNumber,
		OrderDate:     orderDate,
		TrackingCode:  trackingCode,
		BuyerId:       buyerId,
		CarrierId:     carrierId,
		OrderStatusId: orderStatusId,
		WarehouseId:   warehouseId,
	}

	query := sqlInsert

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newPurchaseOrder.OrderNumber,
		&newPurchaseOrder.OrderDate,
		&newPurchaseOrder.TrackingCode,
		&newPurchaseOrder.BuyerId,
		&newPurchaseOrder.CarrierId,
		&newPurchaseOrder.OrderStatusId,
		&newPurchaseOrder.WarehouseId,
	)
	if err != nil {
		return &newPurchaseOrder, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return &newPurchaseOrder, err
	}

	newPurchaseOrder.ID = lastID

	return &newPurchaseOrder, nil
}
