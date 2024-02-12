test:
	go test -p=1 -count=1 ./...

run:
	go run ./...

lint:
	golangci-lint run
