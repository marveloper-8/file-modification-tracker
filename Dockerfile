FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o file-modification-tracker ./cmd/service

EXPOSE 8080

CMD ["./file-modification-tracker"]