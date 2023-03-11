FROM golang:1.20.1

ARG PORT

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/crud-restful-api main.go

EXPOSE ${PORT}

ENTRYPOINT ["/app/crud-restful-api"]