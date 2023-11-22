package models

import "fmt"

type ProblemDetails struct {
	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	Detail string `json:"detail,omitempty"`

	Instance string `json:"instance,omitempty"`

	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	AccessTokenError *AccessTokenErr `json:"accessTokenError,omitempty"`

	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`

	NrfId string `json:"nrfId,omitempty"`
}

func (p *ProblemDetails) GetType() string {
	return p.Type
}
func (p *ProblemDetails) GetStatus() int32 {
	return p.Status
}

func (p *ProblemDetails) Error() string {
	return p.Detail
}
func (p *ProblemDetails) MakeError() error {
	return fmt.Errorf("Code:%d, Detail: %s", p.Status, p.Detail)
}
