package ngapConvert

import (
	"encoding/hex"
	"strings"

	"etrib5gc/sbi/models"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/logger"
	"github.com/free5gc/ngap/ngapType"
)

func TraceDataToModels(traceActivation ngapType.TraceActivation) (traceData models.TraceData) {
	// TODO: finish this function when need
	return
}

func TraceDataToNgap(traceData models.TraceData, trsr string) ngapType.TraceActivation {
	var traceActivation ngapType.TraceActivation

	if len(trsr) != 4 {
		logger.NgapLog.Warningln("Trace Recording Session Reference should be 2 octets")
		return traceActivation
	}

	// NG-RAN Trace ID (left most 6 octet Trace Reference + last 2 octet Trace Recoding Session Reference)
	subStringSlice := strings.Split(traceData.TraceRef, "-")

	if len(subStringSlice) != 2 {
		logger.NgapLog.Warningln("TraceRef format is not correct")
		return traceActivation
	}

	plmnID := models.PlmnId{}
	plmnID.Mcc = subStringSlice[0][:3]
	plmnID.Mnc = subStringSlice[0][3:]
	var traceID []byte
	if traceIDTmp, err := hex.DecodeString(subStringSlice[1]); err != nil {
		logger.NgapLog.Warnf("")
	} else {
		traceID = traceIDTmp
	}

	tmp := PlmnIdToNgap(plmnID)
	traceReference := append(tmp.Value, traceID...)
	var trsrNgap []byte
	if trsrNgapTmp, err := hex.DecodeString(trsr); err != nil {
		logger.NgapLog.Warnf("Decode trsr failed: %+v", err)
	} else {
		trsrNgap = trsrNgapTmp
	}

	nGRANTraceID := append(traceReference, trsrNgap...)

	traceActivation.NGRANTraceID.Value = nGRANTraceID

	// Interfaces To Trace
	var interfacesToTrace []byte
	if interfacesToTraceTmp, err := hex.DecodeString(traceData.InterfaceList); err != nil {
		logger.NgapLog.Warnf("Decode Interface failed: %+v", err)
	} else {
		interfacesToTrace = interfacesToTraceTmp
	}
	traceActivation.InterfacesToTrace.Value = aper.BitString{
		Bytes:     interfacesToTrace,
		BitLength: 8,
	}

	// Trace Collection Entity IP Address
	ngapIP := IPAddressToNgap(traceData.CollectionEntityIpv4Addr, string(traceData.CollectionEntityIpv6Addr))
	traceActivation.TraceCollectionEntityIPAddress = ngapIP

	// Trace Depth
	switch traceData.TraceDepth {
	case models.TRACEDEPTH_MINIMUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMinimum
	case models.TRACEDEPTH_MEDIUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMedium
	case models.TRACEDEPTH_MAXIMUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMaximum
	case models.TRACEDEPTH_MINIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMinimumWithoutVendorSpecificExtension
	case models.TRACEDEPTH_MEDIUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMediumWithoutVendorSpecificExtension
	case models.TRACEDEPTH_MAXIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMaximumWithoutVendorSpecificExtension
	}

	return traceActivation
}
