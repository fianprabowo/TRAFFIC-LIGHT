package machine

import (
	"fmt"
	"time"
)

type YellowState struct {
	machine *Machine
}

func NewYellowState(m *Machine) *YellowState {
	return &YellowState{
		machine: m,
	}
}

func (y *YellowState) Change() error {
	y.machine.SetState(y.machine.RedLight)
	return nil
}

func (y *YellowState) GetColor() string {
	return "🟡 Yellow"
}

func (y *YellowState) Wait() error {
	fmt.Println("Waiting for 2 seconds...")
	time.Sleep(2 * time.Second)
	return nil
}
