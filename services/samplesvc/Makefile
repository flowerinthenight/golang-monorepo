VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v0)
BLDVER = module:$(MODULE),version:$(VERSION),build:$(shell date +"%Y%m%d.%H%M%S.%N.%z")
BASE = $(CURDIR)
MODULE = samplesvc

.PHONY: all $(MODULE)
all: version $(MODULE)

$(MODULE):| $(BASE)
	@go build -v -o $(BASE)/bin/$@

$(BASE):
	@mkdir -p $(dir $@)

.PHONY: custom docker deploy

## The rule that is called by our root Makefile during CI builds.
custom: docker deploy

docker:
	cp $(BASE)/services/samplesvc/dockerfile.samplesvc $(BASE)/
	docker build -f dockerfile.samplesvc --rm -t $(MODULE):$(CIRCLE_SHA1) --build-arg version="$(BLDVER)" .

deploy:
	@chmod +x $(BASE)/services/samplesvc/deploy.sh
	NAME=$(MODULE) $(BASE)/services/samplesvc/deploy.sh

.PHONY: clean version list
clean:
	@rm -rfv bin
	@docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
	@exit 0

version:
	@echo "Version: $(VERSION)"

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
