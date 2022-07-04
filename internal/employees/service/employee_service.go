package service

import (
	"context"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
)

type employeeService struct {
	repository domain.EmployeeRepository
}

func NewEmployeeService(er domain.EmployeeRepository) domain.EmployeeService {
	return &employeeService{repository: er}
}

func (e employeeService) GetAll(ctx context.Context) (*[]domain.Employee, error) {
	employees, err := e.repository.GetAll(ctx)

	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (e employeeService) GetById(ctx context.Context, id int) (*domain.Employee, error) {
	employee, err := e.repository.GetById(ctx, id)

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (e employeeService) Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	employee, err := e.repository.Create(ctx, employee)

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (e employeeService) Update(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	employee, err := e.repository.Update(ctx, employee)

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (e employeeService) Delete(ctx context.Context, id int) error {
	err := e.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
