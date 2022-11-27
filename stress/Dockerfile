FROM golang:latest

WORKDIR /build

ADD ./ /build/

RUN CGO_ENABLED=0 GOOS=linux go build -o go-k8s-stress ./cmd/.

EXPOSE 8080

ENTRYPOINT [ "./go-k8s-stress" ]