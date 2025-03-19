package implementations

import "design_pattern/observer_pattern/interfaces"

type ConcreteSubject struct {
	observers []interfaces.IObserver
}

func (s *ConcreteSubject) Register(o interfaces.IObserver) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Deregister(o interfaces.IObserver) {
	for i, obs := range s.observers {
		if obs == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(message string) {
	for _, obs := range s.observers {
		obs.Update(message)
	}
}
