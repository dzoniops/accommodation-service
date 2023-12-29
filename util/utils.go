package util

import (
	"github.com/dzoniops/accommodation-service/models"
	pb "github.com/dzoniops/common/pkg/accommodation"
	reservationpb "github.com/dzoniops/common/pkg/reservation"
)

func Contains(id int64, available *reservationpb.FilterAvailableResponse) int {
	for idx, a := range available.IdPrices {
		if a.Id == id {
			return idx
		}
	}
	return -1
}

func CreateAccommodationSearchInfo(v models.Accommodation, idPrices *reservationpb.IdPrice, numberOfDays int64) *pb.AccommodationSearchInfo {
	var accommodationInfo = pb.AccommodationSearchInfo{
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
		TotalPrice:       idPrices.Price * numberOfDays,
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
func GenerateSearch(accommodations []models.Accommodation, response *reservationpb.FilterAvailableResponse, numberOfDays int64) *pb.AccommodationSearchResponse {
	var result []*pb.AccommodationSearchInfo
	for _, accommodation := range accommodations {
		if i := Contains(accommodation.ID, response); i != -1 {
			accommodationInfo := CreateAccommodationSearchInfo(accommodation, response.IdPrices[i], numberOfDays)
			result = append(result, accommodationInfo)
		}
	}
	return &pb.AccommodationSearchResponse{AccomomodationList: result}
}
