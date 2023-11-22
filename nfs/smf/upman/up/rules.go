package up

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
)

const (
	RULE_INITIAL RuleState = 0
	RULE_CREATE  RuleState = 1
	RULE_UPDATE  RuleState = 2
	RULE_REMOVE  RuleState = 3
)

type RuleState uint8

// Packet Detection Rule. Table 7.5.2.2-1
type PDR struct {
	PDRID uint16

	Precedence         uint32
	PDI                PDI
	OuterHeaderRemoval *pfcptypes.OuterHeaderRemoval

	FAR *FAR
	URR []*URR
	QER []*QER

	State RuleState
}

func (pdr *PDR) toCreatePdr() *pfcpmsg.CreatePDR {
	cpdr := &pfcpmsg.CreatePDR{
		PDRID: &pfcptypes.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
		Precedence: &pfcptypes.Precedence{
			PrecedenceValue: pdr.Precedence,
		},
		PDI: &pfcpmsg.PDI{
			SourceInterface: &pdr.PDI.SourceInterface,
			LocalFTEID:      pdr.PDI.LocalFTeid,
			NetworkInstance: pdr.PDI.NetworkInstance,
			UEIPAddress:     pdr.PDI.UEIPAddress,
		},
		OuterHeaderRemoval: pdr.OuterHeaderRemoval,
	}

	if pdr.PDI.ApplicationID != "" {
		cpdr.PDI.ApplicationID = &pfcptypes.ApplicationID{
			ApplicationIdentifier: []byte(pdr.PDI.ApplicationID),
		}
	}

	if pdr.PDI.SDFFilter != nil {
		cpdr.PDI.SDFFilter = pdr.PDI.SDFFilter
	}

	if far := pdr.FAR; far != nil {
		cpdr.FARID = &pfcptypes.FARID{
			FarIdValue: far.FARID,
		}
	}
	for _, urr := range pdr.URR {
		cpdr.URRID = append(cpdr.URRID, &pfcptypes.URRID{
			UrrIdValue: urr.URRID,
		})
	}

	for _, qer := range pdr.QER {
		cpdr.QERID = append(cpdr.QERID, &pfcptypes.QERID{
			QERID: qer.QERID,
		})
	}

	return cpdr
}

func (pdr *PDR) toUpdatePdr() *pfcpmsg.UpdatePDR {
	updr := &pfcpmsg.UpdatePDR{
		PDRID: &pfcptypes.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
		Precedence: &pfcptypes.Precedence{
			PrecedenceValue: pdr.Precedence,
		},
		PDI: &pfcpmsg.PDI{
			SourceInterface: &pdr.PDI.SourceInterface,
			LocalFTEID:      pdr.PDI.LocalFTeid,
			NetworkInstance: pdr.PDI.NetworkInstance,
			UEIPAddress:     pdr.PDI.UEIPAddress,
		},
	}

	if pdr.PDI.ApplicationID != "" {
		updr.PDI.ApplicationID = &pfcptypes.ApplicationID{
			ApplicationIdentifier: []byte(pdr.PDI.ApplicationID),
		}
	}

	if pdr.PDI.SDFFilter != nil {
		updr.PDI.SDFFilter = pdr.PDI.SDFFilter
	}

	updr.OuterHeaderRemoval = pdr.OuterHeaderRemoval

	if pdr.FAR != nil {
		updr.FARID = &pfcptypes.FARID{
			FarIdValue: pdr.FAR.FARID,
		}
	}
	return updr
}

func (pdr *PDR) toRemovePdr() *pfcpmsg.RemovePDR {
	return &pfcpmsg.RemovePDR{
		PDRID: &pfcptypes.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
	}
}

// Packet Detection.
// 7.5.2.2-2
type PDI struct {
	SourceInterface pfcptypes.SourceInterface
	LocalFTeid      *pfcptypes.FTEID
	NetworkInstance *pfcptypes.NetworkInstance
	UEIPAddress     *pfcptypes.UEIPAddress
	SDFFilter       *pfcptypes.SDFFilter
	ApplicationID   string
}

// Forwarding Action Rule
type FAR struct {
	FARID uint32

	ApplyAction          pfcptypes.ApplyAction
	ForwardingParameters *ForwardingParameters

	BAR   *BAR
	State RuleState
}

func (far *FAR) toCreateFar() *pfcpmsg.CreateFAR {
	//add createFAR
	cfar := &pfcpmsg.CreateFAR{
		FARID: &pfcptypes.FARID{
			FarIdValue: far.FARID,
		},
		ApplyAction: &pfcptypes.ApplyAction{},
	}

	if far.ForwardingParameters != nil {
		cfar.ApplyAction.Forw = true
	} else {
		//	29.244 v15.3 Table 7.5.2.3-1 Farwarding Parameters IE shall be
		//	present when the Apply-Action requests the packets to be forwarded.
		//	FAR without Farwarding Parameters set Apply Action as Drop instead
		//	of Forward.

		cfar.ApplyAction.Forw = false
		cfar.ApplyAction.Drop = true
	}

	if far.ForwardingParameters != nil {
		cfar.ForwardingParameters = &pfcpmsg.ForwardingParametersIEInFAR{
			DestinationInterface: &far.ForwardingParameters.DestinationInterface,
			NetworkInstance:      far.ForwardingParameters.NetworkInstance,
			OuterHeaderCreation:  far.ForwardingParameters.OuterHeaderCreation,
		}
		if far.ForwardingParameters.ForwardingPolicyID != "" {
			cfar.ForwardingParameters.ForwardingPolicy = &pfcptypes.ForwardingPolicy{
				ForwardingPolicyIdentifierLength: uint8(len(far.ForwardingParameters.ForwardingPolicyID)),
				ForwardingPolicyIdentifier:       []byte(far.ForwardingParameters.ForwardingPolicyID),
			}
		}
	}
	if far.BAR != nil {
		cfar.BARID = &pfcptypes.BARID{
			BarIdValue: far.BAR.BARID,
		}
	}
	return cfar
}

