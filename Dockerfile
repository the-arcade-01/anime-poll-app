# Stage 1: Build the Go application
FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

# Stage 2: Run the Go application
FROM golang:1.23.2

WORKDIR /app

COPY --from=builder /app .

CMD ["make", "run"]