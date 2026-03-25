package database

import (
	"context"
	"fmt"
	"sync"
)

const (
	StateStarting = iota
	StateRunning
	StateTerminating
	StateTerminated
)

type Instance struct {
	configuration *Configuration
	state         int
	stateMutex    sync.RWMutex
}

// setState - sets the state of the database instance
func (instance *Instance) setState(state int) {
	instance.stateMutex.Lock()
	defer instance.stateMutex.Unlock()

	instance.state = state
}

// getState - gets the state of the database instance
func (instance *Instance) getState() int {
	instance.stateMutex.RLock()
	defer instance.stateMutex.RUnlock()

	return instance.state
}

// Run - runs the database instance
func (instance *Instance) Run(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	defer wg.Done()

	for instance.getState() != StateTerminated {
		if ctx.Err() != nil {
			instance.setState(StateTerminating)
		}

		switch instance.getState() {
		case StateTerminating:
			instance.setState(StateTerminated)
		case StateStarting:
			instance.setState(StateRunning)
			fmt.Println("State set to running")
		case StateRunning:
			println("State is running")
		default:
			panic("unhandled default case")
		}
	}

	println("Terminated")

	return nil
}
