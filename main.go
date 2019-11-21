package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
	pb "github.com/muhammadhidayah/users-service/proto/users"
	deliveryGRPC "github.com/muhammadhidayah/users-service/users/delivery/grpc"
	"github.com/muhammadhidayah/users-service/users/repository"
	"github.com/muhammadhidayah/users-service/users/usecase"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf(fmt.Sprintf("Could not connect to DB: %v", err))
	}

	defer db.Close()

	db.AutoMigrate(&pb.User{})

	registry := kubernetes.NewRegistry()

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Registry(registry),
	)

	srv.Init()

	repo := repository.NewUserRepo(db)
	ucase := usecase.NewUsersUsecase(repo)
	handler := deliveryGRPC.UserHandlerGRPC{ucase}

	pb.RegisterUserServiceHandler(srv.Server(), &handler)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
