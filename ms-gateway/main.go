package main

import (
	"log"
	"ms-gateway/controller"
	"ms-gateway/pb"
	route "ms-gateway/router"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	e := echo.New()

	route.InitRoutes(e, u, h, b)

	// e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
