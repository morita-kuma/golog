setup:
		go get github.com/golang/dep/cmd/dep
		go get github.com/golang/lint/golint
		go get golang.org/x/tools/cmd/goimports

lint: setup
		for gopkg in $$(go list ./... | grep -v vendor); do golint $$gopkg;  done

fmt: setup
		for gofile in $$(find . -name "*.go" | grep -v vendor); do goimports -w $$gofile; done;
		go fmt $(shell go list ./... | grep -v vendor)

build: setup
		go build
