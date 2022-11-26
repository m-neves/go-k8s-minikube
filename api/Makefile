all: build load

build:
	docker build -f Dockerfile -t go-k8s-api:1.0 .

load:
	minikube image load go-k8s-api:1.0

run:
	docker run go-k8s-api:1.0