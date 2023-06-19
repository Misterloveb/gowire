FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV GW_MYSQL.IP=gowire_mysql
ENV GW_APP.IP=gowire_app
ENV GW_MYSQL.PORT=3306
ENV GW_REDIS.IP=gowire_redis
ENV GW_REDIS.PORT=6379
ENV GW_REDIS.PASSWORD=654321

COPY . .
RUN go build -v -o /app/gowire ./cmd/server/main.go

CMD ["/app/gowire"]

