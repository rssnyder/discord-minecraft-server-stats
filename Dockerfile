FROM golang:1.18-alpine
LABEL org.opencontainers.image.source https://github.com/rssnyder/discord-minecraft-server-stats

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /discord-minecraft-server-stats

ENTRYPOINT /discord-minecraft-server-stats -token "$TOKEN" -domain "$DOMAIN" -status "$STATUS" -loop "$LOOP" -users "$USERS"
