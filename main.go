package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/muhammadhidayah/users-service/proto/users"
	deliveryGRPC "github.com/muhammadhidayah/users-service/users/delivery/grpc"
	"github.com/muhammadhidayah/users-service/users/repository"
	"github.com/muhammadhidayah/users-service/users/usecase"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf(fmt.Sprintf("Could not connect to DB: %v", err))
	}

	defer db.Close()

	db.AutoMigrate(&pb.User{})

	srv := k8s.NewService(
		micro.Name("inact.srv.user"),
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
