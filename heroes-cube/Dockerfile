<<<<<<< HEAD
FROM golang:1.22 AS builder

WORKDIR /app/

COPY . .

WORKDIR /app/api

RUN go build api.go

EXPOSE 8085

CMD ["./api"]
=======
FROM golang:1.22.2 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o heroes-api api/api.go 

FROM scratch
WORKDIR /
COPY --from=build /app/heroes-api /api
ENTRYPOINT ["/api"]
>>>>>>> 5dbcaa9 (add dockerfile)
