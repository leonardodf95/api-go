package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/leonardodf95/api-go/internal/events/domain"
)

type EventRepository struct {
	events []domain.Event
	spots  []domain.Spot
}

func NewEventRepository(jsonData []byte) (*EventRepository, error) {
	var data struct {
		Events []domain.Event `json:"events"`
		Spots  []domain.Spot  `json:"spots"`
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Errei aqui")
		return nil, err
	}

	const timeFormat = "2006-01-02T15:04:05"

	for i, event := range data.Events {
		data.Events[i].Date, err = time.Parse(timeFormat, event.DateStr)
		if err != nil {
			return nil, errors.New("failed to parse event date: " + event.DateStr)
		}
		data.Events[i].CreatedAt, err = time.Parse(timeFormat, event.CreatedAtStr)
		if err != nil {
			return nil, errors.New("failed to parse event created_at: " + event.CreatedAtStr)
		}
	}

	repo := &EventRepository{
		events: data.Events,
		spots:  data.Spots,
	}
	return repo, nil
}

func (r *EventRepository) ListEvents() ([]domain.Event, error) {
	return r.events, nil
}

func (r *EventRepository) FindEventByID(eventID int) (*domain.Event, error) {
	for _, event := range r.events {
		if event.ID == eventID {
			return &event, nil
		}
	}
	return nil, errors.New("event not found")
}

func (r *EventRepository) FindSpotsByEventID(eventID int) ([]domain.Spot, error) {
	var spots []domain.Spot
	for _, spot := range r.spots {
		if spot.EventID == eventID {
			spots = append(spots, spot)
		}
	}
	return spots, nil
}

func (r *EventRepository) ReserveSpot(eventID int, name string) (*domain.Spot, error) {

	for i, spot := range r.spots {
		if spot.Name == name && spot.EventID == eventID {
			if spot.Status == domain.SpotStatusAvailable {
				r.spots[i].Status = domain.SpotStatusReserved
				return &r.spots[i], nil
			} else {
				return nil, domain.ErrSpotAlreadyReserved
			}
		}
	}
	return nil, domain.ErrSpotNotFound
}
