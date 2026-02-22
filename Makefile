CR_VERSION ?= 1.8.1
CONTROLLERGEN_VERSION ?= 0.20.0
GORELEASER_VERSION ?= 2.12.7
HELM_VERSION ?= 4.0.4
SVU_VERSION ?= 3.3.0

BIN ?= ./.bin
DIST ?= ./.dist

arch := $(shell uname -m)
ifeq ($(arch),aarch64)
    arch := arm64
else
    arch := amd64
endif
bin := $(abspath $(BIN))
dist := $(abspath $(DIST))
project := $(abspath $(dir $(MAKEFILE_LIST)))

include Makefile.include.mk

.PHONY: default
default: list-targets

list-targets:
	@echo "available targets:"
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null \
		| awk -v RS= -F: '/(^|\n)# Files(\n|$$$$)/,/(^|\n)# Finished Make data base/ {if ($$$$1 !~ "^[#.]") {print $$$$1}}' \
		| sort \
		| grep -E -v -e '^[^[:alnum:]]' -e '^$$@$$$$' \
		| sed 's/^/\t/'

.PHONY: build-chart
build-chart:
	@if [ "$(VERSION)" = "" ]; then 1>&2 echo "VERSION unset"; exit 1; fi;
	# remove existing chart package
	rm -rf $(dist)/s3-controller-*.tgz
	# package chart s3-controller
	helm package $(project)/charts/s3-controller --app-version $(VERSION) --destination $(dist) --version $(VERSION)
	# rename chart to include 'helm'
	mv $(dist)/s3-controller-$(VERSION).tgz $(dist)/s3-controller-helm-$(VERSION).tgz

.PHONY: build-docs
build-docs: install-docs
	# build docs
	cd docs && npm run build -- --out-dir=$(dist)

.PHONY: dev-docs
dev-docs: install-docs
	# run docs dev server
	cd docs && npm run start

.PHONY: install-docs
install-docs:
	# install doc dependencies
	cd docs && npm install

.PHONY: install-project
install-project:
	# install project dependencies
	go mod download

.PHONY: install-tools
install-tools:

$(eval $(call tool-from-apt,bsdtar,libarchive-tools))
$(eval $(call tool-from-apt,curl,curl))

cr_arch := $(arch)
cr_url := https://github.com/helm/chart-releaser/releases/download/v$(CR_VERSION)/chart-releaser_$(CR_VERSION)_linux_$(cr_arch).tar.gz
$(eval $(call tool-from-tar-gz,cr,$(cr_url),0))

controllergen_arch := $(arch)
controllergen_url := https://github.com/kubernetes-sigs/controller-tools/releases/download/v$(CONTROLLERGEN_VERSION)/controller-gen-linux-$(controllergen_arch)
$(eval $(call tool-from-url,controller-gen,$(controllergen_url)))

goreleaser_arch := $(arch)
ifeq ($(goreleaser_arch),amd64)
	goreleaser_arch := x86_64
endif
goreleaser_url := https://github.com/goreleaser/goreleaser/releases/download/v$(GORELEASER_VERSION)/goreleaser_Linux_$(goreleaser_arch).tar.gz
$(eval $(call tool-from-tar-gz,goreleaser,$(goreleaser_url),0))

helm_arch := $(arch)
helm_url := https://get.helm.sh/helm-v$(HELM_VERSION)-linux-$(helm_arch).tar.gz
$(eval $(call tool-from-tar-gz,helm,$(helm_url),1))

svu_url := https://github.com/caarlos0/svu/releases/download/v$(SVU_VERSION)/svu_$(SVU_VERSION)_linux_$(arch).tar.gz
$(eval $(call tool-from-tar-gz,svu,$(svu_url),0))

.PHONY: generate
generate:
	# clean up generated dirs
	rm -rf $(project)/charts/s3-controller/generated && mkdir -p $(project)/charts/s3-controller/generated
	# create deepcopy implementations
	controller-gen object paths="$(project)/pkg/api/..."
	# create crd manifests
	controller-gen crd paths="$(project)/pkg/api/..." output:stdout > charts/s3-controller/generated/crds.yaml
	# create rbac manifests
	controller-gen rbac:roleName=__roleName__ paths="$(project)/pkg/api/..." output:stdout > $(project)/charts/s3-controller/generated/rbac.yaml

$(bin):
	# make $(bin) folder
	mkdir -p $(bin)