# builder image
FROM golang:1.15.12-alpine3.13 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/app .

EXPOSE 10000

# executable
ENTRYPOINT [ "./app" ]