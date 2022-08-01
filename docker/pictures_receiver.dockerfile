FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build ./cmd/main.go

FROM alpine:latest as production

COPY --from=build /app/main .

EXPOSE 9090

CMD ["./main"]