package pucore

import "fmt"

var (
	ErrScopeBehaviorNullHandlers = func(name string) error {
		return fmt.Errorf("Behavior '%' has no handlers in scope", name)
	}
)

type Scope struct {
	inj       *Injector
	behaviors *Behaviors

	error error // save error on working chain. if error, stop chain
}

func NewScope(inj *Injector, b *Behaviors) *Scope {
	return &Scope{
		inj:       inj,
		behaviors: b,
	}
}

func (s *Scope) Injector() *Injector {
	return s.inj
}

type scopeContext struct {
	name     string
	handlers []BehaviorHandler
	index    int
	length   int
	inj      *Injector
}

func (sc *scopeContext) Next() error {
	if sc.index >= sc.length {
		return nil
	}
	var err error
	handler := sc.handlers[sc.index]
	if handler != nil {
		if err = handler(sc.inj); err != nil {
			return err
		}
	}
	sc.index++
	return sc.Next()
}

func (s *Scope) Do(b BehaviorHandler, inputs ...interface{}) *Scope {
	if s.error != nil {
		return s
	}
	handlers := s.behaviors.Assemble(b)
	if len(handlers) == 0 {
		s.error = ErrScopeBehaviorNullHandlers(funcName(b))
		return s
	}
	ctx := &scopeContext{
		name:     funcName(b),
		handlers: handlers,
		index:    0,
		length:   len(handlers),
		inj:      s.inj,
	}
	if err := ctx.Next(); err != nil {
		s.error = err
		return s
	}
	return s
}

func (s *Scope) Error() error {
	return s.error
}
