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

	e := echo.New()

	route.InitRoutes(e, u)

	// e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
