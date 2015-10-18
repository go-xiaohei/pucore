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
	ErrInjectorValueNotFound = func(valueName string) error {
		return fmt.Errorf("injector can't find '%s'", valueName)
	}
)

// Injector provides global variables chain as context role.
// The values in injector is unique with its type.
// So injector can't save two values with same type.
type Injector struct {
	data map[string]reflect.Value
}

// NewInjector creates new injector instance with given values.
func NewInjector(values ...interface{}) *Injector {
	inj := &Injector{
		data: make(map[string]reflect.Value),
	}
	inj.Set(values...)
	return inj
}

// Set sets values into injector.
// It needs a pointer.
func (inj *Injector) Set(values ...interface{}) {
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			panic(ErrInjectorSetNeedPointer)
		}
		inj.data[rv.Type().String()] = rv
	}
}

// Get gets values from injector.
// It needs a pointer.
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

// Has checks the values are assigned in injector.
// If a value is not in, return error.
func (inj *Injector) Has(values ...interface{}) error {
	for _, value := range values {
		rtStr := reflect.TypeOf(value).String()
		if _, ok := inj.data[rtStr]; !ok {
			return ErrInjectorValueNotFound(rtStr)
		}
	}
	return nil
}

// Merge merges a injector with another one.
// If override, old value is overwritten.
// Or return error.
func (inj *Injector) Merge(i *Injector, override bool) error {
	for name, v := range i.data {
		if _, ok := inj.data[name]; ok && !override {
			return ErrInjectorSetTwicePointer(v)
		}
		inj.data[name] = v
	}
	return nil
}
