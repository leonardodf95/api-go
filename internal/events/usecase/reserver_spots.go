package usecase

import (
	"github.com/leonardodf95/api-go/internal/events/domain"
)

type ReserveSpotInputDTO struct {
	EventID int      `json:"event_id"`
	Spots   []string `json:"spots"`
}

type ReserveSpotOutputDTO struct {
	EventID int           `json:"event_id"`
	Spots   []domain.Spot `json:"spots"`
}

type ReserveSpotUseCase struct {
	repo domain.EventRepository
}

func NewReserveSpotUseCase(repo domain.EventRepository) *ReserveSpotUseCase {
	return &ReserveSpotUseCase{
		repo: repo,
	}
}

func (uc *ReserveSpotUseCase) Execute(input ReserveSpotInputDTO) (*ReserveSpotOutputDTO, error) {
	// Verifica o evento
	_, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	var spots []domain.Spot

	for _, spot := range input.Spots {
		spotReseved, err := uc.repo.ReserveSpot(input.EventID, spot)
		if err != nil {
			return nil, err
		}
		spots = append(spots, *spotReseved)
	}

	return &ReserveSpotOutputDTO{EventID: input.EventID, Spots: spots}, nil
}
