package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *repository) GetById(ctx context.Context, id int) (*Product, error) {
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
		return &product, errors.New("section id not found")
	}

	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (r *repository) CreateNewProduct(ctx context.Context, product *Product) (*Product, error) {
	newProduct := Product{}
	result, err := r.db.ExecContext(
		ctx,
		sqlStore,
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
	newProduct := Product{}

	result, err := r.db.ExecContext(
		ctx,
		sqlUpdate,
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
		&product.Id,
	)
	if err != nil {
		return &newProduct, err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return &newProduct, errors.New("section id not found")
	}

	if err != nil {
		return &newProduct, err
	}

	return product, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range listOfProducts {
		if listOfProducts[i].Id == int64(id) {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}
	listOfProducts = append(listOfProducts[:index], listOfProducts[index+1:]...)
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
