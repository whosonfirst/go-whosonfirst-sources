CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

LDFLAGS=-s -w

spec:	
	go run cmd/mk-spec/main.go > sources/spec.go
