package payment

import (
	"github.com/Plutokim/sport/src/repos/payment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Read(c *gin.Context, db *gorm.DB) {
	userId := c.Param("user-id")
	paymentRepo := payment.New(db)
	payment, err := paymentRepo.FindByUserId(userId)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": payment})
}
