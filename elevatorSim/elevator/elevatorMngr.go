package elevator

import (
	"context"
	"fmt"
	"log"

	"elevatorSim/api"
	docker "elevatorSim/dockerRun"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	STARTING_PORT = 7000
)

var (
	ORDINAL_EXP = [3]string{"st", "nd", "th"}
)

type ElevatorMngr struct {
	nr_elevator        int
	dockerRun          *docker.DockerRun
	elevatorContainers map[string]*ElevatorContainer
	//eventChan EventWatch
}

func NewElevatorMngr(nr_elevator int) *ElevatorMngr {
	dockerRun := docker.NewDockerRun()
	dockerRun.EnsureImageExists()

	return &ElevatorMngr{
		nr_elevator:        nr_elevator,
		dockerRun:          dockerRun,
		elevatorContainers: make(map[string]*ElevatorContainer),
		//eventChan: NewEventWatch(),
	}
}

func (em *ElevatorMngr) AddElevators() error {

	for i := 0; i < em.nr_elevator; i++ {
		name := fmt.Sprintf("%d%sElevator", i+1, ORDINAL_EXP[map[bool]int{true: 2, false: i}[i > 2]])
		port := fmt.Sprintf("%d", STARTING_PORT+i)

		if err := em.AddElevator(name, port); err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	return nil
}

func (em *ElevatorMngr) AddElevator(name string, port string) error {

	rsp, err := em.dockerRun.CreateContainer(name, port)
	if err != nil || rsp == nil {
		return fmt.Errorf("failed to create container: %v", err)
	}
	glog.V(2).Infof("Start new elevator container %s", rsp.ID)

	if err = em.dockerRun.StartContainer(rsp.ID); err != nil {
		return fmt.Errorf("failed to start sandbox container: %v", err)
	}

	elevatorCon := NewElevatorContainer(name, port, rsp.ID)
	em.elevatorContainers[name] = elevatorCon

	return nil
}

func (em *ElevatorMngr) DeleteElevators() {
	for i := 0; i < em.nr_elevator; i++ {
		name := fmt.Sprintf("%d%sElevator", i, ORDINAL_EXP[map[bool]int{true: 2, false: i}[i > 2]])

		if err := em.DeleteElevator(em.elevatorContainers[name]); err != nil {
			glog.Fatalf("Error while delete container %s: %v", name, err)
		}
	}
}

func (em *ElevatorMngr) DeleteElevator(elevator *ElevatorContainer) error {
	if err := em.dockerRun.StopContainer(elevator.id); err != nil {
		return err
	}
	if err := em.dockerRun.RemoveContainer(elevator.id); err != nil {
		return err
	}
	if err := elevator.RemoveConnection(); err != nil {
		return err
	}
	return nil
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

// TODO:
// 1. get Elevator IP using elevator container name
// 2. dial server with IP:PORT
// 3. add tls connection
func (em *ElevatorMngr) GetElevatorsStatus() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// elevator name, service client
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

// client log
//https://stackoverflow.com/questions/50465273/go-docker-client-get-container-logs-every-seconds-returns-nothing
