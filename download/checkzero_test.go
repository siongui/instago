package igdl

import (
	"testing"
)

func ExampleCheckZero(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.CheckZero("path/to/downloaded/files", "dir/to/move/zero/files")
}
