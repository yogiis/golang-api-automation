FROM golang:1.20.1

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/crud-restful-api main.go

EXPOSE 3000

ENTRYPOINT ["/app/crud-restful-api"]