package service

import(
	"testing"

	"github.com/rs/zerolog"
	"github.com/google/go-cmp/cmp"

	"github.com/lambda-agregation-card-person-worker/internal/core/domain"
	"github.com/lambda-agregation-card-person-worker/internal/repository"
)

var (
	tableName = "agregation_card_person"
	agregationRepository	*repository.AgregationRepository

	card01 = domain.NewCard("CARD-4444.000.000.999",
							"CARD-4444.000.000.999",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"02/26",
							"TENANT-001")
							
)

func TestAddCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddCard Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	result, err := service.AddCard(*card01)
	if err != nil {
		t.Errorf("Error -TestAddCard Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddCard input : %v" , *card01)
	}
}
