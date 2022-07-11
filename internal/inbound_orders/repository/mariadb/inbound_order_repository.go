package mariadb

import (
	"context"
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.InboundOrderRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.InboundOrder, error) {
	var inboundOrders []domain.InboundOrder

	rows, err := m.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &inboundOrders, err
	}

	defer rows.Close()

	for rows.Next() {
		var inboundOrder domain.InboundOrder

		if err := rows.Scan(
			&inboundOrder.ID,
			&inboundOrder.OrderDate,
			&inboundOrder.OrderNumber,
			&inboundOrder.EmployeeId,
			&inboundOrder.ProductBatchId,
			&inboundOrder.WarehouseId,
		); err != nil {
			return &inboundOrders, err
		}

		inboundOrders = append(inboundOrders, inboundOrder)
	}
	return &inboundOrders, nil
}

func (m mariadbRepository) Create(ctx context.Context, inbounOrder *domain.InboundOrder) (*domain.InboundOrder, error) {
	newInboundOrder := domain.InboundOrder{}

	result, err := m.db.ExecContext(
		ctx,
		sqlInsert,
		&inbounOrder.OrderDate,
		&inbounOrder.OrderNumber,
		&inbounOrder.EmployeeId,
		&inbounOrder.ProductBatchId,
		&inbounOrder.WarehouseId,
	)
	if err != nil {
		return &newInboundOrder, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return &newInboundOrder, err
	}

	inbounOrder.ID = lastId

	return inbounOrder, nil
}
