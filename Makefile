
all: flow-go 
	go build --tags=relic ./...

flow-go:
	rm -rf ./flow-go/crypto || true
	go mod download github.com/onflow/flow-go/crypto@v0.24.3
	mkdir -p ./flow-go/crypto
	cp -r $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3/* ./flow-go/crypto
	export _P=$(PWD) && cd ./flow-go/crypto && go generate && go build

