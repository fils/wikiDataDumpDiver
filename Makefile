BINARY := wdparse
#DOCKERVER :=`cat VERSION`
.DEFAULT_GOAL := wdparse

wdparse:
		cd cmd/$(BINARY) ; \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

