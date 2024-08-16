
.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./gen/go/yq*/*.go \
					./gen/go/yq*/*.md \
					./gen/go/yq*/.openapi-generator \
					./gen/go/yq*/api \
					./gen/go/yq*/docs \
					./gen/go/yq*/go.*

.PHONY: generate-go-sdk
generate-go-sdk: | clean
	@openapi-generator-cli generate \
		-i query1.yml \
		-g go \
		-t ./gen/go/templates1 \
		-o ./gen/go/yq1 \
		-c query1-config-go.yml \
		--git-repo-id yahoo-finance-openapi/gen/go \
		--git-user-id S035779
	@openapi-generator-cli generate \
		-i query2.yml \
		-g go \
		-t ./gen/go/templates2 \
		-o ./gen/go/yq2 \
		-c query2-config-go.yml \
		--git-repo-id yahoo-finance-openapi/gen/go \
		--git-user-id S035779

.PHONY: authorize-go-templates
authorize-go-template: | clean
	@openapi-generator-cli author template \
		-g go \
		-o ./gen/go/templates1
	@openapi-generator-cli author template \
		-g go \
		-o ./gen/go/templates2

.PHONY: version
version:
	@cat query1.yml | grep -e "^  version: " | sed "s/  version: //"

.PHONY: check-version-bump
check-version-bump:
	@$(eval VER1 := $(shell cat query1.yml | grep -e "^  version: " | sed "s/  version: //"))
	@$(eval VER2 := $(shell cat query2.yml | grep -e "^  version: " | sed "s/  version: //"))
	@if [ "${VER1}" != "${VER2}" ]; then\
		1>&2 echo "Version mismatch, please make sure API version in query1.yml matches the one in query2.yml";\
		exit 1;\
	fi
	@$(eval LATEST_TAG := $(shell git fetch && git describe --tags $(shell git rev-list --tags --max-count=1) 2> /dev/null | sed "s/^v//"))
	@if [ "${VER1}" = "${LATEST_TAG}" ]; then\
		1>&2 echo "Please bump the API version in the api.yml OpenApi spec";\
	 	exit 2;\
	fi
