package main

import (
	"log"
	"ms-user/config"
	"ms-user/controller"
	"ms-user/pb"
	"ms-user/repository"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectPostgresDB()

	UserRepository := repository.NewUserRepository(db)
	UserController := controller.NewUserController(*UserRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, UserController)

	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Println(err)
	}

	log.Println("gRPC [ms-user] started on :50001")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
