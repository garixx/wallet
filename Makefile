#BINARY=engine
#engine:
#	go build -o ${BINARY} cmd/*.go
test:
	go test ./...

container:
	docker build -t workshop-app .

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

extra-coverage:
	go test -v -short -count=1 -race -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: clean install test build docker run stop vendor lint-prepare lint