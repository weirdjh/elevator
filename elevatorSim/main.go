package main

import (
	docker "elevatorSim/dockerRun"
	util "elevatorSim/elevator"
)

func main() {
	dockerrun := docker.NewDockerRun()
	dockerrun.EnsureImageExists()

	elevMngr := util.NewElevatorMngr()

	elevMngr.GetElevatorsStatus()
	/*
		elevMngr.AddElevator("first")
		elevMngr.AddElevator("second")

		elevMngr.PingElevator()


		go elevMngr.MngElevatorUp("first", 10)

		time.Sleep(time.Second * 3)

		go elevMngr.MngElevatorUp("second", 10)

		elevMngr.PrintElevator()

		time.Sleep(time.Second * 3)

		elevMngr.PrintElevator()
	*/
}

/*
func main() {

	elevator := e.NewElevator("first")
	elevator.PrintElevator()

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	done <- true
	elapsed := time.Since(start)
	fmt.Println("elapsed : %s", elapsed)
}

func run2(done chan bool) {
	start := time.Now()
	time.Sleep(2 * time.Second)
	done <- true
	elapsed := time.Since(start)
	fmt.Println("elapsed : %s", elapsed)
}
*/
