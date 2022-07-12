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
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must create warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				warehouseFake.Address,
				warehouseFake.Telephone,
				warehouseFake.WarehouseCode,
				warehouseFake.MinimumCapacity,
				warehouseFake.MinimumTemperature,
				warehouseFake.LocalityId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		warehousesRepo := NewWarehouseRepository(db)

		wh, err := warehousesRepo.Create(context.TODO(), &warehouseFake)
		assert.NoError(t, err)
		assert.Equal(t, warehouseFake.Address, wh.Address)
	})

	t.Run("Must fail when inserting invalid data", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				0,
				0,
				0,
				0,
				0,
				0,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		warehousesRepo := NewWarehouseRepository(db)

		_, err = warehousesRepo.Create(context.TODO(), &warehouseFake)
		assert.Error(t, err)
	})

	t.Run("Must fail on warehouse creation", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlStore)).
			WithArgs(
				warehouseFake.Address,
				warehouseFake.Telephone,
				warehouseFake.WarehouseCode,
				warehouseFake.MinimumCapacity,
				warehouseFake.MinimumTemperature,
				warehouseFake.LocalityId,
			).WillReturnResult(sqlmock.NewErrorResult(errors.New("fail")))

		warehousesRepo := NewWarehouseRepository(db)

		_, err = warehousesRepo.Create(context.TODO(), &warehouseFake)
		assert.Error(t, err)
	})

}

func TestUpdate(t *testing.T) {
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must update warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).
			WithArgs(
				warehouseFake.WarehouseCode,
				warehouseFake.Address,
				warehouseFake.Telephone,
				warehouseFake.MinimumCapacity,
				warehouseFake.MinimumTemperature,
				warehouseFake.ID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Update(context.TODO(), &warehouseFake)
		assert.NoError(t, err)
	})

	t.Run("Must fail updating with invalid data", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).
			WithArgs(
				0,
				0,
				0,
				0,
				0,
				0,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Update(context.TODO(), &warehouseFake)
		assert.Error(t, err)
	})

}

func TestFindById(t *testing.T) {
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must find warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"address",
				"telephone",
				"warehouse_code",
				"minimum_capacity",
				"minimum_temperature",
				"locality_id",
			},
		).AddRow(
			warehouseFake.ID,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).
			WithArgs(
				warehouseFake.ID,
			).WillReturnRows(fakeRows)

		warehousesRepo := NewWarehouseRepository(db)

		wh, err := warehousesRepo.FindById(context.TODO(), warehouseFake.ID)
		assert.NoError(t, err)
		assert.Equal(t, warehouseFake.Address, wh.Address)
	})

	t.Run("Must fail finding warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).
			WithArgs(
				warehouseFake.ID,
			).WillReturnError(errors.New("fail"))

		warehousesRepo := NewWarehouseRepository(db)

		_, err = warehousesRepo.FindById(context.TODO(), warehouseFake.ID)
		assert.Error(t, err)
	})
}

func TestFindByWarehouseCode(t *testing.T) {
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must find warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"address",
				"telephone",
				"warehouse_code",
				"minimum_capacity",
				"minimum_temperature",
				"locality_id",
			},
		).AddRow(
			warehouseFake.ID,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByWarehouseCode)).
			WithArgs(
				warehouseFake.WarehouseCode,
			).WillReturnRows(fakeRows)

		warehousesRepo := NewWarehouseRepository(db)

		wh, err := warehousesRepo.FindByWarehouseCode(context.TODO(), warehouseFake.WarehouseCode)
		assert.NoError(t, err)
		assert.Equal(t, warehouseFake.Address, wh.Address)
	})

	t.Run("Must fail finding warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByWarehouseCode)).
			WithArgs(
				warehouseFake.WarehouseCode,
			).WillReturnError(errors.New("fail"))

		warehousesRepo := NewWarehouseRepository(db)

		_, err = warehousesRepo.FindByWarehouseCode(context.TODO(), warehouseFake.WarehouseCode)
		assert.Error(t, err)
	})

	t.Run("Must not fail when there are no warehouses", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetByWarehouseCode)).
			WithArgs(
				warehouseFake.WarehouseCode,
			).WillReturnError(sql.ErrNoRows)

		warehousesRepo := NewWarehouseRepository(db)

		wh, err := warehousesRepo.FindByWarehouseCode(context.TODO(), warehouseFake.WarehouseCode)
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Nil(t, wh)
	})
}

func TestGetAll(t *testing.T) {
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must find warehouses", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"address",
				"telephone",
				"warehouse_code",
				"minimum_capacity",
				"minimum_temperature",
				"locality_id",
			},
		).AddRow(
			warehouseFake.ID,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		).AddRow(
			warehouseFake.ID+1,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(fakeRows)

		warehousesRepo := NewWarehouseRepository(db)

		whs, err := warehousesRepo.GetAll(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(*whs))
	})

	t.Run("Must fail finding warehouses", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnError(sql.ErrConnDone)

		warehousesRepo := NewWarehouseRepository(db)

		_, err = warehousesRepo.GetAll(context.TODO())
		assert.Error(t, err)
	})

	t.Run("Must return some warehouse and fail some queries", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		fakeRows := sqlmock.NewRows(
			[]string{
				"id",
				"address",
				"telephone",
				"warehouse_code",
				"minimum_capacity",
				"minimum_temperature",
				"locality_id",
			},
		).AddRow(
			-1,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		).AddRow(
			nil,
			warehouseFake.Address,
			warehouseFake.Telephone,
			warehouseFake.WarehouseCode,
			warehouseFake.MinimumCapacity,
			warehouseFake.MinimumTemperature,
			warehouseFake.LocalityId,
		).RowError(2, errors.New("row error"))

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(fakeRows)

		warehousesRepo := NewWarehouseRepository(db)

		whs, err := warehousesRepo.GetAll(context.TODO())
		assert.Error(t, err)
		assert.NotNil(t, whs)
	})
}

func TestDelete(t *testing.T) {
	warehouseFake := utils.CreateRandomWarehouse()

	t.Run("Must delete warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(warehouseFake.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Delete(context.TODO(), warehouseFake.ID)
		assert.NoError(t, err)
	})

	t.Run("Must fail deleting warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(warehouseFake.ID).WillReturnError(errors.New("fail"))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Delete(context.TODO(), warehouseFake.ID)
		assert.Error(t, err)
	})

	t.Run("Must fail when no warehouse exists", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(warehouseFake.ID).WillReturnResult(sqlmock.NewResult(0, 0))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Delete(context.TODO(), warehouseFake.ID)
		assert.Error(t, err)
	})

	t.Run("Must fail checking if warehouses were affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(warehouseFake.ID).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		warehousesRepo := NewWarehouseRepository(db)

		err = warehousesRepo.Delete(context.TODO(), warehouseFake.ID)
		assert.Error(t, err)
	})
}
