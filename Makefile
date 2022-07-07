
all: flow-go 
	go build --tags=relic ./...

flow-go: clean
	git clone https://github.com/onflow/flow-go.git
	#export _P=$(PWD) && cd ./flow-go/crypto && go generate && go build
	rm -rf ./flow-go/crypto/relic
	cd ./flow-go && make install-tools

clean:
	rm -rf ./flow-go
