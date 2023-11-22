#Cloud-native B5G core network design document

## Scope

The system developed in this project is to show a proof of concept how a B5G
core network can be constructed for deployment on cloud-native platform at
scale. This work is not to create a full-functionality of mobile core systems.
Instead, it will implemente core features to demonstrate scenarios that covers
the main components of the architecture.

##Overall architecure

The B5G architecure is a refinement of the original 5G core architecture with
the main focus is to create a separation between mobile core functionality and
the deployment logic of the network. The original idea for this design was born
from an obsevation that the current 5G core architecure is overly-complex for
deploying at scale. There are too many options for deployments that lead to
ambiguity in software implementation and deployment. There is a need to
simplify the architecture and its signaling procedures so that it is easier for
the developers to implement the functionalities of the system correctly while
not being interfered with deployment aspects. 



Going beyond 5G, the core
network will be cloud-native where network functions are deployed in
heterogeouns clouds (from on-prem, edge cloud, central cloud to public cloud).
The deployment and life-cycle management of the network functions will be
highly dynamics. In addition, the core network may interact with other
services from thire party or public. The core network can be a part of an
ecosystem of diverse mobile applications running on a communication platform.  
The new design for the core network should address these 
not add more complexity but to simplify by a better abstraction that make room
for the system to grow 
However, this
shares a same view with the service mesh concept in the cloud-native
technology.  

The network function layer consist of network functions that handle bussiness
logic of mobile core network. The bussines include UE authenticacation,
connection management, registration management, and session managements.
There are a few network components have been removed from the current 5G
architecture

###Network function identification

In order to separate business logic of network functions from deployments, it
is necessary to define a method to identify network functions. With an
identification scheme, it is clear for a consumer to request services from a
producer by using the identity of the producer as a target to locate the
services.  The current 5G architecture does not have a clear definition of
network function identification. For example, to identify an UDM to serve a UE,
the SUPI of the UE is used as a criteria to search for the UDM. The NRF organiz

###Network Funtcion – Fabric APIs

All network functions have a set of unified APIs to the fabric. From the
perspective of network functions, the fabric is abstracted through a Forwarder
class. The APIs are supported as methods of the class.

The Forwarder abstraction offers two methods for requesting a service:

type Forwarder interface {
	DiscoveryAndRequest(nfquery NfQuery, request SbiRequest) (add NfAddress,
	response SbiResponse, err error)
}

Request(request SbiRequest, addr NfAddress) (response SbiResponse, err error)

The first method should be called when a consumer request a service for the
first time. It should provide an identity of the producer throught an object of
NfQuery. The query object allows the fabric to disovery and select a suitable
producer to handler the request service. The second argument is the request
which is used to request the service. It is meaned to delivered to the selected
producer. The return of this method includes three values: an address of the
selected producer; a response for the requested service, and an error value to
indicate if the request is successfully served.

