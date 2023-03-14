package repository

import(

	"github.com/lambda-agregation-card-person-worker/internal/core/domain"
	"github.com/lambda-agregation-card-person-worker/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

)

func (r *AgregationRepository) AddPerson(person domain.Person) (*domain.Person, error){
	childLogger.Debug().Msg("AddPerson")

	item, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: r.tableName,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrInsert
	}

	_, err = r.client.TransactWriteItems(transaction)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrInsert
	}

	return &person , nil
}

func (r *AgregationRepository) GetPerson(id string) (*domain.Person, error){
	childLogger.Debug().Msg("GetPerson")

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(id)),
		expression.Key("sk").BeginsWith(id),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.QueryInput{
								TableName:                 r.tableName,
								ExpressionAttributeNames:  expr.Names(),
								ExpressionAttributeValues: expr.Values(),
								KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.Query(key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrQuery
	}

	person := []domain.Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &person)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(person) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &person[0], nil
	}
}
