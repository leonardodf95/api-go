package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID int) (*Event, error)
	FindSpotsByEventID(eventID int) ([]Spot, error)
	ReserveSpot(EventID int, spotName string) (*Spot, error)
}
