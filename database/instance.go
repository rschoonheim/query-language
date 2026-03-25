package database

import (
	"context"
	"sync"
)

const (
	StateStarting = iota
	StateRunning
	StateTerminating
	StateTerminated
	stateContextError
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

// hasState - checks if the database instance has the given state
func (instance *Instance) hasState(state int) bool {
	instance.stateMutex.RLock()
	defer instance.stateMutex.RUnlock()

	return instance.state == state
}

// Run - runs the database instance
func (instance *Instance) Run(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	defer wg.Done()

	for !instance.hasState(StateTerminated) {
		instance.handleContextError(ctx)

		switch instance.getState() {
		default:
			panic("unhandled default case")
		case StateStarting:
			instance.handleStateStarting()
		case StateRunning:
			instance.handleStateRunning()
		case StateTerminating:
			instance.handleStateTerminating()
		}
	}

	return nil
}

// handleStateStarting - handles the starting state of the database instance
func (instance *Instance) handleStateStarting() {
	println("State is starting")
	instance.setState(StateRunning)
}

// handleStateRunning - handles the running state of the database instance
func (instance *Instance) handleStateRunning() {
	println("State is running")
}

// handleStateTerminating - handles the terminating state of the database instance
func (instance *Instance) handleStateTerminating() {
	println("State is terminating")
	instance.setState(StateTerminated)
}

// handleContextError - transitions the instance state when the context is done
func (instance *Instance) handleContextError(ctx context.Context) {
	if ctx.Err() == nil {
		return
	}

	switch ctx.Err() {
	case context.Canceled, context.DeadlineExceeded:
		instance.setState(StateTerminating)
	default:
		instance.setState(stateContextError)
	}
}
