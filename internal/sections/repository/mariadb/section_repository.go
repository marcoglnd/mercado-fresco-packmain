package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
)

// var sectionList []Section = []Section{}

type repository struct{ db *sql.DB }

func NewMariaDBRepository(db *sql.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) (*[]domain.Section, error) {
	sections := []domain.Section{}

	rows, err := r.db.QueryContext(ctx, sqlGetAllSections)
	if err != nil {
		return &sections, err
	}

	defer rows.Close()

	for rows.Next() {
		var section domain.Section

		if err := rows.Scan(
			&section.ID,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseId,
			&section.ProductTypeId,
		); err != nil {
			return &sections, err
		}

		sections = append(sections, section)
	}

	return &sections, nil
}

func (r *repository) GetById(ctx context.Context, id int64) (*domain.Section, error) {
	row := r.db.QueryRowContext(ctx, sqlGetSectionById, id)

	section := domain.Section{}

	err := row.Scan(
		&section.ID,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &section, domain.ErrIDNotFound
	}

	if err != nil {
		return &section, err
	}

	return &section, nil
}

func (r *repository) Create(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	newSection := domain.Section{
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseId:        section.WarehouseId,
		ProductTypeId:      section.ProductTypeId,
	}

	result, err := r.db.ExecContext(
		ctx,
		sqlInsertSection,
		&newSection.SectionNumber,
		&newSection.CurrentTemperature,
		&newSection.MinimumTemperature,
		&newSection.CurrentCapacity,
		&newSection.MinimumCapacity,
		&newSection.MaximumCapacity,
		&newSection.WarehouseId,
		&newSection.ProductTypeId,
	)
	if err != nil {
		return &newSection, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return &newSection, err
	}
	newSection.ID = insertedId
	return &newSection, nil
}

func (r *repository) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	newSection := domain.Section{
		ID:                 section.ID,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseId:        section.WarehouseId,
		ProductTypeId:      section.ProductTypeId,
	}

	result, err := r.db.ExecContext(
		ctx,
		sqlUpdateSection,
		&newSection.SectionNumber,
		&newSection.CurrentTemperature,
		&newSection.MinimumTemperature,
		&newSection.CurrentCapacity,
		&newSection.MinimumCapacity,
		&newSection.MaximumCapacity,
		&newSection.WarehouseId,
		&newSection.ProductTypeId,
		&newSection.ID,
	)
	if err != nil {
		return &newSection, err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return &newSection, domain.ErrIDNotFound
	}

	if err != nil {
		return &newSection, err
	}

	return section, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, sqlDeleteSection, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	if err != nil {
		return err
	}

	return nil
}
