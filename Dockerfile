ARG GO_VERSION=1.21rc4

FROM golang:${GO_VERSION}-alpine AS build

RUN apk update && apk upgrade
RUN apk add --no-cache\
    git\
    gcc make libc-dev g++\
    && rm -rf /var/cache/apk/*

ENV APP_HOME=/app
WORKDIR $APP_HOME

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd

FROM scratch AS bin

COPY --from=build /app/main /main

ENTRYPOINT ["/main"]
