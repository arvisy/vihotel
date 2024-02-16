package main

import (
	"log"
	"ms-gateway/controller"
	"ms-gateway/pb"
	route "ms-gateway/router"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title vihotel
// @version 1.0
// @description An e-commerce API for booking hotel and with many features, made using microservices and gRPC.

// @contact.name Arviansyah Nur
// @contact.url http://www.swagger.io/support
// @contact.email arviansyahnur3@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	// user-service
	userConn, err := grpc.Dial(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	u := controller.NewUserController(userService)

	// hotel-service
	hotelConn, err := grpc.Dial(":50002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer hotelConn.Close()

	hotelService := pb.NewHotelServiceClient(hotelConn)
	h := controller.NewHotelController(hotelService)

	// booking-service
	bookingConn, err := grpc.Dial(":50003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer bookingConn.Close()

	bookingService := pb.NewBookingServiceClient(bookingConn)
	b := controller.NewBookingController(userService, hotelService, bookingService)

	// payment-service
	paymentConn, err := grpc.Dial(":50004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer paymentConn.Close()

	paymentService := pb.NewPaymentServiceClient(paymentConn)
	p := controller.NewPaymentController(paymentService, bookingService)

	e := echo.New()

	route.InitRoutes(e, u, h, b, p)

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
