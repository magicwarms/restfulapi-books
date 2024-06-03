FROM golang:1.21.1-alpine AS build-stage

LABEL MAINTAINER="Andhana Utama <andhanautama@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o books-assessment-app .

FROM alpine:3.19 AS final

WORKDIR /app

COPY --from=build-stage /app/books-assessment-app .

COPY .env ./

RUN apk add --no-cache ca-certificates

EXPOSE 9005

CMD ["./books-assessment-app"]

