package services

import (
	"context"

	"github.com/dzoniops/accommodation-service/db"
	"github.com/dzoniops/accommodation-service/models"
	pb "github.com/dzoniops/common/pkg/accommodation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAccommodationServiceServer
}

func (s *Server) CreateAccommodation(
	c context.Context,
	req *pb.AccommodationRequest,
) (*pb.AccommodationResponse, error) {
	var accommodation = models.Accommodation{
		HostID:           req.HostId,
		Name:             req.Name,
		Town:             req.Town,
		Municipality:     req.Municipality,
		Country:          req.Country,
		Amenities:        req.Amenities,
		MinGuests:        int(req.MinGuests),
		MaxGuests:        int(req.MaxGuests),
		Images:           []models.AccommodationImage{},
		PricingModel:     models.PricingModel(req.PricingModel),
		ReservationModel: models.ReservationModel(req.ReservationModel),
	}

	for _, v := range req.Images {
		var image = models.AccommodationImage{B64IMG: v.B64Img}
		accommodation.Images = append(accommodation.Images, image)
	}

	db.DB.Create(&accommodation)
	return &pb.AccommodationResponse{AccommodationId: int64(accommodation.ID)}, status.New(codes.OK, "").Err()
}
