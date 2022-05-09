package main

import (
	"fmt"
	"net"

	"github.com/evgeniy-dammer/simplevaluesgrpc/src/addservice/handlers"
	addservice "github.com/evgeniy-dammer/simplevaluesgrpc/src/addservice/proto"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":1111")

	if err != nil {
		fmt.Println(err)
	}

	defer listen.Close()

	addServ := handlers.AddServiceServer{}
	grpcServer := grpc.NewServer()

	addservice.RegisterAddServiceServer(grpcServer, &addServ)

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}

}
