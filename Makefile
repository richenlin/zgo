.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP = zgo
SERVER_BIN = ./cmd/app/${APP}
RELEASE_ROOT = release
RELEASE_SERVER = release/${APP}

all: start

run: start

dev: debug

mod:
	go mod init ${APP}

install:
	go install ./cmd/app

build:
	go build -ldflags "-w -s" -o $(SERVER_BIN) ./cmd/app

debug:
	go run cmd/app/main.go web -c ./configs/config.toml

start:
	go run cmd/app/main.go web

# go get -u github.com/swaggo/swag/cmd/swag
swagger:
	swag init --generalInfo ./app/swagger.go --output ./app/swagger

# go get -u github.com/google/wire/cmd/wire
wire:
	wire gen ./app/injector

# go get github.com/facebookincubator/ent/cmd/entc
# entc init User
# generate the schema for User under <project>/ent/schema/
entc:
	go generate ./ent

# go get -u github.com/mdempsky/gocode
code:
	gocode -s -debug

test:
	@go test -v $(shell go list ./...)

clean:
	rm -rf data release $(SERVER_BIN) ./app/test/data ./cmd/app/data

pack: build
	rm -rf $(RELEASE_ROOT) && mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && tar -cvf $(APP).tar ${APP} && rm -rf ${APP}
