.PHONY: run test coverage

run:
	go run main.go

test:
	go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils -v

coverage:
	go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils -coverprofile=coverage.out
	go tool cover -func=coverage.out | grep -E "total|service|utils"
