package common

type AmfQuery struct {
	PlmnId     string
	Region     string
	Set        string
	Pointer    string
	InstanceId string
}

type UdmQuery struct {
	PlmnId     string
	Group      string
	InstanceId string
}

type AusfQuery struct {
	PlmnId     string
	Group      string
	InstanceId string
}

type SmfQuery struct {
	PlmnId     string
	Slice      string
	Dnn        string
	InstanceId string
}

type PcfQuery struct {
	PlmnId     string
	Slice      string
	Dnn        string
	InstanceId string
}
