package sbi

import (
	"etrib5gc/sbi/models"
	"net/url"
)

type Request struct {
	Path         string
	Method       string
	Body         interface{}
	HeaderParams map[string]string
	QueryParams  url.Values
	FormParams   url.Values
	FormFileName string
	FileName     string
	FileBytes    []byte
	Callback     models.Callback

	BodyBytes []byte      //encoded body
	Request   interface{} //dataplane protocol-bound request
}

func DefaultRequest() *Request {
	ret := &Request{
		HeaderParams: make(map[string]string),
	}
	ret.HeaderParams["Content-Type"] = "application/json"
	ret.HeaderParams["Accept"] = "application/json;application/problem+json"
	return ret
}

type Response struct {
	Response  interface{}
	BodyBytes []byte

	Body       interface{}
	Status     string
	StatusCode int
}

func (resp *Response) SetBody(code int, body interface{}) {
	resp.Body = body
	resp.StatusCode = code
}

func (resp *Response) SetProblem(prob *models.ProblemDetails) {
	//panic if prob is nil

	resp.StatusCode = int(prob.Status)
	resp.Status = prob.Detail
	resp.Body = prob
}

// Abstraction of a consumer client
//NOTE: a client should call 'Send' method first to get a response. Then he
//should allocate a data structure for decoding the response body. The type of
//the data structure is dependent on the status code of the response
type ConsumerClient interface {
	Send(*Request) (*Response, error)
	DecodeResponse(*Response) error
}

// an abstraction of the context where a request is received at a producer. The
// first handler method (openapi/producers) will inject a correct expected
// body for decoding. The next handler (application layer) will process the
// decoded body then create a response. It then write it to a response writer
// to complete the whole procedure of handling a request.

type RequestContext interface {
	DecodeRequest(body interface{}) error //decode the request to get embeded body
	Param(string) string                  // get a parameter from the request (application handler need it)
	Header(string) string                 // get a header parameter from the request (application handler need it)
}
