server:
	go run ./...

swag:
	swag init -g cmd/server/main.go

test:
	go test -v ./... -covermode=atomic -coverpkg=./... -count=1  -race -timeout=30m -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

dockerup:
	docker compose up

dockerdown:
	docker compose down

mockery:
	mockery --all --keeptree

.PRONY: server swag test coverage dockerup dockerdown mockery