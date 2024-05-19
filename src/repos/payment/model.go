package payment

import "github.com/google/uuid"

type Payment struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID     uuid.UUID `json:"user_id"`
	CardToken  string    `json:"card_token"`
	CardNumber string    `json:"card_number"`
	ExpiryDate string    `json:"expiry_date"`
}

func NewPayment(userID uuid.UUID, cardToken, cardNumber, expiryDate string) *Payment {
	return &Payment{
		ID:         uuid.New(),
		UserID:     userID,
		CardToken:  cardToken,
		CardNumber: cardNumber,
		ExpiryDate: expiryDate,
	}
}
