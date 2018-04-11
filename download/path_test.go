package igdl

import (
	"testing"
)

func TestGetPostFilePath(t *testing.T) {
	path := GetPostFilePath("instagram", "123.mp4", 1520056661)
	if path != "Instagram/instagram/posts/instagram-post-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestGetStoryFilePath(t *testing.T) {
	path := getStoryFilePath("instagram", "123.mp4", 1520056661)
	if path != "Instagram/instagram/stories/instagram-story-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestGetPostLiveFilePath(t *testing.T) {
	path := getPostLiveFilePath("instagram", "123.mp4", "video", 1520056661)
	if path != "Instagram/instagram/postlives/instagram-postlive-video-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestGetPostLiveMergedFilePath(t *testing.T) {
	vpath := "Instagram/instagram/postlives/instagram-postlive-video-2018-03-03T13:57:41+08:00-1520056661.mp4"
	apath := "Instagram/instagram/postlives/instagram-postlive-audio-2018-03-03T13:57:41+08:00-1520056661.mp4"
	path := getPostLiveMergedFilePath(vpath, apath)
	if path != "Instagram/instagram/postlives/instagram-postlive-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestAppendIndexToFilename(t *testing.T) {
	nf := appendIndexToFilename("instagram-post-2018-03-03T13:57:41+08:00-1520056661.mp4", 0)
	if nf != "instagram-post-2018-03-03T13:57:41+08:00-1520056661-0.mp4" {
		t.Error(nf)
	}
	nf = appendIndexToFilename("instagram-post-2018-03-03T13:57:41+08:00-1520056661.mp4", 1)
	if nf != "instagram-post-2018-03-03T13:57:41+08:00-1520056661-1.mp4" {
		t.Error(nf)
	}
}
