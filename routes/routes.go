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

	// Customer Router
	r.GET("/customer", controllers.GetAllCustomer)
	r.POST("/customer", controllers.CreateCustomer)
	r.PATCH("/customer/:id", controllers.UpdateCustomer)
	r.DELETE("/customer/:id", controllers.DeleteCustomer)

	// Cars Router
	r.GET("/cars", controllers.GetAllCars)
	r.POST("/cars", controllers.CreateCars)
	r.PATCH("/cars/:id", controllers.UpdateCars)
	r.DELETE("/cars/:id", controllers.DeleteCars)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
