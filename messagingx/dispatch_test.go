package messagingx

import "fmt"

func ExampleTraceDispatch_Channel() {
	d := NewTraceDispatcher(nil, EmissaryChannel)
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event, "")
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	channel = EmissaryChannel
	d.Dispatch(nil, channel, event, "")
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	//Output:
	//test: Dispatch() -> [channel:]
	//trace -> 2024-11-24T18:40:08.606Z [emissary] [] [<nil>]
	//test: Dispatch() -> [channel:emissary] []

}

func ExampleTraceDispatch_Event() {
	d := NewTraceDispatcher([]string{ShutdownEvent, StartupEvent}, "")
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event, "")
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = ShutdownEvent
	d.Dispatch(nil, channel, event, "")
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = ObservationEvent
	d.Dispatch(nil, channel, event, "")
	fmt.Printf("test: Dispatch() -> [channel:%v] [%v]\n", channel, event)

	//Output:
	//test: Dispatch() -> []
	//trace -> 2024-11-24T18:46:04.697Z [] [event:shutdown] [<nil>]
	//test: Dispatch() -> [event:shutdown]
	//test: Dispatch() -> [channel:] [event:observation]

}
