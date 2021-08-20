#!/bin/bash

# exit when any command fails
set -e

ensure_minikube_is_installed() {
    if ! command -v $MINIKUBE_VERSION_COMMAND &> /dev/null ; then
        echo -e "Minikube is not installed"
        if [[ "$OSTYPE" == "linux-gnu"* ]]; then
            echo -e "Installing minikube"
            curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
            mv minikube-linux-amd64 minikube
        else 
            echo -e "Please install minikube and then run this script"
            exit 1
        fi
    fi
}

ensure_kubectl_is_installed() {
    if ! command -v $KUBECTL_VERSION_COMMAND &> /dev/null ; then
        echo -e "Kubectl is not installed"
        if [[ "$OSTYPE" == "linux-gnu"* ]]; then
            curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
        else 
            echo -e "Please install kubectl and then run this script"
            exit 1  
        fi  
    fi
}

start_minikube() {
    # Check if `minikube status` is Running. If it's not, then create a cluster
    if [[ $(minikube status | grep -o "Running" 2>/dev/null | uniq) != "Running" ]]; then
        echo -e "Creating new minikube cluster"
        minikube start

        # Wait for cluster to start
        echo -e "Waiting for cluster to start"
        while [ "$ready" != "Running" ]; do
            ready=$(minikube status | grep apiserver | grep -o "Running")
            sleep 5
        done
        echo -e "Minikube cluster started\n"
    fi
}

# Variables
MINIKUBE_VERSION_COMMAND="minikube version"
KUBECTL_VERSION_COMMAND="kubectl version"
NAMESPACE=chuck-norris-api
DEPLOYMENT=chuck-norris-api
STATEFULSET=mysql

#################################################
################ Start of script ################
#################################################

echo -e "\n========= Application deployment started ==========\n"

ensure_kubectl_is_installed

echo -e "Kubectl is installed"

ensure_minikube_is_installed

echo -e "Minikube is installed"

start_minikube

echo -e "Minikube cluster started"

# Deploy Ingress Controller
minikube addons enable ingress

# Wait till Ingress Controller is deployed
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=150s

# Create namespace 
kubectl create namespace $NAMESPACE || true

# Create mysql statefulset
kubectl -n $NAMESPACE apply -f deploy/mysql/

# Wait till statefulset is available
kubectl wait -n $NAMESPACE \
  --for=condition=ready pod \
  --selector=app=mysql \
  --timeout=150s


# Create application resources
kubectl -n $NAMESPACE apply -f deploy/api/

# Wait till deployment is available
kubectl wait -n $NAMESPACE \
  --for=condition=ready pod \
  --selector=app=chuck-norris-api \
  --timeout=150s

echo -e "Application deployed"

CLUSTER_IP=$(minikube ip)/banter

echo -e "\n========= Application is accesible at $CLUSTER_IP. Triggering REST request to banter endpoint ==========\n"

sleep 10s
curl $CLUSTER_IP
echo -e "\n\n"
