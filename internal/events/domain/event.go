package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidEvent    = errors.New("invalid event data")
	ErrEventFull       = errors.New("event is full")
	ErrTicketNotFound  = errors.New("ticket not found")
	ErrTicketNotEnough = errors.New("not enough tickets available")
	ErrEventNotFound   = errors.New("event not found")
)

// Rating represents the age restriction for an event.
type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

// Event represents an event with tickets and spots.
type Event struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Organization string    `json:"organization"`
	Date         time.Time `json:"-"`
	Price        float64   `json:"price"`
	Rating       Rating    `json:"rating"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"-"`
	Location     string    `json:"location"`
	DateStr      string    `json:"date"`
	CreatedAtStr string    `json:"created_at"`
}
