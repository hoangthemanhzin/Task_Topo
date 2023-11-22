package ngap

/*
func (r *Ngap) FindAmf(ue *context.UeContext) sbi.ConsumerClient {
	return nil
	//TODO: to be implemented
	//using provided information, try to locate the AMF to handle the UE
	//(by looking for an existing UeContext of all RANs then attach it to
	//the UeContext)
	//if no AMF is found, we should forward this request to a default AMF
	//(create a default UeContext)

	//This part is implemented by free5gc, we may need to do it
	//differently (tqtung)

	if fiveGSTMSI != nil {
		log.Debug("Receive 5G-S-TMSI")

		servedGuami := amf.ServedGuamiList()[0]

		// <5G-S-TMSI> := <AMF Set ID><AMF Pointer><5G-TMSI>
		// GUAMI := <MCC><MNC><AMF Region ID><AMF Set ID><AMF Pointer>
		// 5G-GUTI := <GUAMI><5G-TMSI>
		tmpReginID, _, _ := ngapConvert.AmfIdToNgap(servedGuami.AmfId)
		amfID := ngapConvert.AmfIdToModels(tmpReginID, fiveGSTMSI.AMFSetID.Value, fiveGSTMSI.AMFPointer.Value)
		tmsi := hex.EncodeToString(fiveGSTMSI.FiveGTMSI.Value)
		guti := servedGuami.PlmnId.Mcc + servedGuami.PlmnId.Mnc + amfID + tmsi

		// TODO: invoke Namf_Communication_UEContextTransfer if serving AMF has changed since
		// last Registration Request procedure
		// Described in TS 23.502 4.2.2.2.2 step 4 (without UDSF deployment)

		if amfUe, ok := amf.AmfUeFindByGuti(guti); !ok {
			log.Warnf("Unknown UE [GUTI: %s]", guti)
		} else {
			log.Tracef("find AmfUe [GUTI: %s]", guti)
			log.Debugf("AmfUe Attach UeContext [UeContextNgapID: %d]", ue.UeContextNgapId())
			amfUe.AttachUeContext(ue)
		}
	}
}
*/
