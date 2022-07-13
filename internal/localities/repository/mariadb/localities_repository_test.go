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
	queryInsertLocality               = regexp.QuoteMeta(sqlCreateLocality)
	queryGetLocalityById              = regexp.QuoteMeta(sqlGetLocalityById)
	queryGetQtyOfSellersByLocalityId  = regexp.QuoteMeta(sqlGetQtyOfSellersByLocalityId)
	queryGetAllQtyOfSellersByLocality = regexp.QuoteMeta(sqlGetQtyOfSellersByLocality)
)

var rowsStructLocalityByID = []string{
	"id",
	"locality_name",
	"province_name",
	"country_name",
}

var rowsStructGetQtyOfSellers = []string{
	"locality_id",
	"locality_name",
	"sellers_count",
}

var rowsStructQtyOfSellerByLocalityByID = []string{
	"locality_name",
	"province_name",
	"country_name",
}

func TestCreateLocality(t *testing.T) {
	mockLocality := utils.CreateRandomLocality()

	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertLocality).
			WithArgs(
				mockLocality.LocalityName,
				mockLocality.ProvinceID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		localityId, err := repo.CreateLocality(context.Background(), &mockLocality)
		assert.NoError(t, err)

		assert.Equal(t, int64(1), localityId)
	})

	t.Run("failed to create locality", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertLocality).
			WithArgs(0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.CreateLocality(context.Background(), &mockLocality)

		assert.Error(t, err)
	})
}

func TestLocalityByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockGetLocality := utils.CreateRandomGetLocality()

		rows := sqlmock.NewRows(rowsStructLocalityByID).AddRow(
			mockGetLocality.ID,
			mockGetLocality.LocalityName,
			mockGetLocality.ProvinceName,
			mockGetLocality.CountryName,
		)

		mock.ExpectQuery(queryGetLocalityById).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		result, err := localityRepo.GetLocalityByID(context.Background(), 0)

		assert.NoError(t, err)

		assert.Equal(t, &mockGetLocality, result)
	})

	t.Run("fail to select locality", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetLocalityById).WillReturnError(sql.ErrNoRows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetLocalityByID(context.Background(), 0)
		assert.Error(t, err)
	})

	t.Run("fail to scan locality", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStructLocalityByID).AddRow("", "", "", "")

		mock.ExpectQuery(queryGetLocalityById).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetLocalityByID(context.Background(), 1)
		assert.Error(t, err)
	})
}

func TestAllQtyOfSellers(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockQtyOfSellers := utils.CreateRandomListQtyOfSellers()

		rows := sqlmock.NewRows(rowsStructGetQtyOfSellers)
		for _, mockQtyOfSeller := range mockQtyOfSellers {
			rows.AddRow(
				mockQtyOfSeller.LocalityID,
				mockQtyOfSeller.LocalityName,
				mockQtyOfSeller.SellersCount,
			)
		}

		mock.ExpectQuery(queryGetAllQtyOfSellersByLocality).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		result, err := localityRepo.GetAllQtyOfSellers(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockQtyOfSellers)
	})

	t.Run("fail to scan sellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStructGetQtyOfSellers).AddRow("", "", "")

		mock.ExpectQuery(queryGetAllQtyOfSellersByLocality).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetAllQtyOfSellers(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select sellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAllQtyOfSellersByLocality).WillReturnError(sql.ErrNoRows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetAllQtyOfSellers(context.Background())
		assert.Error(t, err)
	})
}

func TestGetQtyOfSellersByLocalityId(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockGetQtyOfSellers := utils.CreateRandomQtyOfSellers()

		rows := sqlmock.NewRows(rowsStructQtyOfSellerByLocalityByID).AddRow(
			mockGetQtyOfSellers.LocalityID,
			mockGetQtyOfSellers.LocalityName,
			mockGetQtyOfSellers.SellersCount,
		)

		mock.ExpectQuery(queryGetQtyOfSellersByLocalityId).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		result, err := localityRepo.GetQtyOfSellersByLocalityId(context.Background(), 0)

		assert.NoError(t, err)

		assert.Equal(t, &mockGetQtyOfSellers, result)
	})

	t.Run("fail to select QtyOfSellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetQtyOfSellersByLocalityId).WillReturnError(sql.ErrNoRows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetQtyOfSellersByLocalityId(context.Background(), 0)
		assert.Error(t, err)
	})

	t.Run("fail to scan QtyOfSellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsStructQtyOfSellerByLocalityByID).AddRow("", "", "")

		mock.ExpectQuery(queryGetQtyOfSellersByLocalityId).WillReturnRows(rows)

		localityRepo := NewMariaDBRepository(db)

		_, err = localityRepo.GetQtyOfSellersByLocalityId(context.Background(), 1)
		assert.Error(t, err)
	})
}
