#!/usr/bin/env bash

(cd ./1.deploy && kubectl delete -f .)

docker rmi --force "$(docker inspect --format="{{.Id}}" localhost:32000/dapr-go-calculator-adder:registry)"
docker rmi --force "$(docker inspect --format="{{.Id}}" localhost:32000/dapr-go-calculator-react-calculator:registry)"

docker build -f adder/Dockerfile adder/. -t localhost:32000/dapr-go-calculator-adder:registry
docker build -f react-calculator/Dockerfile react-calculator/. -t localhost:32000/dapr-go-calculator-react-frontend:registry

microk8s.ctr images rm localhost:32000/dapr-go-calculator-adder:registry
microk8s.ctr images rm localhost:32000/dapr-go-calculator-react-frontend:registry
microk8s.ctr images rm --sync "$(microk8s.ctr images list -q | grep localhost:32000/dapr-go-calculator-adder@sha256)"
microk8s.ctr images rm --sync "$(microk8s.ctr images list -q | grep localhost:32000/dapr-go-calculator-react-frontend@sha256)"

docker push localhost:32000/dapr-go-calculator-adder:registry
docker push localhost:32000/dapr-go-calculator-react-frontend:registry

(cd ./1.deploy && kubectl apply -f .)
