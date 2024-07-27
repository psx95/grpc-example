package main

import (
	"context"
	"fmt"
	pb "invoiceservice/pb/invoiceservice/gen_proto"
)

type InvoiceGRPCServer struct {
	// need to embed this since this a requirement from the
	// generated protobuf from grpc plugin
	pb.UnimplementedInvoicerServer
}

func (server *InvoiceGRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	fmt.Printf("Create called with request: %v\n", req)
	resp := &pb.CreateResponse{
		Pdf: nil,
	}
	return resp, nil
}

func (server *InvoiceGRPCServer) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	fmt.Println("Fetch called")
	return nil, fmt.Errorf("not implemented")
}
