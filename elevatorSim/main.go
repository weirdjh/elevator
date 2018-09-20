package main

import (
	docker "elevatorSim/dockerRun"
	mng "elevatorSim/elevator"
	"elevatorSim/timer"
	"elevatorSim/traffic"
	"fmt"
	"time"
)

const (
	NR_ELEVATOR  = 2
	TESTDURATION = 10
)

func main() {

	/* Set Timer for this Program */
	GlobalTimer := timer.NewTimer()

	/* Run Docker Container and Put Elevator image on Container */
	dockerrun := docker.NewDockerRun()
	dockerrun.EnsureImageExists()

	elevMngr := mng.NewElevatorMngr(NR_ELEVATOR, GlobalTimer)
	if err := elevMngr.AddElevators(); err != nil {
		fmt.Printf("%v", err)
	}
	defer elevMngr.DeleteElevators()

	/* Test Traffic Queue */
	traffics := traffic.ElevatorTraffic()
	for _, t := range traffics {
		fmt.Println(t)
	}

	/* Run Test */
	elapsedTime := time.Second * 0
	var testDuration time.Duration = time.Second * TESTDURATION

	for elapsedTime < testDuration {

		curTime := GlobalTimer.GetTime()

		if len(traffics) == 0 {
			fmt.Println("No More Traffic")
			break
		}

		traffic := traffics[0].Move.At
		if curTime == traffic {
			go elevMngr.MngElevator(*traffics[0])

			traffics = traffics[1:]
			elapsedTime = curTime.Sub(GlobalTimer.ProgramInitTime)
		}
	}

	fmt.Println(elapsedTime)
	/*
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
	*/
	time.Sleep(time.Second * 10)
}
