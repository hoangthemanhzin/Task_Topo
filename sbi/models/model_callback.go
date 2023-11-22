package models

type Callback string

/*
type Callback struct {
	Url        string //callback url, suppress the next two attributes
	ServiceId  string //name of service
	InstanceId string //identity of running instance
	Extra      string
}

func (cb *Callback) String() string {
	if len(cb.Url) > 0 {
		return cb.Url
	}
	return fmt.Sprintf("%s_%s_%s", cb.ServiceId, cb.InstanceId, cb.Extra)
}

func (cb *Callback) Load(str string) {
	parts := strings.Split(str, "_")
	if len(parts) == 1 {
		cb.Url = str
		return
	}

	if len(parts) >= 2 {
		cb.ServiceId = parts[0]
		cb.InstanceId = parts[1]
	}
	if len(parts) >= 3 {
		cb.Extra = parts[2]
	}
}
*/
