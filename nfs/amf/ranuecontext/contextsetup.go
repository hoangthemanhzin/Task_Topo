package ranuecontext

import (
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/sessioncontext"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

type acceptanceReport struct {
	pdulist    []n2models.DlPduSessionResourceInfo //list of pdu session information to send to RAN
	naspdu     []byte                              //registration acceptance or service acceptance
	naspdulist [][]byte                            //pending n1 message
}

type pduSessionManagementReport struct {
	errpdulist []uint8                             //fail to establish pdu sessions
	errcauses  []uint8                             //and their causes
	reactlist  *[16]bool                           //list of reactivated pdu session
	statuslist *[16]bool                           //
	targetlist []int32                             //target pdu session ids in N1N2 messages
	pdulist    []n2models.DlPduSessionResourceInfo //list of pdu session information to send to RAN
	naspdulist [][]byte                            //list of pending n1 messages in n1n2messages that can not be included in n2 messages to RAN
}

func (ranue *RanUe) handleInitialRegistration(msg *nasMessage.RegistrationRequest) {

	var err error

	cause := nasMessage.Cause5GMMProtocolErrorUnspecified

	/*
		1. update security context (generate keys)
		//should have done it after sec mode establishment
	*/

	//2. Get subscribed Nssai from UDM
	if err = ranue.ue.GetSubscribedNssai(); err != nil {
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(cause)
		return
	}

	//3. Handle requested Nssai (querying NSSF, NRF if needs)
	if err = ranue.ue.HandleRequestedNssai(); err != nil {
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(cause)
		return
	}

	//4. Get Capability5GMM
	if msg.Capability5GMM != nil {
		ranue.cap5gmm = *msg.Capability5GMM
	} else {
		ranue.Error("UE has no Capability5GM")
		ranue.rejectRegistration(cause)
		return
	}

	//5. Warn of not supporting MICO
	if msg.MICOIndication != nil {
		ranue.Warnf("Receive MICO Indication[RAAI: %d], Not Supported", msg.MICOIndication.GetRAAI())
	}
	//6. Store last visted TAI
	ranue.ue.SetLastTai(msg.LastVisitedRegisteredTAI)

	//8. Set DRX
	ranue.setDrx(msg.RequestedDRXParameters)

	// second step

	//. communicate with UDM
	if err = ranue.ue.UpdateWithUdm(); err != nil {
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(cause)
		return
	}
	//. Create AM Policy Association (with PCF)
	if err = ranue.ue.CreateAmPolicyAssociation(); err != nil {
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(cause)
		return
	}

	//. Allocate registration area
	ranue.ue.AllocateRegArea()
	//. Assign LADN information
	ranue.ue.AssignLadnInfo()

	//14. send registration accept (may use ngap in Non-3GPP access)
	ranue.acceptRegistration(nil)

}

func (ranue *RanUe) handleMobilityAndPeriodicRegistrationUpdating(msg *nasMessage.RegistrationRequest) {
	cause := nasMessage.Cause5GMMProtocolErrorUnspecified
	var err error
	//1. Check UpdateType5GS
	if msg.UpdateType5GS != nil && msg.UpdateType5GS.GetNGRanRcu() == nasMessage.NGRanRadioCapabilityUpdateNeeded {
		ranue.ue.ResetUeRadioCap()
	}

	//2. Get subscribed Nssai from UDM
	if err = ranue.ue.GetSubscribedNssai(); err != nil {
		ranue.rejectRegistration(cause)
		return
	}

	//3. TODO: Handle requested Nssai (querying NSSF, NRF if needs)
	if err = ranue.ue.HandleRequestedNssai(); err != nil {
		ranue.rejectRegistration(cause)
		return
	}

	//4. Check Capability5GMM
	if msg.Capability5GMM != nil {
		ranue.cap5gmm = *msg.Capability5GMM
	} else {
		ranue.rejectRegistration(cause)
		return
	}

	//5. Store last visited Tai
	ranue.ue.SetLastTai(msg.LastVisitedRegisteredTAI)

	//6. Warn that MICO is not supported
	if msg.MICOIndication != nil {
		ranue.Warnf("Receive MICO Indication[RAAI: %d], Not Supported", msg.MICOIndication.GetRAAI())
	}

	//7. Set DRX
	ranue.setDrx(msg.RequestedDRXParameters)
	/*
		if ranue.regctx.amfchanged {
			if err = ranue.ue.UpdateWithUdm(); err != nil {
				ranue.sendEvent(RejectEvent, cause)
				return
			}
		}
	*/
	/*
		//TODO:
			12. notify pcf of location changed (if needs)

	*/
	var report pduSessionManagementReport

	//. send SmContextUpdate if there are pdu sessions with uplink data
	if msg.UplinkDataStatus != nil {
		ranue.handleUplinkDataStatus(msg.UplinkDataStatus.Buffer, &report, false)
	}

	//. ask SMF to release sessions which are not available in the Ue
	if msg.PDUSessionStatus != nil {
		ranue.handlePduSessionStatus(msg.PDUSessionStatus.Buffer, &report)
	}

	//. handling pending N1N2Message if any (psi in AllowedPDUSessionStatus )
	if msg.AllowedPDUSessionStatus != nil {
		ranue.processPendingN1N2(msg.AllowedPDUSessionStatus.Buffer, &report)
	} else {
		ranue.processPendingN1N2([]byte{}, &report)
	}

	//. Allocate registration area
	ranue.ue.AllocateRegArea()
	//. Assign LADN information
	ranue.ue.AssignLadnInfo()

	//14. send registration accept
	ranue.acceptRegistration(&report)

	//send pending downlink n1 message from a transfering N1N2Message
	for _, naspdu := range report.naspdulist {
		ranue.sendNas(naspdu)
	}

}

func (ranue *RanUe) handleService(msg *nasMessage.ServiceRequest) {
	cause := nasMessage.Cause5GMMProtocolErrorUnspecified //default cause for rejection
	stype := msg.GetServiceTypeValue()
	var report pduSessionManagementReport

	//. check service type, send accept is is a service signaling
	if stype == nasMessage.ServiceTypeEmergencyServices || stype == nasMessage.ServiceTypeEmergencyServicesFallback {
		ranue.Warnf("emergency service is not supported")
	}

	if stype == nasMessage.ServiceTypeSignalling {
		ranue.acceptService(&report)
		return
	}

	//. handle uplink data status
	if msg.UplinkDataStatus != nil {
		skip := stype == nasMessage.ServiceTypeMobileTerminatedServices
		ranue.handleUplinkDataStatus(msg.UplinkDataStatus.Buffer, &report, skip)
	}

	//handle PDU session status
	if msg.PDUSessionStatus != nil {
		ranue.handlePduSessionStatus(msg.PDUSessionStatus.Buffer, &report)
	}

	switch stype {
	case nasMessage.ServiceTypeMobileTerminatedServices: // Trigger by Network
		if msg.AllowedPDUSessionStatus != nil {
			ranue.processPendingN1N2(msg.AllowedPDUSessionStatus.Buffer, &report)
		} else {
			ranue.processPendingN1N2([]byte{}, &report)
		}

	case nasMessage.ServiceTypeData:
		if !ranue.ue.IsReEstablishPduSessionAllowed() {
			ranue.rejectService(cause, report.statuslist, false)
		}

	default:
		ranue.Errorf("Service Type[%d] is not supported", stype)
		ranue.rejectService(cause, report.statuslist, false)
		return
	}
	ranue.acceptService(&report)

	//send pending downlink n1 message from a transfering N1N2Message
	for _, naspdu := range report.naspdulist {
		ranue.sendNas(naspdu)
	}
}

func (ranue *RanUe) setDrx(params *nasType.RequestedDRXParameters) {
	if params != nil {
		switch params.GetDRXValue() {
		case nasMessage.DRXcycleParameterT32:
			ranue.drx = nasMessage.DRXcycleParameterT32
		case nasMessage.DRXcycleParameterT64:
			ranue.drx = nasMessage.DRXcycleParameterT64
		case nasMessage.DRXcycleParameterT128:
			ranue.drx = nasMessage.DRXcycleParameterT128
		case nasMessage.DRXcycleParameterT256:
			ranue.drx = nasMessage.DRXcycleParameterT256
		case nasMessage.DRXValueNotSpecified:
			fallthrough
		default:
			ranue.drx = nasMessage.DRXValueNotSpecified
		}
	}
}

// TODO: we should go to Registered state immediately, then try to send
// registration acceptance  a few time
func (ranue *RanUe) acceptRegistration(report *pduSessionManagementReport) {
	if report == nil {
		//report can be nil if it is an initial registration
		//create a dummy report with default attributes
		report = &pduSessionManagementReport{}
	}

	if naspdu, err := nas.BuildRegistrationAccept(ranue, report.statuslist,
		report.reactlist, report.errpdulist, report.errcauses, ranue.fillRegistrationAccept); err == nil {
		ranue.acceptance = &acceptanceReport{
			pdulist:    report.pdulist,
			naspdu:     naspdu,
			naspdulist: report.naspdulist,
		}
		ranue.sendEvent(DoneEvent, nil)
	} else {
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(0)
	}
}
func (ranue *RanUe) sendAcceptance4Service() {
	var err error
	if ranue.regctx.IsRequestContext() {
		err = ranue.sendInitialContextSetupRequest(ranue.acceptance.pdulist, ranue.acceptance.naspdu)
	} else {
		if len(ranue.acceptance.pdulist) > 0 {
			err = ranue.sendPduSessionResourceSetupRequest(ranue.acceptance.pdulist, ranue.acceptance.naspdu)
		} else {
			//move to registered state
			err = ranue.sendNas(ranue.acceptance.naspdu)
		}
	}
	if err != nil {
		ranue.Errorf(err.Error())
		//NOTE: don't reject, try a few times to send the acceptance util
		//the registration time is expired
	}

}
func (ranue *RanUe) sendAcceptance4Registration() {

	var err error
	//send registration accept accordingly to the type of RAN and the
	//"Context Request" indicator received in the InitialUeMessage
	if ranue.regctx.IsRequestContext() {
		//TODO: generate keys for RAN
		//need to send InitialContextRequest
		if ranue.access == models.ACCESSTYPE__3_GPP_ACCESS {
			ranue.t3550.Start()
			err = ranue.sendInitialContextSetupRequest(ranue.acceptance.pdulist, ranue.acceptance.naspdu)
		} else {
			err = ranue.sendInitialContextSetupRequest(ranue.acceptance.pdulist, nil)
			//save for sending the RegistrationAccept message later
			//t3550 should starts later
			//ranue.pendingnaspdu = naspdu
		}
		ranue.contextsent = true
	} else {
		if len(ranue.acceptance.pdulist) > 0 {
			err = ranue.sendPduSessionResourceSetupRequest(ranue.acceptance.pdulist, ranue.acceptance.naspdu)
		} else {
			//move to registered state
			err = ranue.sendNas(ranue.acceptance.naspdu)
			ranue.logSendingReport("RegistrationAccept", err)
		}
	}

	if err != nil {
		//ranue.Errorf(err.Error())
		//NOTE: don't reject, try a few times to send the acceptance util
		//the registration time is expired
	}

}

func (ranue *RanUe) acceptService(report *pduSessionManagementReport) {
	//report must not be nil
	if naspdu, err := nas.BuildServiceAccept(ranue, report.statuslist,
		report.reactlist, report.errpdulist, report.errcauses); err == nil {
		ranue.acceptance = &acceptanceReport{
			pdulist:    report.pdulist,
			naspdu:     naspdu,
			naspdulist: report.naspdulist,
		}

	} else {
		ranue.Errorf("Build ServiceAccept failed: %s", err.Error())
		cause := nasMessage.Cause5GMMProtocolErrorUnspecified //default cause for rejection
		ranue.rejectService(cause, report.statuslist, false)
	}
}

// UplinkDataStatus (PDU sessionsto be activated, only associated with the
// related access). If service request is for signaling, the list should not
// be included by UE. Alway-on sessions should be in the list event without
// pending data
func (ranue *RanUe) handleUplinkDataStatus(content []byte, report *pduSessionManagementReport, skip bool) {
	pdumap := nasConvert.PSIToBooleanArray(content)
	ue := ranue.ue

	if !ue.IsReEstablishPduSessionAllowed() {
		for id, hasdat := range pdumap {
			if hasdat {
				report.errpdulist = append(report.errpdulist, uint8(id))
				report.errcauses = append(report.errcauses, nasMessage.Cause5GMMRestrictedServiceArea)
			}
		}
		return
	}
	//get list of PDU session ids from pending N1N2Messages
	var skipids []int32
	skipids = ranue.n1n2man.getTargetIds()

	report.reactlist = new([16]bool) //NOTE: should be created only if there is at least on activated pdu session
	for id, hasdat := range pdumap {
		if !hasdat {
			continue
		}
		//check if we can skip this pdu session (NOTE: are they already
		//activated at SMF?)
		if skip {
			found := false
			for _, skipid := range skipids {
				if skipid == int32(id) {
					found = true
					break
				}
			}
			if found {
				//no need to ask SMF for activation of this session
				continue
			}
		}
		//ask smf for activation
		if sc := findSession(ranue.ue, int32(id)); sc != nil && sc.Access() == models.ACCESSTYPE__3_GPP_ACCESS {
			rsp, ersp, err := sc.ActivateCnxState(ranue.access)
			updateReport(report, sc, rsp, ersp, err)
		}
	}
	return
}

// PDUSessionStatus indicates the sessions available in the UE (so basically
// any other sessions existing in the network should be released
func (ranue *RanUe) handlePduSessionStatus(content []byte, report *pduSessionManagementReport) {
	psilist := nasConvert.PSIToBooleanArray(content)
	report.statuslist = new([16]bool)
	for psi := 1; psi <= 15; psi++ {
		if sc := findSession(ranue.ue, int32(psi)); sc != nil {
			if !psilist[psi] && sc.Access() == ranue.Access() {
				cause := models.CAUSE_PDU_SESSION_STATUS_MISMATCH
				if _, err := sc.ReleaseSmContext(sessioncontext.Causes{Cause: &cause}, "", nil); err != nil {
					ranue.Errorf("Fail to release sc:%s", err.Error())
					//TODO: handle n1/n2 message in the response (if any)
					report.statuslist[psi] = true
				} else {
					report.statuslist[psi] = false //session is inactive now
				}
			} else {
				report.statuslist[psi] = true
			}
		}
	}
}

// content : AllowedPduSessionStatus - a list of PDU session provided by UE as a response
// to a paging/notification (non-3gpp) from network; identifies pdu session that can be
// transferred to 3GPP access (from non-3gpp)
func (ranue *RanUe) processPendingN1N2(content []byte, report *pduSessionManagementReport) {
	ranue.Info("Handling pending N1N2message transfer with AllowedPduSessionStatus element")
	for _, n1n2 := range ranue.n1n2man.pendinglist() {
		sc := n1n2.sc
		sid := sc.Id()

		reqinfo := n1n2.req.JsonData
		n1msg := n1n2.req.BinaryDataN1Message
		n2info := n1n2.req.BinaryDataN2Information

		var err error
		var naspdu []byte
		if len(n2info) == 0 {
			n1type := getN1Type(reqinfo.N1MessageContainer.N1MessageClass)
			if n1type == nasMessage.PayloadContainerTypeN1SMInfo {
				naspdu, err = nas.BuildDLNASTransport(ranue, n1type, n1msg, int32(sid), 0, nil, 0)
			} else {
				naspdu, err = nas.BuildDLNASTransport(ranue, n1type, n1msg, 0, 0, nil, 0)
			}

			if err == nil && len(naspdu) > 0 {
				report.naspdulist = append(report.naspdulist, naspdu)
			}
			if err != nil {
				ranue.Errorf("Fail to build a DLNASTransport: %s", err.Error())
			}
			continue
		}

		sminfo := reqinfo.N2InfoContainer.SmInfo
		//if the smcontext of the N1N2Message tranfer is in Non3Gpp and UE changes
		//its to 3GPP, notify the SMF of the change
		if sc.Access() == models.ACCESSTYPE_NON_3_GPP_ACCESS && len(content) > 0 {
			psilist := nasConvert.PSIToBooleanArray(content)
			if report.reactlist == nil {
				report.reactlist = new([16]bool)
			}
			if psilist[sid] {
				//notify Smf of changing access to 3GPP
				rsp, ersp, err := sc.ChangeAccess3gpp()
				updateReport(report, sc, rsp, ersp, err)
			} else {
				ranue.Warnf("UE was reachable but did not accept to re-activate the PDU Session[%d]", sid)
				//TODO : callback.SendN1N2TransferFailureNotification(ue,models.N1N2MessageTransferCause_UE_NOT_REACHABLE_FOR_SESSION)
			}
		}
		//add the pending N2Information for later sending to the gnB
		if sc.Access() == models.ACCESSTYPE__3_GPP_ACCESS && sminfo.N2InfoContent.NgapIeType == models.NGAPIETYPE_PDU_RES_SETUP_REQ {
			report.pdulist = append(report.pdulist, n2models.DlPduSessionResourceInfo{
				Id:       int64(sid),
				NasPdu:   naspdu,
				Transfer: n2info,
				Snssai:   sminfo.SNssai,
			})
		}
	}
	ranue.n1n2man.clean()
}

func updateReport(report *pduSessionManagementReport, sc *sessioncontext.SessionContext, rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sid := sc.Id()
	if rsp == nil {
		report.reactlist[sid] = true
		report.errpdulist = append(report.errpdulist, uint8(sid))
		cause := nasMessage.Cause5GMMProtocolErrorUnspecified
		if ersp != nil {
			switch ersp.JsonData.Error.Cause {
			case "OUT_OF_LADN_SERVICE_AREA":
				cause = nasMessage.Cause5GMMLADNNotAvailable
			case "PRIORITIZED_SERVICES_ONLY":
				cause = nasMessage.Cause5GMMRestrictedServiceArea
			case "DNN_CONGESTION", "S-NSSAI_CONGESTION":
				cause = nasMessage.Cause5GMMInsufficientUserPlaneResourcesForThePDUSession
			}
		}
		report.errcauses = append(report.errcauses, cause)
	} else {
		report.pdulist = append(report.pdulist, n2models.DlPduSessionResourceInfo{
			Id:       int64(sid),
			NasPdu:   rsp.BinaryDataN1SmMessage,
			Transfer: rsp.BinaryDataN2SmInformation,
			Snssai:   sc.Snssai(), //from smcontext or response?
		})
	}
}

/*
func (ranue *RanUe) updateOldAmf() (err error) {
	//TODO: send a registration status update to old AMF
		configuration := Namf_Communication.NewConfiguration()
		configuration.SetBasePath(ue.TargetAmfUri)
		client := Namf_Communication.NewAPIClient(configuration)

		ueContextId := fmt.Sprintf("5g-guti-%s", ue.Guti)
		res, httpResp, localErr :=
			client.IndividualUeContextDocumentApi.RegistrationStatusUpdate(context.TODO(), ueContextId, request)
		if localErr == nil {
			regStatusTransferComplete = res.RegStatusTransferComplete
		} else if httpResp != nil {
			if httpResp.Status != localErr.Error() {
				err = localErr
				return
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			problemDetails = &problem
		} else {
			err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
		}
		return
	return
}

*/
