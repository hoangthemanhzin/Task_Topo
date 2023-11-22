package sm

import (
	"etrib5gc/logctx"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
)

type Pco struct {
	dnsipv4     bool
	dnsipv6     bool
	pcscf       bool
	ipv4linkmtu bool
}

// is any flag set?
func (pco *Pco) isSet() bool {
	return pco.dnsipv4 || pco.dnsipv6 || pco.pcscf || pco.ipv4linkmtu
}

func (pco *Pco) toNas(smcontext *SmContext) (nasbuf []byte) {
	log := smcontext.LogWriter
	options := nasConvert.NewProtocolConfigurationOptions()

	// IPv4 DNS
	if pco.dnsipv4 {
		if err := options.AddDNSServerIPv4Address(smcontext.tunnel.DnsIpv4Addr()); err != nil {
			log.Warnf("Error while adding DNS IPv4 Addr: ", err)
		}
	}

	// IPv6 DNS
	if pco.dnsipv6 {
		if err := options.AddDNSServerIPv6Address(smcontext.tunnel.DnsIpv6Addr()); err != nil {
			log.Warnf("Error while adding DNS IPv6 Addr: ", err)
		}
	}

	// IPv4 PCSCF (need for ims DNN)
	if pco.pcscf {
		if err := options.AddPCSCFIPv4Address(smcontext.tunnel.PcscfIpv4Addr()); err != nil {
			log.Warnf("Error while adding PCSCF IPv4 Addr: ", err)
		}
	}

	// MTU
	if pco.ipv4linkmtu {
		err := options.AddIPv4LinkMTU(smcontext.tunnel.Ipv4LinkMtu())
		if err != nil {
			log.Warnf("Error while adding MTU: ", err)
		}
	}

	nasbuf = options.Marshal()
	return
}

