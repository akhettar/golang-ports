all: generate_protobuf build test

generate_protobuf:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/port_manager.proto

build:
	golint ./...
	go build -o pps

test:
	go test ./...

run:
	./pps	
