package pucore

import "fmt"

const (
	MODULAR_HOOK_ENABLE_BEFORE = iota + 1
	MODULAR_HOOK_ENABLE_AFTER
	MODULAR_HOOK_DISABLE_BEFORE
	MODULAR_HOOK_DISABLE_AFTER
)

var (
	ErrModuleRegisterTwice = func(mod IModule) error {
		return fmt.Errorf("Module '%s' register twice", mod.Id())
	}
	ErrModuleNotFound = func(mod IModule) error {
		return fmt.Errorf("Module '%s' is not found", mod.Id())
	}
)

type ModuleContext struct {
	Injector *Injector
	Behavior *Behaviors
	Modular  *Modular
}

type IModule interface {
	Id() string
	Prepare(*ModuleContext) error
	Enable(*ModuleContext) error
	Disable(*ModuleContext) error
}

type ModularHook func(m IModule) error

type Modular struct {
	modules  map[string]IModule
	inj      *Injector
	behavior *Behaviors
	hooks    map[int][]ModularHook
}

func NewModular(inj *Injector, be *Behaviors) *Modular {
	return &Modular{
		inj:      inj,
		behavior: be,
		modules:  make(map[string]IModule),
		hooks:    make(map[int][]ModularHook),
	}
}

func (m *Modular) Register(modules ...IModule) error {
	ctx := &ModuleContext{
		Injector: m.inj,
		Behavior: m.behavior,
		Modular:  m,
	}
	for _, mod := range modules {
		if m.Has(mod) {
			panic(ErrModuleRegisterTwice(mod))
		}
		if err := mod.Prepare(ctx); err != nil {
			panic(err)
		}
		m.modules[mod.Id()] = mod
	}
	return nil
}

func (m *Modular) Has(mod IModule) bool {
	return m.modules[mod.Id()] != nil
}

func (m *Modular) Enable(mod IModule) error {
	if !m.Has(mod) {
		return ErrModuleNotFound(mod)
	}
	var err error
	if err = m.emit(MODULAR_HOOK_ENABLE_BEFORE, mod); err != nil {
		return err
	}
	ctx := &ModuleContext{
		Injector: m.inj,
		Behavior: m.behavior,
		Modular:  m,
	}
	if err = m.modules[mod.Id()].Enable(ctx); err != nil {
		return nil
	}
	if err = m.emit(MODULAR_HOOK_ENABLE_AFTER, mod); err != nil {
		return err
	}
	return nil
}

func (m *Modular) Disable(mod IModule) error {
	if !m.Has(mod) {
		return ErrModuleNotFound(mod)
	}
	var err error
	if err = m.emit(MODULAR_HOOK_DISABLE_BEFORE, mod); err != nil {
		return err
	}
	ctx := &ModuleContext{
		Injector: m.inj,
		Behavior: m.behavior,
		Modular:  m,
	}
	if err = m.modules[mod.Id()].Disable(ctx); err != nil {
		return nil
	}
	if err = m.emit(MODULAR_HOOK_DISABLE_AFTER, mod); err != nil {
		return err
	}
	return nil
}

func (m *Modular) EnableAll() error {
	var err error
	for _, mod := range m.modules {
		if err = m.Enable(mod); err != nil {
			return err
		}
	}
	return nil
}

func (m *Modular) DisableAll() error {
	var err error
	for _, mod := range m.modules {
		if err = m.Disable(mod); err != nil {
			return err
		}
	}
	return nil
}

func (m *Modular) Hook(pos int, fn ModularHook) {
	m.hooks[pos] = append(m.hooks[pos], fn)
}

func (m *Modular) emit(pos int, mod IModule) error {
	hooks := m.hooks[pos]
	if len(hooks) > 0 {
		var err error
		for _, hook := range hooks {
			if err = hook(mod); err != nil {
				return err
			}
		}
	}
	return nil
}
