FROM golang:1.19

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.script

# build go app
RUN go mod download
RUN go build -o ./bin/migrate.bin ./migrations/migrate.go
RUN go build -o ./bin/api.bin ./cmd/main.go

ENTRYPOINT [ "./wait-for-postgres.script", "postgres_test", "postgres", "123", "/app/bin/migrate.bin", "/app/bin/api.bin" ]