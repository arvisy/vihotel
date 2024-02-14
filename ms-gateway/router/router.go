package route

import (
	"ms-gateway/controller"
	"ms-gateway/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(
	e *echo.Echo,
	user *controller.UserContoller,
) {
	// public endpoint
	e.POST("/register", user.Register)
	e.POST("/login", user.Login)

	// private endpoint
	customer := e.Group("/api/v1")
	customer.Use(middleware.Authentication, middleware.CustomerAuth)
	{
		customer.GET("/user", user.GetInfoCustomer)
		customer.PUT("/user", user.UpdateCustomer)
		customer.DELETE("/user", user.DeleteCustomer)
		customer.POST("/user/address", user.AddAddress)
		customer.PUT("/user/address", user.UpdateAddress)
	}

	admin := e.Group("/api/v1")
	admin.Use(middleware.Authentication, middleware.AdminAuth)
	{
		admin.GET("/user/admin/:id", user.GetCustomerAdmin)
		admin.GET("/user/admin", user.GetAllCustomerAdmin)
		admin.PUT("/user/admin/:id", user.UpdateCustomerAdmin)
		admin.DELETE("/user/admin/:id", user.DeleteCustomerAdmin)
	}
}
