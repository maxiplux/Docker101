# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

BINARY = AuthServer
GOARCH = amd64
LDFLAGS= -ldflags '-extldflags "-static"'
LOG= api.log
DEPENDENCIES := github.com/gorilla/mux \
								golang.org/x/crypto/bcrypt \
								github.com/lib/pq

.PHONY: clean
clean:
		@if [ -f bin/${BINARY}-* ] ; then rm bin/${BINARY}-* ; fi
		@if [ -f ${LOG} ] ; then rm ${LOG} ; fi

.PHONY: build
build:
	@if [ -f bin/${BINARY}-* ] ; then rm bin/${BINARY}-* ; fi
	+ GOOS=linux CGO_ENABLED=0 GOARCH=${GOARCH} go build ${LDFLAGS} -o ./bin/${BINARY}-linux ./src/

.PHONY: test
test:
		@if [ -f bin/${BINARY}-* ] ; then rm ${BINARY}-* ; fi
		@if [ -f ${LOG} ] ; then rm ${LOG} ; fi
		go run *.go 2> $(LOG)

.PHONY: deps
deps:
			go get $(DEPENDENCIES)
#	EOF!
