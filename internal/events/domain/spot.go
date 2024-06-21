package domain

import (
	"errors"
)

var (
	ErrInvalidSpotNumber   = errors.New("invalid spot number")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

type Spot struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Status  SpotStatus `json:"status"`
	EventID int        `json:"event_id"`
}
