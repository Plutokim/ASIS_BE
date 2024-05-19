package product

import "github.com/google/uuid"

type Product struct {
	ID                uuid.UUID `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name"`
	Price             float64   `json:"price"`
	Description       string    `json:"description"`
	ImageFolderURL    string    `json:"image_folder_url"`
	AvailableQuantity int       `json:"available_quantity"`
}

func NewProduct(name, description, imageFolderURL string, price float64, availableQuantity int) *Product {
	return &Product{
		ID:                uuid.New(),
		Name:              name,
		Price:             price,
		Description:       description,
		ImageFolderURL:    imageFolderURL,
		AvailableQuantity: availableQuantity,
	}
}
