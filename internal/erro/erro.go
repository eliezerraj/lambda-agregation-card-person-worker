package erro

import (
	"errors"

)

var (
	ErrListNotAllowed 	= errors.New("Lista (SCAN) não permitida para o DynamoDB")
	ErrList 			= errors.New("Erro na leitura (LIST)")
	ErrSaveDatabase 	= errors.New("Erro no UPSERT")
	ErrCreateSession	= errors.New("Erro na Criaçao da Sessao AWS")
	ErrOpenDatabase 	= errors.New("Erro na abertura do DB")
	ErrNotFound 		= errors.New("Item não encontrado")
	ErrFunctionNotImpl 	= errors.New("Função não implementada")
	ErrInsert 			= errors.New("Erro na inserção do dado")
	ErrQuery 			= errors.New("Erro na query")
	ErrDelete 			= errors.New("Erro no Delete")
	ErrPutEvent			= errors.New("Erro na notificação PUTEVENT")
	ErrUnmarshal 		= errors.New("Erro na conversão do JSON")
	ErrUnauthorized 	= errors.New("Erro de autorização")
	ErrConvertion 		= errors.New("Erro de conversão de String para Inteiro")
	ErrMethodNotAllowed = errors.New("Metodo não permitido")
	ErrPreparedQuery 	= errors.New("Erro na preparação da Query para o Dynamo")
	ErrQueryEmpty	 	= errors.New("Query string não pode ser vazia")
	ErrEventDetail	 	= errors.New("Evento não suportado")
)