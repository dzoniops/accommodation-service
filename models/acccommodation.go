package models

import (
	"gorm.io/gorm"
)

type Accommodation struct {
	gorm.Model
	ID               int64                `json:"id"                gorm:"primaryKey"`
	HostID           int64                `json:"host_id"`
	Name             string               `json:"name"`
	Town             string               `json:"town"`
	Municipality     string               `json:"municipality"`
	Country          string               `json:"country"`
	Amenities        string               `json:"amenities"`
	MinGuests        int                  `json:"min_guests"`
	MaxGuests        int                  `json:"max_guests"`
	PricingModel     PricingModel         `json:"pricing_model"`
	ReservationModel ReservationModel     `json:"reservation_model"`
	Images           []AccommodationImage `json:"images"`
}

type PricingModel int32

const (
	PRICING_MODEL_UNSPECIFIED PricingModel = 0
	PUPN                      PricingModel = 1
	PGPN                      PricingModel = 2
)

type ReservationModel int32

const (
	RESERVATION_MODEL_UNSPECIFIED ReservationModel = 0
	AUTO                          ReservationModel = 1
	MANUAL                        ReservationModel = 2
)

type AccommodationImage struct {
	ID              int64  `json:"id"`
	B64IMG          string `json:"b64_img"       gorm:"not null"`
	AccommodationID int64  `json:"accommodation" gorm:"not null"`
}
