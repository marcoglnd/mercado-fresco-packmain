package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"

func CreateRandomEmployee() domain.Employee {
	employee := domain.Employee{
		ID:           1,
		CardNumberId: RandomString(3),
		FirstName:    RandomString(6),
		LastName:     RandomString(6),
		WarehouseId:  RandomInt64(),
	}
	return employee
}

func CreateRandomListEmployees() []domain.Employee {
	var listOfEmployees []domain.Employee
	for i := 1; i <= 5; i++ {
		employee := CreateRandomEmployee()
		employee.ID = int64(i)
		listOfEmployees = append(listOfEmployees, employee)
	}
	return listOfEmployees
}
