//https://github.com/kubernetes/kubernetes/search?q=RuntimeServiceServer&unscoped_q=RuntimeServiceServer

package main

import "elevatorContainer/elevator"

// main start a gRPC server and waits for connection

// TODO: Give Elevator Server Name Dynamically
func main() {
	es := elevator.NewElevatorServer()
	es.Start()

}
