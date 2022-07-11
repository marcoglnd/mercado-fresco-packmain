package domain

import "context"

type EmployeeRepository interface {
	GetAll(ctx context.Context) (*[]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	GetByCardNumberId(ctx context.Context, cardNumberId string) (*Employee, error)
	Create(ctx context.Context, employee *Employee) (*Employee, error)
	Update(ctx context.Context, employee *Employee) (*Employee, error)
	Delete(ctx context.Context, id int64) error
	ReportAllInboundOrders(ctx context.Context) (*[]InboundOrder, error)
	ReportInboundOrders(ctx context.Context, employeeId int64) (*InboundOrder, error)
}

type EmployeeService interface {
	GetAll(ctx context.Context) (*[]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, employee *Employee) (*Employee, error)
	Update(ctx context.Context, employee *Employee) (*Employee, error)
	Delete(ctx context.Context, id int64) error
	ReportAllInboundOrders(ctx context.Context) (*[]InboundOrder, error)
	ReportInboundOrders(ctx context.Context, employeeId int64) (*InboundOrder, error)
}
