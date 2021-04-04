# Builds a simple Go environment and runs our app
FROM golang:latest
RUN mkdir /go-api-calendar
ADD . /go-api-calendar
WORKDIR /go-api-calendar
RUN go build -o main .
CMD ["/go-api-calendar/main"]