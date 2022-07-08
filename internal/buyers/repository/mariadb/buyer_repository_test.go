package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsert  = regexp.QuoteMeta(sqlInsert)
	queryGetAll  = regexp.QuoteMeta(sqlGetAll)
	queryGetById = regexp.QuoteMeta(sqlGetById)
	queryUpdate  = regexp.QuoteMeta(sqlUpdate)
	queryDelete  = regexp.QuoteMeta(sqlDelete)
)

var rowsStruct = []string{
	"id",
	"card_number_id",
	"first_name",
	"last_name",
}

var rowsListReportPurchaseOrders = []string{
	"id",
	"card_number_id",
	"first_name",
	"last_name",
	"purchase_orders_count",
}

func TestCreateNewBuyer(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(
				mockBuyer.CardNumberID,
				mockBuyer.FirstName,
				mockBuyer.LastName,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Create(context.Background(), mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)
		assert.NoError(t, err)

		assert.Equal(t, &mockBuyer, sec)
	})

	t.Run("failed to create buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Create(context.Background(), mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockBuyers := utils.CreateRandomListBuyers()

		rows := sqlmock.NewRows(rowsStruct)
		for _, mockBuyer := range mockBuyers {
			rows.AddRow(
				mockBuyer.ID,
				mockBuyer.CardNumberID,
				mockBuyer.FirstName,
				mockBuyer.LastName,
			)
		}

		mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

		buyersRepo := NewMariaDBRepository(db)

		result, err := buyersRepo.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockBuyers)
	})

	t.Run("fail to scan buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "")

		mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.GetAll(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAll).WillReturnError(sql.ErrNoRows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.GetAll(context.Background())
		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockBuyer := utils.CreateRandomBuyer()

		rows := sqlmock.NewRows(rowsStruct).AddRow(
			mockBuyer.ID,
			mockBuyer.CardNumberID,
			mockBuyer.FirstName,
			mockBuyer.LastName,
		)

		mock.ExpectQuery(queryGetById).WillReturnRows(rows)

		buyersRepo := NewMariaDBRepository(db)

		result, err := buyersRepo.GetById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockBuyer)
	})

	t.Run("fail to scan buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "")

		mock.ExpectQuery(queryGetById).WillReturnRows(rows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.GetById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.GetById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestUpdateBuyer(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(
				mockBuyer.CardNumberID,
				mockBuyer.FirstName,
				mockBuyer.LastName,
				mockBuyer.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Update(context.Background(), mockBuyer.ID, mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)
		assert.NoError(t, err)

		assert.Equal(t, &mockBuyer, sec)
	})

	t.Run("fail to update buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), mockBuyer.ID, mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)
		assert.Error(t, err)
	})

	t.Run("Buyer not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(
				mockBuyer.CardNumberID,
				mockBuyer.FirstName,
				mockBuyer.LastName,
				mockBuyer.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), mockBuyer.ID, mockBuyer.CardNumberID, mockBuyer.FirstName, mockBuyer.LastName)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}

func TestDeleteBuyer(t *testing.T) {
	mockBuyer := utils.CreateRandomBuyer()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDelete).
			WithArgs(
				mockBuyer.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		err = repo.Delete(context.Background(), mockBuyer.ID)
		assert.NoError(t, err)
	})

	t.Run("fail to delete buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDelete).
			WithArgs(0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockBuyer.ID)
		assert.Error(t, err)
	})

	t.Run("Buyer not deleted", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDelete).
			WithArgs(mockBuyer.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockBuyer.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}

func TestGetQtyOfAllRecords(t *testing.T) {
	// t.Run("success", func(t *testing.T) {
	// 	db, mock, err := sqlmock.New()
	// 	assert.NoError(t, err)
	// 	defer db.Close()

	// 	mockListOfReportPurchaseOrders := utils.CreateRandomListReportPurchaseOrder()

	// 	rows := sqlmock.NewRows(rowsListReportPurchaseOrders)
	// 	for _, mockReport := range mockListOfReportPurchaseOrders {
	// 		rows.AddRow(
	// 			mockReport.ID,
	// 			mockReport.CardNumberID,
	// 			mockReport.FirstName,
	// 			mockReport.LastName,
	// 			mockReport.PurchaseOrdersCount,
	// 		)
	// 	}

	// 	mock.ExpectQuery(sqlFindAllPurchaseOrders).WillReturnRows(rows)

	// 	buyersRepo := NewMariaDBRepository(db)

	// 	result, err := buyersRepo.ReportAllPurchaseOrders(context.Background())
	// 	assert.NoError(t, err)

	// 	assert.Equal(t, result, &mockListOfReportPurchaseOrders)
	// })

	t.Run("fail to scan qty of all records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsListReportPurchaseOrders).AddRow("", "", "", "", "")

		mock.ExpectQuery(sqlFindAllPurchaseOrders).WillReturnRows(rows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.ReportAllPurchaseOrders(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select qty of all records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(sqlFindAllPurchaseOrders).WillReturnError(sql.ErrNoRows)

		buyersRepo := NewMariaDBRepository(db)

		_, err = buyersRepo.ReportAllPurchaseOrders(context.Background())
		assert.Error(t, err)
	})
}
