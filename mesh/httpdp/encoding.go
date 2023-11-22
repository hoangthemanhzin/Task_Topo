package httpdp

import (
	"bytes"
	"encoding/json"
	"etrib5gc/sbi"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// unmarshaling the request body bytes (given a http request embeded in the
// given sbi's request)
func DecodeRequest(req *sbi.Request) (err error) {
	httpreq := req.Request.(*http.Request)
	req.Path = httpreq.URL.Path
	req.Method = httpreq.Method

	//var header map[string]string
	req.HeaderParams = make(map[string]string)
	for key, element := range httpreq.Header {
		req.HeaderParams[key] = element[0]
	}

	if err = json.Unmarshal(req.BodyBytes, req.Body); err != nil {
		err = fmt.Errorf("Decode Request failed: %+v", err)
	}
	return
}

// unmarshaling the response body bytes (give a http response embeded in the
// given sbi's response)
func DecodeResponse(rsp *sbi.Response) (err error) {
	if rsp.Body != nil {
		if err = json.Unmarshal(rsp.BodyBytes, rsp.Body); err != nil {
			err = fmt.Errorf("Decode Response failed: %+v", err)
		}
	}
	return
}

// marshaling the request body into bytes arrays the prepare a http request
func EncodeRequest(req *sbi.Request) (err error) {
	var body io.Reader
	if req.Body != nil {
		if req.BodyBytes, err = json.Marshal(req.Body); err != nil {
			return
		}
		body = bytes.NewBuffer(req.BodyBytes)
	}

	var link *url.URL
	if link, err = url.Parse(req.Path); err != nil {
		return
	}
	query := link.Query()
	for k, v := range req.QueryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	link.RawQuery = query.Encode()
	var httpreq *http.Request
	if httpreq, err = http.NewRequest(req.Method, link.String(), body); err != nil {
		return
	}
	//log.Tracef("send a request [url=%s]", link.String())
	if len(req.HeaderParams) > 0 {
		headers := http.Header{}
		for h, v := range req.HeaderParams {
			/*
				if h == "Callback" {
					log.Infof("Have a callback:%s", v)
				}
			*/
			headers.Set(h, v)
		}
		httpreq.Header = headers
	}
	req.Request = httpreq
	return
}

// marshaling the response body into an byte array. No need to create an http
// response since gin.Context will write the response
func EncodeResponse(rsp *sbi.Response) (err error) {
	rsp.BodyBytes, err = json.Marshal(rsp.Body)
	return
}
