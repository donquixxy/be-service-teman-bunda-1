# builder image
FROM golang:1.18.1-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o be-service-teman-bunda .

# generate clean, final image for end users
FROM alpine
RUN apk update && apk add ca-certificates && apk add tzdata
COPY --from=builder /build .
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
EXPOSE 9000

CMD ./be-service-teman-bunda