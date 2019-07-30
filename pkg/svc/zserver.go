package svc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	"github.com/tcvem/backend/pkg/pb"
)

const (
	// version is the current version of the service
	version = "0.0.1"
)

// Default implementation of the Backend server interface
type server struct{ db *gorm.DB }

// GetVersion returns the current version of the service
func (server) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

// NewBasicServer returns an instance of the default server interface
func NewBasicServer(database *gorm.DB) (pb.BackendServer, error) {
	return &server{db: database}, nil
}
