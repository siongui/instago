# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export PATH := $(GOROOT)/bin:$(PATH)

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

default: smart_test

highlight: fmt
	@echo "\033[92mDownload story highlights of all following users ...\033[0m"
	@go run example/download.go -downloadtype=highlight

story: fmt
	@echo "\033[92mDownload stories and post lives ...\033[0m"
	@go run example/download.go -downloadtype=story

timeline: fmt
	@echo "\033[92mDownload timeline ...\033[0m"
	@go run example/download.go -downloadtype=timeline

allposts:
	@echo "\033[92mDownload all user posts ...\033[0m"
	@go run example/allposts.go -id=$(id)

allpostsnologin:
	@echo "\033[92mDownload all user posts ...\033[0m"
	@go run example/allpostsnologin.go

post:
	@echo "\033[92mDownload single post $(code) ...\033[0m"
	@go run example/singlepost.go -code=$(code)

batchusers:
	@echo "\033[92mDownload all posts of multiple users ...\033[0m"
	@go run example/batchusers.go

allsaved:
	@echo "\033[92mDownload all your saved posts ...\033[0m"
	@go run example/download.go -downloadtype=saved

saved:
	@echo "\033[92mDownload latest 12 your saved posts ...\033[0m"
	@go run example/saved.go -num=12

userstory: fmt
	@echo "\033[92mDownload user $(id) unexpired stories ...\033[0m"
	@go run example/userstory.go -id=$(id)

userstorypostlive: fmt
	@echo "\033[92mDownload user $(id) unexpired stories and postlives ...\033[0m"
	@go run example/userstorypostlive.go -id=$(id)

userstory2layer: fmt
	@echo "\033[92mDownload user $(id) unexpired stories and stories of reel mentions...\033[0m"
	@go run example/userstory2layer.go -id=$(id)

userstorybatch: fmt
	@echo "\033[92mDownload unexpired stories of multiple users ...\033[0m"
	@go run example/userstorybatch.go

userstoryhighlight: fmt
	@echo "\033[92mDownload user $(id) story highlights ...\033[0m"
	@go run example/userstoryhighlight.go -id=$(id)

allunexpiredstories: fmt
	@echo "\033[92mDownload all unexpired stories ...\033[0m"
	@go run example/userstoryunexpired.go

post_test: fmt
	@echo "\033[92mDownload post test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) post_test.go

story_test: fmt
	@echo "\033[92mDownload user stories test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) story_test.go

smart_test: fmt
	@echo "\033[92mSmart downloading test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) smart_test.go

timeline_test: fmt
	@echo "\033[92mDownload timeline test ...\033[0m"
	# -count=1 disable test cache
	@go test -v -count=1 $(ALL_GO_SOURCES) timeline_test.go

postlive_test: fmt
	@echo "\033[92mDownload postlives test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) postlive_test.go

test_highlight: fmt
	@echo "\033[92mDownload user highlight test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) highlight_test.go

test_saved: fmt
	@echo "\033[92mSaved posts test ...\033[0m"
	# -count=1 disable test cache
	@go test -v -count=1 $(ALL_GO_SOURCES) saved_test.go

test_follow: fmt
	@echo "\033[92mTest follow methods ...\033[0m"
	# -count=1 disable test cache
	@go test -v -count=1 $(ALL_GO_SOURCES) follow_test.go

test_diffuser: fmt
	@echo "\033[92mTest users difference ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) diffuser_test.go

test_checkzero: fmt
	@echo "\033[92mTest check zero files and re-download ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) checkzero_test.go

test_timelimit: fmt
	@echo "\033[92mTest check zero files and re-download ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) timelimit_test.go

test_reelmedia: fmt
	@echo "\033[92mDownload reelmedia test ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) reelmedia_test.go

test_screenshot: fmt
	@echo "\033[92mTest screenshot feature ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) screenshot_test.go

example_zero: fmt
	@echo "\033[92mFind files with 0 size ...\033[0m"
	@go run example/zero.go -root=Instagram

example_find_post:
	@echo "\033[92mFind files with -post- in filename ...\033[0m"
	@find Instagram -name "*-post-*"

example_1dayrm: fmt
	@echo "\033[92mRemove files older than one day ...\033[0m"
	@go run example/1dayrm.go -root=Instagram -todir=${HOME}/Downloads/

example_rmemptydir: fmt
	@echo "\033[92mRemove empty dir ...\033[0m"
	@go run example/rmemptydir.go -root=Instagram

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

### example Makefile for simonedegiacomi/gphotosuploader
### Don't forget to prepare auth file of gphotosuploader
#upload:
#	while [ true ] ; do \
#		cd auth; gphotosuploader --upload path/to/Instagram/ --maxConcurrent 4 ; \
#		sleep 10 ; \
#	done
#
#install_gphotosuploader:
#	go get github.com/simonedegiacomi/gphotosuploader
