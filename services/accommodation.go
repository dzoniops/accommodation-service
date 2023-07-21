package services

import (
	"context"
	"fmt"

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

func (s *Server) GetAccommodationById(
	c context.Context,
	req *pb.AccommodationResponse,
) (*pb.AccommodationInfo, error) {

	var id int64 = req.AccommodationId
	var accommo models.Accommodation
	if result := db.DB.Preload("Images").Where(models.Accommodation{ID: id}).First(&accommo); result.Error != nil {
		fmt.Println(result.Error)
	}
	var accommodation = pb.AccommodationInfo{
		Id:               accommo.ID,
		HostId:           accommo.HostID,
		Name:             accommo.Name,
		Town:             accommo.Town,
		Municipality:     accommo.Municipality,
		Country:          accommo.Country,
		Amenities:        accommo.Amenities,
		MinGuests:        int64(accommo.MinGuests),
		MaxGuests:        int64(accommo.MaxGuests),
		PricingModel:     string(accommo.PricingModel),
		ReservationModel: string(accommo.ReservationModel),
		Images:           []*pb.AccommodationImageResponse{},
	}

	for _, v := range accommo.Images {
		var image = pb.AccommodationImageResponse{
			B64Img:          v.B64IMG,
			ImageId:         v.ID,
			AccommodationId: v.AccommodationID,
		}
		accommodation.Images = append(accommodation.Images, &image)
	}

	return &accommodation, status.New(codes.OK, "").Err()
}

func (s *Server) AccommodationSearch(
	c context.Context,
	req *pb.AccommodationSearchRequest,
) (*pb.AccommodationSearchResponse, error) {
	var town string = req.Town
	var municipality string = req.Municipality
	var country string = req.Country
	var guestCount int = int(req.GuestCount)

	var accommo []models.Accommodation
	if result := db.DB.Preload("Images").Where("Min_Guests <= ? AND Max_Guests >= ? AND town LIKE ? AND municipality LIKE ? AND country LIKE ?", guestCount, guestCount, "%"+town+"%", "%"+municipality+"%", "%"+country+"%").Find(&accommo); result.Error != nil {
		fmt.Println(result.Error)
	}
	var searchResult = pb.AccommodationSearchResponse{AccomomodationList: []*pb.AccommodationInfo{}}

	for _, v := range accommo {
		var accommodation = pb.AccommodationInfo{
			Id:               v.ID,
			HostId:           v.HostID,
			Name:             v.Name,
			Town:             v.Town,
			Municipality:     v.Municipality,
			Country:          v.Country,
			Amenities:        v.Amenities,
			MinGuests:        int64(v.MinGuests),
			MaxGuests:        int64(v.MaxGuests),
			PricingModel:     string(v.PricingModel),
			ReservationModel: string(v.ReservationModel),
			Images:           []*pb.AccommodationImageResponse{},
		}
		for _, v2 := range v.Images {
			var image = pb.AccommodationImageResponse{
				B64Img:          v2.B64IMG,
				ImageId:         v2.ID,
				AccommodationId: v2.AccommodationID,
			}
			accommodation.Images = append(accommodation.Images, &image)
		}
		searchResult.AccomomodationList = append(searchResult.AccomomodationList, &accommodation)
	}

	return &searchResult, status.New(codes.OK, "").Err()
}
