# Lambda-agregation-card-person-worker

POC Lambda for technical purposes

Lambda read notification from EventBridge and link CARD data and PERSON data together

Diagrama Flow

    Source(Lambda-Person) => EventBridge (person) <= Lambda => DynamoDB (agregation_card_person)
    Source(Lambda-Card) => EventBridge (card)  <= Lambda => DynamoDB (agregation_card_person)


## Compile

   Manually compile the function

    GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

    zip -jrm ../build/main.zip ../build/main

    aws lambda update-function-code \
    --function-name lambda-agregation-card-person-worker \
    --zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-agregation-card-person-worker/build/main.zip \
    --publish

## Endpoint

+ Worker (no endpoint)

## Event

+ eventTypePerson =  "personCreated"
    
    Call EventPersonAggregation

+ eventTypeCard 	= 	"cardCreated"

    Call EventCardAggregation

## Pipeline

Prerequisite: 

Lambda function already created

+ buildspec.yml: build the main.go and move to S3
+ buildspec-update.yml: update the lambda-function using S3 build and prepare to canary deploy
+ appspec: execute the canary deploy

## DynamoDB

    PERSON-100 PERSON-100 M Eliezer Antunes

    CARD-8888.000.100.001 PERSON:PERSON-100 Eliezer R A unior 8888.000.100.001

    AGREGATION-8888.000.100.001 PERSON:PERSON-100 8888.000.100.001 2023-07-12T18:54:40.758065589Z ACTIVE TENANT-001 2023-07-12T18:54:41.362029704Z

