package main

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	
	//"elevatorSim/clock"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	
	switch m.Name {
	case "start":
		// Unmarshal payload
		var passedTime string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err := json.Unmarshal(m.Payload, &passedTime); err != nil {
				payload := err.Error()
				return payload, err
			}
		}
	}
	return nil, nil
}