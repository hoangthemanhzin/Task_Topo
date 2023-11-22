package ngap

import (
	"etrib5gc/nfs/pran/ran"

	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleUplinkRanConfigurationTransfer(ran *ran.Ran, uplinkRANConfigurationTransfer *ngapType.UplinkRANConfigurationTransfer) {
	log.Warnf("Uplink Ran Configuration Transfer is not implemented")
}

func SendDownlinkRanConfigurationTransfer(ran *ran.Ran, transfer *ngapType.SONConfigurationTransfer) (err error) {
	log.Warnf("Sending Downlink Ran Configuration Transfer is not implemented")
	return
}
