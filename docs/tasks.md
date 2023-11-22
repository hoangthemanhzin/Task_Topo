## UDR implementation, generating UE configuration and UDR's UE subscriber data

 - Current status: for testing the system , we use a single UE profile + UE
   subsriber data is hard-coded into UDM (only authorization information).

 - UE subscribe data has three parts (as I recalled): authorization
   information (for AMF); session management data (for SMF); and policy data
(for PCF). Currently we only use the authorization data (needed by AMF), but we
will need other data as well in other testing scenarios.

 - UE configuration: authorization information (SUCI, encryption scheme). You
   will need to understand how SUCI is constructed/verify by the network (check 
UDM source code) in order to generate the data.

 - UDR skeleton code has been implemented. You will need to implement methods
   to expose API to UDM for retrieving relevant information; add a backend
database (mongodb) to store the UE's subscribe data.

 - develop a commandline tool to generate UE configuration files (for UERANSIM)
   and inject UE's subscription data into the UDR's database. (for supporting
large-scale experiments in the future)

 - you may expose an web interface at UDR for create UE profile (adding UE to
   the database and generating UE configuration as the same time) for
testing conviniency.


## Kubenetes deployment

 - Contenerization: there are many tutorial on how to build a docker image for
   go-lang application. Try to keep application images as small-sized as possible.

 - Create a new project that build docker images for all network functions in
   etrib5gc. Add UERANSIM (UE+GnB); UPMF; modified free5gc's UPF. Refer to the
[free5gc docker compose repository](https://github.com/free5gc/free5gc-compose)
, you can reuse their script. A single
command execution should build all the images at once. 
 - Setting up the system without Kubernetes, make sure it works in simple test
   (UE can register and establish a PDU session to UPF)

 - Now move this setting into Kubernetes: write Kubenetes deployment manifest
   for all NFs, use Helm charts for automation. You will need to go through
several Kubernetes deployment tutorials and Helm charts tutorial. There are
many Youtube tutorial videos also. Make sure you can re-produce the previous
test on K8s-based system.


## UPFM integration

 - Fork a new git branch from etrb5gc repo for integration

 - Replace Pfcp at SMF anh UPF (free5gc)
 
 - Add logic for handling multiple UPFs at SMF

 - Add UPMF to the etr5gc repo 

 - test the multiple-UPFs senario



## Some small tasks

 - Implement supi, suci generation methods (needed for generating UE profiles).
   package etrib5gc/utils/suci has implemented the suci recovery (from supi)
method. In this same package you should implement two methods:
GenerateSupi(plmnid) to return a supi; and Supi2Suci(Profile) that return a
suci from supi with Profile is a security profile of a network provider


