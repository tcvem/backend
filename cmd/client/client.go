package main

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
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

func (m *TcvemClient) CreateCertficate(host, port, notes string) (*pb.CreateCertficateResponse, error) {

	// We can now create stubs that wrap conn:
	stub := pb.NewCertificateServiceClient(m.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uuid := uuid.NewV4()
	certficateORM := pb.CertficateORM{
		Id:    uuid.String(),
		Host:  host,
		Port:  port,
		Notes: notes,
	}

	cert, err := certficateORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	req := &pb.CreateCertficateRequest{Payload: &cert}
	resp, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
