package faas

import "sync"

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

	//TODO check function type
	// func(ctx *context.Context, request interface{}) (response interface{}, err error)

	serviceMap[serviceName] = function
}

func GetServiceMap() map[string]interface{} {
	lock.Lock()
	defer lock.Unlock()

	result := make(map[string]interface{})
	for k, v := range serviceMap {
		result[k] = v
	}
}
