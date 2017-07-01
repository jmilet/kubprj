VERSION=v.0.0.1
RUN_GOOS=darwin
RUN_GOARCH=amd64
GOOS=linux
GOARCH=amd64


# Run server
run:
	GOOS=$(RUN_GOOS) GOARCH=$(RUN_GOARCH) go run $(KUBPRJ)/go/src/jmi/cmd/countersrv/main.go

install:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go install jmi/cmd/countersrv

test:
	go test jmi/randomstring

container: install
	(cd $(KUBPRJ)/go; docker build -t jmilet/countersrv -f countersrv.df .)

runcontainer: container
	docker run --rm -p 9092:9092 jmilet/countersrv

pushcontainer:
	docker push jmilet/countersrv:$(VERSION)

clean:
	rm -v $(GOPATH)/bin/countersrv
	rm -vrf $(GOPATH)/bin/linux_amd64
	rm -vrf $(GOPATH)/bin/freebsd_arm
