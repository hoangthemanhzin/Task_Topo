package up

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
)

func inState(s RuleState, states []RuleState) bool {
	for _, state := range states {
		if s == state {
			return true
		}
	}
	return false
}

//get rule lists for create/update/delete
func getRuleList(pdrs []*PDR, states []RuleState) (pdrlist []*PDR, farlist []*FAR, barlist []*BAR, qerlist []*QER, urrlist []*URR) {
	qers := make(map[uint32]bool)
	urrs := make(map[uint32]bool)
	for _, pdr := range pdrs {
		if inState(pdr.State, states) {
			pdrlist = append(pdrlist, pdr)
		}
		if far := pdr.FAR; far != nil {
			if inState(far.State, states) {
				farlist = append(farlist, far)
			}
			if bar := far.BAR; bar != nil && inState(bar.State, states) {
				barlist = append(barlist, bar)
			}
		}

		for _, urr := range pdr.URR {
			if _, ok := urrs[urr.URRID]; !ok && inState(urr.State, states) {
				urrlist = append(urrlist, urr)
				urrs[urr.URRID] = true
			}
		}
		for _, qer := range pdr.QER {
			if _, ok := qers[qer.QERID]; !ok && inState(qer.State, states) {
				qerlist = append(qerlist, qer)
				qers[qer.QERID] = true
			}
		}
	}
	return
}

func (session *PfcpSession) FillDeletionRequest(msg *pfcpmsg.PFCPSessionDeletionRequest) {
	//not implmented yet
}
func (session *PfcpSession) FillEstablishmentRequest(msg *pfcpmsg.PFCPSessionEstablishmentRequest) {
	isv4 := true

	msg.CPFSEID = &pfcptypes.FSEID{
		V4:          isv4,
		V6:          !isv4,
		Seid:        session.localseid,
		Ipv4Address: session.upf.ip,
	}

	pdrlist, farlist, barlist, qerlist, urrlist := getRuleList(session.pdrs, []RuleState{RULE_INITIAL})

	for _, pdr := range pdrlist {
		msg.CreatePDR = append(msg.CreatePDR, pdr.toCreatePdr())
		pdr.State = RULE_CREATE
	}

	for _, far := range farlist {
		msg.CreateFAR = append(msg.CreateFAR, far.toCreateFar())
		far.State = RULE_CREATE
	}

	for _, bar := range barlist {
		msg.CreateBAR = append(msg.CreateBAR, bar.toCreateBar())
		bar.State = RULE_CREATE
	}

	for _, qer := range qerlist {
		msg.CreateQER = append(msg.CreateQER, qer.toCreateQer())
		qer.State = RULE_CREATE
	}
	for _, urr := range urrlist {
		msg.CreateURR = append(msg.CreateURR, urr.toCreateUrr())
		urr.State = RULE_CREATE
	}

	msg.PDNType = &pfcptypes.PDNType{
		PdnType: pfcptypes.PDNTypeIpv4,
	}
	return
}

func (session *PfcpSession) FillModificationRequest(msg *pfcpmsg.PFCPSessionModificationRequest) {
	v4 := true
	msg.CPFSEID = &pfcptypes.FSEID{
		V4:          v4,
		V6:          !v4,
		Seid:        session.localseid,
		Ipv4Address: session.upf.ip,
	}

	pdrlist, farlist, barlist, qerlist, urrlist := getRuleList(session.pdrs, []RuleState{RULE_INITIAL, RULE_UPDATE, RULE_REMOVE})

	for _, pdr := range pdrlist {
		switch pdr.State {
		case RULE_INITIAL:
			msg.CreatePDR = append(msg.CreatePDR, pdr.toCreatePdr())
		case RULE_UPDATE:
			msg.UpdatePDR = append(msg.UpdatePDR, pdr.toUpdatePdr())
		case RULE_REMOVE:
			msg.RemovePDR = append(msg.RemovePDR, pdr.toRemovePdr())
		}
		pdr.State = RULE_CREATE
	}

	for _, far := range farlist {
		switch far.State {
		case RULE_INITIAL:
			msg.CreateFAR = append(msg.CreateFAR, far.toCreateFar())
		case RULE_UPDATE:
			msg.UpdateFAR = append(msg.UpdateFAR, far.toUpdateFar())
		case RULE_REMOVE:
			msg.RemoveFAR = append(msg.RemoveFAR, far.toRemoveFar())
		}
		far.State = RULE_CREATE
	}

	for _, bar := range barlist {
		switch bar.State {
		case RULE_INITIAL:
			msg.CreateBAR = append(msg.CreateBAR, bar.toCreateBar())
		case RULE_UPDATE:
			msg.UpdateBAR = bar.toUpdateBar() //TODO: should re-check this one
		case RULE_REMOVE:
			msg.RemoveBAR = append(msg.RemoveBAR, bar.toRemoveBar())
		}

		bar.State = RULE_CREATE
	}

	for _, qer := range qerlist {
		switch qer.State {
		case RULE_INITIAL:
			msg.CreateQER = append(msg.CreateQER, qer.toCreateQer())
		case RULE_UPDATE:
			msg.UpdateQER = append(msg.UpdateQER, qer.toUpdateQer())
		case RULE_REMOVE:
			msg.RemoveQER = append(msg.RemoveQER, qer.toRemoveQer())
		}

		qer.State = RULE_CREATE
	}
	for _, urr := range urrlist {
		switch urr.State {
		case RULE_INITIAL:
			msg.CreateURR = append(msg.CreateURR, urr.toCreateUrr())
		case RULE_UPDATE:
			msg.UpdateURR = append(msg.UpdateURR, urr.toUpdateUrr())
		case RULE_REMOVE:
			msg.RemoveURR = append(msg.RemoveURR, urr.toRemoveUrr())
		}
		urr.State = RULE_CREATE
	}
}
