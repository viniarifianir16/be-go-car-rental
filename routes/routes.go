package routes

import (
	"be-go-car-rental/controllers"

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
	r.GET("/customers", controllers.GetAllCustomers)
	r.POST("/customers", controllers.CreateCustomers)
	r.PATCH("/customers/:id", controllers.UpdateCustomers)
	r.DELETE("/customers/:id", controllers.DeleteCustomers)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
