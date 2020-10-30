package instago

import (
	"testing"
)

func TestGetPostFilename(t *testing.T) {
	filename := GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.jpg?abc=1", 1520056661, nil)
	if filename != "instagram-25025320-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.jpg" {
		t.Error(filename)
		return
	}

	username1 := IGTaggedUser{Id: "12345", Username: "testuser"}
	username2 := IGTaggedUser{Id: "123456", Username: "testuser111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}
	username3 := IGTaggedUser{Id: "25025320", Username: "instagram"}
	username4 := IGTaggedUser{Id: "12345", Username: "testuser"}

	taggedusers := []IGTaggedUser{username1}
	filename = GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, taggedusers)
	if filename != "instagram-25025320-testuser-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test username more than filename length 256
	taggedusers2 := []IGTaggedUser{username2}
	filename = GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.jpg", 1520056661, taggedusers2)
	if filename != "instagram-25025320-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.jpg" {
		t.Error(filename)
		return
	}

	// test username more than filename length 256
	taggedusers3 := []IGTaggedUser{username2, username1}
	filename = GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, taggedusers3)
	if filename != "instagram-25025320-testuser-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test duplicate username
	taggedusers4 := []IGTaggedUser{username3, username2, username1}
	filename = GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, taggedusers4)
	if filename != "instagram-25025320-testuser-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test duplicate username
	taggedusers5 := []IGTaggedUser{username3, username2, username1, username4}
	filename = GetPostFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, taggedusers5)
	if filename != "instagram-25025320-testuser-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}
}

func TestBuildStoryFilename(t *testing.T) {
	filename := BuildStoryFilename("123.jpg?abc=1", "instagram", "25025320", 1520056661)
	if filename != "instagram-25025320-story-2018-03-03T13:57:41+08:00-1520056661.jpg" {
		t.Error(filename)
		return
	}
}

func TestGetStoryFilename(t *testing.T) {
	filename := GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, nil)
	if filename != "instagram-25025320-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	user1 := IGUser{Pk: 12345, Username: "testuser"}
	user2 := IGUser{Pk: 123456, Username: "testuser111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}
	user3 := IGUser{Pk: 25025320, Username: "instagram"}
	user4 := IGUser{Pk: 12345, Username: "testuser"}

	rms := []ItemReelMention{{User: user1}}
	filename = GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, rms)
	if filename != "instagram-25025320-testuser-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test username more than filename length 256
	rms2 := []ItemReelMention{{User: user2}}
	filename = GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, rms2)
	if filename != "instagram-25025320-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test username more than filename length 256
	rms3 := []ItemReelMention{{User: user2}, {User: user1}}
	filename = GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, rms3)
	if filename != "instagram-25025320-testuser-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test duplicate username
	rms4 := []ItemReelMention{{User: user3}, {User: user2}, {User: user1}}
	filename = GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, rms4)
	if filename != "instagram-25025320-testuser-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}

	// test duplicate username
	rms5 := []ItemReelMention{{User: user3}, {User: user2}, {User: user1}, {User: user4}}
	filename = GetStoryFilename("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661, rms5)
	if filename != "instagram-25025320-testuser-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(filename)
		return
	}
}

func TestAppendIndexToFilename(t *testing.T) {
	nf := AppendIndexToFilename("instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661.mp4", 0)
	if nf != "instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661-0.mp4" {
		t.Error(nf)
	}
	nf = AppendIndexToFilename("instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661.mp4", 1)
	if nf != "instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661-1.mp4" {
		t.Error(nf)
	}
}

func TestGetRFC3339String(t *testing.T) {
	s := "testuser-1234567-post-2018-06-10T09:41:35+08:00--j03-x3BXga-1528594895.jpg"
	sp := GetRFC3339String(s)
	if sp != "2018-06-10T09:41:35+08:00" {
		t.Error(sp)
		return
	}
}

func TestExtractPostCodeFromFilename(t *testing.T) {
	s := "testuser-1234567-post-2018-06-10T09:41:35+08:00--j03-x3BXga-1528594895.jpg"
	code := ExtractPostCodeFromFilename(s)
	if code != "-j03-x3BXga" {
		t.Error(code)
		return
	}

	s2 := "instagram-25025320-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4"
	code = ExtractPostCodeFromFilename(s2)
	if code != "Bh7kySfDYq8" {
		t.Error(code)
		return
	}

	s3 := "instagram-25025320-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq--1520056661-2.mp4"
	code = ExtractPostCodeFromFilename(s3)
	if code != "Bh7kySfDYq-" {
		t.Error(code)
		return
	}
}

func TestExtractUsernameIdFromFilename(t *testing.T) {
	s := "instagram-25025320-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.jpg"
	username, id := ExtractUsernameIdFromFilename(s)
	if username != "instagram" || id != "25025320" {
		t.Error(username, id)
		return
	}
}
