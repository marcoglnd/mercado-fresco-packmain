package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
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

var rowsStruct = []string{
	"id",
	"cid",
	"company_name",
	"address",
	"telephone",
	"locality_id",
}

func TestGetAll(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockSellers := utils.CreateRandomListSeller()

		rows := sqlmock.NewRows(rowsStruct)
		for _, mockSeller := range mockSellers {
			rows.AddRow(
				mockSeller.ID,
				mockSeller.Cid,
				mockSeller.Company_name,
				mockSeller.Address,
				mockSeller.Telephone,
				mockSeller.LocalityID,
			)
		}

		mock.ExpectQuery(queryGetAllSellers).WillReturnRows(rows)

		sellersRepo := NewMariaDBRepository(db)

		result, err := sellersRepo.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockSellers)
	})

	t.Run("fail to scan seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(queryGetAllSellers).WillReturnRows(rows)

		sellersRepo := NewMariaDBRepository(db)

		_, err = sellersRepo.GetAll(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAllSellers).WillReturnError(sql.ErrNoRows)

		sellersRepo := NewMariaDBRepository(db)

		_, err = sellersRepo.GetAll(context.Background())
		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockSeller := utils.CreateRandomSeller()

		rows := sqlmock.NewRows(rowsStruct).AddRow(
			mockSeller.ID,
			mockSeller.Cid,
			mockSeller.Company_name,
			mockSeller.Address,
			mockSeller.Telephone,
			mockSeller.LocalityID,
		)

		mock.ExpectQuery(queryGetSellerById).WillReturnRows(rows)

		sellersRepo := NewMariaDBRepository(db)

		result, err := sellersRepo.GetByID(context.Background(), 0)

		assert.NoError(t, err)

		assert.Equal(t, result, &mockSeller)
	})

	t.Run("fail to scan seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(queryGetSellerById).WillReturnRows(rows)

		sellersRepo := NewMariaDBRepository(db)

		_, err = sellersRepo.GetByID(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetSellerById).WillReturnError(sql.ErrNoRows)

		sellersRepo := NewMariaDBRepository(db)

		_, err = sellersRepo.GetByID(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestCreate(t *testing.T) {
	mockSeller := utils.CreateRandomSeller()

	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSeller).
			WithArgs(
				mockSeller.Cid,
				mockSeller.Company_name,
				mockSeller.Address,
				mockSeller.Telephone,
				mockSeller.LocalityID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		seller, err := repo.Create(context.Background(), &mockSeller)
		assert.NoError(t, err)

		assert.Equal(t, &mockSeller, seller)
	})

	t.Run("fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSeller).
			WithArgs(0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Create(context.Background(), &mockSeller)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mockSeller := utils.CreateRandomSeller()

	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSeller).
			WithArgs(
				mockSeller.Cid,
				mockSeller.Company_name,
				mockSeller.Address,
				mockSeller.Telephone,
				mockSeller.LocalityID,
				mockSeller.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		sellersRepo := NewMariaDBRepository(db)

		sec, err := sellersRepo.Update(context.Background(), &mockSeller)
		assert.NoError(t, err)

		assert.Equal(t, &mockSeller, sec)
	})

	t.Run("fail to update seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSeller).
			WithArgs(0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sellersRepo := NewMariaDBRepository(db)
		_, err = sellersRepo.Update(context.Background(), &mockSeller)
		assert.Error(t, err)
	})

	t.Run("Seller not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSeller).
			WithArgs(
				mockSeller.Cid,
				mockSeller.Company_name,
				mockSeller.Address,
				mockSeller.Telephone,
				mockSeller.LocalityID,
				mockSeller.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sellersRepo := NewMariaDBRepository(db)
		_, err = sellersRepo.Update(context.Background(), &mockSeller)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}

func TestDelete(t *testing.T) {
	mockSeller := utils.CreateRandomSeller()

	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSeller).
			WithArgs(
				mockSeller.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		sellersRepo := NewMariaDBRepository(db)

		err = sellersRepo.Delete(context.Background(), mockSeller.ID)
		assert.NoError(t, err)
	})

	t.Run("fail to delete seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSeller).
			WithArgs(0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sellersRepo := NewMariaDBRepository(db)
		err = sellersRepo.Delete(context.Background(), mockSeller.ID)
		assert.Error(t, err)
	})

	t.Run("Seller not deleted", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSeller).
			WithArgs(mockSeller.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sellersRepo := NewMariaDBRepository(db)
		err = sellersRepo.Delete(context.Background(), mockSeller.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}
