package group

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils"
	"net/http"
)

// sbi producer handler for GetNfGroupIDs
func OnGetNfGroupIDs(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	nfTypeStr := ctx.Param("nf-type")
	if len(nfTypeStr) == 0 {
		//nfType is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "nfType is required",
		})
		return
	}
	var nfType []models.NFType
	var nfTypeErr error
	if nfType, nfTypeErr = utils.String2ArrayOfNFType(nfTypeStr); nfTypeErr != nil {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: nfTypeErr.Error(),
		})
		return
	}

	subscriberId := ctx.Param("subscriberId")
	if len(subscriberId) == 0 {
		//subscriberId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriberId is required",
		})
		return
	}

	if rsp, prob := prod.GROUP_HandleGetNfGroupIDs(nfType, subscriberId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, rsp)
	}
	return
}

type Producer interface {
	GROUP_HandleGetNfGroupIDs(nfType []models.NFType, subscriberId string) (map[string]string, *models.ProblemDetails)
}
