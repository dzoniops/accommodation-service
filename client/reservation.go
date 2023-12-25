package client

import (
	"fmt"
	"github.com/dzoniops/accommodation-service/models"
	pb "github.com/dzoniops/common/pkg/reservation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type ReservationClient struct {
	client pb.ReservationServiceClient
}

func InitReservationClient(url string) *ReservationClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	client := pb.NewReservationServiceClient(conn)
	return &ReservationClient{client: client}
}

func (r *ReservationClient) filterAccommodations(startDate, endDate time.Time, accommodations []models.Accommodation) {

}
