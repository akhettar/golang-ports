# Ports Parser / Manager
This project demonstrates the design and the implementation of Ports manger. This project consists of two services:

1. [ports-parser-service](ports-parser-service)

This is service exposes REST API whereby you can parse a large file containing Port details entry and push them to the `ports-manager-servie` through gRPC protocol

2. [ports-manager-service](ports-manager-service)

This service is receive a given Port details and persist the details into a persistent store

## How to interact with the serivces

1. Run the services using using [docker compose file](docker-compose.yml)

2. Send the request below to process ports - there are alread sample files packaged within the docker image

Process Ports Rquest:
```
curl -X POST http://localhost:8080/ports -d '{"File":"/tmp/ports_small.json","Size":212111}'

curl -X POST http://localhost:8080/ports -d '{"File":"/tmp/ports.json","Size":45444545}'       
```
Process Ports Response:

```
{"file":"ports_small.json","number_of_records":7}
```

3. Query a given port

```  
 curl http://localhost:8080/ports/AEAJM
```

```
{
  "port_code": "AEAJM",
  "name": "Ajman",
  "city": "Ajman",
  "country": "United Arab Emirates",
  "alias": null,
  "regions": null,
  "coordinates": [
    55.5136433,
    25.4052165
  ],
  "province": "Ajman",
  "timezone": "Asia/Dubai",
  "unlocs": null,
  "code": "52000"
}
```