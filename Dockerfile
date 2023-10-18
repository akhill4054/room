FROM golang:1.19.5-alpine

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /bin/room-webserver

RUN go install github.com/silenceper/gowatch@latest

EXPOSE $PORT

ENTRYPOINT ["/bin/room-webserver"]
