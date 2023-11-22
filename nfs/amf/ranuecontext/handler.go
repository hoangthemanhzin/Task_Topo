package ranuecontext

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/sbi/models/n2models"
	"fmt"

	libnas "github.com/free5gc/nas"
)

// receive external events for handling
func (ranue *RanUe) HandleEvent(ev *common.EventData) (err error) {
	switch ev.EvType {
	case events.INIT_UE_CONTEXT:
		err = ranue.sendEvent(InitEvent, ev.Content)
	case events.NAS_NON_DELIVERY:
		err = ranue.sendEvent(NasNonDeliveryEvent, ev.Content)
	case events.NAS_UL_TRANSPORT:
		msg, _ := ev.Content.(*n2models.UlNasTransport)
		var nasMsg libnas.Message
		if nasMsg, err = nas.Decode(ranue, msg.NasPdu); err != nil {
			ranue.Errorf("Decode ULTransportNas message failed: %s", err.Error())
			return
		}
		ranue.ue.SetLocation(msg.Loc)
		err = ranue.sendEvent(N1MsgEvent, nasMsg.GmmMessage)
	case events.UECTX_REL_REQ:
		err = ranue.sendEvent(UectxRelReqEvent, ev.Content)
	case events.PDU_NOTIFY:
		err = ranue.sendEvent(PduNotifyEvent, ev.Content)
	case events.PDU_MOD_IND:
		err = ranue.sendEvent(PduModIndEvent, ev.Content)

	case events.SEND_N1N2_TRANSFER: //from UeContext
		err = ranue.sendEvent(N1N2TransferEvent, ev.Content)

	case events.SEND_NOTIFICATION: //from UeContext
		err = ranue.sendEvent(NotificationEvent, ev.Content)

	case events.SEND_PAGING: //from UeContext
		err = ranue.sendEvent(PagingEvent, ev.Content)

	case events.REGISTRATION_GRANTED: //from UeContext
		err = ranue.sendEvent(RegistrationRequestEvent, ev.Content)

	case events.REGISTRATION_REJECTED: //from UeContext
		regctx, _ := ev.Content.(*events.RegistrationContext)
		if regctx.RegistrationRequest() != nil {
			//TODO: send a registration reject
			ranue.Infof("Registration procedure pending")
		} else {
			//TODO: send a service reject
			ranue.Infof("Service procedure pending")
		}
	default:
		err = fmt.Errorf("Unknown event")
	}
	if err != nil {
		err = common.WrapError("HandleEvent failed", err)
		ranue.Error(err.Error())
	}
	return
}
