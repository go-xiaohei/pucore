package pucore

type Core struct {
	*Injector
	*Modular
}

func NewCore() *Core {
	inj := NewInjector()
	modular := NewModular(inj)
	return &Core{
		Injector: inj,
		Modular:  modular,
	}
}
