# Ports Parser Service

This is a rest service that accepts REST call to download a Json file containing Port details and stream process all these details to another Service usihng GRPC protocol

## Dependencies 
This service has a dependency on [ports-manager-service](../ports-manager-service)

## How to run this service

1. Run the  [ports-manager-service](../ports-manager-service) first

2. Build and run
```
go build && ./ports-parser-service
```

3. Send Request
Rquest:
```
curl -X POST http://localhost:8080/ports -d '{"File":"ports_small.json","Size":212111}'
       
```

Response:

```
{"file":"ports_small.json","number_of_records":7}
```

Alternatively, you can use [docker-compose](../docker-compose.yml)

