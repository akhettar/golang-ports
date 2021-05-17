package api

import (
	"context"
	"errors"
	pb "ports-manager-service/grpc"
	m "ports-manager-service/model"
	repo "ports-manager-service/repo"
	repository "ports-manager-service/repo"
	"reflect"
	"testing"
)

func TestPortServer_AddPort(t *testing.T) {
	type fields struct {
		Repository repo.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.PortRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.PortResponse
		wantErr bool
	}{
		{
			name:    "Add success",
			fields:  fields{repository.NewInMemoryRepository()},
			args:    args{ctx: context.Background(), req: &pb.PortRequest{PortCode: "AEAJM", Name: "Ajman"}},
			want:    &pb.PortResponse{Message: "Port with code: AEAJM successfully added"},
			wantErr: false,
		},
		{
			name:    "Add failure",
			fields:  fields{NewFakeRepository()},
			args:    args{ctx: context.Background(), req: &pb.PortRequest{PortCode: "AEAJM", Name: "Ajman"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PortServer{
				Repository: tt.fields.Repository,
			}
			got, err := ps.AddPort(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PortServer.AddPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PortServer.AddPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortServer_GetPort(t *testing.T) {
	// Given port exist
	rep := repo.NewInMemoryRepository()
	rep.AddPort(m.Port{PortCode: "AEAJM"})

	type fields struct {
		Repository repo.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.PortCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Port
		wantErr bool
	}{
		{
			name:    "Get Port for a given code returns Port",
			fields:  fields{rep},
			args:    args{ctx: context.Background(), req: &pb.PortCode{Code: "AEAJM"}},
			want:    &pb.Port{PortCode: "AEAJM"},
			wantErr: false,
		},
		{
			name:    "Get Port for unknown code returns not found",
			fields:  fields{rep},
			args:    args{ctx: context.Background(), req: &pb.PortCode{Code: "FAKE"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PortServer{
				Repository: tt.fields.Repository,
			}
			got, err := ps.GetPort(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PortServer.GetPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PortServer.GetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

type FakeRepository struct{}

func NewFakeRepository() repo.Repository {
	return FakeRepository{}
}

func (r FakeRepository) AddPort(p m.Port) error {
	return errors.New("failed to persist the port")
}

func (r FakeRepository) GetPort(code string) (*m.Port, error) {
	return nil, nil
}
