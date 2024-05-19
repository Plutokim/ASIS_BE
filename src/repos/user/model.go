package user

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}

func NewUser(fullName, phoneNumber, email, password string) *User {
	return &User{
		ID:          uuid.New(),
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    password,
	}
}
