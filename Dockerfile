# syntax=docker/dockerfile:1

#####################################################
# Build
#####################################################
FROM golang:1.23-alpine AS builder

RUN apk --update add \
    ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app

#####################################################
# Application
#####################################################
FROM gcr.io/distroless/static:nonroot AS app

WORKDIR /
COPY --from=builder /app /app

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
