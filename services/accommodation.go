package services

import (
	"context"
	"github.com/dzoniops/accommodation-service/client"
	"github.com/dzoniops/accommodation-service/db"
	"github.com/dzoniops/accommodation-service/models"
	pb "github.com/dzoniops/common/pkg/accommodation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedAccommodationServiceServer
	ReservationClient client.ReservationClient
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
	return &pb.AccommodationResponse{AccommodationId: accommodation.ID}, status.New(codes.OK, "").Err()
}

func (s *Server) GetAccommodationById(
	c context.Context,
	req *pb.AccommodationResponse,
) (*pb.AccommodationInfo, error) {

	var accommodation models.Accommodation
	if res := db.DB.Preload("Images").Where(models.Accommodation{ID: req.AccommodationId}).First(&accommodation); res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}
	var accommodationInfo = pb.AccommodationInfo{
		Id:               accommodation.ID,
		HostId:           accommodation.HostID,
		Name:             accommodation.Name,
		Town:             accommodation.Town,
		Municipality:     accommodation.Municipality,
		Country:          accommodation.Country,
		Amenities:        accommodation.Amenities,
		MinGuests:        int64(accommodation.MinGuests),
		MaxGuests:        int64(accommodation.MaxGuests),
		PricingModel:     pb.PricingModel(accommodation.PricingModel),
		ReservationModel: pb.ReservationModel(accommodation.ReservationModel),
		Images:           []*pb.AccommodationImageResponse{},
	}

	for _, v := range accommodation.Images {
		var image = pb.AccommodationImageResponse{
			B64Img:          v.B64IMG,
			ImageId:         v.ID,
			AccommodationId: v.AccommodationID,
		}
		accommodationInfo.Images = append(accommodationInfo.Images, &image)
	}

	return &accommodationInfo, status.New(codes.OK, "").Err()
}

func (s *Server) AccommodationSearch(c context.Context, req *pb.AccommodationSearchRequest) (*pb.AccommodationSearchResponse, error) {
	var town = req.Town
	var municipality = req.Municipality
	var country = req.Country
	var guestCount = int(req.GuestCount)

	var accommodations []models.Accommodation
	if result := db.DB.Preload("Images").
		Where("Min_Guests <= ? AND Max_Guests >= ? AND (town LIKE ? OR municipality LIKE ? OR country LIKE ?)",
			guestCount, guestCount, "%"+town+"%", "%"+municipality+"%", "%"+country+"%").Find(&accommodations); result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	//TODO: calculate price for each search accommodation

	searchResult, err := s.ReservationClient.FilterAccommodations(c, req.StartDate, req.EndDate, accommodations)
	if err != nil {
		return nil, err
	}
	return searchResult, status.New(codes.OK, "").Err()
}

func (s *Server) DeleteByHost(c context.Context, req *pb.IdRequest) (*emptypb.Empty, error) {
	if res := db.DB.Where("host_id = ?", req.Id).Delete(&models.Accommodation{}); res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}
	return &emptypb.Empty{}, nil
}
