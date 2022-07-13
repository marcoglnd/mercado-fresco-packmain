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
	queryGetRecordsById = regexp.QuoteMeta(sqlGetRecord)

	queryGetQtyOfRecordsById = regexp.QuoteMeta(sqlGetQtyOfRecordsById)
	queryGetQtyOfAllRecords  = regexp.QuoteMeta(sqlGetQtyOfRecords)

	queryInsertBatch    = regexp.QuoteMeta(sqlCreateBatch)
	queryGetBatchesById = regexp.QuoteMeta(sqlGetBatch)

	queryGetQtdProductsBySectionId = regexp.QuoteMeta(sqlGetQtdProductsBySectionId)
	queryGetQtdProductsInSection  = regexp.QuoteMeta(sqlGetQtdProductsInSection)
)

var rowsProductStruct = []string{
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

var rowsProductRecordStruct = []string{
	"last_update_date",
	"purchase_price",
	"sale_price",
	"product_id",
}

var rowsQtyOfRecordsStruct = []string{
	"id",
	"description",
	"records_count",
}

var rowsProductBatchesStruct = []string{
	"batch_number",
	"current_quantity",
	"current_temperature",
	"due_date",
	"initial_quantity",
	"manufacturing_date",
	"manufacturing_hour",
	"minimum_temperature",
	"product_id",
	"section_id",
}

var rowsQtdOfProducts = []string{
	"id",
	"section_number",
	"products_count",
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

		rows := sqlmock.NewRows(rowsProductStruct)
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

		rows := sqlmock.NewRows(rowsProductStruct).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

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

		rows := sqlmock.NewRows(rowsProductStruct).AddRow(
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

		rows := sqlmock.NewRows(rowsProductStruct).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

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

func TestGetProductRecordsById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockProductRecords := utils.CreateRandomProductRecords()

		rows := sqlmock.NewRows(rowsProductRecordStruct).AddRow(
			mockProductRecords.LastUpdateDate,
			mockProductRecords.PurchasePrice,
			mockProductRecords.SalePrice,
			mockProductRecords.ProductId,
		)

		mock.ExpectQuery(queryGetRecordsById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetProductRecordsById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockProductRecords)
	})

	t.Run("fail to scan product records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsProductRecordStruct).AddRow("", "", "", "")

		mock.ExpectQuery(queryGetRecordsById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetProductRecordsById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select product records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetRecordsById).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetProductRecordsById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestGetQtyOfRecordsById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockQtyOfRecords := utils.CreateRandomQtyOfRecords()

		rows := sqlmock.NewRows(rowsQtyOfRecordsStruct).AddRow(
			mockQtyOfRecords.ProductId,
			mockQtyOfRecords.Description,
			mockQtyOfRecords.RecordsCount,
		)

		mock.ExpectQuery(queryGetQtyOfRecordsById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetQtyOfRecordsById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockQtyOfRecords)
	})

	t.Run("fail to scan qty Of records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsQtyOfRecordsStruct).AddRow("", "", "")

		mock.ExpectQuery(queryGetQtyOfRecordsById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtyOfRecordsById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select qty Of records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetQtyOfRecordsById).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtyOfRecordsById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestGetQtyOfAllRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockQtyOfAllRecords := utils.CreateRandomListQtyOfRecords()

		rows := sqlmock.NewRows(rowsQtyOfRecordsStruct)
		for _, mockQtyOfAllRecord := range mockQtyOfAllRecords {
			rows.AddRow(
				mockQtyOfAllRecord.ProductId,
				mockQtyOfAllRecord.Description,
				mockQtyOfAllRecord.RecordsCount,
			)
		}

		mock.ExpectQuery(queryGetQtyOfAllRecords).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetQtyOfAllRecords(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockQtyOfAllRecords)
	})

	t.Run("fail to scan qty of all records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsQtyOfRecordsStruct).AddRow("", "", "")

		mock.ExpectQuery(queryGetQtyOfAllRecords).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtyOfAllRecords(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select qty of all records", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetQtyOfAllRecords).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtyOfAllRecords(context.Background())
		assert.Error(t, err)
	})
}

