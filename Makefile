SVC=ipsvc
VERSION=0.0.1

LDFLAGS=-ldflags '-s -w -extldflags "-static"'

.PHONY: default
default: bin

.PHONY: bin
bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${SVC}

.PHONEY: clean
clean:
	if [ -f ${SVC} ]; then rm ${SVC}; fi
