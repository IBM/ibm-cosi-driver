#!/bin/bash

eval $(minikube docker-env)
#kubectl delete ns ibm-cosi-driver
# minikube cache delete ibm/ibm-cosi-driver:latest
docker rmi ibm-cosi-driver -f
docker rmi ibm/ibm-cosi-driver -f

make clean
make build
make container
docker tag ibm-cosi-driver:latest ibm/ibm-cosi-driver:latest
# minikube cache add ibm/ibm-cosi-driver:latest
# kubectl apply -k .
