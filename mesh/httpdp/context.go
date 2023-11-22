package httpdp

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// a wrapper for gin.Context. It implements the sbi.RequestContext
// abstraction
type ginContextEx struct {
	context  *gin.Context
	request  *sbi.Request
	response sbi.Response
}

func newGinContextEx(ctx *gin.Context) *ginContextEx {
	ret := &ginContextEx{
		context: ctx,
		request: &sbi.Request{
			Request: ctx.Request,
		},
	}
	return ret
}

func (c *ginContextEx) Param(key string) string {
	return c.context.Param(key)
}

func (c *ginContextEx) Header(key string) string {
	return c.context.Request.Header.Get(key)
}

// must be called by an sbi's producer handler
func (c *ginContextEx) DecodeRequest(body interface{}) (err error) {
	//1. read the request body
	if c.request.BodyBytes, err = c.context.GetRawData(); err != nil {
		return
	} else {
		//2. decode the body (suppose that a right type of body has been pupulated
		c.request.Body = body
		if err = DecodeRequest(c.request); err != nil {
			return
		}
	}
	return nil
}

func (c *ginContextEx) writeResponse() {
	//encode the body of the response into bytes
	if err := EncodeResponse(&c.response); err != nil {
		//write the poblem in application/json format
		c.context.JSON(http.StatusInternalServerError, &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	} else {
		//TODO: how to set the Content-Type?
		//write to the response writer
		c.context.Data(c.response.StatusCode, "application/json", c.response.BodyBytes)
	}
}

// build a service handler for gin
// the first parameter is the sbi handler, the second parameter is an
// NF specific producer implementation.
func CreateGinHandler(sbiFn sbi.SbiProducerHandler, handler interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		ctx := newGinContextEx(context)
		//call the sbi producer handler to decode request and call
		//application handler to process the request
		ctx.response = sbiFn(ctx, handler)
		//write the response or error to the dataplane
		ctx.writeResponse()
	}
}

func MakeHttpRoutes(routes sbi.SbiRoutes, handler interface{}) HttpRoutes {
	ret := make([]HttpRoute, len(routes), len(routes))
	for i, r := range routes {
		ret[i] = HttpRoute{
			Name:        r.Label,
			Method:      r.Method,
			Pattern:     r.Path,
			HandlerFunc: CreateGinHandler(r.Handler, handler),
		}
	}
	return ret
}
