package main

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	pb "github.com/ruandao/micro-shippy-user-service/proto/user"
	"log"
	"os"
)

const (
	defaultDB = "database:5432"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
	)
	srv.Init()

	dbUri := os.Getenv("DB_HOST")
	if dbUri == "" {
		dbUri = defaultDB
	}

	db, err := gorm.Open(dbUri, "user")
	if err != nil {
		log.Fatalf("connect database err: %v", err)
	}
	defer db.Close()

	repository := &UserRepository{db}
	h := &handler{repository}
	pb.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("user service err: %v", err)
	}
}
