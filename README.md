# Ports Parser / Manager
This project demonstrates the design and the implementation of Ports manger. This project consists of two services:

1. [ports-parser-service](ports-parser-service)

This is service exposes REST API whereby you can parse a large file containing Port details entry and push them to the `ports-manager-servie` through gRPC protocol

2. [ports-manager-service](ports-manager-service)

This service is receive a given Port details and persist the details into a persistent store

## How to interact with the serivce

There is docker-compose.yml which you can run. But unfortunately, the network link isn't working with gRPC. Fix is in progress :). Alternatively, you can run the two services independtly see the REDAME files for each service

Rquest:
```
curl -X POST http://localhost:8080/ports -d '{"File":"ports_small.json","Size":212111}'
       
```

Response:

```
{"file":"ports_small.json","number_of_records":7}
```

