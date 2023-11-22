package ngap

import (
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
)

func sbiCritDiag(critical *ngapType.CriticalityDiagnostics) *n2models.CritDiag {
	//TODO
	return nil
}

func locConvert(ngaploc *ngapType.UserLocationInformation) *models.UserLocation {
	//TO DO: to be implemente
	return nil
}
func recRanNodeListConvert(list []ngapType.RecommendedRANNodeItem) (rannodes []n2models.RecRanNode) {
	//TO DO: to be implemente
	return
}
func recCellListConvert(list []ngapType.RecommendedCellItem) (cells []n2models.RecCell) {
	//TO DO: to be implemente
	return
}

func causeConvert(cause *ngapType.Cause) (n2cause n2models.Cause) {
	if cause == nil {
		n2cause.Value = 0
		return
	}
	n2cause.Present = uint8(cause.Present)
	switch cause.Present {
	case ngapType.CausePresentRadioNetwork:
		n2cause.Value = uint8(cause.RadioNetwork.Value)
	case ngapType.CausePresentTransport:
		n2cause.Value = uint8(cause.Transport.Value)
	case ngapType.CausePresentNas:
		n2cause.Value = uint8(cause.Nas.Value)
	case ngapType.CausePresentProtocol:
		n2cause.Value = uint8(cause.Protocol.Value)
	case ngapType.CausePresentMisc:
		n2cause.Value = uint8(cause.Misc.Value)
	default:
		n2cause.Value = 0
	}
	return
}
func dummy(dat interface{}) {
	log.Info("a dummy function")
}

func printCriticalityDiagnostics(ran *ran.Ran, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	if criticalityDiagnostics.ProcedureCriticality != nil {
		switch criticalityDiagnostics.ProcedureCriticality.Value {
		case ngapType.CriticalityPresentReject:
			log.Info("Procedure Criticality: Reject")
		case ngapType.CriticalityPresentIgnore:
			log.Info("Procedure Criticality: Ignore")
		case ngapType.CriticalityPresentNotify:
			log.Info("Procedure Criticality: Notify")
		}
	}

	if criticalityDiagnostics.IEsCriticalityDiagnostics != nil {
		for _, ieCriticalityDiagnostics := range criticalityDiagnostics.IEsCriticalityDiagnostics.List {
			log.Infof("IE ID: %d", ieCriticalityDiagnostics.IEID.Value)

			switch ieCriticalityDiagnostics.IECriticality.Value {
			case ngapType.CriticalityPresentReject:
				log.Info("Criticality Reject")
			case ngapType.CriticalityPresentNotify:
				log.Info("Criticality Notify")
			}

			switch ieCriticalityDiagnostics.TypeOfError.Value {
			case ngapType.TypeOfErrorPresentNotUnderstood:
				log.Info("Type of error: Not understood")
			case ngapType.TypeOfErrorPresentMissing:
				log.Info("Type of error: Missing")
			}
		}
	}
}

func buildCriticalityDiagnostics(
	procedureCode *int64,
	triggeringMessage *aper.Enumerated,
	procedureCriticality *aper.Enumerated,
	iesCriticalityDiagnostics *ngapType.CriticalityDiagnosticsIEList) (
	criticalityDiagnostics ngapType.CriticalityDiagnostics) {
	if procedureCode != nil {
		criticalityDiagnostics.ProcedureCode = new(ngapType.ProcedureCode)
		criticalityDiagnostics.ProcedureCode.Value = *procedureCode
	}

	if triggeringMessage != nil {
		criticalityDiagnostics.TriggeringMessage = new(ngapType.TriggeringMessage)
		criticalityDiagnostics.TriggeringMessage.Value = *triggeringMessage
	}

	if procedureCriticality != nil {
		criticalityDiagnostics.ProcedureCriticality = new(ngapType.Criticality)
		criticalityDiagnostics.ProcedureCriticality.Value = *procedureCriticality
	}

	if iesCriticalityDiagnostics != nil {
		criticalityDiagnostics.IEsCriticalityDiagnostics = iesCriticalityDiagnostics
	}

	return criticalityDiagnostics
}

func buildCriticalityDiagnosticsIEItem(ieCriticality aper.Enumerated, ieID int64, typeOfErr aper.Enumerated) (
	item ngapType.CriticalityDiagnosticsIEItem) {
	item = ngapType.CriticalityDiagnosticsIEItem{
		IECriticality: ngapType.Criticality{
			Value: ieCriticality,
		},
		IEID: ngapType.ProtocolIEID{
			Value: ieID,
		},
		TypeOfError: ngapType.TypeOfError{
			Value: typeOfErr,
		},
	}

	return item
}
func printAndGetCause(ran *ran.Ran, cause *ngapType.Cause) (present int, value aper.Enumerated) {
	present = cause.Present
	switch cause.Present {
	case ngapType.CausePresentRadioNetwork:
		log.Warnf("Cause RadioNetwork[%d]", cause.RadioNetwork.Value)
		value = cause.RadioNetwork.Value
	case ngapType.CausePresentTransport:
		log.Warnf("Cause Transport[%d]", cause.Transport.Value)
		value = cause.Transport.Value
	case ngapType.CausePresentProtocol:
		log.Warnf("Cause Protocol[%d]", cause.Protocol.Value)
		value = cause.Protocol.Value
	case ngapType.CausePresentNas:
		log.Warnf("Cause Nas[%d]", cause.Nas.Value)
		value = cause.Nas.Value
	case ngapType.CausePresentMisc:
		log.Warnf("Cause Misc[%d]", cause.Misc.Value)
		value = cause.Misc.Value
	default:
		log.Errorf("Invalid Cause group[%d]", cause.Present)
	}
	return
}
