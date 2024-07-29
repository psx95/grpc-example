### Interacting with the Server

This file contains information on how to interact with the server and the various ways to do it.

#### Using CLIs like grpcurl

grpcurl can be used to issue requests to the server.

##### Getting list of services available on the server:
```
grpcurl -plaintext localhost:9997 list
```

##### Getting information about a particular service on the server:
```
grpcurl -plaintext localhost:9997 describe invoiceservice.Invoicer
```

#### Calling a service method:
```
grpcurl -plaintext localhost:9997 invoiceservice.Invoicer/Create
```

#### Calling a service method with request parameters:
```
grpcurl -plaintext -d '{"from": "psx", "to": "someone", "total":{"currency": "INR", "amount": "1240.67"}}' localhost:9997 invoiceservice.Invoicer/Create
```

#### Calling a service method with request parameters defined in a file:
```
grpcurl -plaintext -d "`cat sample_create_request.json`" localhost:9997 invoiceservice.Invoicer/Create
```
OR
```
grpcurl -plaintext -d localhost:9997 invoiceservice.Invoicer/Fetch <sample_fetch_request.json
```

