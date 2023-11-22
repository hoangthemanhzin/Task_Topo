package producer

import (
	"etrib5gc/sbi/models"
)

func (p *Producer) SDM_HandleGetAmData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp *models.AccessAndMobilitySubscriptionData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	log.Info("SDM_HandleGetAmData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSupiOrGpsi(ueId string, supportedFeatures string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (rsp *models.IdTranslationResult, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SDM_HandleGetSupiOrGpsi has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmfSelData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp *models.SmfSelectionSubscriptionData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	log.Info("SDM_HandleGetSmfSelData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmData(supi string, supportedFeatures string, singleNssai *models.Snssai, dnn string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp []models.SessionManagementSubscriptionData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SDM_HandleGetSmData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInAmfData(supi string, supportedFeatures string) (rsp *models.UeContextInAmfData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SDM_HandleGetUeCtxInAmfData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInSmfData(supi string, supportedFeatures string) (rsp *models.UeContextInSmfData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	log.Info("SDM_HandleGetUeCtxInSmfData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInSmsfData(supi string, supportedFeatures string) (rsp *models.UeContextInSmsfData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SDM_HandleGetUeCtxInSmsfData has not been implemented")
	return
}
