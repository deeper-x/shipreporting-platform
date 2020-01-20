export GOPATH=${HOME}/go
export GOBIN=${GOPATH}/bin
export GO111MODULE=on

build:
	go get -d ./...
	@echo "export to ${GOBIN}..."
	go build -o ${GOBIN}

