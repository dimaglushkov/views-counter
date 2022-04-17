FROM golang:1.18.1-buster

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy

RUN go build -o main .

CMD ["./main"]