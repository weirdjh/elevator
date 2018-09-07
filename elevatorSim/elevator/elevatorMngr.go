package elevator

import (
	"context"
	"log"

	"elevatorSim/api"

	"google.golang.org/grpc"
)

type ElevatorMngr struct {
	//elevators []*Elevator
	//eventChan EventWatch
}

func NewElevatorMngr() *ElevatorMngr {
	return &ElevatorMngr{
		//elevators: []*Elevator{},
		//eventChan: NewEventWatch(),
	}
}

/*
func (em *ElevatorMngr) MngElevator() {

	start := time.Now()
		for time.Since(start) < 10*time.Second {
			select {
			case Event <- em.eventChan:
				fmt.Printf("%s", Event.ElevName)
			}
		}

		elapsed := time.Since(start)

}

func (em *ElevatorMngr) GetElevator(name string) (*Elevator, error) {

	for _, elevator := range em.elevators {
		if elevator.name == name {
			return elevator, nil
		}
	}
	return nil, fmt.Errorf("elevator not found")
}

func (em *ElevatorMngr) AddElevator(name string) {
	elevator := NewElevator(name)
	em.elevators = append(em.elevators, elevator)
}

func (em *ElevatorMngr) MngElevatorUp(name string, dst int) error {
	elevator, err := em.GetElevator(name)
	if err != nil {
		return err
	}
	elevator.ElevatorUp(dst)
	return nil
}
*/
/*
func (em *ElevatorMngr) PingElevator() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
*/

func (em *ElevatorMngr) GetElevatorsStatus() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewElevatorServiceClient(conn)

	response, err := c.GetElevatorStatus(context.Background(), &api.GetElevatorStatusRequest{})
	if err != nil {
		log.Fatalf("Error while get status: %s", err)
	}
	log.Println(response)
}

/*
func (em *ElevatorMngr) PrintElevator() {
	for _, elevator := range em.elevators {
		fmt.Printf("%s : %d ", elevator.name, elevator.floor)
	}
	fmt.Printf("\n")
}
*/
