# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Build binary for API
build: fmt vet
	go build -o api ./cmd/api/main.go

# Run API
run: fmt vet
	go run ./cmd/api/main.go