package main

import (
	"context"
	"fmt"

	addservice "github.com/evgeniy-dammer/simplevaluesgrpc/src/addservice/proto"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:1111", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	addServ := addservice.NewAddServiceClient(connection)

	response, err := addServ.Add(context.Background(), &addservice.AddRequest{A: 3, B: 5})

	if err != nil {
		fmt.Println(err)
	} else {
		result := response.Result
		fmt.Println("Result: ", result)
	}
}
