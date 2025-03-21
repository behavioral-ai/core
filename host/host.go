package host

import (
	"context"
	"github.com/behavioral-ai/core/access"
	http2 "github.com/behavioral-ai/core/http"
	"net/http"
	"time"
)

const (
	Route    = "host"
	EtcRoute = "etc"
	XFrom    = "X-From"
	XTo      = "X-To"
	XRoute   = "X-Route"
)

var (
	//exchangeProxy = aspect.NewExchangeProxy()
	hostDuration = time.Second * 5
	authExchange http2.Exchange
	okFunc       = func(code int) bool { return code == http.StatusOK }
)

func init() {
	resp, _ := http2.NewResponse(http.StatusOK, nil, nil)
	authExchange = func(_ *http.Request) (*http.Response, error) {
		return resp, nil
	}
}

func SetHostTimeout(d time.Duration) {
	hostDuration = d
}

func SetAuthExchange(h http2.Exchange, ok func(int) bool) {
	if h != nil {
		authExchange = h
		if ok != nil {
			okFunc = ok
		}
	}
}

func Exchange(w http.ResponseWriter, r *http.Request, handler http2.Exchange) {
	controllerCode := ""
	start := time.Now().UTC()
	var resp *http.Response
	var err error

	http2.AddRequestId(r)
	resp, err = authExchange(r)
	if !okFunc(resp.StatusCode) {
		w.WriteHeader(resp.StatusCode)
		access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Routing{From: "", Route: Route, To: "", Percent: -1}, access.Controller{Timeout: hostDuration, RateLimit: 0, RateBurst: 0, Code: controllerCode})
		return
	}
	from := r.Header.Get(XFrom)
	if from == "" {
		r.Header.Set(XFrom, Route)
	}
	r.Header.Set(XFrom, Route)
	// TODO: Need to create a new request with the appropriate timeout and host name, using an intermediary
	if hostDuration > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), hostDuration)
		defer cancel()
		r2 := r.Clone(ctx)
		resp, err = handler(r2)
	} else {
		resp, err = handler(r)
	}
	resp.Header.Del(XRoute)
	if err != nil && err.Error() == "httpx.StatusGatewayTimeout" {
		controllerCode = access.ControllerTimeout
	}
	resp.ContentLength = http2.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
	r.Header.Set(XTo, Route)
	access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Routing{From: from, Route: Route, To: "", Percent: -1}, access.Controller{Timeout: hostDuration, RateLimit: 0, RateBurst: 0, Code: controllerCode})
}
