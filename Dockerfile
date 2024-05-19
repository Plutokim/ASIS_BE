FROM golang:1.21-bullseye
WORKDIR /app
RUN go install github.com/cosmtrek/air@v1.49.0
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .