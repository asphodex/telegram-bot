FROM golang:latest

COPY . .

RUN go mod download

RUN go build -o raketa-bot ./cmd/telegram-bot/main.go

CMD ["/go/bin/telegram-bot"]

