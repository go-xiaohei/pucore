package pucore

import (
	"reflect"
	"runtime"
	"strings"
)

// get function name
func funcName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	nameData := strings.Split(name, "/")
	if len(nameData) > 2 {
		nameData = nameData[len(nameData)-1:]
	}
	name = strings.TrimSuffix(strings.Join(nameData, "."), "-fm")
	return strings.TrimSuffix(name, "·fm")
}

type BehaviorHandler func(*Injector) error

type Behaviors struct {
	behaviors  map[string]BehaviorHandler
	middleware map[string][]BehaviorHandler
}

func NewBehaviors(bs ...BehaviorHandler) *Behaviors {
	b := &Behaviors{
		behaviors:  make(map[string]BehaviorHandler),
		middleware: make(map[string][]BehaviorHandler),
	}
	b.Add(bs...)
	return b
}

func (b *Behaviors) Add(bs ...BehaviorHandler) {
	for _, be := range bs {
		name := funcName(be)
		b.behaviors[name] = be
	}
}

func (b *Behaviors) Remove(bs ...BehaviorHandler) {
	for _, be := range bs {
		name := funcName(be)
		delete(b.behaviors, name)
	}
}

func (b *Behaviors) Listen(be BehaviorHandler, hooks ...BehaviorHandler) {
	name := funcName(be)
	b.middleware[name] = append(b.middleware[name], hooks...)
}

func (b *Behaviors) Assemble(be BehaviorHandler) []BehaviorHandler {
	name := funcName(be)
	handlers := make([]BehaviorHandler, 0)
	if len(b.middleware[name]) > 0 {
		handlers = b.middleware[name]
	}
	return append(handlers, b.behaviors[name])
}
