all: build load

build:
	docker build -f Dockerfile -t go-k8s-srv:1.0 .

load:
	minikube image load go-k8s-srv:1.0

run:
	docker run -p 8080:8080 go-k8s-srv:1.0