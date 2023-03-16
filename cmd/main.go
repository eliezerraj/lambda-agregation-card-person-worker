package main

import(
	"os"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lambda-agregation-card-person-worker/internal/adapter/handler"
	"github.com/lambda-agregation-card-person-worker/internal/repository"
	"github.com/lambda-agregation-card-person-worker/internal/service"
	"github.com/lambda-agregation-card-person-worker/internal/erro"
)

var (
	logLevel		=	zerolog.DebugLevel // InfoLevel DebugLevel
	tableName		=	"agregation_card_person"
	version			=	"lambda-agregation-card-person-worker (github) version 1.5"
	eventTypePerson =  "personCreated"
	eventTypeCard 	= 	"cardCreated"
	agregationRepository	*repository.AgregationRepository
	agregationService		*service.AgregationService
	workerHandler			*handler.WorkerHandler
)

func getEnv(){
	if os.Getenv("TABLE_NAME") !=  "" {
		tableName = os.Getenv("TABLE_NAME")
	}
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.DebugLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init() {
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main() {
	log.Debug().Msg("main lambda-agregation-card-person-worker (go) v 1.5")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	agregationRepository, err := repository.NewAgregationRepository(tableName)
	if err != nil{
		return
	}
	agregationService 	= service.NewAgregationService(*agregationRepository)
	workerHandler 		= handler.NewWorkerHandler(*agregationService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, event events.CloudWatchEvent) ( error) {
	log.Debug().Msg("handler")
	log.Debug().Msg("**************************")
	log.Debug().Str("event.DetailType ", event.DetailType).
				Msg("----")
	log.Debug().Msg("-*******************")

	if (event.DetailType == eventTypeCard) {
		err := workerHandler.EventCard(event)
		if err != nil {
			return erro.ErrEventDetail
		}
	}else if (event.DetailType == eventTypePerson) {
		err := workerHandler.EventPerson(event)
		if err != nil {
			return erro.ErrEventDetail
		}
	}else {
		return erro.ErrEventDetail
	}

	return nil
}