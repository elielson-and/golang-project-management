
FROM golang:1.21-alpine

RUN apk add --no-cache git curl

# Compiledaemon p hot reload
RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#RUN swag init -g cmd/api/main.go


COPY . .

CMD ["CompileDaemon", "--build=go build -o main ./cmd/api/main.go", "--command=./main"]
