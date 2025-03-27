package httpx

import (
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"net/http"
)

func ExampleCopy() {
	h := make(http.Header)
	h.Add("key-1", "value-1")
	h.Add("key-2", "value-2")
	h.Add("key-3", "value-3")

	h2 := CloneHeader(h)

	fmt.Printf("test: CloneHeader() -> %v\n", h2)

	//Output:
	//test: CloneHeader() -> map[Key-1:[value-1] Key-2:[value-2] Key-3:[value-3]]

}

func ExampleCloneHeaderWithEncoding() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	h := make(http.Header)
	h.Add("key-1", "value-1")
	h.Add("key-2", "value-2")
	h.Add("key-3", "value-3")
	req.Header = h

	h2 := CloneHeaderWithEncoding(req)
	fmt.Printf("test: CloneHeaderWithEncoding() -> %v\n", h2.Get(iox.AcceptEncoding))

	//Output:
	//test: CloneHeaderWithEncoding() -> gzip

}
