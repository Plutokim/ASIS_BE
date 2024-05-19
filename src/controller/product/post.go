package product

import (
	"github.com/Plutokim/sport/src/repos/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Description       string  `json:"description"`
	ImageFolderURL    string  `json:"image_folder_url"`
	AvailableQuantity int     `json:"available_quantity"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	repo := product.New(db)
	_, err := repo.Create(body.Name, body.Description, body.ImageFolderURL, body.Price, body.AvailableQuantity)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
