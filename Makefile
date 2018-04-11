# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
export GOCACHE=off

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

default: userstory

userstory: fmt
	@echo "\033[92mTest Getting user unexpired stories ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userstory_test.go

storyhighlights: fmt
	@echo "\033[92mTest Getting Story Highlights ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userstoryhighlight_test.go

follow: fmt
	@echo "\033[92mTest following and followers ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) follow_test.go

userinfo: fmt
	@echo "\033[92mTest user info from .../username/?__a=1 ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userinfo_test.go

postlive: fmt
	@echo "\033[92mTest Post Live Data in Reels Tray Feed ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) typepostlive_test.go

timeline: fmt
	@echo "\033[92mTest Timeline Feed ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) timeline_test.go

userreelmedia: fmt
	@echo "\033[92mTest User Reel Media Feed ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userreelmedia_test.go

reelstray: fmt
	@echo "\033[92mTest Reels Tray Feed ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) reelstray_test.go

toplive: fmt
	@echo "\033[92mTest top live API...\033[0m"
	@go test -v $(ALL_GO_SOURCES) toplive_test.go

post: fmt
	@echo "\033[92mTest getting post information...\033[0m"
	@go test -v $(ALL_GO_SOURCES) post_test.go

topsearch: fmt
	@echo "\033[92mTest Topsearch on web interface...\033[0m"
	@go test -v $(ALL_GO_SOURCES) topsearch_test.go

test: fmt
	@echo "\033[92mTest ...\033[0m"
	@go test -v

save_reels_tray_json:
	@echo "\033[92mSave Reels Tray Feed JSON ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) devsave_test.go -args -jsonKind="reels_tray"

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
