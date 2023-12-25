# Choose whatever you want, version >= 1.16
FROM golang:1.21-alpine

EXPOSE 8099

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]