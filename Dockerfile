FROM golang:1.18.1-alpine
RUN apk update && apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

RUN cp /app/main .

CMD ["/app/main"]