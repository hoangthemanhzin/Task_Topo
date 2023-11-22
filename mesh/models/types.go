package models

import (
	"strings"
)

type ServiceName string

type Labels map[string]string
type Selectors map[string]string

func (s Selectors) Match(l Labels) bool {
	if len(s) == 0 || len(s) == 0 {
		//no selector or no label
		return false
	}
	for k, v := range s {
		if v1, ok := l[k]; !ok {
			//key not exits
			return false
		} else if strings.Compare(v, v1) != 0 {
			//values are not the same
			return false
		}
	}
	return true
}

type Endpoint struct {
	Id      string
	Ip      string
	SbiPort int
	Labels  Labels
}

type Endpoints []Endpoint

type RouteMatch map[string]string

func NewRouteMatch() RouteMatch {
	return make(map[string]string)
}

func (m RouteMatch) Set(key string, value string) {
	m[key] = value
}

type MatchRule struct {
	Headers map[string]string `json:"headers"`
}

type Destination struct {
	GroupId string `json:"group"`
	Weight  int    `json:"weight"`
	Lb      uint8  `json:"lb"`
}

type RouteRule struct {
	//	Id           string         //rule identity
	Match        MatchRule     `json:"match"`        //how to match
	Destinations []Destination `json:"destinations"` //where to send requests
}

func (r *RouteRule) IsMatched(match RouteMatch) bool {
	if r.isEmptyMatch() { //no matching rule
		//always match
		return true
	}

	//empty matching hint
	if len(match) == 0 {
		//no match
		return false
	}

	//the hint must includes all k-v pairs in the matching rule
	for k, v := range r.Match.Headers {
		if v1, ok := match[k]; !ok { //key not exist
			return false
		} else {
			if strings.Compare(v, v1) != 0 { //value not matched
				return false
			}
		}
	}
	return true
}

// check if a matching rule is available
func (r *RouteRule) isEmptyMatch() bool {
	return len(r.Match.Headers) == 0
}

type EndpointGroup struct {
	//	Id       string   //group identity (in a service)
	Selectors Selectors `json:"selectors"` //for selecting endpoints in a service group
}

type Service struct {
	Id        ServiceName              `json:"id"`        //service name
	Selectors Selectors                `json:"selectors"` //for selecting service endpoints
	Groups    map[string]EndpointGroup `json:"groups"`    //pre-defined groups  of endpoints
	Routes    []RouteRule              `json:"routes"`    //for route matching
}
