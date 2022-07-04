package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
)

type repository struct{ db *sql.DB }

func NewMariaDBRepository(db *sql.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) (*[]domain.Product, error) {
	products := []domain.Product{}

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.ExpirationRate,
			&product.FreezingRate,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.ProductCode,
			&product.RecommendedFreezingTemperature,
			&product.Width,
			&product.ProductTypeId,
			&product.SellerId,
		); err != nil {
			return &products, err
		}

		products = append(products, product)
	}

	return &products, nil
}

func (r *repository) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	row := r.db.QueryRowContext(ctx, sqlGetById, id)

	product := domain.Product{}

	err := row.Scan(
		&product.Id,
		&product.Description,
		&product.ExpirationRate,
		&product.FreezingRate,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.ProductCode,
		&product.RecommendedFreezingTemperature,
		&product.Width,
		&product.ProductTypeId,
		&product.SellerId,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &product, domain.ErrIDNotFound
	}

	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (r *repository) CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	newProduct := domain.Product{
		Description:                    product.Description,
		ExpirationRate:                 product.ExpirationRate,
		FreezingRate:                   product.FreezingRate,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ProductCode:                    product.ProductCode,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		Width:                          product.Width,
		ProductTypeId:                  product.ProductTypeId,
		SellerId:                       product.SellerId,
	}
	result, err := r.db.ExecContext(
		ctx,
		sqlInsert,
		&newProduct.Description,
		&newProduct.ExpirationRate,
		&newProduct.FreezingRate,
		&newProduct.Height,
		&newProduct.Length,
		&newProduct.NetWeight,
		&newProduct.ProductCode,
		&newProduct.RecommendedFreezingTemperature,
		&newProduct.Width,
		&newProduct.ProductTypeId,
		&newProduct.SellerId,
	)
	if err != nil {
		return &newProduct, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return &newProduct, err
	}
	newProduct.Id = insertedId
	return &newProduct, nil
}

func (r *repository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	newProduct := domain.Product{
		Id:                             product.Id,
		Description:                    product.Description,
		ExpirationRate:                 product.ExpirationRate,
		FreezingRate:                   product.FreezingRate,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ProductCode:                    product.ProductCode,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		Width:                          product.Width,
		ProductTypeId:                  product.ProductTypeId,
		SellerId:                       product.SellerId,
	}

	result, err := r.db.ExecContext(
		ctx,
		sqlUpdate,
		&newProduct.Description,
		&newProduct.ExpirationRate,
		&newProduct.FreezingRate,
		&newProduct.Height,
		&newProduct.Length,
		&newProduct.NetWeight,
		&newProduct.ProductCode,
		&newProduct.RecommendedFreezingTemperature,
		&newProduct.Width,
		&newProduct.ProductTypeId,
		&newProduct.SellerId,
		&newProduct.Id,
	)
	if err != nil {
		return &newProduct, err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return &newProduct, domain.ErrIDNotFound
	}

	if err != nil {
		return &newProduct, err
	}

	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, sqlDelete, id)
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

func (r *repository) CreateProductRecords(ctx context.Context, record *domain.ProductRecords) (int64, error) {
	newRecord := domain.ProductRecords{
		PurchasePrice: record.PurchasePrice,
		SalePrice:     record.SalePrice,
		ProductId:     record.ProductId,
	}
	result, err := r.db.ExecContext(
		ctx,
		sqlCreateRecord,
		&newRecord.PurchasePrice,
		&newRecord.SalePrice,
		&newRecord.ProductId,
	)
	if err != nil {
		return 0, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}


func (r *repository) GetProductRecords(ctx context.Context, id int64) (*domain.ProductRecords, error) {
	row := r.db.QueryRowContext(ctx, sqlGetRecord, id)

	record := domain.ProductRecords{}

	err := row.Scan(
		&record.LastUpdateDate,
		&record.PurchasePrice,
		&record.SalePrice,
		&record.ProductId,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &record, domain.ErrIDNotFound
	}

	if err != nil {
		return &record, err
	}

	return &record, nil
}

func (r *repository) GetQtyOfRecords(ctx context.Context, id int64) (*domain.QtyOfRecords, error) {
	row := r.db.QueryRowContext(ctx, sqlGetQtyOfRecords, id)

	report := domain.QtyOfRecords{}

	err := row.Scan(
		&report.ProductId,
		&report.Description,
		&report.RecordsCount,

	)
	if errors.Is(err, sql.ErrNoRows) {
		return &report, domain.ErrIDNotFound
	}

	if err != nil {
		return &report, err
	}

	return &report, nil
}