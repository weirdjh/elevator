package elevator

import (
	"context"
	"elevatorSim/api"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type ElevatorStatus struct {
	Name  string
	State string
	Floor int32
}

type ElevatorContainer struct {
	Id         string
	Name       string
	Conn       *grpc.ClientConn
	ServiceCli api.ElevatorServiceClient
	Port       string
}

func NewElevatorContainer(name string, port string, containerID string) *ElevatorContainer {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	sc := api.NewElevatorServiceClient(conn)
	return &ElevatorContainer{
		Id:         containerID,
		Name:       name,
		Conn:       conn,
		ServiceCli: sc,
		Port:       port,
	}
}

func (e *ElevatorContainer) RemoveConnection() error {
	return e.Conn.Close()
}

func (e *ElevatorContainer) GetElevatorStatus() (*ElevatorStatus, error) {
	response, err := e.ServiceCli.GetElevatorStatus(context.Background(), &api.GetElevatorStatusRequest{})
	if err != nil {
		return nil, fmt.Errorf("Error while get status: %s", err)
	}
	status := &ElevatorStatus{
		Name:  response.Name,
		State: response.State,
		Floor: response.Floor,
	}
	return status, nil
}

func (e *ElevatorContainer) ElevatorUp(dest int32) (*api.ElevatorUpResponse, error) {
	response, err := e.ServiceCli.ElevatorUp(context.Background(), &api.ElevatorUpRequest{Destination: dest})
	if err != nil {
		return nil, fmt.Errorf("Error while elevator up: %s", err)
	}

	if response.Done {
		fmt.Printf("%s -> Up Done : ", e.Name)
	}

	return response, nil
}

func (e *ElevatorContainer) ElevatorDown(dest int32) (*api.ElevatorDownResponse, error) {
	response, err := e.ServiceCli.ElevatorDown(context.Background(), &api.ElevatorDownRequest{Destination: dest})
	if err != nil {
		return nil, fmt.Errorf("Error while elevator up: %s", err)
	}

	if response.Done {
		fmt.Printf("%s -> Down Done : ", e.Name)
	}
	return response, nil
}
