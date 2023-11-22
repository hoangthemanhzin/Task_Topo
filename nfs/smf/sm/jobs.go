package sm

import (
	"etrib5gc/nas"
	"etrib5gc/sbi/models"
	"net/http"
)

const (
	POST_SMCONTEXTS uint8 = iota
	UPDATE_SMCONTEXT
	RELEASE_SMCONTEXT
)

type UpdateSmContextJob struct {
	Req  *models.UpdateSmContextRequest
	Rsp  *models.UpdateSmContextResponse
	Ersp *models.UpdateSmContextErrorResponse

	n1msgid    string              //identity of n1msg for responding
	n2infoid   string              //identity of n2info for responding
	n2infotype models.N2SmInfoType //type of n2info for responding
	n1msgbyte  []byte              //encoded n1msg for responding
	n2info     []byte              //n2info for responding (either normal or error)
	err        *models.ExtProblemDetails
	UpCnxState models.UpCnxState
	rulechange bool
}

func (j *UpdateSmContextJob) buildResponse() {
	if j.err != nil {
		j.Ersp = &models.UpdateSmContextErrorResponse{
			BinaryDataN1SmMessage:     j.n1msgbyte,
			BinaryDataN2SmInformation: j.n2info,
			JsonData: models.SmContextUpdateError{
				Error: j.err,
				N1SmMsg: models.RefToBinaryData{
					ContentId: j.n1msgid,
				},
				N2SmInfo: models.RefToBinaryData{
					ContentId: j.n2infoid,
				},
				N2SmInfoType: j.n2infotype,
			},
		}
	} else { //make sure there must be a n1msg or a n2info
		j.Rsp = &models.UpdateSmContextResponse{
			BinaryDataN1SmMessage:     j.n1msgbyte,
			BinaryDataN2SmInformation: j.n2info,
			JsonData: models.SmContextUpdatedData{
				N1SmMsg: models.RefToBinaryData{
					ContentId: j.n1msgid,
				},
				N2SmInfo: models.RefToBinaryData{
					ContentId: j.n2infoid,
				},
				N2SmInfoType: j.n2infotype,
			},
		}
	}
}

type PostSmContextsJob struct {
	Callback *models.Callback
	Req      *models.PostSmContextsRequest
	Rsp      *models.PostSmContextsResponse
	Ersp     *models.PostSmContextsErrorResponse
}

func (j *PostSmContextsJob) setN1Error(err error, n1cause uint8, smctx *SmContext) {
	j.Rsp = nil
	if n1msg, newerr := nas.BuildPduSessionEstablishmentReject(smctx, n1cause); newerr == nil {
		smctx.Infof("Build PduSessionEstablishmentReject")
		j.Ersp = &models.PostSmContextsErrorResponse{
			JsonData: models.SmContextCreateError{
				Error: &models.ExtProblemDetails{
					Status: http.StatusInternalServerError,
					Detail: err.Error(),
					Type:   "N1SmError",
				},
				N1SmMsg: models.RefToBinaryData{
					ContentId: "N1SmMsg",
				},
			},
			BinaryDataN1SmMessage: n1msg,
		}
	} else {
		smctx.Infof("Build PduSessionEstablishmentReject failed: %+v", newerr)
		j.Ersp = &models.PostSmContextsErrorResponse{
			JsonData: models.SmContextCreateError{
				Error: &models.ExtProblemDetails{
					Status: http.StatusInternalServerError,
					Detail: newerr.Error(),
					Type:   "N1SmMsg encoding error",
				},
			},
		}

	}
}

type ReleaseSmContextJob struct {
}
