# chuck-norris-api

Golang REST API with MYSQL for persistence

## Application

This is a sample application using the [chuck-norris-api](https://api.chucknorris.io/) package. This application is a demonstration of a basic REST API in Golang which interacts with a MYSQL database as a persistent storage back-end.

### Dependencies

1. `net/http` stdlib is used for the http server and routing for the sake of simplicity.
2. `GORM`, an ORM library for Golang, is used for interacting with the MYSQL database.

### Endpoints

- [x] `GET /banter` returns a list of jokes
- [x] `GET /health/live` checks for the liveness of application
- [x] `GET /health/ready` checks for the readiness of application to accept traffic. Ensures that all backing services are up and running.

### Configuration

The following tables lists the configurable parameters for the application and their default values:

| Parameter      | Default     | Description                             |
| -------------- | ----------- | --------------------------------------- |
| `SQL_HOST`     | `localhost` | Host name for SQL database              |
| `SQL_PORT`     | `3036`      | Access port for SQL                     |
| `SQL_USERNAME` | `root`      | Username for accessing the SQL database |
| `SQL_PASSWORD` |             | Password for accessing the SQL database |
| `SQL_DATABASE` | `banter`    | Database name                           |

**NOTE:** The `SQL_PASSWORD` parameter has no default value for security reasons.

### Logging strategy

Following [12 factor app](https://12factor.net) methodologies, we will rely on `stdout` for logs of this application.

## Deployment

Deployment is divided into the following components:

- Simple storage server i.e. `MYSQL`
- REST API server that exposes `/banter` endpoint to retrieve list of all jokes from the storage server
- Reverse proxy that serves the REST API server

As a best practice, we are not using default port `3306` for the SQL server and instead relying on `3313` as the serving port.

### Requirements

1. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
2. [minikube](https://minikube.sigs.k8s.io/docs/start/)

### Deploy

For end-to-end deployment, run `make release-and-deploy` this command will:

- Build image against latest code
- Push image to remote registry
- Update image reference in [deployment](./deploy/api/deployment.yaml)
- Run deployment script, which performs the following action
  - Ensure that kubectl and minikube are installed. This script will install them automatically on "linux-gnu" OS.
  - Start a minikube cluster
  - Deploy nginx-ingress
  - Deploy MYSQL
  - Deploy the REST API server
  - Print out the URL to access the application

Else, you can run `make deploy` to just deploy the application without creating a new release(build, push docker image)

## Local Development

### Requirements

- golang 1.16

### Execution

- Run `docker-compose up -d` to start the local development environment i.e. mysql; the backing service
- Run `SQL_PASSWORD=root make run` to run the application that serves on port 8080
