# Ports Manager Service

This is a gRPC service that accepts call to persist Port details in a content store

## How to run this service

1. Build and run
```
go build && ./ports-manager-service
```

2. simply run: `make` this will generate protobuf files, build, lint, test and run the server

You can run this service and port manager service with [docker-compose](../docker-compose.yml)



