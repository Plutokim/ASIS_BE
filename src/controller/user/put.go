package user

import (
	"github.com/Plutokim/sport/src/repos/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context, db *gorm.DB) {
	id := c.Param("user-id")
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	repo := user.New(db)
	err := repo.Update(id, body.FullName, body.PhoneNumber, body.Email, body.Password)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
