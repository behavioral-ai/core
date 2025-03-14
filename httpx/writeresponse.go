package httpx

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/iox"
	"net/http"
)

// WriteResponse - write a http.Response, utilizing the content, status code, and headers
// Content types supported: []byte, string, error, io.Reader, io.ReadCloser. Other types will be treated as JSON and serialized, if
// the headers content type is JSON. If not JSON, then an error will be raised.
func WriteResponse(w http.ResponseWriter, headers any, statusCode int, content any, reqHeader http.Header) (contentLength int64) {
	var e aspect.Log

	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	SetHeaders(w, headers)
	if content == nil {
		w.WriteHeader(statusCode)
		return 0
	}
	if len(w.Header().Get(ContentEncoding)) != 0 {
		reqHeader.Set(AcceptEncoding, "")
	}
	writer, err := iox.NewEncodingWriter(w, reqHeader)
	if err != nil {
		status0 := aspect.NewStatusError(aspect.StatusIOError, err)
		e.Handle(status0.WithRequestId(w.Header()))
		w.WriteHeader(status0.HttpCode())
		return 0
	}
	if writer.ContentEncoding() != iox.NoneEncoding {
		w.Header().Add(ContentEncoding, writer.ContentEncoding())
	}
	w.WriteHeader(statusCode)
	var status0 *aspect.Status
	contentLength, status0 = writeContent(writer, content, w.Header().Get(ContentType))
	_ = writer.Close()
	if !status0.OK() {
		e.Handle(status0.WithRequestId(w.Header()))
	}
	return contentLength
}

func CreateAcceptEncodingHeader() http.Header {
	out := make(http.Header)
	out.Add(AcceptEncoding, AcceptEncodingValue)
	return out
}
