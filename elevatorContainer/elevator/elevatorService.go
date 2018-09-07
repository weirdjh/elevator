package elevator

import (
	"context"
	"elevatorContainer/api"
	"time"
)

type ElevatorState string

const (
	Idle ElevatorState = "IDLE"
	Stop ElevatorState = "STOP"
	Move ElevatorState = "MOVE"
)

type Elevator struct {
	name string
	//people    []Person
	state     ElevatorState
	direction string
	floor     int32

	//eventChan EventWatch
	//spec      map[string]string
}

func NewElevator(name string) *Elevator {
	return &Elevator{
		name:      name,
		state:     Idle,
		direction: "",
		floor:     0,
	}
}

func (e *Elevator) GetElevatorStatus(cxt context.Context, req *api.GetElevatorStatusRequest) (*api.GetElevatorStatusResponse, error) {
	resp := &api.GetElevatorStatusResponse{
		Name:      e.name,
		Direction: e.direction,
		Floor:     e.floor,
	}
	return resp, nil
}

func (e *Elevator) ElevatorUp(cxt context.Context, req *api.ElevatorUpRequest) (*api.ElevatorUpResponse, error) {
	e.state = Move
	resp := &api.ElevatorUpResponse{}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor += 1
	}
	return resp, nil
}

func (e *Elevator) ElevatorDown(cxt context.Context, req *api.ElevatorDownRequest) (*api.ElevatorDownResponse, error) {
	e.state = Move
	resp := &api.ElevatorDownResponse{}

	for e.floor < req.GetDestination() {
		time.Sleep(time.Second * 1)
		e.floor -= 1
	}
	return resp, nil
}

func (e *Elevator) ElevatorStop(cxt context.Context, req *api.ElevatorStopRequest) (*api.ElevatorStopResponse, error) {
	e.state = Stop
	resp := &api.ElevatorStopResponse{}

	time.Sleep(time.Second * 2)
	return resp, nil
}
