package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadUserStoryByNameLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	mgr.DownloadUserStoryByNameLayer("instagram", 2)
}

func ExampleDownloadUserStoryPostlive(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	err = mgr.DownloadUserStoryPostlive(25025320)
	if err != nil {
		panic(err)
	}
}

func ExampleDownloadUserStoryPostliveByName(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	err = mgr.DownloadUserStoryPostliveByName("instagram")
	if err != nil {
		panic(err)
	}
}

func ExampleDownloadUserStoryLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadUserStoryLayer(25025320, 2)
}

func ExampleDownloadStoryForever(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadStoryForever(90, 90, true, false)
}
