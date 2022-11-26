FROM golang:latest

WORKDIR /build

ADD ./ /build/

RUN CGO_ENABLED=0 GOOS=linux go build -o go-k8s-api ./cmd/.

ENTRYPOINT [ "./go-k8s-api" ]