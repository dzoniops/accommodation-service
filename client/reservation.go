package client

import (
	"fmt"
	pb "github.com/dzoniops/common/pkg/reservation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReservationClient struct {
	client pb.ReservationServiceClient
}

func InitClient(url string) *ReservationClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	client := pb.NewReservationServiceClient(conn)
	return &ReservationClient{client: client}
}
