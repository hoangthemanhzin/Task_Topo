controller:
	go build -o bin/controller mesh/controller/cmd/main.go 
amf:
	go build -o bin/amf nfs/amf/cmd/main.go	
damf:
	go build -o bin/damf nfs/damf/cmd/main.go	
udm:
	go build -o bin/udm nfs/udm/cmd/main.go	
ausf:
	go build -o bin/ausf nfs/ausf/cmd/main.go	
smf:
	go build -o bin/smf nfs/smf/cmd/main.go	
pcf:
	go build -o bin/pcf nfs/pcf/cmd/main.go	
udr:
	go build -o bin/udr nfs/udr/cmd/main.go	
pran:
	go build -o bin/pran nfs/pran/cmd/main.go	
upmf:
	go build -o bin/upmf nfs/upmf/cmd/main.go	
upf:
	go build -o bin/upf nfs/upf/cmd/main.go	
docker:
	cd docker; make
clean:
	rm bin/*
docker-clean:
	cd docker; make clean

.PHONY: docker
.DEFAULT_GOAL := all
all: controller pran damf amf udm ausf smf pcf udr upmf upf
