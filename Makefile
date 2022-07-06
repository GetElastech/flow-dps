
all:
	go build --tags=relic ./...

crypto:
	rm -rf ./flow-go/crypto || true
	go mod download github.com/onflow/flow-go/crypto@v0.24.3
	export _P=$(PWD) && cd $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate && go build && cd $(_P)
	mkdir -p flow-go/crypto
	cp -r $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3/* ./flow-go/crypto
