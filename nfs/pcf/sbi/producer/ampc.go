package producer

import (
	"etrib5gc/sbi/models"
)

func (p *Producer) AMPC_HandleCreateIndividualAMPolicyAssociation(body models.PolicyAssociationRequest) (result models.PolicyAssociation, prob *models.ProblemDetails) {
	//TODO: to be implemented
	log.Warnf("Receive a CreateIndividualAMPolicyAssociation from AMF for ue[SUPI=%s], return a dummy PolicyAssociation", body.Supi)
	return
}
func (p *Producer) AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId string) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("AMPC_HandleDeleteIndividualAMPolicyAssociation has not been implemented")
	return
}
func (p *Producer) AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId string) (result models.PolicyAssociation, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("AMPC_HandleReadIndividualAMPolicyAssociation has not been implemented")
	return
}
func (p *Producer) AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (result models.PolicyUpdate, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation has not been implemented")
	return
}
