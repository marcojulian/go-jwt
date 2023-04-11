FROM golang:1.20-alpine3.17 AS build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.17
COPY --from=build-env /app/main /app/main
COPY wait-for-db.sh /app/wait-for-db.sh
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY .env ./
EXPOSE 8080
CMD ["/bin/sh", "/app/wait-for-db.sh", "db", "5432", "/app/main"]