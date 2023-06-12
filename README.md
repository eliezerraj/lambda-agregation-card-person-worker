# lambda-agregation-card-person-worker

POC Lambda for technical purposes

Lambda read notification from EventBridge and link CARD data and PERSON data together

Diagrama Flow

    EventBridge (person) ==> Lambda ==> DynamoDB (agregation_card_person)
    EventBridge (card)  ==> Lambda ==> DynamoDB (agregation_card_person)


## Compile

    GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

    zip -jrm ../build/main.zip ../build/main

    aws lambda update-function-code \
    --function-name lambda-agregation-card-person-worker \
    --zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-agregation-card-person-worker/build/main.zip \
    --publish

## Endpoint

Worker (no endpoint)
