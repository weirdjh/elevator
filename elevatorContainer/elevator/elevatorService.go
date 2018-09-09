package elevator

import (
	"context"
	"elevatorContainer/api"
	"fmt"
	"time"
)

const (
	Idle     = "IDLE"
	Stop     = "STOP"
	MoveUp   = "MOVEUP"
	MoveDown = "MOVEDOWN"
)

type ElevatorService struct {
	name string
	//people    []Person
	state string
	floor int32

	//eventChan EventWatch
	//spec      map[string]string
}

func NewElevatorService(name string) *ElevatorService {
	return &ElevatorService{
		name:  name,
		state: Idle,
		floor: 0,
	}
}

func (e *ElevatorService) GetElevatorStatus(cxt context.Context, req *api.GetElevatorStatusRequest) (*api.GetElevatorStatusResponse, error) {
	resp := &api.GetElevatorStatusResponse{
		Name:  e.name,
		State: e.state,
		Floor: e.floor,
	}
	return resp, nil
}

func (e *ElevatorService) ElevatorUp(cxt context.Context, req *api.ElevatorUpRequest) (*api.ElevatorUpResponse, error) {
	e.state = MoveUp
	resp := &api.ElevatorUpResponse{Done: true}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor += 1
		fmt.Printf("Currnet floor : %d\n", e.floor)
	}
	e.state = Idle
	return resp, nil
}

func (e *ElevatorService) ElevatorDown(cxt context.Context, req *api.ElevatorDownRequest) (*api.ElevatorDownResponse, error) {
	e.state = MoveDown
	resp := &api.ElevatorDownResponse{Done: true}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor -= 1
		fmt.Printf("Currnet floor : %d\n", e.floor)
	}
	e.state = Idle
	return resp, nil
}

func (e *ElevatorService) ElevatorStop(cxt context.Context, req *api.ElevatorStopRequest) (*api.ElevatorStopResponse, error) {
	e.state = Stop
	resp := &api.ElevatorStopResponse{Done: true}

	time.Sleep(time.Second * 2)
	e.state = Idle
	return resp, nil
}