func (far *FAR) toUpdateFar() *pfcpmsg.UpdateFAR {
	ufar := &pfcpmsg.UpdateFAR{
		FARID: &pfcptypes.FARID{
			FarIdValue: far.FARID,
		},
		ApplyAction: &pfcptypes.ApplyAction{
			Forw: far.ApplyAction.Forw,
			Buff: far.ApplyAction.Buff,
			Nocp: far.ApplyAction.Nocp,
			Dupl: far.ApplyAction.Dupl,
			Drop: far.ApplyAction.Drop,
		},
	}
	if far.BAR != nil {
		ufar.BARID = &pfcptypes.BARID{
			BarIdValue: far.BAR.BARID,
		}
	}

	if far.ForwardingParameters != nil {
		ufar.UpdateForwardingParameters = &pfcpmsg.UpdateForwardingParametersIEInFAR{
			DestinationInterface: &far.ForwardingParameters.DestinationInterface,
			NetworkInstance:      far.ForwardingParameters.NetworkInstance,
			OuterHeaderCreation:  far.ForwardingParameters.OuterHeaderCreation,
			PFCPSMReqFlags: &pfcptypes.PFCPSMReqFlags{
				Sndem: far.ForwardingParameters.SendEndMarker,
			},
		}
		if far.ForwardingParameters.ForwardingPolicyID != "" {
			ufar.UpdateForwardingParameters.ForwardingPolicy = &pfcptypes.ForwardingPolicy{
				ForwardingPolicyIdentifierLength: uint8(len(far.ForwardingParameters.ForwardingPolicyID)),
				ForwardingPolicyIdentifier:       []byte(far.ForwardingParameters.ForwardingPolicyID),
			}
		}
	}

	return ufar
}

func (far *FAR) toRemoveFar() *pfcpmsg.RemoveFAR {

	return &pfcpmsg.RemoveFAR{
		FARID: &pfcptypes.FARID{
			FarIdValue: far.FARID,
		},
	}
}

// Forwarding Parameters.
type ForwardingParameters struct {
	DestinationInterface pfcptypes.DestinationInterface
	NetworkInstance      *pfcptypes.NetworkInstance
	OuterHeaderCreation  *pfcptypes.OuterHeaderCreation
	ForwardingPolicyID   string
	SendEndMarker        bool
}

// Buffering Action Rule
type BAR struct {
	BARID uint8

	DownlinkDataNotificationDelay  pfcptypes.DownlinkDataNotificationDelay
	SuggestedBufferingPacketsCount pfcptypes.SuggestedBufferingPacketsCount

	State RuleState
}

func (bar *BAR) toCreateBar() *pfcpmsg.CreateBAR {
	return &pfcpmsg.CreateBAR{
		BARID: &pfcptypes.BARID{
			BarIdValue: bar.BARID,
		},
		//DownlinkDataNotificationDelay:  &pfcptypes.DownlinkDataNotificationDelay{},
		//SuggestedBufferingPacketsCount: &pfcptypes.SuggestedBufferingPacketsCount{},
	}
}

func (bar *BAR) toUpdateBar() *pfcpmsg.UpdateBARPFCPSessionModificationRequest {
	return nil
}

func (bar *BAR) toRemoveBar() *pfcpmsg.RemoveBAR {
	return &pfcpmsg.RemoveBAR{
		BARID: &pfcptypes.BARID{
			BarIdValue: bar.BARID,
		},
	}
}

// QoS Enhancement Rule
type QER struct {
	QERID uint32

	QFI pfcptypes.QFI

	GateStatus *pfcptypes.GateStatus
	MBR        *pfcptypes.MBR
	GBR        *pfcptypes.GBR

	State RuleState
}

func (qer *QER) toCreateQer() *pfcpmsg.CreateQER {
	return &pfcpmsg.CreateQER{
		QERID: &pfcptypes.QERID{
			QERID: qer.QERID,
		},
		GateStatus:        qer.GateStatus,
		QoSFlowIdentifier: &qer.QFI,
		MaximumBitrate:    qer.MBR,
		GuaranteedBitrate: qer.GBR,
	}
}

func (qer *QER) toUpdateQer() *pfcpmsg.UpdateQER {
	return nil
}

func (qer *QER) toRemoveQer() *pfcpmsg.RemoveQER {
	return &pfcpmsg.RemoveQER{
		QERID: &pfcptypes.QERID{
			QERID: qer.QERID,
		},
	}
}

// Usage Report Rule
type URR struct {
	URRID uint32
	//TODO: Add more atributes
	State RuleState
}

func (urr *URR) toCreateUrr() *pfcpmsg.CreateURR {
	return &pfcpmsg.CreateURR{
		URRID: &pfcptypes.URRID{
			UrrIdValue: urr.URRID,
		},
	}
}

func (urr *URR) toUpdateUrr() *pfcpmsg.UpdateURR {
	return &pfcpmsg.UpdateURR{
		URRID: &pfcptypes.URRID{
			UrrIdValue: urr.URRID,
		},
		//add more attributes
	}
}

func (urr *URR) toRemoveUrr() *pfcpmsg.RemoveURR {
	return &pfcpmsg.RemoveURR{
		URRID: &pfcptypes.URRID{
			UrrIdValue: urr.URRID,
		},
	}
}
