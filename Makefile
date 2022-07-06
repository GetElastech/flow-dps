
all: crypto
	go build --tags=relic ./...

crypto:
	rm -rf ./flow-go/crypto || true
	go mod download github.com/onflow/flow-go/crypto@v0.24.3
	mkdir -p ./crypto
	cp -r $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3/* ./crypto
	export _P=$(PWD) && cd ./crypto && go generate && go build

