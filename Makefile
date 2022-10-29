define LD_FLAGS
"-X 'main.Commit=$$(git rev-parse HEAD)'\
 -X 'main.Tag=$$(git name-rev --tags --name-only HEAD)'"
endef
.PHONY: build

build:
	mkdir -p ./build/bin
	@go build -ldflags=${LD_FLAGS} -a -o ./build/bin/gopkgsearch ./...

test.build:
	@go build -o /dev/null ./...