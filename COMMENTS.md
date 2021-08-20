# Thoughts behind things

## Architecture

Wanted to keep the architecture as simple as possible and tried to use least amount of external libraries. That's why I went with `net/http` package for the server. Although `database/sql` stdlib would have been sufficient for interacting with database but I wanted an ORM type library, so chose [gorm](https://github.com/go-gorm/gorm).

## Bootstrapping

In future, we can use something like [liquibase](https://www.liquibase.org/) to make sure that the database is in the right state. Although it's preferred for schematic changes but we can maybe insert dummy data using that as well.

For now, we just check if the database is empty and then add dummy values into it.

## Testing

Testing has not been implemented yet although I am planning on adding concrete tests. Will use [testify](https://github.com/stretchr/testify) for assertions.

## Secrets

Right now, for demonstration purposes, Kubernetes secrets are stored in git repository as they are. This is a bad practice and should be avoided at all costs. We can use tools like [vault](https://www.vaultproject.io/) for storing secrets in a dedicated server or use [sealed-secrets](https://github.com/bitnami-labs/sealed-secrets) for storing encrypted secrets in git repository.

## Deployment Automation

The plan was to write a script that can create a cluster using [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/) and then deploy the application. However, for now, I am just going to use [minikube](https://minikube.sigs.k8s.io/) for deployment.

## Manifests

Respective helm charts can be used to deploy api, server and reverse proxy. This would minimize the amount of code(yaml) that we have to manage.
