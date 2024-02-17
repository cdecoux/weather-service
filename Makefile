.PHONY: generators



# Make sure $GOPATH/bin is in your PATH. Default $GOPATH is $HOME/go (usually)
generate:
	go generate ./...