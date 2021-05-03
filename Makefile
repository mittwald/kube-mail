.PHONY: generate
generate:
	controller-gen paths=go/... crd:trivialVersions=true object output:crd:artifacts:config=./go/crd/
