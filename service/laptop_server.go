package service

import (
	"context"
	"errors"
	"log"

	"github.com/akagami-harsh/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is a server that provice laptop related services.
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopSever creates a new laptop server.
func NewLaptopSever() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop creates a unary RPC to create a new laptop.
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s ", laptop.Id)

	if len(laptop.Id) > 0 {
		//check if it is valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid laptop id is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "can not generate a new UUID: %v", err)
		}
		laptop.Id = id.String()
	}

	// save the laptop to in-memory store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}

	log.Printf("saved laptop with id: %s ", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
