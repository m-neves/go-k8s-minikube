all: build load

build:
	docker build -f Dockerfile -t go-k8s-stress:1.0 .

load:
	minikube image load go-k8s-stress:1.0

run:
	docker run -p 8080:8080 go-k8s-stress:1.0

deploy:
	kubectl apply -f k8s/deploy.yaml
	kubectl apply -f k8s/load_balancer.yaml

# minikube tunnel to expose load balancer external ip
# kubectl get service to view port