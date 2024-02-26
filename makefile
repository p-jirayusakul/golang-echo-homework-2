include .env
export

test:
	go test -v -cover -short ./...

server:
	go run main.go

swagger:
	swag init

mock:
	mockgen -package mockup -destination test/mockup/store.go github.com/p-jirayusakul/golang-echo-homework-2/database Store

.PHONY: test server swagger mock