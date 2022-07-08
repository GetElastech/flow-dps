
all: crypto 
	go build --tags=relic ./...

flow-go: clean
	git clone https://github.com/onflow/flow-go.git
	rm -rf ./flow-go/crypto/relic
	cd ./flow-go && git checkout v0.26.14-test-synchronization && make install-tools

crypto:
	export GOPATH=$(PWD)/gopath && go mod download github.com/onflow/flow-go/crypto@v0.24.3
	export GOPATH=$(PWD)/gopath && export _P=$(PWD) && cd $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate && go build && cd $(_P)
	mkdir -p ./flow-go/crypto
	export GOPATH=$(PWD)/gopath && cp -r $(GOPATH)/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3/* ./flow-go/crypto

clean:
	rm -rf ./flow-go
	rm -rf ./crypto
