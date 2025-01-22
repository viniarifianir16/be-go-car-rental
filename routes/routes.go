package routes

import (
	"be-go-car-rental/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Customers Router
	r.GET("/customers", controller.GetAllCustomers)
	r.POST("/customers", controller.CreateCustomers)
	r.PATCH("/customers/:id", controller.UpdateCustomers)
	r.DELETE("/customers/:id", controller.DeleteCustomers)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
