package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dzoniops/accommodation-service/db"
	"github.com/dzoniops/accommodation-service/services"
	pb "github.com/dzoniops/common/pkg/accommodation"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccommodationServiceServer(s, &services.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
