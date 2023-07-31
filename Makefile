lint:
	golangci-lint run

go-mod:
	go mod download && go mod tidy

pre-push: go-mod lint