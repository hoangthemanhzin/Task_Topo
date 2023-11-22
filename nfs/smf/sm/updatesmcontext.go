package sm

func (smctx *SmContext) handleUpdateSmContext(job *UpdateSmContextJob) {
	defer job.buildResponse()
	smctx.handleN1Msg(job)

	if job.err != nil {
		return
	}

	smctx.handleN2Info(job)

	if job.err != nil {
		return
	}
	if job.Req.JsonData.AnTypeCanBeChanged {
		smctx.handleAnType(job)
	} else if len(job.Req.JsonData.HoState) > 0 {
		smctx.handleHoState(job)
	} else if len(job.Req.JsonData.Cause) > 0 {
		smctx.handleCause(job)
	} else if len(job.Req.JsonData.UpCnxState) > 0 {
		smctx.handleUpCnxState(job)
	}
	return
}

func (smctx *SmContext) handleAnType(job *UpdateSmContextJob) {
	smctx.Warnf("AnType change is not handled")
}
