package fasthttputil

import (
	"encoding/json"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

type Test struct {
	Foo string `form:"foo" json:"foo,omitempty"`
	Bar int    `form:"bar" json:"bar,omitempty"`
}

func confirmResponse(t *testing.T, expected *Test, res *fasthttp.Response) bool {
	actual := &Test{}
	err := json.Unmarshal(res.Body(), actual)
	return !assert.NoError(t, err) || !assert.EqualValues(t, expected, actual)
}

func isOK(t *testing.T, err error, res *fasthttp.Response) bool {
	return !assert.NoError(t, err) || !assert.Equal(t, fasthttp.StatusOK, res.StatusCode(), string(res.Body()))
}

func TestBind(t *testing.T) {
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			st := Test{}
			err := Bind(ctx, &st)
			if err != nil {
				ctx.SetBodyString(err.Error())
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return
			}

			ctx.SetContentType("application/json")
			ctx.SetStatusCode(fasthttp.StatusOK)
			err = json.NewEncoder(ctx).Encode(st)
			if err != nil {
				t.Fatal(err)
			}
		},
	}

	ln := fasthttputil.NewInmemoryListener()
	go s.Serve(ln)

	c := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.SetRequestURI("http://localhost")
	req.URI().QueryArgs().Set("foo", "bar")
	err := c.Do(req, res)
	if !isOK(t, err, res) || !confirmResponse(t, &Test{Foo: "bar"}, res) {
		return
	}

	req.URI().QueryArgs().Set("bar", "99")
	err = c.Do(req, res)
	if !isOK(t, err, res) || !confirmResponse(t, &Test{Foo: "bar", Bar: 99}, res) {
		return
	}

	req.URI().QueryArgs().Del("foo")
	err = c.Do(req, res)
	if !isOK(t, err, res) || !confirmResponse(t, &Test{Bar: 99}, res) {
		return
	}

	req.URI().QueryArgs().Del("bar")
	req.Header.SetMethod(http.MethodPost)
	req.SetBodyString("foo=bar")
	err = c.Do(req, res)
	if !isOK(t, err, res) || !confirmResponse(t, &Test{Foo: "bar"}, res) {
		return
	}

	req.SetBodyString("{\"foo\": \"bar\"}")
	req.Header.SetContentType("application/json")
	err = c.Do(req, res)
	if !isOK(t, err, res) || !confirmResponse(t, &Test{Foo: "bar"}, res) {
		return
	}
}
