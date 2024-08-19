FROM golang:1.22-alpine as build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /build/main ./cmd

FROM alpine:latest as runtime

WORKDIR /app

COPY --from=build /build/main . 

EXPOSE 3000

CMD ["./main"]
