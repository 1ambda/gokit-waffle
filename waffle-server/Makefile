GOARCH  = amd64
GOOS    = $(OS)

ifeq ($(GOOS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
endif

GOCMD		= go
GOBUILD	= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOTEST	= ginkgo -r -cover
GOTOOL	= $(GOCMD) tool cover
GOGET		= $(GOCMD) get -u -v
GODEP		= glide
RM			= rm -rf

APP			  = WaffleServer
DIST_DIR  = dist

# build flags
IMPORTPATH	= "github.com/1ambda/gokit-waffle/waffle-server"
VERSION		= "0.0.1"
GITSHA		= $(shell git rev-parse HEAD)
BUILDTIME	= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS		= -ldflags="-X ${IMPORTPATH}/config.GitHash=${GITSHA} -X ${IMPORTPATH}/config.BuildTime=${BUILDTIME} -X ${IMPORTPATH}/config.Version=${VERSION}"

# build commands
.PHONY: all install build build-all test run clean fclean

all: fclean install build-all test

install:
	$(GOGET) github.com/Masterminds/glide # for dep
	$(GOGET) github.com/onsi/ginkgo/ginkgo # for test
	$(GOGET) github.com/franciscocpg/gox # for cross compile
	$(GODEP) install

build: clean
	$(GOBUILD) $(LDFLAGS) -o ${DIST_DIR}/${GOOS}-${GOARCH}/${APP} -v ./main.go

build-all: clean
	gox -verbose \
	$(LDFLAGS) \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="${DIST_DIR}/{{.OS}}-{{.Arch}}/${APP}" .

test:
	$(GOTEST)

run:
	./${DIST_DIR}/${GOOS}-${GOARCH}/${APP}

clean:
	find . -type f -name '*.coverprofile' -exec rm {} +
	$(RM) ${DIST_DIR}

fclean: clean
	$(RM) vendor
