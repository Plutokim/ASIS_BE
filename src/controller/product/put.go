package product

import (
	"github.com/Plutokim/sport/src/repos/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context, db *gorm.DB) {
	id := c.Param("product-id")
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	productRepo := product.New(db)
	err := productRepo.Update(id, body.Name, body.Description, body.ImageFolderURL, body.Price, body.AvailableQuantity)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
