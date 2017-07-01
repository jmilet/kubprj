# Run server
run:
	go run $(KUBPRJ)/go/src/jmi/cmd/countersrv/main.go

install:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go install jmi/cmd/countersrv

test:
	go test jmi/randomstring

container: install
	(cd $(KUBPRJ)/go; docker build -t countersrv -f countersrv.df .)

runcontainer: container
	docker run --rm -p 9092:9092 countersrv

clean:
	rm -v $(GOPATH)/bin/countersrv
	rm -v $(GOPATH)/bin/linux_amd64
	rm -v $(GOPATH)/bin/freebsd_arm
