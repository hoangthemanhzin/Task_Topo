package n42

type MessageCode uint8

// const (
// 	UPF_REGISTER_SUCCESS MessageCode = 1
// 	UPF_REGISTER_FAIL    MessageCode = 2
// )

type RegistrationRequest struct {
	UpfId    string                  `json:"upfid"`
	Slices   []string                `json:"slices"`
	Ip       string                  `json:"ip"`
	SbiPort  int                     `json:"sbiport"`
	PfcpPort int                     `json:"pfcpport"`
	Infs     map[string]NetInfConfig `json:"infs"`
	Time     int64                   `json:"time"`
}

type RegistrationResponse struct {
	Time   int64
	Status MessageCode
}

type NetInfConfig struct {
	Addr    string         `json:"addr"`
	DnnInfo *DnnInfoConfig `json:"dnninfo,omitempty"`
}
type DnnInfoConfig struct {
	Cidr string `json:"cidr"`
}
