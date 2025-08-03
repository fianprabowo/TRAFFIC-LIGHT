package machine

import (
	"fmt"
	"time"
)

type RedState struct {
	machine *Machine
}

func NewRedState(m *Machine) *RedState {
	return &RedState{
		machine: m,
	}
}

func (r *RedState) Change() error {
	r.machine.SetState(r.machine.GreenLight)
	return nil
}

func (r *RedState) GetColor() string {
	return "🔴 Red"
}

func (r *RedState) Wait() error {
	fmt.Println("Waiting for 7 seconds...")
	time.Sleep(7 * time.Second)
	return nil
}
