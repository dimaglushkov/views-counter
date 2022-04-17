FROM golang:1.18.1-buster

WORKDIR /app

ENV DB_DSN=assets/storage.db
ENV PORT=13004

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy

VOLUME ./assets/:./assets/

RUN go build -o main .

EXPOSE 13004

CMD ["./main"]