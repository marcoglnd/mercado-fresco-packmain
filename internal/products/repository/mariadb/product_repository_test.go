package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsert = regexp.QuoteMeta(sqlInsert)
	queryGetAll = regexp.QuoteMeta(sqlGetAll)
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
	mockSection := utils.CreateRandomProduct()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(
				mockSection.Description,
				mockSection.ExpirationRate,
				mockSection.FreezingRate,
				mockSection.Height,
				mockSection.Length,
				mockSection.NetWeight,
				mockSection.ProductCode,
				mockSection.RecommendedFreezingTemperature,
				mockSection.Width,
				mockSection.ProductTypeId,
				mockSection.SellerId,
			).WillReturnResult(sqlmock.NewResult(1, 1)) // last id, // rows affected

		repo := NewMariaDBRepository(db)

		sec, err := repo.CreateNewProduct(context.TODO(), &mockSection)
		assert.NoError(t, err)

		assert.Equal(t, &mockSection, sec)
	})

	t.Run("failed to create product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.CreateNewProduct(context.TODO(), &mockSection)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
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
}

func TestGetAllFailScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(rowsStruct).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	productsRepo := NewMariaDBRepository(db)

	_, err = productsRepo.GetAll(context.Background())
	assert.Error(t, err)
}

func TestGetAllFailSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetAll).WillReturnError(sql.ErrNoRows)

	productsRepo := NewMariaDBRepository(db)

	_, err = productsRepo.GetAll(context.Background())
	assert.Error(t, err)
}
