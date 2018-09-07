package elevator

import (
	"elevatorContainer/api"
	"fmt"
	"log"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type ElevatorServer struct {
	service api.ElevatorServiceServer
	server  *grpc.Server
}

func NewElevatorServer() *ElevatorServer {
	return &ElevatorServer{
		service: NewElevator("tmp"),
	}
}

func (s *ElevatorServer) Start() error {

	glog.V(2).Infof("Start elevator grpc server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.server = grpc.NewServer()

	api.RegisterElevatorServiceServer(s.server, s.service)

	// start the server
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return nil
}
