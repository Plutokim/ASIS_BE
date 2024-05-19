package payment

import (
	"github.com/Plutokim/sport/src/repos/payment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context, db *gorm.DB) {
	userId := c.Param("user-id")
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	repo := payment.New(db)
	err := repo.Update(userId, body.CardToken, body.CardNumber, body.ExpiryDate)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
