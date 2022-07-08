
all: crypto 
	go build --tags=relic ./...

flow-go: clean
	git clone https://github.com/onflow/flow-go.git
	rm -rf ./flow-go/crypto/relic
	cd ./flow-go && git checkout v0.26.14-test-synchronization && make install-tools

crypto:
	GOPATH=$(PWD)/.gopath go mod download github.com/onflow/flow-go/crypto@v0.24.3
	cd $(PWD)/.gopath/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate && go build
	mkdir -p ./flow-go/crypto
	cp -r $(PWD)/.gopath/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3/* ./flow-go/crypto

clean:
	rm -rf ./flow-go
	rm -rf ./crypto
	rm -rf ./.gopath
