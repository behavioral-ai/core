package http

import (
	"encoding/json"
	"errors"
	"fmt"
	iox "github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"io"
	"reflect"
	"strings"
)

const (
	jsonToken = "json"
)

func writeContent(w io.Writer, content any, contentType string) (length int64, status *messaging.Status) {
	var err error
	var cnt int

	if content == nil {
		return 0, messaging.StatusOK()
	}
	switch ptr := (content).(type) {
	case []byte:
		cnt, err = w.Write(ptr)
	case string:
		cnt, err = w.Write([]byte(ptr))
	case error:
		cnt, err = w.Write([]byte(ptr.Error()))
	case io.Reader:
		var buf []byte
		var err1 error

		buf, err1 = iox.ReadAll(ptr, nil)
		if err1 != nil {
			status = messaging.NewStatusError(messaging.StatusIOError, err, "")
			return 0, status
		}
		status = messaging.StatusOK()
		cnt, err = w.Write(buf)
	case io.ReadCloser:
		var buf []byte
		var err1 error

		buf, err1 = iox.ReadAll(ptr, nil)
		_ = ptr.Close()
		if err1 != nil {
			status = messaging.NewStatusError(messaging.StatusIOError, err, "")
			return 0, status
		}
		status = messaging.StatusOK()
		cnt, err = w.Write(buf)
	default:
		if strings.Contains(contentType, jsonToken) {
			var buf []byte

			buf, err = json.Marshal(content)
			if err != nil {
				status = messaging.NewStatusError(messaging.StatusJsonEncodeError, err, "")
				if !status.OK() {
					return
				}
			}
			cnt, err = w.Write(buf)
		} else {
			return 0, messaging.NewStatusError(messaging.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))), "")
		}
	}
	if err != nil {
		return 0, messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return int64(cnt), messaging.StatusOK()
}
