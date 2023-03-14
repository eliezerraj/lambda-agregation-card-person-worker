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
						
)

func TestAddCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddCard Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	card01 := domain.NewCard("CARD-4444.000.000.001",
							"PERSON:PERSON-001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"12/28",
							"TENANT-001")

	result, err := service.AddCard(*card01)
	if err != nil {
		t.Errorf("Error -TestAddCard Access DynanoDB %v err:%v ", tableName, err)
	}

	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddCard input : %v" , *card01)
	}
}

func TestAddPerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPerson Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	person := domain.NewPerson("PERSON-001","PERSON-001","Mr Cookie","F")
	result, err := service.AddPerson(*person)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access DynanoDB %v err: %v", tableName, err)
	}

	if (cmp.Equal(person, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPerson input : %v" , *person)
	}
}
