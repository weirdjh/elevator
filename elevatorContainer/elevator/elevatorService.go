package elevator

import (
	"context"
	"elevatorContainer/api"
	"fmt"
	"time"
)

type ElevatorState string

const (
	Idle ElevatorState = "IDLE"
	Stop ElevatorState = "STOP"
	Move ElevatorState = "MOVE"
)

type ElevatorService struct {
	name string
	//people    []Person
	state     ElevatorState
	direction string
	floor     int32

	//eventChan EventWatch
	//spec      map[string]string
}

func NewElevatorService(name string) *ElevatorService {
	return &ElevatorService{
		name:      name,
		state:     Idle,
		direction: "",
		floor:     0,
	}
}

func (e *ElevatorService) GetElevatorStatus(cxt context.Context, req *api.GetElevatorStatusRequest) (*api.GetElevatorStatusResponse, error) {
	resp := &api.GetElevatorStatusResponse{
		Name:      e.name,
		Direction: e.direction,
		Floor:     e.floor,
	}
	return resp, nil
}

func (e *ElevatorService) ElevatorUp(cxt context.Context, req *api.ElevatorUpRequest) (*api.ElevatorUpResponse, error) {
	e.state = Move
	resp := &api.ElevatorUpResponse{}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor += 1
		fmt.Printf("Currnet floor : %d\n", e.floor)
	}
	return resp, nil
}

func (e *ElevatorService) ElevatorDown(cxt context.Context, req *api.ElevatorDownRequest) (*api.ElevatorDownResponse, error) {
	e.state = Move
	resp := &api.ElevatorDownResponse{}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor -= 1
		fmt.Printf("Currnet floor : %d\n", e.floor)
	}
	return resp, nil
}

func (e *ElevatorService) ElevatorStop(cxt context.Context, req *api.ElevatorStopRequest) (*api.ElevatorStopResponse, error) {
	e.state = Stop
	resp := &api.ElevatorStopResponse{}

	time.Sleep(time.Second * 2)
	return resp, nil
}
