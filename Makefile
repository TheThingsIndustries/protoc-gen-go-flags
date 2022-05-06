# Copyright © 2021 The Things Industries B.V.
# SPDX-License-Identifier: Apache-2.0

.PHONY: default

default: build test

.PHONY: clean

clean:
	rm -f ./annotations/*.pb.go
	rm -f ./test/*/*.pb.go

.dev/protoc-gen-go-flags/annotations.proto: annotations.proto
	mkdir -p $(shell dirname $@)
	cp $< $@

annotations/annotations.pb.go: .dev/protoc-gen-go-flags/annotations.proto .dev/golangproto/bin/protoc .dev/golangproto/bin/protoc-gen-go
	PATH="$$PWD/.bin:$$PWD/.dev/golangproto/bin:$$PATH" protoc -I .dev --go_opt=module=github.com/TheThingsIndustries/protoc-gen-go-flags --go_out=./ $<

internal/gogoproto/gogo.pb.go: internal/gogoproto/gogo.proto .dev/golangproto/bin/protoc .dev/golangproto/bin/protoc-gen-go
	PATH="$$PWD/.bin:$$PWD/.dev/golangproto/bin:$$PATH" protoc -I . --go_opt=paths=source_relative --go_out=./ ./internal/gogoproto/gogo.proto

internal/flagsplugin/annotations.pb.go: internal/flagsplugin/annotations.proto
	protoc -I . --go_opt=paths=source_relative --go_out=./ ./internal/flagsplugin/annotations.proto

BINARY_DEPS = annotations/annotations.pb.go internal/gogoproto/gogo.pb.go $(wildcard cmd/protoc-gen-go-flags/*.go) $(wildcard internal/gen/*.go)

VERSION ?= 0.0.0-dev

LDFLAGS = -X github.com/TheThingsIndustries/protoc-gen-go-flags/internal/gen.Version=$(VERSION)

.bin/protoc-gen-go-flags: $(BINARY_DEPS)
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/protoc-gen-go-flags

.bin/protoc-gen-go-flags-linux-amd64: $(BINARY_DEPS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/protoc-gen-go-flags

.bin/protoc-gen-go-flags-linux-arm64: $(BINARY_DEPS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/protoc-gen-go-flags

REPLACES = Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types

.PHONY: build

build: .bin/protoc-gen-go-flags .bin/protoc-gen-go-flags-linux-amd64 .bin/protoc-gen-go-flags-linux-arm64

.PHONY: watch

watch:
	ls annotations.proto cmd/protoc-gen-go-flags/*.go internal/gen/*.go test/*.proto | entr make build test

OS :=
ifeq ($(shell uname),Linux)
	OS = linux
endif
ifeq ($(shell uname),Darwin)
	OS = osx
endif

.dev/golangproto/bin/protoc:
	mkdir -p .dev/golangproto/bin
	curl -sSL -o .dev/golangproto/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-$(OS)-x86_64.zip
	unzip -o .dev/golangproto/protoc.zip -d .dev/golangproto/

.dev/gogoproto/bin/protoc:
	mkdir -p .dev/gogoproto/bin
	curl -sSL -o .dev/gogoproto/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-$(OS)-x86_64.zip
	unzip -o .dev/gogoproto/protoc.zip -d .dev/gogoproto/
	curl -sSL -o .dev/gogoproto/gogoproto.zip https://github.com/gogo/protobuf/archive/refs/heads/master.zip
	unzip -o .dev/gogoproto/gogoproto.zip protobuf-master/protobuf/google/protobuf/*.proto -d .dev/gogoproto
	mv .dev/gogoproto/protobuf-master/protobuf/google/protobuf/*.proto .dev/gogoproto/include/google/protobuf/

.dev/golangproto/bin/protoc-gen-go:
	go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

.dev/gogoproto/bin/protoc-gen-gogo:
	go build -o $@ github.com/gogo/protobuf/protoc-gen-gogo

.PHONY: testprotos

testprotos: build .dev/golangproto/bin/protoc .dev/gogoproto/bin/protoc .dev/golangproto/bin/protoc-gen-go .dev/gogoproto/bin/protoc-gen-gogo
	PATH="$$PWD/.bin:$$PWD/.dev/golangproto/bin:$$PATH" protoc -I ./test -I . \
	  --go_opt=paths=source_relative --go_out=./test/golang \
	  --go-flags_opt=paths=source_relative --go-flags_out=./test/golang \
	  ./test/*.proto
	PATH="$$PWD/.bin:$$PWD/.dev/gogoproto/bin:$$PATH" protoc -I ./test -I . \
	  --gogo_opt=paths=source_relative --gogo_opt=$(REPLACES) --gogo_out=./test/gogo \
	  --go-flags_opt=paths=source_relative --go-flags_opt=$(REPLACES) --go-flags_opt=lang=gogo --go-flags_out=./test/gogo \
	  ./test/*.proto

.PHONY: test

test: testprotos
	go test ./flagsplugin ./test/gogo ./test/golang
