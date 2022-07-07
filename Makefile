
all: flow-go 
	go build --tags=relic ./...

flow-go: clean
	git clone https://github.com/onflow/flow-go.git
	rm -rf ./flow-go/crypto/relic
	cd ./flow-go && git checkout master && make install-tools

clean:
	rm -rf ./flow-go
