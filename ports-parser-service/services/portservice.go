package services

import (
	"context"
	"google.golang.org/grpc"
	"log"
	gc "ports-parser-service/grpc"
)

// PortService type
type PortService interface {
	Add(context.Context, *gc.PortRequest) (*gc.PortResponse, error)
	Get(context.Context, *gc.PortCode) (*gc.Port, error)
}

// PortManager type
type PortManager struct {
	Addr string
}

// NewPortService creates an instance of the PortService
func NewPortService(addr string) PortService {
	if addr == "" {
		addr = "localhost:10000"
	}
	return &PortManager{Addr: addr}
}

// Add port
func (pm *PortManager) Add(ctx context.Context, req *gc.PortRequest) (*gc.PortResponse, error) {

	// Attempt to create a connection to the Port Manager Server
	conn, err := grpc.Dial(pm.Addr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Println("failed to create connecto the Port Manager server")
		return nil, err
	}

	defer conn.Close()

	c := gc.NewPortManagerClient(conn)

	// fire add port request
	log.Printf("sending add port with code %s", req.PortCode)
	return c.AddPort(ctx, req)

}

// Get port
func (pm *PortManager) Get(ctx context.Context, code *gc.PortCode) (*gc.Port, error) {

	// Attempt to create a connection to the Port Manager Server
	conn, err := grpc.Dial(pm.Addr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Println("failed to create connecto the Port Manager server")
		return nil, err
	}

	defer conn.Close()

	c := gc.NewPortManagerClient(conn)

	// send  fertch port for given code
	return c.GetPort(ctx, code)
}
