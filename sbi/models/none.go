package models

type NfProfile struct{}

type SubscribedData struct{}

type SubscribedSnssai struct {
	SubscribedSnssai  *Snssai `json:"subscribedSnssai" bson:"subscribedSnssai"`
	DefaultIndication bool    `json:"defaultIndication,omitempty" bson:"defaultIndication"`
}

type AuthorizedNetworkSliceInfo struct {
	AllowedNssaiList []AllowedNssai `json:"allowedNssaiList,omitempty" bson:"allowedNssaiList"`

	ConfiguredNssai []ConfiguredSnssai `json:"configuredNssai,omitempty" bson:"configuredNssai"`

	TargetAmfSet string `json:"targetAmfSet,omitempty" bson:"targetAmfSet"`

	CandidateAmfList []string `json:"candidateAmfList,omitempty" bson:"candidateAmfList"`

	RejectedNssaiInPlmn []Snssai `json:"rejectedNssaiInPlmn,omitempty" bson:"rejectedNssaiInPlmn"`

	RejectedNssaiInTa []Snssai `json:"rejectedNssaiInTa,omitempty" bson:"rejectedNssaiInTa"`

	NsiInformation *NsiInformation `json:"nsiInformation,omitempty" bson:"nsiInformation"`

	SupportedFeatures string `json:"supportedFeatures,omitempty" bson:"supportedFeatures"`

	NrfAmfSet string `json:"nrfAmfSet,omitempty" bson:"nrfAmfSet"`
}
