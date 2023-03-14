package service

import (
	"time"
	"strings"

	"github.com/lambda-agregation-card-person-worker/internal/core/domain"

)

func (s *AgregationService) AddCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("AddCard")
	
	var agregation_status = ""
	var time_create_at = time.Now()
	var time_update_at = time.Time{}
	var sk = "PERSON:"

	// Add new CARD
	_, err := s.agregationRepository.AddCard(card)
	if err != nil {
		return nil, err
	}

	//Get Person
	sk_str := strings.Replace(card.SK, "PERSON:", "", -1)
	person, err := s.agregationRepository.GetPerson(sk_str)
	if err != nil {
		agregation_status = "WAITING-PERSON-DATA - " + sk_str
	} else {
		agregation_status = "ACTIVE"
		time_update_at = time.Now()
		sk = "PERSON:" + person.ID
	}

	// Add new card
	id := "AGREGATION-" + card.CardNumber
	agregation := domain.NewAgregationCardPerson(id,
												 sk,
												 card.CardNumber,
												 "",
												 agregation_status,
												 &time_create_at,
												 &time_update_at,
												 "TENANT-001")

	_, err = s.agregationRepository.AddAgregation(*agregation)
	if err != nil {
		return nil, err
	}

	if (agregation_status == "ACTIVE"){
		agregation_delete := &domain.AgregationCardPerson{ ID: id, SK:"PERSON:" }
		err = s.agregationRepository.DeleteAgregation(*agregation_delete)
		if err != nil {
			return nil, err
		}
	}

	return &card, nil
}

func (s *AgregationService) AddPerson(person domain.Person) (*domain.Person, error){
	childLogger.Debug().Msg("AddPerson")

	// Add new PERSON
	_, err := s.agregationRepository.AddPerson(person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
