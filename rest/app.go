package rest

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type RESTHandler func(*Request, *Response)

func NewRESTHandler(handlers ...RESTHandler) func (http.ResponseWriter,  *http.Request){
	return func (rw http.ResponseWriter, req *http.Request){
		var (
			appReq = &Request{}
			appRes = &Response{}
		)

		handlersStack := []RESTHandler{
			readHTTPRequest(req),
		}

		for _, h := range append(handlersStack, handlers...){
			if appRes.Failed() {
				break
			}
			h(appReq, appRes)
		}

		var respBody []byte
		if appRes.ResponseBody != nil {
			b, err := json.Marshal(appRes.ResponseBody)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(http.StatusText(http.StatusInternalServerError)))
				return
			}

			respBody = b
		}

		rw.Header().Set("Content-Type", "application/json")
		for k,v := range appRes.headers {
			rw.Header().Add(k, v)
		}

		// in case someone forgot to add the 200 - OK status
		if appRes.StatusCode == 0{
			appRes.StatusCode = http.StatusOK
		}

		rw.WriteHeader(appRes.StatusCode)
		rw.Write(respBody)
	}
}


func readHTTPRequest(req *http.Request) RESTHandler {
	return func(appRequest *Request, appResponse *Response){
		appRequest.req = req
		appRequest.CallBody = make(map[string]interface{})

		if req.Body == nil {
			return
		}

		defer req.Body.Close()

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			appResponse.Terminate(
				http.StatusBadRequest,
				NewError("failed to process request. empty body", nil),
			)
			return
		}

		if len(reqBody) == 0 {
			return
		}

		appRequest.reqBody = reqBody

		var jsonObject interface{}
		err = json.Unmarshal(reqBody, &jsonObject)
		if err != nil {

			appResponse.Terminate(
				http.StatusBadRequest,
				NewError("failed to unmarshal request body to json", nil),
			)
			return
		}

		switch obj := jsonObject.(type) {
		case map[string]interface{}:
			appRequest.CallBody = obj
		case []interface{}:
			appRequest.CallBody = make(map[string]interface{}) // the initial payload is available on appRequest.HTTPRequest()
		default:
			appResponse.Terminate(
				http.StatusBadRequest,
				NewError("", nil),
			)
			return

		}
		return
	}
}


