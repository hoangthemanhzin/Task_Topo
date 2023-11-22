package uecontext

func (uectx *UeContext) CreateAmPolicyAssociation() (err error) {
	/*
		ue.Infof("Send an AmPolicyAssociation request to PCF")
		req := &models.PolicyAssociationRequest{
			AccessType: ranstate.Access(),
		}
		ranstate.ue.SetAmPolRequest(req)
		var ampol *models.PolicyAssociation
		if pcfcli := ranstate.ue.Pcf(); pcfcli == nil {
			err = fmt.Errorf("PCF not found")
		} else {
			if ampol, err = ampc.CreateIndividualAMPolicyAssociation(pcfcli, *req); err != nil {
				return
			}
			ranstate.ue.SetAmPol(ampol)
			//TODO: get the AmPolicyUri and AmPolicyAssociationId from the response
			//header. For the meantime we may change the PolicyAssociation data
			//structure to include them.
		}
	*/
	return
}

// allocate registration area for a RanState
func (uectx *UeContext) AllocateRegArea() {
	uectx.Warn("Allocate registration areas not implemented")
	//TODO:
	return
}

// asiggn LADN information for RanState
func (uectx *UeContext) AssignLadnInfo() {
	uectx.Warn("Assign ladn information not implemented")
	//TODO:
	return
}
