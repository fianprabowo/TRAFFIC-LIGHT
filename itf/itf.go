package itf

type State interface {
	Change() error
	GetColor() string
}
