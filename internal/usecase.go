package internal

type State interface {
	Change() error
	Wait() error
	GetColor() string
}
