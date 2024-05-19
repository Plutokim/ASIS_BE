package user

import (
	"github.com/Plutokim/sport/src/repos/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context, db *gorm.DB) {
	userRepo := user.New(db)
	users, err := userRepo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": users})
}
