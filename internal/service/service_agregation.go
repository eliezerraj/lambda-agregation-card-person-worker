package service

import (
	"time"

	"github.com/lambda-agregation-card-person-worker/internal/core/domain"

)

func (s *AgregationService) AddCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("AddCard")

	// Add new card
	id := "AGREGATION-" + card.CardNumber
	sk := "AGREGATION-" + card.CardNumber

	agregation := domain.NewAgregationCardPerson(id,
												 sk,
												 card.CardNumber,
												 "",
												 "ACTIVE",
												time.Now(),
												 "TENANT-001")

	_, err := s.agregationRepository.AddAgregation(*agregation)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (s *AgregationService) GetAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("GetAgregation")

	c, err := s.agregationRepository.GetAgregation(agregation)
	if err != nil {
		return nil, err
	}
	return c, nil
}
