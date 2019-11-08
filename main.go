package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
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

	srv := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	srv.Init()

	repo := repository.NewUserRepo(db)
	ucase := usecase.NewUsersUsecase(repo)
	handler := deliveryGRPC.UserHandlerGRPC{ucase}

	pb.RegisterUsersServiceHandler(srv.Server(), &handler)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
