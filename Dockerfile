# Etapa de build
FROM golang:1.22.4 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-test

FROM scratch

WORKDIR /

COPY --from=build /api-test /api-test

EXPOSE 8080

ENTRYPOINT ["/api-test"]
