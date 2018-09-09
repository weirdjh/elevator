package main

import (
	docker "elevatorSim/dockerRun"
	mng "elevatorSim/elevator"
	"elevatorSim/traffic"
	"fmt"
	"time"
)

const (
	nrElevator = 2
)

func main() {

	t := traffic.ElevatorTraffic()
	fmt.Println(t[0])

	dockerrun := docker.NewDockerRun()
	dockerrun.EnsureImageExists()

	elevMngr := mng.NewElevatorMngr(nrElevator)
	if err := elevMngr.AddElevators(); err != nil {
		fmt.Printf("%v", err)
	}
	defer elevMngr.DeleteElevators()

	// docker attach 1stElevator
	elevator1 := elevMngr.GetElevator("1stElevator")
	go elevator1.ElevatorUp(10)

	time.Sleep(time.Millisecond * 2500)
	elevator2 := elevMngr.GetElevator("2ndElevator")
	go elevator2.ElevatorUp(6)

	time.Sleep(time.Second * 3)
	status, _ := elevator1.GetElevatorStatus()
	fmt.Println(status)
	status, _ = elevator2.GetElevatorStatus()
	fmt.Println(status)

	time.Sleep(time.Second * 3)
	status, _ = elevator1.GetElevatorStatus()
	fmt.Println(status)
	status, _ = elevator2.GetElevatorStatus()
	fmt.Println(status)

	time.Sleep(time.Second * 10)
}
