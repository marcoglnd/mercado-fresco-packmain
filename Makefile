server:
	go run ./...

test:
	go test -v ./... -covermode=atomic -coverpkg=./... -count=1  -race -timeout=30m -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

dockerup:
	docker compose up

dockerdown:
	docker compose down

.PRONY: server test coverage dockerup dockerdown