package product

import (
	"github.com/Plutokim/sport/src/repos/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context, db *gorm.DB) {
	productRepo := product.New(db)
	products, err := productRepo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": products})
}
