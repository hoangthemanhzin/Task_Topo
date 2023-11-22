package sm

import (
	"etrib5gc/nas"
	amfcomm "etrib5gc/sbi/amf/comm"
	"etrib5gc/sbi/models"
)

// send pfcp session establishment requests to UPFs
func (smctx *SmContext) establishPfcpSessions() (err error) {
	smctx.Infof("Establish pfcp session with UPF(s)")
	if err = smctx.tunnel.Update(); err != nil {
		smctx.Errorf("Pfcp session establishment failed: %+v", err)
		return
	}
	return
}

func (smctx *SmContext) acceptPduSessionEstablishment() (err error) {
	n1n2req := models.N1N2MessageTransferRequest{}
	if n1n2req.BinaryDataN1Message, err = nas.BuildPduSessionEstablishmentAccept(smctx); err != nil {
		smctx.Errorf("Build GSM PDUSessionEstablishmentAccept failed: %s", err)
		return
	}
	smctx.Info("PduSessionEstablishmentAccept built")
	if n1n2req.BinaryDataN2Information, err = smctx.buildPduSessionResourceSetupRequestTransfer(); err != nil {
		smctx.Errorf("Build PDUSessionResourceSetupRequestTransfer failed: %s", err)
		return
	}
	smctx.Info("PDUSessionResourceSetupRequestTransfer built")

	n1n2req.JsonData = models.N1N2MessageTransferReqData{
		PduSessionId: int32(smctx.sid),
		N1MessageContainer: &models.N1MessageContainer{
			N1MessageClass:   "SM",
			N1MessageContent: models.RefToBinaryData{ContentId: "GSM_NAS"},
		},
		N2InfoContainer: &models.N2InfoContainer{
			N2InformationClass: models.N2INFORMATIONCLASS_SM,
			SmInfo: models.N2SmInformation{
				PduSessionId: int32(smctx.sid),
				N2InfoContent: models.N2InfoContent{
					NgapIeType: models.NGAPIETYPE_PDU_RES_SETUP_REQ,
					NgapData: models.RefToBinaryData{
						ContentId: "N2SmInformation",
					},
				},
				//SNssai: *smctx.snssai,
			},
		},
	}

	//send n1n2 here
	if smctx.amfcli == nil {
		smctx.Errorf("Unexpected Error: Empty AMF consumer for the session")
	}

	_, _, err = amfcomm.N1N2MessageTransfer(smctx.amfcli, smctx.Supi(), n1n2req)
	smctx.logSendingReport("N1N2MessageTransfer", err)
	//TODO: handle the response
	return
}

func (smctx *SmContext) rejectPduSessionEstablishment(n1cause uint8) (err error) {
	n1n2req := models.N1N2MessageTransferRequest{}
	if n1n2req.BinaryDataN1Message, err = nas.BuildPduSessionEstablishmentReject(smctx, n1cause); err != nil {
		smctx.Errorf("Build GSM PDUSessionEstablishmentReject failed: %+v", err)
		return
	}
	smctx.Info("PduSessionEstablishmentReject built")
	n1n2req.JsonData = models.N1N2MessageTransferReqData{
		PduSessionId: int32(smctx.sid),
		N1MessageContainer: &models.N1MessageContainer{
			N1MessageClass:   "SM",
			N1MessageContent: models.RefToBinaryData{ContentId: "GSM_NAS"},
		},
	}

	_, _, err = amfcomm.N1N2MessageTransfer(smctx.amfcli, smctx.Supi(), n1n2req)
	smctx.logSendingReport("N1N2MessageTransfer", err)
	//TODO: handle the response
	return
}
