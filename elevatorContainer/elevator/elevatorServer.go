package elevator

import (
	"elevatorContainer/api"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	defaultElevatorPort = 7777
)

type ElevatorServer struct {
	service api.ElevatorServiceServer
	server  *grpc.Server
}

func NewElevatorServer() *ElevatorServer {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &ElevatorServer{
		service: NewElevatorService(hostname),
	}
}

func (s *ElevatorServer) Start() error {

	glog.V(2).Infof("Start elevator grpc server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", defaultElevatorPort))
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
