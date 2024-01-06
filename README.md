# PCBook gRPC Service

This project implements a simple gRPC service for managing a catalog of laptops. It includes a server that provides gRPC endpoints for creating and searching for laptops, and a client that can communicate with the server using gRPC.

## Getting Started

To get started with the PCBook gRPC service, follow these instructions:

### Prerequisites

- Go (Golang) programming language installed (version 1.15 or higher recommended)
- Protocol Buffer (Protobuf) compiler installed (`protoc`)
- gRPC-Go plugin for the Protocol Compiler

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/akagami-harsh/pcbook.git
cd pcbook
```

Generating gRPC Code
Before running the server or client, you need to generate the gRPC code from the .proto files:

```sh
make gen
```
This command will generate Go code for the gRPC service and messages based on the .proto files located in the proto directory.

Running the Server
To start the gRPC server, use the following command:

```sh
make server
```
By default, the server will start on port 8080. You can specify a different port by modifying the server target in the Makefile.


Running the Client
To run the gRPC client, use the following command:
```sh
make client
```
The client will attempt to connect to the server at 0.0.0.0:8080. You can specify a different address by modifying the client target in the Makefile.


Cleaning Up Generated Files
If you want to clean up the generated files, you can use the following command:
```sh
make clean
```
This will remove all generated .go files in the pb directory.

Running Tests
To run tests for the project, use the following command:

```sh
make test
```


This will execute all tests in the project and provide a coverage report.

### Project Structure
- `cmd/server/main.go`: The entry point for the gRPC server.
- `cmd/client/main.go`: The entry point for the gRPC client.
- `pb`: Directory containing generated Go files for gRPC.
- `service`: Go package that implements the laptop service logic.
- `proto`: Directory containing Protobuf definitions.
