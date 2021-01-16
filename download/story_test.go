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

	mgr.DownloadUserStoryByNameLayer("instagram", 2, 12)
}

func ExampleDownloadUserStory(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	err = mgr.DownloadUserStory("25025320")
	if err != nil {
		panic(err)
	}
}

func ExampleDownloadUserStoryByName(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	err = mgr.DownloadUserStoryByName("instagram")
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

	mgr.DownloadUserStoryLayer("25025320", 2, 12)
}
