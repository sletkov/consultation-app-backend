FROM golang:1.22.1-alpine3.19

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata git

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main cmd/main.go

CMD ["./main"]