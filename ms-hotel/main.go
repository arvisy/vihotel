package main

import (
	"log"
	"ms-hotel/config"
	"ms-hotel/controller"
	"ms-hotel/pb"
	"ms-hotel/repository"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectPostgresDB()

	HotelRepository := repository.NewHotelRepository(db)
	HotelController := controller.NewHotelController(*HotelRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterHotelServiceServer(grpcServer, HotelController)

	listen, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Println(err)
	}

	log.Println("gRPC [ms-user] started on :50002")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
