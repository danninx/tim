testable := ./internal/{plate,conf,actions}

coverage:
	go test -coverprofile=test/coverage.out $(testable)
	go tool cover -html=test/coverage.out

PHONY: coverage
