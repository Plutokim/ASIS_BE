package user

import (
	"gorm.io/gorm"
)

const (
	TableName = "users"
)

type Repo struct {
	Conn *gorm.DB
}

func New(conn *gorm.DB) *Repo {
	return &Repo{Conn: conn}
}

func (r *Repo) Create(fullName, phoneNumber, email, password string) (*User, error) {
	user := NewUser(fullName, phoneNumber, email, password)
	err := r.Conn.Table(TableName).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) FindById(id string) (*User, error) {
	user := &User{}
	err := r.Conn.Table(TableName).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) FindAll() ([]User, error) {
	var users []User
	err := r.Conn.Table(TableName).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repo) Delete(id string) error {
	return r.Conn.Table(TableName).Where("id = ?", id).Delete(&User{}).Error
}

func (r *Repo) Update(id, fullName, phoneNumber, email, password string) error {
	dbUser, err := r.FindById(id)
	if err != nil {
		return err
	}
	dbUser.FullName = fullName
	dbUser.PhoneNumber = phoneNumber
	dbUser.Email = email
	dbUser.Password = password
	err = r.Conn.Table(TableName).Where("id = ?", id).Updates(dbUser).Error
	if err != nil {
		return err
	}
	return nil
}
