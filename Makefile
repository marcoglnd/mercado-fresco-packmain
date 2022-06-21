coverage:
	go test -v ./... -covermode=atomic -coverpkg=./... -count=1  -race -timeout=30m -coverprofile=coverage.out && go tool cover -html=coverage.out

.PRONY: coverage