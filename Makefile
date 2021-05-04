.PHONY: generate
generate:
	cd go && controller-gen paths=./... crd:trivialVersions=true object output:crd:artifacts:config=crd/
	cd go/generate && go run types.go
	cp go/crd/*.yaml deploy/helm-chart/kube-mail/crds/

