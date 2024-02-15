package main

import (
	"log"
	"ms-booking/config"
	"ms-booking/controller"
	"ms-booking/pb"
	"ms-booking/repository"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectMongoDB()

	BookingRepository := repository.NewBookingRepository(db)
	BookingController := controller.NewBookingController(BookingRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, BookingController)

	listen, err := net.Listen("tcp", ":50003")
	if err != nil {
		log.Println(err)
	}

	log.Println("gRPC [ms-user] started on :50003")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
