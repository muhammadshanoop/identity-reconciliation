
# Stage 1: Build the Go app
FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Stage 2: Run the app
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "./main" ]
