package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/akagami-harsh/pcbook/pb"
	"github.com/akagami-harsh/pcbook/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the sever port")
	flag.Parse()
	log.Printf("start the server on port: %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
