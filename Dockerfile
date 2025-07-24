FROM golang:alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

# Create the /tmp/certs directory and write the certificates from build arguments
ARG CA_CERT
ARG CLIENT_CERT
ARG CLIENT_KEY
ARG APP_ENV

EXPOSE 3000

CMD ["./main"]