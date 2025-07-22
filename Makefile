# Go parameters
BINARY_NAME=bff-server
CMD_PATH=./cmd/server

.PHONY: all run build test tidy lint clean

all: build

run:
	@echo "Running the application..."
	go run ${CMD_PATH}

build:
	@echo "Building the binary..."
	go build -o ${BINARY_NAME} ${CMD_PATH}

test:
	@echo "Running tests..."
	go test ./... -v

tidy:
	@echo "Tidying go modules..."
	go mod tidy

lint:
	@echo "Running linter..."
	go vet ./...

clean:
	@echo "Cleaning up..."
	rm -f ${BINARY_NAME}

swaginit:
	@echo "Swagger install dependencies..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/files
	go get github.com/swaggo/gin-swagger
	go get github.com/swaggo/swag
	go mod tidy

swagdoc:
	@echo "Swagger init doc..."
	swag init -g ./cmd/server/main.go

testinit:
	@echo "Testify install dependencies..."
	go get github.com/stretchr/testify

test:
	@echo "Testify running..."
	go test ./test/... -v