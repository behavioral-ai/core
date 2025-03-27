package httpx

import (
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"strings"
)

const (
	ContentTypeJson     = "application/json"
	ContentType         = "Content-Type"
	ContentEncoding     = "Content-Encoding"
	AcceptEncoding      = "Accept-Encoding"
	AcceptEncodingValue = "gzip, deflate, br"
	ContentLength       = "Content-Length"
	ContentEncodingGzip = "gzip"
	ContentTypeTextHtml = "text/html"
	ContentTypeText     = "text/plain charset=utf-8"
	ContentLocation     = "Content-Location"
	ExchangeOverride    = "X-Exchange-Override"
	ContentResolver     = "X-Content-Resolver"
	ResolverSeparator   = "->"
)

func SetHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Set(name, value)
	return h
}

func SetHeaders(w http.ResponseWriter, headers any) {
	if headers == nil {
		return
	}
	if pairs, ok := headers.([]Attr); ok {
		for _, pair := range pairs {
			w.Header().Set(strings.ToLower(pair.Key), pair.Value)
		}
		return
	}
	if h, ok := headers.(http.Header); ok {
		for k, v := range h {
			if len(v) > 0 {
				w.Header().Set(strings.ToLower(k), v[0])
			}
		}
	}
}

func CloneHeader(hdr http.Header) http.Header {
	clone := hdr.Clone()
	if clone == nil {
		clone = make(http.Header)
	}
	return clone
}

func CloneHeaderWithEncoding(req *http.Request) http.Header {
	if req == nil {
		return make(http.Header)
	}
	h := CloneHeader(req.Header)
	if req.Method == http.MethodGet && h.Get(iox.AcceptEncoding) == "" {
		h.Add(iox.AcceptEncoding, iox.GzipEncoding)
	}
	return h
}
