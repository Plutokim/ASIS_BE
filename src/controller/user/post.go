package user

import (
	"github.com/Plutokim/sport/src/repos/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func Create(c *gin.Context, db *gorm.DB) {
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	repo := user.New(db)
	_, err := repo.Create(body.FullName, body.PhoneNumber, body.Email, body.Password)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
