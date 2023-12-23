FROM golang:1.20.7 as modules
RUN echo test
RUN echo test
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20.7 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /grpc-server
WORKDIR /grpc-server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/server ./main.go
EXPOSE 8080
ENTRYPOINT ["grpc-server"]