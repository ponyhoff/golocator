package rest

type Response struct {
	StatusCode   int
	ResponseBody map[string]interface{}

	shouldTerminate bool
	headers         map[string]string
}

type RESTError struct {
	ErrorCode *string
	Message   string
}

func NewError(message string, errorCode *string) RESTError {
	return RESTError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func (r *Response) Failed() bool {
	return r.shouldTerminate
}

func (r *Response) End() {
	r.shouldTerminate = true
}

func (r *Response) AppendHeader(key, value string) {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = value
}

func (r *Response) Terminate(statusCode int, e RESTError) {
	r.StatusCode = statusCode
	r.ResponseBody = e.ResponseBody()
	r.End()
}

func (err RESTError) ResponseBody() map[string]interface{} {
	b := map[string]interface{}{
		"message": err.Message,
	}
	if err.ErrorCode != nil {
		b["error_code"] = err.ErrorCode
	}

	return b
}
