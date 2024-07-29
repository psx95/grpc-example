package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "invoiceservice/pb/invoiceservice/gen_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 9997, "The server port")
)

func main() {
	flag.Parse()
	fmt.Printf("Starting invoice server at port: %v\n", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterInvoicerServer(server, &InvoiceGRPCServer{})

	// Register reflection service on the server
	// allows the clients to get information about the server, such as the supported methods
	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
