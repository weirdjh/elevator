package elevator

import (
	"context"
	"elevatorSim/api"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type ElevatorStatus struct {
	name      string
	direction string
	floor     int32
}

type ElevatorContainer struct {
	id         string
	name       string
	conn       *grpc.ClientConn
	serviceCli api.ElevatorServiceClient
	port       string
}

func NewElevatorContainer(name string, port string, containerID string) *ElevatorContainer {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	sc := api.NewElevatorServiceClient(conn)
	return &ElevatorContainer{
		id:         containerID,
		name:       name,
		conn:       conn,
		serviceCli: sc,
		port:       port,
	}
}

func (e *ElevatorContainer) RemoveConnection() error {
	return e.conn.Close()
}

func (e *ElevatorContainer) GetElevatorStatus() (*ElevatorStatus, error) {
	response, err := e.serviceCli.GetElevatorStatus(context.Background(), &api.GetElevatorStatusRequest{})
	if err != nil {
		return nil, fmt.Errorf("Error while get status: %s", err)
	}
	status := &ElevatorStatus{
		name:      response.Name,
		direction: response.Direction,
		floor:     response.Floor,
	}
	return status, nil
}

func (e *ElevatorContainer) ElevatorUp(dest int32) (*api.ElevatorUpResponse, error) {
	response, err := e.serviceCli.ElevatorUp(context.Background(), &api.ElevatorUpRequest{Destination: dest})
	if err != nil {
		return nil, fmt.Errorf("Error while elevator up: %s", err)
	}
	return response, nil
}

func (e *ElevatorContainer) ElevatorDown(dest int32) (*api.ElevatorDownResponse, error) {
	response, err := e.serviceCli.ElevatorDown(context.Background(), &api.ElevatorDownRequest{Destination: dest})
	if err != nil {
		return nil, fmt.Errorf("Error while elevator up: %s", err)
	}
	return response, nil
}
