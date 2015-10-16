package pucore

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrInjectorSetNeedPointer    = errors.New("injector needs set a pointer")
	ErrInjectorSetUnknownPointer = func(rv reflect.Value) error {
		return fmt.Errorf("injector can't assign to %s", rv.Type().String())
	}
	ErrInjectorSetTwicePointer = func(rv reflect.Value) error {
		return fmt.Errorf("injector can't set %s twice", rv.Type().String())
	}
)

type Injector struct {
	data map[string]reflect.Value
}

func NewInjector(values ...interface{}) *Injector {
	inj := &Injector{
		data: make(map[string]reflect.Value),
	}
	inj.Set(values...)
	return inj
}

func (inj *Injector) Set(values ...interface{}) {
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			panic(ErrInjectorSetNeedPointer)
		}
		inj.data[rv.Type().String()] = rv
	}
}

func (inj *Injector) Get(values ...interface{}) {
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			panic(ErrInjectorSetNeedPointer)
		}
		resRv, ok := inj.data[rv.Type().String()]
		if !ok {
			panic(ErrInjectorSetUnknownPointer(rv))
		}
		rv.Elem().Set(resRv.Elem())
	}
}

func (inj *Injector) Has(value interface{}) bool {
	rtStr := reflect.TypeOf(value).String()
	_, ok := inj.data[rtStr]
	return ok
}

func (inj *Injector) Merge(i *Injector, override bool) error {
	for name, v := range i.data {
		if _, ok := inj.data[name]; ok && !override {
			return ErrInjectorSetTwicePointer(v)
		}
		inj.data[name] = v
	}
	return nil
}
