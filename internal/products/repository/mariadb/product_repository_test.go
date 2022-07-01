package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsert  = regexp.QuoteMeta(sqlInsert)
	queryGetAll  = regexp.QuoteMeta(sqlGetAll)
	queryGetById = regexp.QuoteMeta(sqlGetById)
	queryUpdate  = regexp.QuoteMeta(sqlUpdate)
)

var rowsStruct = []string{
	"id",
	"description",
	"expiration_rate",
	"freezing_rate",
	"height",
	"length",
	"net_weight",
	"product_code",
	"recommended_freezing_temperature",
	"width",
	"product_type_id",
	"seller_id",
}

func TestCreateNewProduct(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(
				mockProduct.Description,
				mockProduct.ExpirationRate,
				mockProduct.FreezingRate,
				mockProduct.Height,
				mockProduct.Length,
				mockProduct.NetWeight,
				mockProduct.ProductCode,
				mockProduct.RecommendedFreezingTemperature,
				mockProduct.Width,
				mockProduct.ProductTypeId,
				mockProduct.SellerId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.CreateNewProduct(context.Background(), &mockProduct)
		assert.NoError(t, err)

		assert.Equal(t, &mockProduct, sec)
	})

	t.Run("failed to create product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.CreateNewProduct(context.Background(), &mockProduct)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockProducts := utils.CreateRandomListProduct()

		rows := sqlmock.NewRows(rowsStruct)
		for _, mockProduct := range mockProducts {
			rows.AddRow(
				mockProduct.Id,
				mockProduct.Description,
				mockProduct.ExpirationRate,
				mockProduct.FreezingRate,
				mockProduct.Height,
				mockProduct.Length,
				mockProduct.NetWeight,
				mockProduct.ProductCode,
				mockProduct.RecommendedFreezingTemperature,
				mockProduct.Width,
				mockProduct.ProductTypeId,
				mockProduct.SellerId,
			)
		}

		mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockProducts)
	})

	t.Run("fail to scan product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetAll(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAll).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetAll(context.Background())
		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockProduct := utils.CreateRandomProduct()

		rows := sqlmock.NewRows(rowsStruct).AddRow(
			mockProduct.Id,
			mockProduct.Description,
			mockProduct.ExpirationRate,
			mockProduct.FreezingRate,
			mockProduct.Height,
			mockProduct.Length,
			mockProduct.NetWeight,
			mockProduct.ProductCode,
			mockProduct.RecommendedFreezingTemperature,
			mockProduct.Width,
			mockProduct.ProductTypeId,
			mockProduct.SellerId,
		)

		mock.ExpectQuery(queryGetById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockProduct)
	})

	t.Run("fail to scan product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(queryGetById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(
				mockProduct.Description,
				mockProduct.ExpirationRate,
				mockProduct.FreezingRate,
				mockProduct.Height,
				mockProduct.Length,
				mockProduct.NetWeight,
				mockProduct.ProductCode,
				mockProduct.RecommendedFreezingTemperature,
				mockProduct.Width,
				mockProduct.ProductTypeId,
				mockProduct.SellerId,
				mockProduct.Id,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Update(context.Background(), &mockProduct)
		assert.NoError(t, err)

		assert.Equal(t, &mockProduct, sec)
	})

	t.Run("fail to update product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), &mockProduct)
		assert.Error(t, err)
	})

	t.Run("Product not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdate).
			WithArgs(
				mockProduct.Description,
				mockProduct.ExpirationRate,
				mockProduct.FreezingRate,
				mockProduct.Height,
				mockProduct.Length,
				mockProduct.NetWeight,
				mockProduct.ProductCode,
				mockProduct.RecommendedFreezingTemperature,
				mockProduct.Width,
				mockProduct.ProductTypeId,
				mockProduct.SellerId,
				mockProduct.Id,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), &mockProduct)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}
