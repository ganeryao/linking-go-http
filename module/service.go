package module

import (
	"reflect"
)

var (
	handlers = make(map[string]*Service, 0)
)

type (
	Service struct {
		Name    string    // name of service
		Handler Component // Handler of the Component
		Options options   // options
	}
)

func GetService(name string) (*Service, bool) {
	handler, ok := handlers[name]
	return handler, ok
}

func NewService(comp Component, opts []Option) *Service {
	s := &Service{
		Handler: comp,
	}
	// apply options
	for i := range opts {
		opt := opts[i]
		opt(&s.Options)
	}
	if name := s.Options.name; name != "" {
		s.Name = name
	} else {
		s.Name = reflect.Indirect(reflect.ValueOf(comp)).Type().Name()
	}
	return s
}

func Register(comp Component, options ...Option) {
	s := NewService(comp, options)
	handlers[s.Name] = s
}
