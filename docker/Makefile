PRJ = 'b5gc'
IMAGE_BASE = 'base'
TAG = 'latest'

.PHONY: base ueransim

all:  base amf damf pran ausf pcf smf udm controller upf ueransim #udr

base:
	./copybuild.sh
	docker build --target b5gc-base -t ${PRJ}-${IMAGE_BASE}:${TAG} -f nf.Dockerfile .
	docker image ls ${PRJ}-${IMAGE_BASE}:${TAG}

smf: base
	docker build --build-arg B5GC_MODULE=smf -t ${PRJ}-smf:${TAG} -f nf.Dockerfile .
amf: base
	docker build --build-arg B5GC_MODULE=amf -t ${PRJ}-amf:${TAG} -f nf.Dockerfile .
damf: base
	docker build --build-arg B5GC_MODULE=damf -t ${PRJ}-damf:${TAG} -f nf.Dockerfile .
pran: base
	docker build --build-arg B5GC_MODULE=pran -t ${PRJ}-pran:${TAG} -f nf.Dockerfile .
#upf: base
#	docker build --build-arg B5GC_MODULE=upf -t ${PRJ}-upf:${TAG} -f nf.Dockerfile .
udr: base
	docker build --build-arg B5GC_MODULE=udr -t ${PRJ}-udr:${TAG} -f nf.Dockerfile .
udm: base
	docker build --build-arg B5GC_MODULE=udm -t ${PRJ}-udm:${TAG} -f nf.Dockerfile .
pcf: base
	docker build --build-arg B5GC_MODULE=pcf -t ${PRJ}-pcf:${TAG} -f nf.Dockerfile .
ausf: base
	docker build --build-arg B5GC_MODULE=ausf -t ${PRJ}-ausf:${TAG} -f nf.Dockerfile .
controller: base
	docker build --build-arg B5GC_MODULE=controller -t ${PRJ}-controller:${TAG} -f nf.Dockerfile .

upf: base
	docker build -t ${PRJ}-upf:${TAG} -f upf.Dockerfile .

ueransim:
	echo Build UERANSIM
	docker build -t ueransim:${TAG} -f ueransim.Dockerfile .



clean:
	rm -rf b5gc
	-docker system prune --force
	-docker rmi b5gc-amf:latest
	-docker rmi b5gc-damf:latest
	-docker rmi b5gc-pran:latest
	-docker rmi b5gc-pcf:latest
	-docker rmi b5gc-smf:latest
	-docker rmi b5gc-udm:latest
	#-docker rmi b5gc-udr:latest
	-docker rmi b5gc-ausf:latest
	-docker rmi b5gc-controller:latest
	-docker rmi b5gc-base:latest
	-docker rmi b5gc-upf:latest
	-docker rmi ueransim:latest
