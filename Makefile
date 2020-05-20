## the binary name
ARTIFACT_NAME = mercurius

## for the module itself
MODULE_PATH = github.com/worldiety/mercurius

## the path which contains the main package to execute
MAIN_PATH = github.com/worldiety/mercurius/cmd/mercurius

## for ldflags replacement
BUILD_FILE_PATH = ${MODULE_PATH}

## which linter version to use?
GOLANGCI_LINT_VERSION = v1.24.0

## better keep hands off below this line

currentDate := $(shell date +"%Y.%m.%d-%H:%M:%S")
currentCommit := $(shell git rev-parse HEAD)

LDFLAGS = -X $(MODULE_PATH)/build.Time=${currentDate} -X $(MODULE_PATH)/build.Commit=${currentCommit}

TMP_DIR = $(TMPDIR)/$(MODULE_PATH)
BUILD_DIR = $(TMP_DIR)/build
TOOLSDIR = $(TMP_DIR)
GO = go
GOLANGCI_LINT = ${TOOLSDIR}/golangci-lint
GOLINT = ${TOOLSDIR}/golint
TMP_GOPATH = $(CURDIR)/${TOOLSDIR}/.gopath

GOROOT = $(shell ${GO} env GOROOT)

all: generate lint test build ## Runs lint, test and build

clean: ## Removes any temporary and output files
	rm -rf ${BUILD_DIR}

lint: ## Executes all linters
	${GOLANGCI_LINT} run --enable-all --exclude-use-default=false

test: ## Executes the tests
	${GO} test -race ./...

.PHONY: build generate setup

build: ## Performs a build and puts everything into the build directory
	${GO} build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${ARTIFACT_NAME} ${MAIN_PATH}


buildWasm: generate ## Performs a wasm frontend build
	GOOS=js GOARCH=wasm ${GO} build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/app.wasm ${MAIN_PATH}

	cp "${GOROOT}/misc/wasm/wasm_exec.js" ${BUILD_DIR}
	cp -r static/. ${BUILD_DIR}

run: clean generate build ## Starts the compiled program
	${BUILD_DIR}/${ARTIFACT_NAME}

help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## Executes go generate
	${GO} generate ./...

setup: installGolangCi  ## Installs golangci-lint
	${GO} mod tidy


installGolangCi:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLSDIR) $(GOLANGCI_LINT_VERSION)


.DEFAULT_GOAL := help

