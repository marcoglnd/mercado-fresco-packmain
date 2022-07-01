package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.BuyerRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	var buyers []domain.Buyer = []domain.Buyer{}

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM buyers")
	if err != nil {
		return &buyers, err
	}

	defer rows.Close()

	for rows.Next() {
		var buyer domain.Buyer

		if err := rows.Scan(
			&buyer.ID,
			&buyer.CardNumberID,
			&buyer.FirstName,
			&buyer.LastName,
		); err != nil {
			return &buyers, err
		}

		buyers = append(buyers, buyer)
	}

	return &buyers, nil
}

func (m mariadbRepository) GetById(ctx context.Context, id int64) (*domain.Buyer, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM buyers WHERE ID = ?", id)

	var buyer domain.Buyer

	err := row.Scan(
		&buyer.ID,
		&buyer.CardNumberID,
		&buyer.FirstName,
		&buyer.LastName,
	)
	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &buyer, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &buyer, err
	}

	return &buyer, nil
}

func (m mariadbRepository) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	var newBuyer = domain.Buyer{
		CardNumberID: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}

	query := `INSERT INTO buyers 
	(card_number_id, first_name, last_name) VALUES (?, ?, ?)`

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newBuyer.CardNumberID,
		&newBuyer.FirstName,
		&newBuyer.LastName,
	)
	if err != nil {
		return &newBuyer, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return &newBuyer, err
	}

	newBuyer.ID = lastID

	return &newBuyer, nil
}

func (m mariadbRepository) Update(ctx context.Context, id int64, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	var newBuyer = domain.Buyer{
		ID:           id,
		CardNumberID: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}

	query := `UPDATE buyers SET 
	card_number_id=?, first_name=?, last_name=? WHERE id=?`

	result, err := m.db.ExecContext(
		ctx,
		query,
		&newBuyer.CardNumberID,
		&newBuyer.FirstName,
		&newBuyer.LastName,
		&newBuyer.ID,
	)
	if err != nil {
		return &newBuyer, err
	}

	affectedRows, err := result.RowsAffected()
	// ID not found
	if affectedRows == 0 {
		return &newBuyer, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &newBuyer, err
	}

	return &newBuyer, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, "DELETE FROM buyers WHERE id=?", id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return err
	}

	return nil
}
