package handler

import(

	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lambda-agregation-card-person-worker/internal/erro"
	"github.com/lambda-agregation-card-person-worker/internal/service"
	"github.com/lambda-agregation-card-person-worker/internal/core/domain"

)

var childLogger = log.With().Str("handler", "WorkerHandler").Logger()

type WorkerHandler struct {
	agregationService service.AgregationService
}

func NewWorkerHandler(agregationService service.AgregationService) *WorkerHandler{
	childLogger.Debug().Msg("NewAgregationHandler")
	return &WorkerHandler{
		agregationService: agregationService,
	}
}

func (h *WorkerHandler) EventCard(event events.CloudWatchEvent) error {
	childLogger.Debug().Msg("EventCard")

    var card domain.Card
    if err := json.Unmarshal([]byte(event.Detail), &card); err != nil {
        return erro.ErrUnmarshal
    }

	_, err := h.agregationService.AddCard(card)
	if err != nil {
		return err
	}

	return nil
}

func (h *WorkerHandler) EventPerson(event events.CloudWatchEvent) error {
	childLogger.Debug().Msg("EventPerson")

    var person domain.Person
    if err := json.Unmarshal([]byte(event.Detail), &person); err != nil {
        return erro.ErrUnmarshal
    }

	//response, err := h.agregationService.AddAgregation(agregation)
	//if err != nil {
	//	return return erro.ErrUnmarshal
	//}

	return nil
}