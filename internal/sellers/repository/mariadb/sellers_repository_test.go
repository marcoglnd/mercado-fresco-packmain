package mariadb

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsertSeller  = regexp.QuoteMeta(sqlInsertSeller)
	queryGetAllSellers = regexp.QuoteMeta(sqlGetAllSellers)
	queryGetSellerById = regexp.QuoteMeta(sqlGetSellerById)
	queryUpdateSeller  = regexp.QuoteMeta(sqlUpdateSeller)
	queryDeleteSeller  = regexp.QuoteMeta(sqlDeleteSeller)
)

func TestCreate(t *testing.T) {
	mockSeller := utils.CreateRandomSeller()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSeller).
			WithArgs(
				mockSeller.Cid,
				mockSeller.Company_name,
				mockSeller.Address,
				mockSeller.Telephone,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		sellersRepo := NewMariaDBRepository(db)

		seller, err := sellersRepo.Create(context.TODO(), &mockSeller)
		assert.Error(t, err)

		assert.Equal(t, &mockSeller, seller)
	})

	t.Run("failed to create", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSeller).
			WithArgs(0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sellersRepo := NewMariaDBRepository(db)
		_, err = sellersRepo.Create(context.TODO(), &mockSeller)

		assert.Error(t, err)
	})
}
