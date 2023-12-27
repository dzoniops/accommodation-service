package client

import (
	"context"
	"fmt"
	"github.com/dzoniops/accommodation-service/models"
	"github.com/dzoniops/accommodation-service/util"
	"github.com/dzoniops/common/pkg/accommodation"
	pb "github.com/dzoniops/common/pkg/reservation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (r *ReservationClient) FilterAccommodations(c context.Context, startDate, endDate *timestamppb.Timestamp, accommodations []models.Accommodation) (*accommodation.AccommodationSearchResponse, error) {
	var filterAccommodations pb.FilterAccommodationsRequest
	for _, a := range accommodations {
		accommodationRequest := pb.AccommodationRequest{
			AccommodationId: a.ID,
			StartDate:       startDate,
			EndDate:         endDate,
		}
		filterAccommodations.Accommodations = append(filterAccommodations.Accommodations, &accommodationRequest)
	}
	available, err := r.client.FilterAvailableForAccommodations(c, &filterAccommodations)
	if err != nil {
		return nil, err
	}
	result := util.GenerateSearch(accommodations, available)
	return result, nil
}
