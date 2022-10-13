FROM golang:1.18.2
RUN go get github.com/codegangsta/gin
RUN mkdir /app
WORKDIR /app