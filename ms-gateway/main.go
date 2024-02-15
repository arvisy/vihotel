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
	userConn, err := grpc.Dial(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	u := controller.NewUserController(userService)

	hotelConn, err := grpc.Dial(":50002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	hotelService := pb.NewHotelServiceClient(hotelConn)
	h := controller.NewHotelController(hotelService)

	e := echo.New()

	route.InitRoutes(e, u, h)

	// e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
