package usecase

type SpotDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	EventID  int    `json:"event_id"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
}
