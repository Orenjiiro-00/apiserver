# Build stage
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Enable CGO for sqlite3
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /app/metrics-server .

# Run stage
FROM debian

WORKDIR /app

COPY --from=build /app/metrics-server /app/metrics-server

EXPOSE 8080

CMD ["/app/metrics-server"]
