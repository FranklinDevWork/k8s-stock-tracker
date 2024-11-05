FROM golang:1.23.2

ENV PORT=8080

WORKDIR /go/src/k8s-stock-tracker

COPY api /go/src/k8s-stock-tracker/api/
COPY go.mod /go/src/k8s-stock-tracker/
COPY go.sum /go/src/k8s-stock-tracker/

RUN go build api/main.go

EXPOSE ${PORT}

ENTRYPOINT [ "./main" ]