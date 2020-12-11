package module

import (
	"reflect"
)

var (
	handlers = make(map[string]*Service, 0)
)

type (
	Service struct {
		Name     string        // name of service
		Type     reflect.Type  // type of the receiver
		Receiver reflect.Value // receiver of methods for the service
		Options  options       // options
	}
)

func ContainsHandler(name string) bool {
	_, ok := handlers[name]
	return ok
}

func GetHandler(name string) (*Service, bool) {
	handler, ok := handlers[name]
	return handler, ok
}

func NewService(comp Component, opts []Option) *Service {
	s := &Service{
		Type:     reflect.TypeOf(comp),
		Receiver: reflect.ValueOf(comp),
	}
	// apply options
	for i := range opts {
		opt := opts[i]
		opt(&s.Options)
	}
	if name := s.Options.name; name != "" {
		s.Name = name
	} else {
		s.Name = reflect.Indirect(s.Receiver).Type().Name()
	}
	return s
}

func Register(comp Component, options ...Option) {
	s := NewService(comp, options)
	handlers[s.Name] = s
}
