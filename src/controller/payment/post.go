package payment

import (
	"github.com/Plutokim/sport/src/repos/payment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	CardToken  string `json:"card_token"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
}

func Create(c *gin.Context, db *gorm.DB) {
	userId := c.Param("user-id")
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	paymentRepo := payment.New(db)
	_, err := paymentRepo.Create(userId, body.CardToken, body.CardNumber, body.ExpiryDate)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
