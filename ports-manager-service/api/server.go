package api

import (
	"context"
	"fmt"

	pb "ports-manager-service/grpc"
	m "ports-manager-service/model"
	repo "ports-manager-service/repo"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// PortServer type
type PortServer struct {
	repo.Repository
}

// NewServer creates an instance of the PortServer
func NewServer() *PortServer {
	repo := repo.NewInMemoryRepository()
	return &PortServer{repo}
}

func (ps *PortServer) AddPort(ctx context.Context, req *pb.PortRequest) (*pb.PortResponse, error) {
	if err := ps.Repository.AddPort(m.ToPort(req)); err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}
	return &pb.PortResponse{Message: fmt.Sprintf("Port with code: %s successfully added", req.PortCode)}, nil
}
func (ps *PortServer) GetPort(ctx context.Context, req *pb.PortCode) (*pb.Port, error) {
	found, err := ps.Repository.GetPort(req.Code)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}

	if found == nil {
		return nil, status.Errorf(codes.Unimplemented, fmt.Sprintf("Port with code: %s not found", req.Code))
	}
	return m.FromPort(found), nil
}
