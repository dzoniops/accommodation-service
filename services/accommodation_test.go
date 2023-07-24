package services

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dzoniops/accommodation-service/db"
	"github.com/dzoniops/accommodation-service/models"
	"github.com/dzoniops/common/pkg/accommodation"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()
}
func teardown() {
	db.DB.Delete(&models.Accommodation{}, "name = ?", "Test Residence")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
func TestAddAccommodationSuccess(t *testing.T) {
	var req = accommodation.AccommodationRequest{
		HostId:           1,
		Name:             "Test Residence",
		Town:             "Zrenjanin",
		Municipality:     "Zrenjanin",
		Country:          "Serbia",
		Amenities:        "wifi;tv",
		MinGuests:        1,
		MaxGuests:        6,
		PricingModel:     accommodation.PricingModel_PRICING_MODEL_PUPN,
		ReservationModel: accommodation.ReservationModel_RESERVATION_MODEL_AUTO}
	info, err := (&Server{}).CreateAccommodation(context.TODO(), &req)

	require.NoError(t, err)
	require.NotEqual(t, info.AccommodationId, 0)
}
