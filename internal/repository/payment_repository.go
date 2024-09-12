package repository

import (
	"database/sql"
	"tourist-app/internal/models"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) GetAll() ([]models.Payment, error) {
	rows, err := r.db.Query("SELECT id, booking_id, amount, method, status, paid_at FROM payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var p models.Payment
		err := rows.Scan(&p.ID, &p.BookingID, &p.Amount, &p.Method, &p.Status, &p.PaidAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}

func (r *PaymentRepository) GetByID(id int) (*models.Payment, error) {
	var p models.Payment
	err := r.db.QueryRow("SELECT id, booking_id, amount, method, status, paid_at FROM payments WHERE id = $1", id).
		Scan(&p.ID, &p.BookingID, &p.Amount, &p.Method, &p.Status, &p.PaidAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
	_, err := r.db.Exec("INSERT INTO payments (booking_id, amount, method, status, paid_at) VALUES ($1, $2, $3, $4, $5)",
		payment.BookingID, payment.Amount, payment.Method, payment.Status, payment.PaidAt)
	return err
}

func (r *PaymentRepository) Update(payment *models.Payment) error {
	_, err := r.db.Exec("UPDATE payments SET booking_id = $1, amount = $2, method = $3, status = $4, paid_at = $5 WHERE id = $6",
		payment.BookingID, payment.Amount, payment.Method, payment.Status, payment.PaidAt, payment.ID)
	return err
}

func (r *PaymentRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM payments WHERE id = $1", id)
	return err
}
