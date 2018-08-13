package faas

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var _namespace string

var (
	serviceMap = make(map[string]interface{})
	lock       = new(sync.Mutex)
)

// default is configured via control panel
func SetGlobalServiceNameSpace(namespace string) {
	_namespace = namespace
}

func GetGlobalServiceNameSpace() string {
	return _namespace
}

func Register(serviceName string, function interface{}) {
	lock.Lock()
	defer lock.Unlock()

	// check function type
	// func(ctx *context.Context, request interface{}) (response interface{}, err error)
	checkFunction(function)

	serviceMap[serviceName] = function
}

func checkFunction(function interface{}) {
	t := reflect.TypeOf(function)
	fmt.Println(t)
	if t.Kind() != reflect.Func {
		panic(errors.New("type of object is not function"))
	}
	if t.NumIn() != 2 {
		panic(errors.New("params length should be 2"))
	}
	if t.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		panic(errors.New("first param should be context.Context"))
	}
	tp2 := t.In(1)
	if !(tp2.Kind() == reflect.Ptr && tp2.Elem().Kind() == reflect.Struct) && !(tp2.Kind() == reflect.Struct) {
		panic(errors.New("second param should be *struct type or struct type"))
	}
	if t.NumOut() != 2 {
		panic(errors.New("returns length should be 2"))
	}
	tr1 := t.Out(0)
	if !(tr1.Kind() == reflect.Ptr && tr1.Elem().Kind() == reflect.Struct) && !(tr1.Kind() == reflect.Struct) {
		panic(errors.New("first return should be *struct type or struct type"))
	}
	if t.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		panic(errors.New("second return should be error"))
	}
}

func GetServiceMap() map[string]interface{} {
	lock.Lock()
	defer lock.Unlock()

	result := make(map[string]interface{})
	for k, v := range serviceMap {
		result[k] = v
	}

	return result
}
