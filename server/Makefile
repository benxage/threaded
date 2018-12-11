PROJ_DIR=$(pwd)

default: run

run: depends
	go run server.go

depends:
	dep ensure

build: clean
	go build server.go

clean:
	rm -rf server
