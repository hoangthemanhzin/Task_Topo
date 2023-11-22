package models

type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`

	NrLocation *NrLocation `json:"nrLocation,omitempty"`

	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`

	UtraLocation *UtraLocation `json:"utraLocation,omitempty"`

	GeraLocation *GeraLocation `json:"geraLocation,omitempty"`
}
