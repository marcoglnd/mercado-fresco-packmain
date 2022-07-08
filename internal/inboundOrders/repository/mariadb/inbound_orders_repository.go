package mariadb

import (
	"context"
	"database/sql"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/inboundOrders/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.InboundOrdersRepository {
	return mariadbRepository{db: db}
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
