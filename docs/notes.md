## AMF configuration:

- An Amf is identified with a PLMN identity and an AMF identity. The Amf
  identity is composed of Region, Set, and identity. The identity is just
telling who is operating this AMF, and it is the way to tell the location of
the AMF in the network operated by this operator. It does not mean that the
Amf is to handle the UE of the operator, that is the Amf can be configured to
handle roaming UE (of other operators).

- In free5gc, AMF configuration includes a supported tracking area list which
  is composed of TACs that connected gnB can support. In out design, the AMF
should not be configured with this list explitly. Instead, it is open to the
operators to define how TACs are organizing in AMF collection structuring. It
can be done by organizing the AMF using a mapping of the network slices and
TAcs to AMF identification (region, set, id). There should be a separated
network functions (aka NSSF) that holds the mapping and answer to gnB when they
needs to look for a righ AMF to handle a registration request. (That is the UE
and the gnB should provide the NFs with slices and Tacs).

- Supported slices:  the AMF should have a list of slices that it can support
  (and their mappings). A slice is identified with its identification and the
PlmnId where it is defined.


## PRAN:

- PRAN aka proxy gnB is supposed to be a CU in the open RAN architecture. It is a
part of SBA, thus NGAP connection to AMFs is deprecated.

- a PRAN is identitied by its Plmn identity and an unique identity within its
  Plmn.

- a PRAN should have a list of TAList


- For an UeContext existed in an AMF, the AMF needs to keep track of PRANs
  which are associated with this UEContext. In 5G, the AMF has NGAP (N2)
connections to the gnBs. In the new architecure where PRAN is also a part of
SBA, there should be a way to identify the PRAN, and the identification method
should be realized through registration and discovery mechanisms.

- When an AMF realize that it needs to page an UE, it look at the current
  location of the UE in its context then find the gnB (PRAN in our case) to
send a paging signal. In SBA, there is no ngap connection, thus CM-CONNECTED
and CM-IDLE states are not defined, only UE's location that matters. Basing on
the last location the AMF should predict a few gnB to page (using TAC
information)


## Service discovery and registration

### About heartbeat:

Should a consumer must manage the status of service instances that it wants to
request Or should the management is supported by the controller?


- In case the consumer manage the status, either it must send a heartbeat
  request or the  service instances have to send liveness notification. The
later approach clearly too complicate because the service instances have to
manage the consumer. So it is better that the consumer must be active in manage
the instances's liveness. 

- If the controller managed the liveness, then we may have two approaches: 1)
  the controller send heartbeat request to managed instances. 2) the managed
instance send liveness notifications. Both approaches seems to work well.
However, we may have the controller being active in this aspect becuase it may
be difficult to prevent a DDOS on the controller if it is to waited for
notification of liveness from service instances.

- So, coming here we have a conclusion that the service instances should be
  passive in liveness management. That is, either the consumer of the
controller should be sending heartbeat requests


The the next question is what kind of information
should be carried by a heartbeat response? There should be a timestamp (so that
the consumer may infer the latency. Surely there should be a time
synchronization mechanism for trustworthyness of the latency measurement). In
addition, a requested instance may send their 

- During registration an NF should send following information to the controller:

  + Service(s) that the NF offer, including service identity and an indication of its statefulness
  + A list of services that it wants to receive updates. The list consists of service identities
	
