.PHONY: all docker docker-clean publish-docker

REPO=github.com/xtao/metaview
#VERSION?=$(shell git describe HEAD | sed s/^v//)
VERSION?=1.1
DATE?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

#LDFLAGS=-X ${REPO}/version.Version=${VERSION} -X ${REPO}/version.BuildDate=${DATE}
TOOLS=${GOPATH}/bin/go-bindata \
      ${GOPATH}/bin/go-bindata-assetfs \
      ${GOPATH}/bin/goconvey
SRC=$(shell find . -name '*.go')
TESTFLAGS="-v"

PACKAGES=$(shell go list ./... | grep -v /vendor/)
METAVIEW_VERSION=$(shell git rev-parse --short HEAD)

#${TOOLS}:
#	go get github.com/jteeuwen/go-bindata/...
#	go get github.com/elazarl/go-bindata-assetfs/...
#	go get github.com/smartystreets/goconvey


vet:
	go vet ${PACKAGES}

lint:
	go list ./... | grep -v /vendor/ | grep -v assets | xargs -L1 golint -set_exit_status

rpm/metaview/metaview: ${SRC}
	CGO_ENABLED=0 GOOS=linux go build -ldflags "${LDFLAGS}" -a -installsuffix cgo -o $@ github.com/xtao/metaview/cmd/metaview

rpm/metacli/metacli: ${SRC}
	CGO_ENABLED=0 GOOS=linux go build -ldflags "${LDFLAGS}" -a -installsuffix cgo -o $@ github.com/xtao/metaview/cmd/metacli


rpm: rpm/metaview/metaview rpm/metacli/metacli
	cd rpm && ./autogen.sh
	cd rpm && ./configure
	sed -i 's/XTAO_GIT_METAVIEW_VERSION/$(METAVIEW_VERSION)/g' rpm/Makefile
	cd rpm && make clean && make rpm

clean:
	rm -rf rpm/metaview/metaview
	rm -rf rpm/metacli/metacli
	rm -rf cmd/metaview/metaview
	rm -rf cmd/metacli/metacli
	rm -rf cmd/metaview/*.log
	rm -rf cmd/metacli/*.log

	rm -rf rpm/rpmbuild
	rm -rf rpm/config.*
	rm -rf rpm/configure
	rm -rf rpm/install.sh
	rm -rf rpm/Makefile
	rm -rf rpm/Makefile.in
	rm -rf rpm/*.tar.gz
