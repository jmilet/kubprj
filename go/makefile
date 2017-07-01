# Run server
run:
	go run src/jmi/cmd/countersrv/server.go

install:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go install jmi/cmd/countersrv

clean:
	rm -v $(GOPATH)/bin/server
