package product

import (
	"gorm.io/gorm"
)

const (
	TableName = "product"
)

type Repo struct {
	Conn *gorm.DB
}

func New(conn *gorm.DB) *Repo {
	return &Repo{Conn: conn}
}

func (r *Repo) Create(name, description, imageFolderURL string, price float64, availableQuantity int) (*Product, error) {
	product := NewProduct(name, description, imageFolderURL, price, availableQuantity)
	err := r.Conn.Table(TableName).Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repo) FindById(id string) (*Product, error) {
	product := &Product{}
	err := r.Conn.Table(TableName).Where("id = ?", id).First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repo) FindAll() ([]Product, error) {
	var products []Product
	err := r.Conn.Table(TableName).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repo) Delete(id string) error {
	return r.Conn.Table(TableName).Where("id = ?", id).Delete(&Product{}).Error
}

func (r *Repo) Update(id, name, description, imageFolderURL string, price float64, availableQuantity int) error {
	dbProduct, err := r.FindById(id)
	if err != nil {
		return err
	}
	dbProduct.Name = name
	dbProduct.Description = description
	dbProduct.ImageFolderURL = imageFolderURL
	dbProduct.Price = price
	dbProduct.AvailableQuantity = availableQuantity
	err = r.Conn.Table(TableName).Where("id = ?", id).Updates(dbProduct).Error
	if err != nil {
		return err
	}
	return nil
}
