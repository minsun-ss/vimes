FROM golang:1.23.4-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o vimes ./cmd

EXPOSE 24680
ENTRYPOINT ["/app/vimes"]
