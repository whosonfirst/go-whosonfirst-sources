CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

spec:
	curl -o sources/spec.json https://raw.githubusercontent.com/whosonfirst/whosonfirst-sources/main/data/sources-spec-latest.json
