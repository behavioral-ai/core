package access

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"net/http"
)

func ExampleNewRequest() {
	h := make(http.Header)
	h.Add(aspect.XDomain, "github/advanced-go/search")
	req := RequestImpl{Method: http.MethodPatch, Url: "https://www.google.com/search?q=golang", Header: h}

	fmt.Printf("test: NewRequest() -> [method:%v] [url:%v] [h:%v]\n", req.Method, req.Url, req.Header)

	//Output:
	//test: NewRequest() -> [method:PATCH] [url:https://www.google.com/search?q=golang] [h:map[X-Domain:[github/advanced-go/search]]]

}
