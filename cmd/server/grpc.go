package main

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/infobloxopen/atlas-app-toolkit/auth"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer(logger *logrus.Logger, db *gorm.DB) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
		// logging middleware
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

		// Request-Id interceptor
		requestid.UnaryServerInterceptor(),

		// validation middleware
		grpc_validator.UnaryServerInterceptor(),

		auth.LogrusUnaryServerInterceptor(),

		// collection operators middleware
		gateway.UnaryServerInterceptor(),
	}
	return CreateServer(logger, db, interceptors)
}