func pcoFromNas(content []byte) *Pco {
	log := logctx.WithFields(logctx.Fields{"mod": "sm-pco"})
	pco := &Pco{}

	options := nasConvert.NewProtocolConfigurationOptions()
	if err := options.UnMarshal(content); err != nil {
		log.Errorf("Parsing PCO failed: %s", err.Error())
		return nil
	}
	log.Tracef("Protocol Configuration Options: %v", options)

	for _, container := range options.ProtocolOrContainerList {
		log.Tracef("Container ID: %d", container.ProtocolOrContainerID)
		log.Tracef("Container Length: %d", container.LengthOfContents)
		switch container.ProtocolOrContainerID {
		case nasMessage.PCSCFIPv6AddressRequestUL:
			log.Trace("Didn't Implement container type PCSCFIPv6AddressRequestUL")
		case nasMessage.IMCNSubsystemSignalingFlagUL:
			log.Trace("Didn't Implement container type IMCNSubsystemSignalingFlagUL")
		case nasMessage.DNSServerIPv6AddressRequestUL:
			pco.dnsipv6 = true
		case nasMessage.NotSupportedUL:
			log.Trace("Didn't Implement container type NotSupportedUL")
		case nasMessage.MSSupportOfNetworkRequestedBearerControlIndicatorUL:
			log.Trace("Didn't Implement container type MSSupportOfNetworkRequestedBearerControlIndicatorUL")
		case nasMessage.DSMIPv6HomeAgentAddressRequestUL:
			log.Trace("Didn't Implement container type DSMIPv6HomeAgentAddressRequestUL")
		case nasMessage.DSMIPv6HomeNetworkPrefixRequestUL:
			log.Trace("Didn't Implement container type DSMIPv6HomeNetworkPrefixRequestUL")
		case nasMessage.DSMIPv6IPv4HomeAgentAddressRequestUL:
			log.Trace("Didn't Implement container type DSMIPv6IPv4HomeAgentAddressRequestUL")
		case nasMessage.IPAddressAllocationViaNASSignallingUL:
			log.Trace("Didn't Implement container type IPAddressAllocationViaNASSignallingUL")
		case nasMessage.IPv4AddressAllocationViaDHCPv4UL:
			log.Trace("Didn't Implement container type IPv4AddressAllocationViaDHCPv4UL")
		case nasMessage.PCSCFIPv4AddressRequestUL:
			pco.pcscf = true
		case nasMessage.DNSServerIPv4AddressRequestUL:
			pco.dnsipv6 = true
		case nasMessage.MSISDNRequestUL:
			log.Trace("Didn't Implement container type MSISDNRequestUL")
		case nasMessage.IFOMSupportRequestUL:
			log.Trace("Didn't Implement container type IFOMSupportRequestUL")
		case nasMessage.IPv4LinkMTURequestUL:
			pco.ipv4linkmtu = true
		case nasMessage.MSSupportOfLocalAddressInTFTIndicatorUL:
			log.Trace("Didn't Implement container type MSSupportOfLocalAddressInTFTIndicatorUL")
		case nasMessage.PCSCFReSelectionSupportUL:
			log.Trace("Didn't Implement container type PCSCFReSelectionSupportUL")
		case nasMessage.NBIFOMRequestIndicatorUL:
			log.Trace("Didn't Implement container type NBIFOMRequestIndicatorUL")
		case nasMessage.NBIFOMModeUL:
			log.Trace("Didn't Implement container type NBIFOMModeUL")
		case nasMessage.NonIPLinkMTURequestUL:
			log.Trace("Didn't Implement container type NonIPLinkMTURequestUL")
		case nasMessage.APNRateControlSupportIndicatorUL:
			log.Trace("Didn't Implement container type APNRateControlSupportIndicatorUL")
		case nasMessage.UEStatus3GPPPSDataOffUL:
			log.Trace("Didn't Implement container type UEStatus3GPPPSDataOffUL")
		case nasMessage.ReliableDataServiceRequestIndicatorUL:
			log.Trace("Didn't Implement container type ReliableDataServiceRequestIndicatorUL")
		case nasMessage.AdditionalAPNRateControlForExceptionDataSupportIndicatorUL:
			log.Trace("Didn't Implement container type AdditionalAPNRateControlForExceptionDataSupportIndicatorUL")
		case nasMessage.PDUSessionIDUL:
			log.Trace("Didn't Implement container type PDUSessionIDUL")
		case nasMessage.EthernetFramePayloadMTURequestUL:
			log.Trace("Didn't Implement container type EthernetFramePayloadMTURequestUL")
		case nasMessage.UnstructuredLinkMTURequestUL:
			log.Trace("Didn't Implement container type UnstructuredLinkMTURequestUL")
		case nasMessage.I5GSMCauseValueUL:
			log.Trace("Didn't Implement container type 5GSMCauseValueUL")
		case nasMessage.QoSRulesWithTheLengthOfTwoOctetsSupportIndicatorUL:
			log.Trace("Didn't Implement container type QoSRulesWithTheLengthOfTwoOctetsSupportIndicatorUL")
		case nasMessage.QoSFlowDescriptionsWithTheLengthOfTwoOctetsSupportIndicatorUL:
			log.Trace("Didn't Implement container type QoSFlowDescriptionsWithTheLengthOfTwoOctetsSupportIndicatorUL")
		case nasMessage.LinkControlProtocolUL:
			log.Trace("Didn't Implement container type LinkControlProtocolUL")
		case nasMessage.PushAccessControlProtocolUL:
			log.Trace("Didn't Implement container type PushAccessControlProtocolUL")
		case nasMessage.ChallengeHandshakeAuthenticationProtocolUL:
			log.Trace("Didn't Implement container type ChallengeHandshakeAuthenticationProtocolUL")
		case nasMessage.InternetProtocolControlProtocolUL:
			log.Trace("Didn't Implement container type InternetProtocolControlProtocolUL")
		default:
			log.Trace("Unknown Container ID [%d]", container.ProtocolOrContainerID)
		}
	}
	if pco.isSet() {
		log.Trace("PCO is set")
		return pco
	}
	return nil
}
