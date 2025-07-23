coverage:
	go test -coverprofile=coverage.out ./internal/* .
	go tool cover -html=coverage.out

PHONY: coverage
