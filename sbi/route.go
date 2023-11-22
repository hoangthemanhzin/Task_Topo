package sbi

// abstraction for openapi producer handlers where correct expected
// data structures should be populated for being ready to decode from a received response
type SbiProducerHandler func(RequestContext, interface{}) Response

type SbiRoute struct {
	Label   string
	Method  string
	Path    string
	Handler SbiProducerHandler
}

type SbiRoutes []SbiRoute

type SbiService struct {
	Group   string
	Routes  SbiRoutes
	Handler interface{}
}
