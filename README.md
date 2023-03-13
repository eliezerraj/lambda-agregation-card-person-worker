GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

zip -jrm ../build/main.zip ../build/main

aws lambda update-function-code \
--function-name lambda-agregation-card-person-worker \
--zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-agregation-card-person-worker/build/main.zip \
--publish