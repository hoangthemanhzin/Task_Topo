package sm

func (smctx *SmContext) logSendingReport(msg string, err error) {
	if err != nil {
		smctx.Errorf("Send %s failed: %+v", msg, err)
	} else {
		smctx.Infof("%s sent", msg)
	}
}
