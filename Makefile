.PHONY: clean all

clean:
	@rm --force bin/* *.json 2>&1 ; echo 'done'

all: bin/create-array-linux bin/create-array-windows bin/create-array-darwin

bin/create-array-% : cmd/cli.go
	GOOS="$*" GOARCH="amd64" go build -o $@ $^
