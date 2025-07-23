testable := ./internal/{plate,conf,actions}

coverage:
	go test -coverprofile=coverage.out $(testable)
	go tool cover -html=coverage.out

PHONY: coverage
