package machine

import (
	"github.com/fianprabowo/TRAFFIC-LIGHT/internal"
)

type Machine struct {
	GreenLight   internal.State
	YellowLight  internal.State
	RedLight     internal.State
	currentState internal.State
}

func NewMachine() *Machine {
	m := &Machine{}
	m.GreenLight = NewGreenState(m)
	m.YellowLight = NewYellowState(m)
	m.RedLight = NewRedState(m)
	m.SetState(m.RedLight) // Start with Red light
	return m
}

func (m *Machine) SetState(state internal.State) {
	m.currentState = state
}

func (m *Machine) Change() error {
	if m.currentState == nil {
		return nil
	}
	m.currentState.Wait()
	return m.currentState.Change()
}

func (m *Machine) GetColor() string {
	if m.currentState == nil {
		return ""
	}
	return m.currentState.GetColor()
}
