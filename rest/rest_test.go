package rest

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleBuildExchangeChain() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildExchangeChain([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChainExchange_Link() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChainExchange_Chainable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Exchange{}, do2Exchange{}, do3Exchange{}})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChainExchange_Any() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Exchange{}, do2ExchangeFn, do3Exchange{}, do4ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do4-Exchange() -> request
	//test: Do4-Exchange() -> response
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChainExchange_Abbreviated() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFailFn, do4ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange-Fail() -> request
	//test: Do3-Exchange-Fail() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChain_Exchange() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func ExampleBuildChainExchange_Exchangeable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func do1ExchangeFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do1-Exchange() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do1-Exchange() -> response\n")
		return
	}
}

func do2ExchangeFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do2-Exchange() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do2-Exchange() -> response\n")
		return
	}
}

func do3ExchangeFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3-Exchange() -> request\n")
		//fmt.Printf("test: Do3() -> response\n")
		//return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3-Exchange() -> response\n")
		return
	}
}

func do3ExchangeFailFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3-Exchange-Fail() -> request\n")
		fmt.Printf("test: Do3-Exchange-Fail() -> response\n")
		return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3-Exchange-Fail() -> response\n")
		return
	}
}

func do4ExchangeFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do4-Exchange() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do4-Exchange() -> response\n")
		return
	}
}

/*
func do5ExchangeFn(req *http.Request) (resp *http.Response, err error) {
	fmt.Printf("test: Do5-Exchange() -> request\n")
	resp = &http.Response{StatusCode: http.StatusOK}
	fmt.Printf("test: Do5-Exchange() -> response\n")
	return
}


*/

type do1Exchange struct{}

func (d do1Exchange) Link(next Exchange) Exchange {
	return do1ExchangeFn(next)
}

type do2Exchange struct{}

func (d do2Exchange) Link(next Exchange) Exchange {
	return do2ExchangeFn(next)
}

type do3Exchange struct{}

func (d do3Exchange) Link(next Exchange) Exchange {
	return do3ExchangeFn(next)
}

type do3FailExchange struct{}

func (d do3FailExchange) Link(next Exchange) Exchange {
	return do3ExchangeFailFn(next)
}

type do4Exchange struct{}

func (d do4Exchange) Link(next Exchange) Exchange {
	return do4ExchangeFn(next)
}

/*
type do5 struct{}

func (d do5) Exchange(r *http.Request) (*http.Response, error) {
	return do5ExchangeFn(r)
}

*/

func _ExampleBuildChain_Panic_Type() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Exchange{}, req, do3Exchange{}, do4ExchangeFn})
	ex(req)

	//Output:
	//fail
}

func _ExampleBuildChain_Panic_Nil() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Exchange{}, nil, do3Exchange{}, do4ExchangeFn})
	ex(req)

	//Output:
	//fail
}

func _ExampleBuildChain_Exchange_Panic() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1ExchangeFn, do2ExchangeFn, do3ExchangeFn})
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do5-Exchange() -> request
	//test: Do5-Exchange() -> response
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response

}

func ExampleBuildChain_Chainable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Exchange{}, do2Exchange{}, do3Exchange{}})
	ex(req)

	//Output:
	//test: Do1-Exchange() -> request
	//test: Do2-Exchange() -> request
	//test: Do3-Exchange() -> request
	//test: Do3-Exchange() -> response
	//test: Do2-Exchange() -> response
	//test: Do1-Exchange() -> response

}

func do1HandlerFn(next messaging.Handler) messaging.Handler {
	return func(m *messaging.Message) {
		fmt.Printf("test: Do1-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do2HandlerFn(next messaging.Handler) messaging.Handler {
	return func(m *messaging.Message) {
		fmt.Printf("test: Do2-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do3HandlerFn(next messaging.Handler) messaging.Handler {
	return func(m *messaging.Message) {
		fmt.Printf("test: Do3-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do4HandlerFn(next messaging.Handler) messaging.Handler {
	return func(m *messaging.Message) {
		fmt.Printf("test: Do4-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

type do1Handler struct{}

func (d do1Handler) Link(next messaging.Handler) messaging.Handler {
	return do1HandlerFn(next)
}

type do2Handler struct{}

func (d do2Handler) Link(next messaging.Handler) messaging.Handler {
	return do2HandlerFn(next)
}

type do3Handler struct{}

func (d do3Handler) Link(next messaging.Handler) messaging.Handler {
	return do3HandlerFn(next)
}

type do4Handler struct{}

func (d do4Handler) Link(next messaging.Handler) messaging.Handler {
	return do4HandlerFn(next)
}

func ExampleBuildHandlerChain() {
	rec := BuildMessagingChain([]any{do1HandlerFn, do2HandlerFn, do3HandlerFn, do4HandlerFn})
	rec(messaging.ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Func() {
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]]([]any{do1HandlerFn, do2HandlerFn, do3HandlerFn, do4HandlerFn})
	rec(messaging.ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Chainable() {
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]]([]any{do1Handler{}, do2Handler{}, do3Handler{}, do4Handler{}})
	rec(messaging.ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Any() {
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]]([]any{do1Handler{}, do2HandlerFn, do3Handler{}, do4HandlerFn})
	rec(messaging.ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive
}

type do1Combined struct{}

func (d do1Combined) Link(next messaging.Handler) messaging.Handler {
	return do1HandlerFn(next)
}

func (d do1Combined) doExchange(next Exchange) Exchange {
	return do1ExchangeFn(next)
}

func ExampleBuildChain_Combined() {
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]]([]any{do1Combined{}})
	rec(messaging.ShutdownMessage)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)

	// This will panic as do1Combined is not of type Chainable[Exchange]
	//ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Combined{}})
	//ex(req)

	// This works
	ex := BuildChain[Exchange, Chainable[Exchange]]([]any{do1Combined{}.doExchange})
	ex(req)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do1-Exchange() -> request
	//test: Do1-Exchange() -> response

}

func _ExampleBuildChain_Empty() {
	// This will panic
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]](nil)
	fmt.Printf("test: BuildChain_Empty() -> %v\n", rec)

	//Output:
	//test: BuildChain_Empty() -> <nil>

}

func ExampleBuildChain_Invalid_Link() {
	// This will panic
	rec := BuildChain[messaging.Handler, Chainable[messaging.Handler]]([]any{"test string"})
	fmt.Printf("test: BuildChain_Empty() -> %v\n", rec)

	//Output:
	//test: BuildChain_Empty() -> <nil>

}
