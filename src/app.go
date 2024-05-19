package main

import (
	"embed"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/Plutokim/sport/src/controller/health"
	"github.com/Plutokim/sport/src/controller/payment"
	"github.com/Plutokim/sport/src/controller/product"
	"github.com/Plutokim/sport/src/controller/user"
	"github.com/Plutokim/sport/src/db"
	"github.com/Plutokim/sport/src/middleware"
	"github.com/Plutokim/sport/src/utils/env"
)

var PORT = env.GetDefault("PORT", "8080")
var PREFIX = env.GetDefault("PREFIX", "/api")
var DBSTRING = env.Get("DBSTRING")

//go:embed migrations/*.sql
var MIGRATION_FS embed.FS

func main() {
	err := applyMigrations(DBSTRING, MIGRATION_FS)
	if err != nil {
		slog.Error("Error applying migrations: %v", err)
	}
	conn := db.New(DBSTRING)

	server := gin.New()
	server.Use(middleware.CORSMiddleware())

	rootGroup := server.Group(fmt.Sprintf("%v", PREFIX))
	rootGroup.GET("/health", health.Get)

	userGroup := rootGroup.Group("/user")
	userGroup.POST("", func(c *gin.Context) { user.Create(c, conn) })
	userGroup.GET("/:user-id", func(c *gin.Context) { user.Read(c, conn) })
	userGroup.PUT("/:user-id", func(c *gin.Context) { user.Update(c, conn) })
	userGroup.DELETE("/:user-id", func(c *gin.Context) { user.Delete(c, conn) })
	userGroup.GET("", func(c *gin.Context) { user.List(c, conn) })

	rootGroup.GET("/payment", func(c *gin.Context) { payment.List(c, conn) })
	paymentGroup := userGroup.Group("/:user-id/payment")
	paymentGroup.POST("", func(c *gin.Context) { payment.Create(c, conn) })
	paymentGroup.GET("", func(c *gin.Context) { payment.Read(c, conn) })
	paymentGroup.PUT("", func(c *gin.Context) { payment.Update(c, conn) })
	paymentGroup.DELETE("", func(c *gin.Context) { payment.Delete(c, conn) })

	productGroup := rootGroup.Group("/product")
	productGroup.POST("", func(c *gin.Context) { product.Create(c, conn) })
	productGroup.GET("/:product-id", func(c *gin.Context) { product.Read(c, conn) })
	productGroup.PUT("/:product-id", func(c *gin.Context) { product.Update(c, conn) })
	productGroup.DELETE("/:product-id", func(c *gin.Context) { product.Delete(c, conn) })
	productGroup.GET("", func(c *gin.Context) { product.List(c, conn) })

	server.Run(fmt.Sprintf(":%v", PORT))
}
