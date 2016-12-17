GOARCH		= amd64

ifeq ($(OS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
else
  GOOS		= $(OS)
endif

GOCMD		= GOOS=$(GOOS) GOARCH=$(GOARCH) go
GOBUILD	= $(GOCMD) build
GOTEST	= ginkgo -r -cover
GOTOOL	= go tool cover
GOGET		= go get
GODEP		= glide
RM			= rm -rf
APP			= WaffleServer

# build flags
VERSION   = "0.0.1"
GITSHA		= $(shell git rev-parse HEAD)
BUILDDATE	= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS   = -ldflags="-X main.GitHash=${GITSHA} -X main.BuildTime=${BUILDDATE} -X main.Version=${VERSION}"


# build commands
.PHONY		: all install build test clean fclean

all		: fclean install build test

install :
		curl https://glide.sh/get | sh
		$(GODEP) install
		$(GOGET) github.com/onsi/ginkgo/ginkgo

build		:
		$(GOBUILD) $(LDFLAGS) -o $(APP) -v ./main.go

test		:
		$(GOTEST)

clean		:
		find . -type f -name '*.coverprofile' -exec rm {} +
		$(RM) $(APP)

fclean  : clean
		$(RM) vendor
