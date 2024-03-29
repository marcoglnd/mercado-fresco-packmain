package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	carrierFake := utils.CreateRandomCarrier()

	t.Run("Must create carrier", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				&carrierFake.Cid,
				&carrierFake.CompanyName,
				&carrierFake.Address,
				&carrierFake.Telephone,
				&carrierFake.LocalityId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		carriersRepo := NewCarrierRepository(db)

		ca, err := carriersRepo.Create(context.TODO(), &carrierFake)
		assert.NoError(t, err)
		assert.Equal(t, carrierFake.Address, ca.Address)
	})

	t.Run("Must fail on carrier context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				&carrierFake.Cid,
				&carrierFake.CompanyName,
				&carrierFake.Address,
				&carrierFake.Telephone,
				&carrierFake.LocalityId,
			).WillReturnError(errors.New("fail db"))

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.Create(context.TODO(), &carrierFake)
		assert.Error(t, err)
	})

	t.Run("Must fail on getting last insert id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				&carrierFake.Cid,
				&carrierFake.CompanyName,
				&carrierFake.Address,
				&carrierFake.Telephone,
				&carrierFake.LocalityId,
			).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.Create(context.TODO(), &carrierFake)
		assert.Error(t, err)
	})

}

func TestFindById(t *testing.T) {
	carrierFake := utils.CreateRandomCarrier()

	t.Run("Must find carrier", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"cid",
				"company_name",
				"address",
				"telephone",
				"locality_id",
			},
		).AddRow(
			carrierFake.ID,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).
			WithArgs(
				carrierFake.ID,
			).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		ca, err := carriersRepo.FindById(context.TODO(), carrierFake.ID)
		assert.NoError(t, err)
		assert.Equal(t, carrierFake.Address, ca.Address)
	})

	t.Run("Must fail finding carrier", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).
			WithArgs(
				carrierFake.ID,
			).WillReturnError(errors.New("fail"))

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.FindById(context.TODO(), carrierFake.ID)
		assert.Error(t, err)
	})
}

func TestFindByCid(t *testing.T) {
	carrierFake := utils.CreateRandomCarrier()

	t.Run("Must find carrier", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"cid",
				"company_name",
				"address",
				"telephone",
				"locality_id",
			},
		).AddRow(
			carrierFake.ID,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByCid)).
			WithArgs(
				carrierFake.Cid,
			).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		ca, err := carriersRepo.FindByCid(context.TODO(), carrierFake.Cid)
		assert.NoError(t, err)
		assert.Equal(t, carrierFake.Address, ca.Address)
	})

	t.Run("Must not fail when there are no carriers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByCid)).
			WithArgs(
				carrierFake.Cid,
			).WillReturnError(sql.ErrNoRows)

		carriersRepo := NewCarrierRepository(db)

		ca, err := carriersRepo.FindByCid(context.TODO(), carrierFake.Cid)
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Nil(t, ca)
	})

	t.Run("Must not fail when there are no carriers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByCid)).
			WithArgs(
				carrierFake.Cid,
			).WillReturnError(errors.New("fail"))

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.FindByCid(context.TODO(), carrierFake.Cid)
		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	carrierFake := utils.CreateRandomCarrier()

	t.Run("Must find carriers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"cid",
				"company_name",
				"address",
				"telephone",
				"locality_id",
			},
		).AddRow(
			carrierFake.ID,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		).AddRow(
			carrierFake.ID+1,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		cas, err := carriersRepo.GetAll(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(*cas))
	})

	t.Run("Must fail finding carriers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnError(sql.ErrConnDone)

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.GetAll(context.TODO())
		assert.Error(t, err)
	})

	t.Run("Must return some carrier and fail some queries", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"cid",
				"company_name",
				"address",
				"telephone",
				"locality_id",
			},
		).AddRow(
			-1,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		).AddRow(
			nil,
			carrierFake.Cid,
			carrierFake.CompanyName,
			carrierFake.Address,
			carrierFake.Telephone,
			carrierFake.LocalityId,
		).RowError(2, errors.New("row error"))

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		cas, err := carriersRepo.GetAll(context.TODO())
		assert.Error(t, err)
		assert.NotNil(t, cas)
	})
}

func TestGetAllReports(t *testing.T) {
	reportFake := utils.CreateRandomCarrierReport()

	t.Run("Must find carriers report", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"locality_name",
				"carriers_count",
			},
		).AddRow(
			reportFake.LocalityId,
			reportFake.LocalityName,
			reportFake.CarriersCount,
		).AddRow(
			reportFake.LocalityId+1,
			reportFake.LocalityName,
			reportFake.CarriersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlCarriersCountAll)).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		rep, err := carriersRepo.GetAllCarriersReport(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(*rep))
	})

	t.Run("Must fail finding carriers report", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlCarriersCountAll)).WillReturnError(sql.ErrConnDone)

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.GetAllCarriersReport(context.TODO())
		assert.Error(t, err)
	})

	t.Run("Must return some carrier report and fail some queries", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"locality_name",
				"carriers_count",
			},
		).AddRow(
			-1,
			reportFake.LocalityName,
			reportFake.CarriersCount,
		).AddRow(
			nil,
			reportFake.LocalityName,
			reportFake.CarriersCount,
		).RowError(2, errors.New("row error"))

		mock.ExpectQuery(regexp.QuoteMeta(sqlCarriersCountAll)).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		rep, err := carriersRepo.GetAllCarriersReport(context.TODO())
		assert.Error(t, err)
		assert.NotNil(t, rep)
	})
}

func TestGetAllReportsById(t *testing.T) {
	reportFake := utils.CreateRandomCarrierReport()

	t.Run("Must find carrier report", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"locality_name",
				"carriers_count",
			},
		).AddRow(
			reportFake.LocalityId,
			reportFake.LocalityName,
			reportFake.CarriersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlCarriersCountById)).
			WithArgs(
				reportFake.LocalityId,
			).WillReturnRows(fakeRows)

		carriersRepo := NewCarrierRepository(db)

		ca, err := carriersRepo.GetCarriersReportById(context.TODO(), reportFake.LocalityId)
		assert.NoError(t, err)
		assert.Equal(t, reportFake.CarriersCount, ca.CarriersCount)
	})

	t.Run("Must fail getting a report", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlCarriersCountById)).
			WithArgs(
				reportFake.LocalityId,
			).WillReturnError(errors.New("fail"))

		carriersRepo := NewCarrierRepository(db)

		_, err = carriersRepo.GetCarriersReportById(context.TODO(), reportFake.LocalityId)
		assert.Error(t, err)
	})
}
