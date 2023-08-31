package controller

import (
	"context"
	"netxd_grpc_mongo/netxd_dal/interfaces"
	"netxd_grpc_mongo/netxd_dal/models"
	pro "netxd_grpc_mongo/netxd_customer"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	db := &models.Customer{Customer_ID: req.Customer_ID}
	result, err := CustomerService.CreateCustomer(db)
	if err != nil {
		return nil, err
	}
	responseCustomer := &pro.CustomerResponse{
		Customer_ID: result.Customer_ID,
		
	}
	
	return responseCustomer, nil
}