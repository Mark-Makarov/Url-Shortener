FROM golang:1.21

WORKDIR /app

RUN apt-get update && apt-get install -y gcc sqlite3

COPY go.mod .
COPY go.sum .

COPY . .

WORKDIR /app/cmd/url-shortener

RUN go build -o server

EXPOSE 9129:9129

CMD ["./server"]