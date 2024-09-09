package repository

import (
	"database/sql"
	"tourist-app/internal/models"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) GetPayments() ([]models.Payment, error) {
	rows, err := r.DB.Query("SELECT * FROM public.payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		if err := rows.Scan(&payment.PaymentID, &payment.BookingID, &payment.Date, &payment.Status, &payment.CreatedAt, &payment.UpdatedAt); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (r *PaymentRepository) GetPayment(id int) (*models.Payment, error) {
	row := r.DB.QueryRow("SELECT * FROM public.payments WHERE payment_id =$1", id)

	var payment models.Payment
	if err := row.Scan(&payment.PaymentID, &payment.BookingID, &payment.Amount, &payment.Date, &payment.Status, &payment.CreatedAt, &payment.UpdatedAt); err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) CreatePayment(payment *models.Payment) error {
	_, err := r.DB.Exec("INSERT INTO public.payments (booking_id, amount, date, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		payment.BookingID, payment.Amount, payment.Date, payment.Status, payment.CreatedAt, payment.UpdatedAt)
	return err
}

func (r *PaymentRepository) UpdatePayment(payment *models.Payment) error {
	_, err := r.DB.Exec("UPDATE public.payments SET booking_id = $1, amount = $2, date = $3, status = $4, updated_at = $5 WHERE payment_id = $6",
		payment.BookingID, payment.Amount, payment.Date, payment.Status, payment.UpdatedAt, payment.PaymentID)
	return err
}

func (r *PaymentRepository) DeletePayment(id int) error {
	_, err := r.DB.Exec("DELETE FROM public.payments WHERE payment_id = $1", id)
	return err
}
