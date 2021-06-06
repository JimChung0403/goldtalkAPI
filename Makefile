


#SUB_PACKAGE  := $(subst $(shell git rev-parse --show-toplevel),,$(CURDIR))
SUB_PACKAGE := "goldtalkAPI"
APP      := $(shell basename $(SUB_PACKAGE))


#export GO111MODULE=on
OUTPUT     = $(CURDIR)/output
CONF       = $(CURDIR)/conf

.DEFAULT: all
all: build

build: clean prepare
	go build -o "$(OUTPUT)/bin/$(SUB_PACKAGE)" "$(CURDIR)/main.go"

clean:
	echo "====clean $(OUTPUT)"
	rm -rf "$(OUTPUT)"

prepare:
	mkdir -p "$(OUTPUT)/log"
	cp -vr "$(CONF)" "$(OUTPUT)"
	cp -v "$(CURDIR)/control.sh" "$(OUTPUT)"

fmt:
	go fmt ./...

run:
	cd "$(OUTPUT)" && bin/$(APP)

swag:
	swag init


.PHONY: all build fmt clean prepare run test init upgrade
$(VERBOSE).SILENT:
