package util

import (
	"github.com/dzoniops/accommodation-service/models"
	pb "github.com/dzoniops/common/pkg/accommodation"
	reservationpb "github.com/dzoniops/common/pkg/reservation"
)

func Contains(id int64, available *reservationpb.FilterAvailableResponse) bool {
	for _, a := range available.IdPrices {
		if a.Id == id {
			return true
		}
	}
	return false
}

func CreateAccommodationInfo(v models.Accommodation) *pb.AccommodationInfo {
	var accommodationInfo = pb.AccommodationInfo{
		Id:               v.ID,
		HostId:           v.HostID,
		Name:             v.Name,
		Town:             v.Town,
		Municipality:     v.Municipality,
		Country:          v.Country,
		Amenities:        v.Amenities,
		MinGuests:        int64(v.MinGuests),
		MaxGuests:        int64(v.MaxGuests),
		PricingModel:     pb.PricingModel(v.PricingModel),
		ReservationModel: pb.ReservationModel(v.ReservationModel),
		Images:           []*pb.AccommodationImageResponse{},
	}
	for _, v2 := range v.Images {
		var image = pb.AccommodationImageResponse{
			B64Img:          v2.B64IMG,
			ImageId:         v2.ID,
			AccommodationId: v2.AccommodationID,
		}
		accommodationInfo.Images = append(accommodationInfo.Images, &image)
	}
	return &accommodationInfo
}
func GenerateSearch(accommodations []models.Accommodation, response *reservationpb.FilterAvailableResponse) *pb.AccommodationSearchResponse {
	var result []*pb.AccommodationInfo
	for _, accommodation := range accommodations {
		if Contains(accommodation.ID, response) {
			accommodationInfo := CreateAccommodationInfo(accommodation)
			result = append(result, accommodationInfo)
		}
	}
	return &pb.AccommodationSearchResponse{AccomomodationList: result}
}
