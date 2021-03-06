setup: ## Install all the build and lint dependencies
	go get -u github.com/golang/dep/...
	dep ensure
	go test

install: ## Install to $GOPATH/src
	go install
