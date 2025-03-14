package messaging

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

func ExampleNewStatus_OK() {
	s := StatusOK()

	path := reflect.TypeOf(Status{}).PkgPath()
	path += "/" + reflect.TypeOf(Status{}).Name()
	fmt.Printf("test: NewStatus() -> [status:%v] [type:%v]\n", s, path)

	//Output:
	//test: NewStatus() -> [status:OK] [type:github.com/behavioral-ai/core/messaging/Status]

}

func ExampleNewStatus_Teapot() {
	s := NewStatus(http.StatusTeapot)
	fmt.Printf("test: NewStatus() -> [status:%v]\n", s)

	//Output:
	//test: NewStatus() -> [status:I'm A Teapot]

}

func ExampleNewStatusError() {
	s := NewStatusError(http.StatusGatewayTimeout, errors.New("rate limited"), "test:agent") //"resiliency:agent/operative/agent1#us-west")
	fmt.Printf("test: NewStatusError() -> [%v]\n", s)

	if _, ok := any(s).(Event); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatusError() -> [Timeout [err:rate limited] [agent:test:agent]]
	//test: Event() -> [true]

}

func ExampleNewStatusMessage() {
	s := NewStatusMessage(http.StatusOK, "successfully change ticker duration", "test:agent")
	fmt.Printf("test: NewStatusMessage() -> [%v]\n", s)

	if _, ok := any(s).(Event); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatusMessage() -> [OK [msg:successfully change ticker duration] [agent:test:agent]]
	//test: Event() -> [true]

}
