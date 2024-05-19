package payment

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TableName = "payment"
)

type Repo struct {
	Conn *gorm.DB
}

func New(conn *gorm.DB) *Repo {
	return &Repo{Conn: conn}
}

func (r *Repo) Create(userID string, cardToken, cardNumber, expiryDate string) (*Payment, error) {
	dbPayment, err := r.FindByUserId(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if dbPayment != nil {
		return nil, errors.New("user already has payment method")
	}
	parsedId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	payment := NewPayment(parsedId, cardToken, cardNumber, expiryDate)
	err = r.Conn.Table(TableName).Create(payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *Repo) FindByUserId(id string) (*Payment, error) {
	payment := &Payment{}
	err := r.Conn.Table(TableName).Where("user_id = ?", id).First(payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *Repo) FindAll() ([]Payment, error) {
	var payments []Payment
	err := r.Conn.Table(TableName).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *Repo) DeleteByUserId(id string) error {
	return r.Conn.Table(TableName).Where("user_id = ?", id).Delete(&Payment{}).Error
}

func (r *Repo) Update(userID string, cardToken, cardNumber, expiryDate string) error {
	dbPayment, err := r.FindByUserId(userID)
	if err != nil {
		return err
	}
	dbPayment.CardToken = cardToken
	dbPayment.CardNumber = cardNumber
	dbPayment.ExpiryDate = expiryDate
	err = r.Conn.Table(TableName).Where("id = ?", dbPayment.ID).Updates(dbPayment).Error
	if err != nil {
		return err
	}
	return nil
}
