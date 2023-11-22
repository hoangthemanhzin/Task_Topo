package upman

/*
func (upman *UpManager) HandleSessionReportRequest(upfid string, localseid uint64, req *pfcpmsg.PFCPSessionReportRequest) (rsp *pfcpmsg.PFCPSessionReportResponse, remoteseid uint64, err error) {
	cause := pfcptypes.Cause{
		CauseValue: pfcptypes.CauseRequestAccepted,
	}
	rsp = &pfcpmsg.PFCPSessionReportResponse{
		Cause: &cause,
	}
	if node := upman.findNode(upfid); node != nil {
		if session := node.upf.FindSession(localseid); session != nil {
			//TODO: handling bussiness logic goes here

			remoteseid = session.RemoteSeid()
			return
		}
	}

	err = fmt.Errorf("Upf not found")
	cause.CauseValue = pfcptypes.CauseRequestRejected
	return
}
*/
