SVC=ipsvc
VERSION=1.0

LDFLAGS=-ldflags '-s -w -extldflags "-static"'

.PHONY: default
default: bin

.PHONY: bin
bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${SVC}

.PHONY: docker
docker:
	$(MAKE) bin
	docker build -t ${SVC}:${VERSION} -f Dockerfile .

.PHONY: clean
clean:
	if [ -f ${SVC} ]; then rm ${SVC}; fi
