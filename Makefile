all: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/go-websocket-example_linux_amd64 -ldflags '-s -w -X "main.mode=prod"' main.go

build-linux-arm64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/go-websocket-example_linux_arm64  -ldflags '-s -w -X "main.mode=prod"' main.go

build-docker-image:
	docker buildx build --platform linux/amd64,linux/arm64 -t caryqy2/go-websocket-example . --push

gox-linux:
	gox -ldflags '-s -w -X "main.mode=prod"' -osarch="linux/amd64 linux/arm64" -output="build/go-websocket-example_{{.OS}}_{{.Arch}}"

gox-all:
	gox -ldflags '-s -w -X "main.mode=prod"' -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="build/go-websocket-example_{{.OS}}_{{.Arch}}"

clean:
	rm -f build/go-websocket-example_*