package main

import (
	"flag"
	"fmt"
	"github.com/dzoniops/accommodation-service/client"
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

	var reservationClient = client.InitReservationClient(
		fmt.Sprintf("localhost:%s", os.Getenv("RESERVATION_PORT")),
	)
	s := grpc.NewServer()
	pb.RegisterAccommodationServiceServer(s, &services.Server{ReservationClient: *reservationClient})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
