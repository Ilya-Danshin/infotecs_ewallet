# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app
COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-ewallet ./cmd/ewallet/main.go

EXPOSE 9445
CMD ["/docker-go-ewallet"]
