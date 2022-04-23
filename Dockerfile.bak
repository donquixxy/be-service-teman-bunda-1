FROM golang:1.18.1-alpine

WORKDIR /app

COPY . .

RUN go build -o be-service-teman-bunda

EXPOSE 9000

CMD ./be-service-teman-bunda
