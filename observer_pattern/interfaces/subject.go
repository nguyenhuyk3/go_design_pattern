package interfaces

type Subject interface {
	Register(o IObserver)
	Deregister(o IObserver)
	Notify(message string)
}
