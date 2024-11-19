OS := $(shell uname)

ifeq ($(OS), windows)
    ifeq ($(CLCOSTA_PATH),)
        CLCOSTA_PATH ?= ${USERPROFILE}/.clcosta
    endif
else
    ifeq ($(CLCOSTA_PATH),)
        CLCOSTA_PATH ?= ${HOME}/.clcosta
    endif
endif

.PHONY: run
run: build
	@./build/clcosta

.PHONY: build
build:
	@rm -rf ./build
	@go build -o ./build/clcosta ./main.go
	@cp -r ./templates $(CLCOSTA_PATH)

.PHONY: dev
dev:
	@go run main.go
