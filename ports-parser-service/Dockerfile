# builder image
FROM golang:1.15.12-alpine3.13 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/app .
COPY --from=builder /build/ports.json /tmp/ports.json
COPY --from=builder /build/ports_small.json /tmp/ports_small.json

EXPOSE 8080

# executable
ENTRYPOINT [ "./app" ]