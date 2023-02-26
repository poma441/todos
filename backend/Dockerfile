FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o todos.bin ./cmd/main.go
RUN go build -o migrations.bin ./migrations/migrate.go