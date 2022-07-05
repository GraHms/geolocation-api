FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod download && go mod verify
RUN go build main.go

ENV PORT=:8080
EXPOSE $PORT

CMD ["/app/main"]