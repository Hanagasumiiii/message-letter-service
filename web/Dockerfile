FROM golang:1.21 as builder
LABEL authors="admin"

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

CMD ["./main"]
