
all:
	go mod download github.com/onflow/flow-go/crypto@v0.24.3
	cd $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate && go build
	mkdir -p flow-go
	cp -r $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 ./flow-go/crypto
	go build --tags=relic ./...

