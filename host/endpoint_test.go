package host

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleExchangeHandler() {
	e := NewEndpoint(nil)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	e.Exchange(rec, req)

	fmt.Printf("test: ExchangeHandler() -> [%v]\n", req.URL.String())

	//Output:
	//test: ExchangeHandler() -> [https://www.google.com/search?q=golang]

}
