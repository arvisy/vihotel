package main

import (
	"log"
	"ms-payment/config"
	"ms-payment/controller"
	"ms-payment/pb"
	"ms-payment/repository"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectMongoDB()

	PaymentRepository := repository.NewPaymentRepository(db)
	PaymentController := controller.NewPaymentController(PaymentRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, PaymentController)

	listen, err := net.Listen("tcp", ":50004")
	if err != nil {
		log.Println(err)
	}

	log.Println("gRPC [ms-user] started on :50004")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
