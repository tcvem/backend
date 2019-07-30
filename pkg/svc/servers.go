package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/tcvem/backend/pkg/pb"
)

// NewCertificateServer returns an instance of the default Certificate server interface
func NewCertificateServer(database *gorm.DB) (*pb.CertificateServiceDefaultServer, error) {
	return &pb.CertificateServiceDefaultServer{DB: database}, nil
}
