FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o customerapi customer-api/cmd/main.go

CMD ["./customerapi"]