# builder image
FROM golang:1.18.1-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o be-service-teman-bunda .

# generate clean, final image for end users
FROM alpine
COPY --from=builder /build/be-service-teman-bunda .

EXPOSE 9000

CMD ./be-service-teman-bunda