.PHONY: cover start test test-integration

cover:
	go tool cover -html=cover.out

start:
	go run cmd/server/*.go

test:
	go test -v -coverprofile=cover.out -short -count=1 ./...

test-integration:
	go test -v -coverprofile=cover.out -p 1 ./...

