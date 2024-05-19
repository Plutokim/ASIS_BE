package payment

import (
	"github.com/Plutokim/sport/src/repos/payment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context, db *gorm.DB) {
	userId := c.Param("user-id")
	paymentRepo := payment.New(db)
	err := paymentRepo.DeleteByUserId(userId)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
