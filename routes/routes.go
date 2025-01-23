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
	r.GET("/bookingdiscount", controllers.GetAllBookingWithDiscount)

	// Cars Router
	r.GET("/cars", controllers.GetAllCars)
	r.POST("/cars", controllers.CreateCars)
	r.PATCH("/cars/:id", controllers.UpdateCars)
	r.DELETE("/cars/:id", controllers.DeleteCars)

	// Booking Router
	r.GET("/booking", controllers.GetAllBooking)
	r.GET("/booking/:id", controllers.GetBookingbyID)
	r.GET("/booking/discount", controllers.GetAllBookingWithDiscount)
	r.GET("/booking/:id/discount", controllers.GetBookingbyIDWithDiscount)
	r.POST("/booking", controllers.CreateBooking)
	r.PATCH("/booking/:id", controllers.UpdateBooking)
	r.DELETE("/booking/:id", controllers.DeleteBooking)

	// Membership Router
	r.GET("/membership", controllers.GetAllMembership)
	r.POST("/membership", controllers.CreateMembership)
	r.PATCH("/membership/:id", controllers.UpdateMembership)
	r.DELETE("/membership/:id", controllers.DeleteMembership)

	// Booking Type Router
	r.GET("/bookingtype", controllers.GetAllBookingType)
	r.POST("/bookingtype", controllers.CreateBookingType)
	r.PATCH("/bookingtype/:id", controllers.UpdateBookingType)
	r.DELETE("/bookingtype/:id", controllers.DeleteBookingType)

	// Driver Incentive Router
	r.GET("/driverincentive", controllers.GetAllDriverIncentive)
	r.POST("/driverincentive", controllers.CreateDriverIncentive)
	r.PATCH("/driverincentive/:id", controllers.UpdateDriverIncentive)
	r.DELETE("/driverincentive/:id", controllers.DeleteDriverIncentive)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
