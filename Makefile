export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

handlers := $(shell find . -name '*main.go')

install:
	@echo "\nInstalling dependencies"
	go get ./...

clean:
	@echo "\nRemoving old builds"
	rm -rf bin

define build-cloudwatchalarmtrigger
	@echo "\nBuilding cloud watch alarm trigger"
	- cd handler/cloudwatchalarmtrigger/ && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../bin/cloudwatchalarmtrigger *.go
	- chmod +x bin/cloudwatchalarmtrigger
	- cd bin/ && zip -j cloudwatchalarmtrigger.zip cloudwatchalarmtrigger
	@echo "Finished building cloud watch alarm trigger"
endef

build:
	${build-cloudwatchalarmtrigger}