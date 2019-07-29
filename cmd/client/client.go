package main

import (
	"context"
	"time"

	"github.com/tcvem/backend/pkg/pb"
	"google.golang.org/grpc"
)

type TcvemClient struct {
	// The tcvem backend Server
	Host string
	// The tcvem backend Server Connection
	Conn *grpc.ClientConn
}

func NewTcvemClient(host string) (*TcvemClient, error) {
	c := TcvemClient{}

	err := c.GetConn(host)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *TcvemClient) GetConn(host string) error {
	// First we create the connection:
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	m.Conn = conn
	return nil
}

func (m *TcvemClient) GetListCertficate() (*pb.ListCertficateResponse, error) {

	// We can now create stubs that wrap conn:
	stub := pb.NewCertificateServiceClient(m.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListCertficateRequest{}
	resp, err := stub.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
