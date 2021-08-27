FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
WORKDIR /app/cmd
RUN pwd
RUN go build -o main

EXPOSE 5555

CMD ["./main"]