package elevator

import (
	"elevatorSim/api"
	"log"

	"google.golang.org/grpc"
)

type ElevatorContainer struct {
	id         string
	name       string
	conn       *grpc.ClientConn
	serviceCli api.ElevatorServiceClient
	port       string
}

func NewElevatorContainer(name string, port string, containerID string) *ElevatorContainer {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1"+port, grpc.WithInsecure())
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
