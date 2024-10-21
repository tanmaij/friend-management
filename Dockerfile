# syntax=docker/dockerfile:1
FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go install github.com/volatiletech/sqlboiler/v4@v4.16.2 \
    && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest \
    && go install github.com/vektra/mockery/v2@v2.46.3 \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1
