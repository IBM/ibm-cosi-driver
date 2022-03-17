DRIVER_NAME=ibm-cosi-driver
.PHONY: build-% build container-% container clean

REVISION=$(shell git describe --long --tags --match='v*' --dirty 2>/dev/null || git rev-list -n1 HEAD)

build-%:
	@echo ">  Building driver binary..."
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.version=$(REVISION) -extldflags "-static"' -o ./bin/$* ./cmd/$*
container-%: build-%
	docker build -t $*:latest -f $(shell if [ -e ./cmd/$*/Dockerfile ]; then echo ./cmd/$*/Dockerfile; else echo Dockerfile; fi) --label revision=$(REVISION) .

build: $(DRIVER_NAME:%=build-%)
container: $(DRIVER_NAME:%=container-%)

clean:
	-rm -rf bin
