package models

import (
	"gorm.io/gorm"
)

type Accommodation struct {
	gorm.Model
	ID               int64                `json:"id"`
	HostID           int64                `json:"hostId"`
	Name             string               `json:"name"`
	Town             string               `json:"town"`
	Municipality     string               `json:"municipality"`
	Country          string               `json:"country"`
	Amenities        string               `json:"amenities"`
	MinGuests        int                  `json:"minGuests"`
	MaxGuests        int                  `json:"maxGuests"`
	PricingModel     PricingModel         `json:"pricingModel"`
	ReservationModel ReservationModel     `json:"reservationModel"`
	Images           []AccommodationImage `json:"images"`
}

type PricingModel string

const (
	PUPN PricingModel = "pupn"
	PGPN PricingModel = "pgpn"
)

type ReservationModel string

const (
	AUTO   ReservationModel = "auto"
	MANUAL ReservationModel = "manual"
)

type AccommodationImage struct {
	ID              int64  `json:"id"`
	B64IMG          string `json:"b64img" gorm:"not null"`
	AccommodationID int64  `json:"accommodation" gorm:"not null"`
}
