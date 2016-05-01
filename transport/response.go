package transport

import "net/http"

//Response is just a convenience for tcp response handling
type Response struct {
	StatusCode int
	Data       []byte
	Headers    http.Header
}

func (r *Response) Write(b []byte) (int, error) {
	//set status to 200 if not set
	if r.StatusCode == 0 {
		r.StatusCode = 200
	}
	if r.Data == nil {
		r.Data = b
	} else {
		r.Data = append(r.Data, b...)
	}
	return len(b), nil
}

func (r *Response) Header() http.Header {
	return r.Headers
}

func (r *Response) WriteHeader(status int) {
	r.StatusCode = status
}
