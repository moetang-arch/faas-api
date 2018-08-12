# faas-api
faas api

# 1. mode 1

just functions to run

* create a package and service
* import "github.com/moetang-arch/faas-api" for initialization
* ONLY the packages in the allowed list can be used

# 2. mode 2

Not Yet Implement

# 3. Example - mode 1

```
package demo

import (
	"context"
	"errors"

	"github.com/moetang-arch/faas-api"
)

func init() {
	faas.SetGlobalServicePrefix("set.if.need")
	faas.Register("demo", HandleRequest)
}

type Request struct {
	Name string
}

type Response struct {
	Result string
}

func HandleRequest(ctx context.Context, request *Request) (response *Response, err error) {
	if len(request.Name) == 0 {
		return nil, errors.New("name is empty")
	}
	resp := new(Response)
	resp.Result = "Hello, " + request.Name
	return resp, nil
}
```
