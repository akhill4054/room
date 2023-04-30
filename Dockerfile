FROM golang:1.19.5-alpine

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /room-web-server

EXPOSE $PORT

ENTRYPOINT ["/room-web-server"]
