FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app/api

RUN go build -v -o main ./...

CMD ["/app/api/main"]
