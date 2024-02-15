package route

import (
	"ms-gateway/controller"
	"ms-gateway/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(
	e *echo.Echo,
	user *controller.UserContoller,
	hotel *controller.HotelController,
	booking *controller.BookingController,
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

		customer.POST("/booking", booking.CreateBooking)
		customer.GET("/booking/:booking_id", booking.GetBooking)
		customer.PUT("/booking/:booking_id", booking.UpdateBooking)
		customer.GET("/booking", booking.ListBookings)
		customer.DELETE("/booking/:booking_id", booking.DeleteBooking)
	}

	admin := e.Group("/api/v1")
	admin.Use(middleware.Authentication, middleware.AdminAuth)
	{
		admin.GET("/user/admin/:id", user.GetCustomerAdmin)
		admin.GET("/user/admin", user.GetAllCustomerAdmin)
		admin.PUT("/user/admin/:id", user.UpdateCustomerAdmin)
		admin.DELETE("/user/admin/:id", user.DeleteCustomerAdmin)

		admin.POST("/hotel", hotel.CreateHotel)
		admin.PUT("/hotel", hotel.UpdateHotel)
		admin.DELETE("/hotel/:id", hotel.DeleteHotel)
		admin.GET("/hotel/:id", hotel.GetHotel)
		admin.GET("/hotel", hotel.ListHotels)

		admin.POST("/hotel/room", hotel.CreateRoom)
		admin.PUT("/hotel/room", hotel.UpdateRoom)
		admin.DELETE("/hotel/room/:id", hotel.DeleteRoom)
		admin.GET("/hotel/room/:id", hotel.GetRoom)
		admin.GET("/hotel/room", hotel.ListRooms)
	}
}
