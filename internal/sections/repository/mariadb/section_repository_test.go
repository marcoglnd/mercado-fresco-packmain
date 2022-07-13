package mariadb

import (
	"context"
	"database/sql"

	// "fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsertSection  = regexp.QuoteMeta(sqlInsertSection)
	queryGetAllSections = regexp.QuoteMeta(sqlGetAllSections)
	queryGetSectionById = regexp.QuoteMeta(sqlGetSectionById)
	queryUpdateSection  = regexp.QuoteMeta(sqlUpdateSection)
	queryDeleteSection  = regexp.QuoteMeta(sqlDeleteSection)
)

var rowsSectionStruct = []string{
	"id",
	"section_number",
	"current_temperature",
	"minimum_temperature",
	"current_capacity",
	"minimum_capacity",
	"maximum_capacity",
	"warehouse_id",
	"product_type_id",
}

func TestCreateNewSection(t *testing.T) {
	mockSection := utils.CreateRandomSection()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSection).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseId,
				mockSection.ProductTypeId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		section, err := repo.Create(context.Background(), &mockSection)
		assert.NoError(t, err)

		assert.Equal(t, &mockSection, section)
	})

	t.Run("failed to create section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsertSection).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Create(context.Background(), &mockSection)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockSections := utils.CreateRandomListSection()

		rows := sqlmock.NewRows(rowsSectionStruct)
		for _, mockSection := range mockSections {
			rows.AddRow(
				mockSection.ID,
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseId,
				mockSection.ProductTypeId,
			)
		}

		mock.ExpectQuery(queryGetAllSections).WillReturnRows(rows)

		sectionsRepo := NewMariaDBRepository(db)

		result, err := sectionsRepo.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result, &mockSections)
	})

	t.Run("fail to scan section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsSectionStruct).AddRow("", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(queryGetAllSections).WillReturnRows(rows)

		sectionsRepo := NewMariaDBRepository(db)

		_, err = sectionsRepo.GetAll(context.Background())
		assert.Error(t, err)
	})

	t.Run("fail to select section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetAllSections).WillReturnError(sql.ErrNoRows)

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

		mockSection := utils.CreateRandomSection()

		rows := sqlmock.NewRows(rowsSectionStruct).AddRow(
			mockSection.ID,
			mockSection.SectionNumber,
			mockSection.CurrentTemperature,
			mockSection.MinimumTemperature,
			mockSection.CurrentCapacity,
			mockSection.MinimumCapacity,
			mockSection.MaximumCapacity,
			mockSection.WarehouseId,
			mockSection.ProductTypeId,
		)

		mock.ExpectQuery(queryGetSectionById).WillReturnRows(rows)

		sectiontsRepo := NewMariaDBRepository(db)

		result, err := sectiontsRepo.GetById(context.Background(), 0)
		assert.NoError(t, err)

		assert.Equal(t, result, &mockSection)
	})

	t.Run("fail to scan section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows(rowsSectionStruct).AddRow("", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(queryGetSectionById).WillReturnRows(rows)

		sectionsRepo := NewMariaDBRepository(db)

		_, err = sectionsRepo.GetById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("fail to select section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(queryGetSectionById).WillReturnError(sql.ErrNoRows)

		sectionsRepo := NewMariaDBRepository(db)

		_, err = sectionsRepo.GetById(context.Background(), 0)
		assert.Error(t, err)
	})
}

func TestUpdateSection(t *testing.T) {
	mockSection := utils.CreateRandomSection()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSection).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseId,
				mockSection.ProductTypeId,
				mockSection.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Update(context.Background(), &mockSection)
		assert.NoError(t, err)

		assert.Equal(t, &mockSection, sec)
	})

	t.Run("fail to update section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSection).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), &mockSection)
		assert.Error(t, err)
	})

	t.Run("Section not updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryUpdateSection).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseId,
				mockSection.ProductTypeId,
				mockSection.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		_, err = repo.Update(context.Background(), &mockSection)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}

func TestDeleteSection(t *testing.T) {
	mockSection := utils.CreateRandomSection()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSection).
			WithArgs(
				mockSection.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)

		err = repo.Delete(context.Background(), mockSection.ID)
		assert.NoError(t, err)
	})

	t.Run("fail to delete section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSection).
			WithArgs(0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockSection.ID)
		assert.Error(t, err)
	})

	t.Run("Section not update", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryDeleteSection).
			WithArgs(mockSection.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewMariaDBRepository(db)
		err = repo.Delete(context.Background(), mockSection.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrIDNotFound, err)
	})
}
