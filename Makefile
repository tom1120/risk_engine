GOCMD=GO111MODULE=on CGO_ENABLED=0 go
GOBUILD=$(GOCMD) build

all: build
build:
	rm -rf dist/
	mkdir -p dist/conf dist/bin
	cp cmd/risk_engine/config.yaml dist/conf
	cp demo dist/demo -r
	$(GOBUILD) -o dist/bin/risk_engine cmd/risk_engine/engine.go

clean:
	rm -rf dist/

run:
	cd dist/; nohup bin/risk_engine -c conf/config.yaml >nohup.out 2>nohup.out &

stop:
	pkill -f bin/risk_engine
