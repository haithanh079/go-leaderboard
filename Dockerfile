FROM golang:alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-redis/redis
RUN go get github.com/joho/godotenv

ADD . /go/src/github.com/haithanh079/go-leaderboard

WORKDIR /go/src/github.com/haithanh079/go-leaderboard

RUN go install

EXPOSE 8000

ENTRYPOINT /go/bin/go-leaderboard
