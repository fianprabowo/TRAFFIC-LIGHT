package machine

import (
	"fmt"
	"time"
)

type GreenState struct {
	machine *Machine
}

func NewGreenState(m *Machine) *GreenState {
	return &GreenState{
		machine: m,
	}
}

func (g *GreenState) Change() error {
	g.machine.SetState(g.machine.YellowLight)
	return nil
}

func (g *GreenState) GetColor() string {
	return "🟢 Green"
}

func (g *GreenState) Wait() error {
	fmt.Println("Waiting for 5 seconds...")
	time.Sleep(5 * time.Second)
	return nil
}
