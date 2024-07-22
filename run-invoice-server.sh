#! /bin/bash
# Script to build and run the invoice service on a specified port.
# Port on which to run the server can be passed as positional argument.

# defualt-value syntax with positional arguments
# Sample command: ./run-invoice-server.sh 8080
# The above sample command would attempt to start server at port 8080. 
PORT=${1:-9997}

echo "Attempting to start invoice service at port $PORT"
go build -C invoiceservice

./invoiceservice/invoiceservice --port=$PORT
