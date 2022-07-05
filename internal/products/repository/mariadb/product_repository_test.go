package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsertProduct  = regexp.QuoteMeta(sqlInsertProduct)
	queryGetAllProducts = regexp.QuoteMeta(sqlGetAllProducts)
	queryGetProductById = regexp.QuoteMeta(sqlGetProductById)
	queryUpdateProduct  = regexp.QuoteMeta(sqlUpdateProduct)
	queryDeleteProduct  = regexp.QuoteMeta(sqlDeleteProduct)
	queryInsertRecord   = regexp.QuoteMeta(sqlCreateRecord)
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

		mock.ExpectExec(queryInsertProduct).
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

		product, err := repo.CreateNewProduct(context.Background(), &mockProduct)
		assert.NoError(t, err)

		assert.Equal(t, &mockProduct, product)
	})

	t.Run("failed to create product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertProduct).
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

		mock.ExpectQuery(queryGetAllProducts).WillReturnRows(rows)

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

		mock.ExpectQuery(queryGetAllProducts).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetAll(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAllProducts).WillReturnError(sql.ErrNoRows)

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

		mock.ExpectQuery(queryGetProductById).WillReturnRows(rows)

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

		mock.ExpectQuery(queryGetProductById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetProductById).WillReturnError(sql.ErrNoRows)

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

		mock.ExpectExec(queryUpdateProduct).
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
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Update(context.Background(), &mockProduct)
		assert.NoError(t, err)

		assert.Equal(t, &mockProduct, sec)
	})

	t.Run("fail to update product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateProduct).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), &mockProduct)
		assert.Error(t, err)
	})

	t.Run("Product not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateProduct).
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

func TestDeleteProduct(t *testing.T) {
	mockProduct := utils.CreateRandomProduct()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteProduct).
			WithArgs(
				mockProduct.Id,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		err = repo.Delete(context.Background(), mockProduct.Id)
		assert.NoError(t, err)
	})

	t.Run("fail to delete product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteProduct).
			WithArgs(0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockProduct.Id)
		assert.Error(t, err)
	})

	t.Run("Product not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteProduct).
			WithArgs(mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockProduct.Id)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}

func TestCreateProductRecords(t *testing.T) {
	mockProductRecords := utils.CreateRandomProductRecords()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertRecord).
			WithArgs(
				mockProductRecords.PurchasePrice,
				mockProductRecords.SalePrice,
				mockProductRecords.ProductId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		recordId, err := repo.CreateProductRecords(context.Background(), &mockProductRecords)
		assert.NoError(t, err)

		assert.True(t, fmt.Sprintf("%T", recordId) == "int64" && recordId > 0)
	})

	t.Run("failed to create product record", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertRecord).
			WithArgs(0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.CreateProductRecords(context.Background(), &mockProductRecords)

		assert.Error(t, err)
	})
}
