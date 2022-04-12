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



kubectl delete bucketclass --all
kubectl delete bucketrequest --all
kubectl delete bucketaccessclass --all
kubectl delete bucketaccessrequest --all
kubectl delete pods awscli

kubectl create -f examples/bucketclass.yaml
kubectl create -f examples/bucketrequest.yaml
kubectl create -f examples/bucketaccessclass.yaml
kubectl create -f examples/bucketaccessrequest.yaml
kubectl create -f examples/demopod.yaml