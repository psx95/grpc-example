#! /bin/bash
echo "Generating Go code from proto files"

mkdir -p invoiceservice/pb
protoc --proto_path=invoiceservice/proto/ invoiceservice/proto/*.proto --go-grpc_out=invoiceservice/pb   

echo "Code generated from proto files"
