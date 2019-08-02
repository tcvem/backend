package cc

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/tcvem/backend/pkg/pb"
	"google.golang.org/grpc"
)

type TcvemClient struct {
	conn *grpc.ClientConn
}

func NewTcvemClient(conn *grpc.ClientConn) *TcvemClient {
	return &TcvemClient{conn: conn}
}

func (m *TcvemClient) CreateCertficate(host, port, notes string) (*pb.CreateCertficateResponse, error) {

	client := pb.NewCertificateServiceClient(m.conn)
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
	resp, err := client.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *TcvemClient) GetListCertficate() (*pb.ListCertficateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListCertficateRequest{}
	client := pb.NewCertificateServiceClient(m.conn)
	resp, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
