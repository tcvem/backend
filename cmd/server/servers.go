package main

import (
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tcvem/backend/pkg/pb"
	"github.com/tcvem/backend/pkg/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func CreateServer(logger *logrus.Logger, db *gorm.DB, interceptors []grpc.UnaryServerInterceptor) (*grpc.Server, error) {
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))

	// register all of our services into the grpcServer
	s, err := svc.NewBasicServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterBackendServer(grpcServer, s)

	certificate, err := svc.NewCertificateServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterCertificateServiceServer(grpcServer, certificate)

	return grpcServer, nil
}
