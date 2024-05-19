package user

import (
	"github.com/Plutokim/sport/src/repos/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context, db *gorm.DB) {
	id := c.Param("user-id")
	userRepo := user.New(db)
	err := userRepo.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