func TestCreateProductBatches(t *testing.T) {
	mockProductBatches := utils.CreateRandomProductBatches()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertBatch).
			WithArgs(
				mockProductBatches.BatchNumber,
				mockProductBatches.CurrentQuantity,
				mockProductBatches.CurrentTemperature,
				mockProductBatches.DueDate,
				mockProductBatches.InitialQuantity,
				mockProductBatches.ManufacturingDate,
				mockProductBatches.ManufacturingHour,
				mockProductBatches.MinimumTemperature,
				mockProductBatches.ProductId,
				mockProductBatches.SectionId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		batchId, err := repo.CreateProductBatches(context.Background(), &mockProductBatches)
		assert.NoError(t, err)

		assert.True(t, fmt.Sprintf("%T", batchId) == "int64" && batchId > 0)
	})

	t.Run("failed to create product batch", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertBatch).
			WithArgs(0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.CreateProductBatches(context.Background(), &mockProductBatches)

		assert.Error(t, err)
	})
}

func TestGetProductBatchesById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockProductBatches := utils.CreateRandomProductBatches()

		rows := sqlmock.NewRows(rowsProductBatchesStruct).AddRow(
			mockProductBatches.BatchNumber,
			mockProductBatches.CurrentQuantity,
			mockProductBatches.CurrentTemperature,
			mockProductBatches.DueDate,
			mockProductBatches.InitialQuantity,
			mockProductBatches.ManufacturingDate,
			mockProductBatches.ManufacturingHour,
			mockProductBatches.MinimumTemperature,
			mockProductBatches.ProductId,
			mockProductBatches.SectionId,
		)

		mock.ExpectQuery(queryGetBatchesById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetProductBatchesById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockProductBatches)
	})

	t.Run("fail to scan product batch", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsProductBatchesStruct).AddRow("", "", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(queryGetBatchesById).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetProductBatchesById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select product batches", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetBatchesById).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetProductBatchesById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestGetQtdProductsBySectionId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockQtdOfProducts := utils.CreateRandomQtdOfProducts()

		rows := sqlmock.NewRows(rowsQtdOfProducts).AddRow(
			mockQtdOfProducts.SectionId,
			mockQtdOfProducts.SectionNumber,
			mockQtdOfProducts.ProductsCount,
		)

		mock.ExpectQuery(queryGetQtdProductsBySectionId).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetQtdProductsBySectionId(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockQtdOfProducts)
	})

	t.Run("fail to scan qtd Of reports", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsQtdOfProducts).AddRow("", "", "")

		mock.ExpectQuery(queryGetQtdProductsBySectionId).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtdProductsBySectionId(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select qtd Of reports", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetQtdProductsBySectionId).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtdProductsBySectionId(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestGetQtdProductsInSection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockQtdProductsInSection := utils.CreateRandomListQtdOfProducts()

		rows := sqlmock.NewRows(rowsQtyOfRecordsStruct)
		for _, mockQtyOfAllReport := range mockQtdProductsInSection {
			rows.AddRow(
				mockQtyOfAllReport.SectionId,
				mockQtyOfAllReport.SectionNumber,
				mockQtyOfAllReport.ProductsCount,
			)
		}

		mock.ExpectQuery(queryGetQtdProductsInSection).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		result, err := productsRepo.GetQtdOfAllProducts(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockQtdProductsInSection)
	})

	t.Run("fail to scan qtd of all reports", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsQtdOfProducts).AddRow("", "", "")

		mock.ExpectQuery(queryGetQtdProductsInSection).WillReturnRows(rows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtdOfAllProducts(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select qtd of all reports", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetQtdProductsInSection).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)

		_, err = productsRepo.GetQtdOfAllProducts(context.Background())
		assert.Error(t, err)
	})
}