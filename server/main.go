package main

import (
	"context"
	"fmt"
	"net"

	pro "netxd_grpc_mongo/netxd_customer" // Import the generated Go code
	"netxd_grpc_mongo/netxd_customer_config/config"
	"netxd_grpc_mongo/netxd_customer_config/constants"
	"netxd_grpc_mongo/netxd_dal/services"
	controller "netxd_grpc_mongo/netxd_customer_controller/contoller"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "bankdb", "profiles")
	controller.CustomerService = services.InitCustomerService(CustomerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() //get the pointer reference of the server
	pro.RegisterCustomerServiceServer(s, &controller.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
