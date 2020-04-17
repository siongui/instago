package instago

import (
	"os"
	"testing"
)

func ExampleGetPostInfo(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	// photos and videos in a post
	em, err := mgr.GetPostInfo("B-KwveQgsdr")
	if err != nil {
		t.Error(err)
		return
	}

	err = PrintPostItem(&em)
	if err != nil {
		t.Error(err)
		return
	}
	PrintTaggedUsers(em.EdgeMediaToTaggedUser)
	for _, child := range em.EdgeSidecarToChildren.Edges {
		PrintTaggedUsers(child.Node.EdgeMediaToTaggedUser)
	}

	// IGTV video
	em, err = mgr.GetPostInfo("B-NK7ZYgntC")
	if err != nil {
		t.Error(err)
		return
	}

	err = PrintPostItem(&em)
	if err != nil {
		t.Error(err)
		return
	}

	// single video
	em, err = mgr.GetPostInfo("B-H13uwgNRK")
	if err != nil {
		t.Error(err)
		return
	}

	err = PrintPostItem(&em)
	if err != nil {
		t.Error(err)
		return
	}

	// single photo
	em, err = mgr.GetPostInfo("B94Z25jgy9w")
	if err != nil {
		t.Error(err)
		return
	}

	err = PrintPostItem(&em)
	if err != nil {
		t.Error(err)
		return
	}
}

func ExampleGetPostInfoNoLogin(t *testing.T) {
	em, err := GetPostInfoNoLogin(os.Getenv("IG_TEST_CODE"))
	if err != nil {
		t.Error(err)
		return
	}
	err = PrintPostItem(&em)
	if err != nil {
		t.Error(err)
		return
	}
}
