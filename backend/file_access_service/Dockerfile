FROM golang:1.25-alpine

WORKDIR /app

RUN apk update && apk add --no-cache \
    build-base \
    sqlite \
    sqlite-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main ./cmd

EXPOSE 8081

CMD ["./bin/main"]