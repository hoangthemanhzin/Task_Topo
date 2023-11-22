package sm

import (
	"etrib5gc/common"
	"fmt"
)

// receive external events for handling
func (smctx *SmContext) HandleEvent(ev *common.EventData) (err error) {
	switch ev.EvType {
	case POST_SMCONTEXTS:
		job, _ := ev.Content.(*common.AsyncJob)
		err = smctx.sendEvent(PostSmContextsEvent, job)
	case UPDATE_SMCONTEXT:
		job, _ := ev.Content.(*common.AsyncJob)
		err = smctx.sendEvent(UpdateSmContextEvent, job)
	case RELEASE_SMCONTEXT:
		job, _ := ev.Content.(*common.AsyncJob)
		err = smctx.sendEvent(ReleaseSmContextEvent, job)
	default:
		err = fmt.Errorf("Unknown event")
	}
	if err != nil {
		smctx.Errorf("HandleEvent failed: %s", err.Error())
	}
	return
}
