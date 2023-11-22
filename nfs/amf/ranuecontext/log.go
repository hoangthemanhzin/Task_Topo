package ranuecontext

func (ranue *RanUe) logSendingReport(msg string, err error) {
	if err != nil {
		ranue.Errorf("Send %s failed: %s", msg, err.Error())
	} else {
		ranue.Infof("%s sent", msg)
	}
}
