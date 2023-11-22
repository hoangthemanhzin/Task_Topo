package ue

func (uectx *UeContext) logSendingReport(msgname string, err error) {
	if err == nil {
		uectx.Infof("Message %s sent", msgname)
	} else {
		uectx.Errorf("Send %s failed: %s", err.Error())
	}
}
