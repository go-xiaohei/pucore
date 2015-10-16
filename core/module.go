package pucore

import (
	"errors"
	"fmt"
)

var (
	ErrModularRegisterTwice = errors.New("modular-register-twice")
	ErrModularNotFound      = func(name string) error {
		return fmt.Errorf("modular '%s' not found", name)
	}
)

type Modular struct {
	modules  map[string]IModule
	injector *Injector
}

func NewModular(inj *Injector, ms ...IModule) *Modular {
	modular := &Modular{
		injector: inj,
		modules:  make(map[string]IModule),
	}
	modular.Register(ms...)
	return modular
}

func (mod *Modular) SetInjector(inj *Injector) {
	mod.injector = inj
}

func (mod *Modular) Register(ms ...IModule) {
	var (
		err error
	)
	for _, m := range ms {
		if mod.HasModular(m) {
			panic(ErrModularRegisterTwice)
		}
		if err = m.Depends(mod); err != nil {
			panic(err)
		}
		if err = m.Bootstrap(mod.injector); err != nil {
			panic(err)
		}
		mod.modules[m.Name()] = m
	}
}

func (mod *Modular) has(name string) bool {
	_, ok := mod.modules[name]
	return ok
}

func (mod *Modular) HasModular(m interface{}) bool {
	if name, ok := m.(string); ok {
		return mod.has(name)
	}
	if module, ok := m.(IModule); ok {
		return mod.has(module.Name())
	}
	return false
}

func (mod *Modular) Merge(m *Modular) {
	mod.injector.Merge(m.injector, true)
	for name, v := range m.modules {
		mod.modules[name] = v
	}
}

func (mod *Modular) Enable(name ...string) error {
	var err error
	if len(name) == 0 {
		for _, m := range mod.modules {
			if err = m.Enable(); err != nil {
				return err
			}
		}
	}
	for _, n := range name {
		if m, ok := mod.modules[n]; ok {
			if err = m.Enable(); err != nil {
				return err
			}
		} else {
			return ErrModularNotFound(n)
		}
	}
	return nil
}

func (mod *Modular) Disable(name ...string) error {
	var err error
	if len(name) == 0 {
		for _, m := range mod.modules {
			if err = m.Disable(); err != nil {
				return err
			}
		}
	}
	for _, n := range name {
		if m, ok := mod.modules[n]; ok {
			if err = m.Disable(); err != nil {
				return err
			}
		} else {
			return ErrModularNotFound(n)
		}
	}
	return nil
}

type IModule interface {
	Name() string
	Bootstrap(inj *Injector) error
	Depends(m *Modular) error

	Enable() error
	Disable() error
}
