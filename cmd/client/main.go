package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	atlas_rpc "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tcvem/backend/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	address := fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port"))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	bc := pb.NewBackendClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := bc.GetVersion(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not version: %v", err)
	}
	log.Printf("Version: %s", r.Version)

	uuid := uuid.NewV4()
	certficateORM := pb.CertficateORM{
		Id:    uuid.String(),
		Host:  "www.google.com",
		Port:  "443",
		Notes: "",
	}

	cert, err := certficateORM.ToPB(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Printf("cert: %+v", cert)
	cc := pb.NewCertificateServiceClient(conn)
	_, err = cc.Create(ctx, &pb.CreateCertficateRequest{Payload: &cert})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	certficateORM = pb.CertficateORM{
		Id:    uuid.String(),
		Name:  "gmail",
		Host:  "smtp.gmail.com",
		Port:  "587",
		Notes: "starttls",
	}

	cert, err = certficateORM.ToPB(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Printf("cert: %+v", cert)
	_, err = cc.Update(ctx, &pb.UpdateCertficateRequest{Payload: &cert})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resourceID := atlas_rpc.Identifier{ResourceId: uuid.String()}
	resp, err := cc.Delete(ctx, &pb.DeleteCertficateRequest{Id: &resourceID})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v", resp)
	respList, err := cc.List(ctx, &pb.ListCertficateRequest{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v", resp)
	for _, result := range respList.Results {
		fmt.Printf("host: %s\nport: %s\n", result.Host, result.Port)
	}
}

func init() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(viper.GetString("config.source"))
	if viper.GetString("config.file") != "" {
		log.Printf("Serving from configuration file: %s", viper.GetString("config.file"))
		viper.SetConfigName(viper.GetString("config.file"))
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("cannot load configuration: %v", err)
		}
	} else {
		log.Printf("Serving from default values, environment variables, and/or flags")
	}
	resource.RegisterApplication(viper.GetString("app.id"))
	resource.SetPlural()
}
