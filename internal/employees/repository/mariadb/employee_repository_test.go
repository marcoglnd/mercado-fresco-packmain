package mariadb

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewEmployee(t *testing.T) {
	mockEmployee := utils.CreateRandomEmployee()

	t.Run("In case of success", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		newEmployee, err := repository.Create(context.Background(), &mockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, &mockEmployee, newEmployee)
	})

	t.Run("In case of error", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			"123", "Liz", "Souza", 42,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		_, err = repository.Create(context.Background(), &mockEmployee)

		assert.Error(t, err)
	})
}
