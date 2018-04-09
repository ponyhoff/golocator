package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Request struct {
	CallBody map[string]interface{}
	Params   map[string]interface{}

	req     *http.Request
	reqBody []byte
}

func (r *Request) HTTPRequest() *http.Request {
	if len(r.reqBody) > 0 {
		r.req.Body = ioutil.NopCloser(bytes.NewReader(r.reqBody))
	}
	return r.req
}


func (r *Request) AddParams(key, value string) {
	if r.Params == nil {
		r.Params = make(map[string]interface{})
	}
	r.Params[key] = value
}