package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
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
	c := pb.NewBackendClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetVersion(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not version: %v", err)
	}
	log.Printf("Version: %s", r.Version)
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
