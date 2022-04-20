FROM golang:1.18.1-alpine
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]