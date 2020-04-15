# make sure we use Go modules for dependency management
export GO111MODULE := auto

# set up working directories
BUILD_TOOL := cmd
TOOLS := cmd build

CMD := ./cmd
BUILD_FUNC := github.com/AnNiran/word-chains/build

test:
	go test ${CMD}
	go test ${BUILD_FUNC}

mod:
	go mod tidy

# cover:

	cd cmd && go build && cd
