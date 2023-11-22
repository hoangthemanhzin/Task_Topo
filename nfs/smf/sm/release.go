package sm

func (smctx *SmContext) releaseSmContext() (err error) {
	//TODO: should we need to send N1,N2 messages?
	state := smctx.CurrentState()
	if state == SM_ACTIVE {
		smctx.Infof("Release tunnel")
		err = smctx.tunnel.Release()
	}
	return
}
