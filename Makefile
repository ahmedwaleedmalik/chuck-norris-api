# Image URL to use all building/pushing image targets
VERSION ?= 0.0.1
IMAGE_TAG_BASE ?= ahmedwaleedmalik/chuck-norris-api
IMG ?= $(IMAGE_TAG_BASE):v$(VERSION)

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

# Run tests
test:
	go test -v ./... -coverprofile cover.out

# Build docker image
docker-build: test 
	docker build -t ${IMG} . --build-arg SERVE_PORT=$(SERVE_PORT)

# Push docker image
docker-push:
	docker push ${IMG}

# Update image value in deployment
update-deployment-image: 
	sed -i "s@image:.*@image: $(IMG)@" deploy/deployment.yaml

# Deploy application to kubernetes
deploy-manifests:
	kubectl apply -f deploy/
	
# Create a new release
release: docker-build docker-push update-deployment-image
	@echo "Release successful"

# Deploy latest changes to a cluster
deploy: release
	bash scripts/deploy.sh