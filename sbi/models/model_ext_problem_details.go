package models

type ExtProblemDetails struct {
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

	RemoteError bool `json:"remoteError,omitempty"`
}
