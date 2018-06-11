# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

default: highlight

highlight: fmt
	@echo "\033[92mDownload story highlights of all following users ...\033[0m"
	@go run example/download.go -downloadtype=highlight

story: fmt
	@echo "\033[92mDownload stories and post lives ...\033[0m"
	@go run example/download.go -downloadtype=story

timeline: fmt
	@echo "\033[92mDownload timeline ...\033[0m"
	@go run example/download.go -downloadtype=timeline

allpostsnologin:
	@echo "\033[92mDownload all user posts ...\033[0m"
	@go run example/allpostsnologin.go

saved:
	@echo "\033[92mDownload all your saved posts ...\033[0m"
	@go run example/download.go -downloadtype=saved

userstory:
	@echo "\033[92mDownload user $(id) unexpired stories ...\033[0m"
	@go fmt example/downloaduserstory.go
	@go run example/downloaduserstory.go -id=$(id)

userstoryhighlight:
	@echo "\033[92mDownload user $(id) story highlights ...\033[0m"
	@go fmt example/downloaduserstoryhighlight.go
	@go run example/downloaduserstoryhighlight.go -id=$(id)

post_test: fmt
	@echo "\033[92mDownload timeline ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) post_test.go

profile_pic_test: fmt
	@echo "\033[92mDownload timeline ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) profilepic_test.go

test: fmt
	@echo "\033[92mRun testing code...\033[0m"
	@go test -v path.go path_test.go util.go util_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt example/*.go

install:
	go get -u github.com/siongui/instago/download
