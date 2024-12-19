FROM golang:1.23.1-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

COPY .env .env
COPY ./schema /schema

RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]