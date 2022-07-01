package products

import (
	"context"
	"database/sql"
	"errors"
)

type repository struct{ db *sql.DB }

func (r *repository) GetAll(ctx context.Context) (*[]Product, error) {
	products := []Product{}

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return &products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

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

func (r *repository) GetById(ctx context.Context, id int64) (*Product, error) {
	row := r.db.QueryRowContext(ctx, sqlGetById, id)

	product := Product{}

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
		return &product, ErrIDNotFound
	}

	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (r *repository) CreateNewProduct(ctx context.Context, product *Product) (*Product, error) {
	newProduct := Product{
		Description: product.Description,
		ExpirationRate: product.ExpirationRate,
		FreezingRate: product.FreezingRate,
		Height: product.Height,
		Length: product.Length,
		NetWeight: product.NetWeight,
		ProductCode: product.ProductCode,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		Width: product.Width,
		ProductTypeId: product.ProductTypeId,
		SellerId: product.SellerId,
	}
	result, err := r.db.ExecContext(
		ctx,
		sqlStore,
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

func (r *repository) Update(ctx context.Context, product *Product) (*Product, error) {
	newProduct := Product{
		Description: product.Description,
		ExpirationRate: product.ExpirationRate,
		FreezingRate: product.FreezingRate,
		Height: product.Height,
		Length: product.Length,
		NetWeight: product.NetWeight,
		ProductCode: product.ProductCode,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		Width: product.Width,
		ProductTypeId: product.ProductTypeId,
		SellerId: product.SellerId,
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
	)
	if err != nil {
		return &newProduct, err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return &newProduct, ErrIDNotFound
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
		return ErrIDNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
