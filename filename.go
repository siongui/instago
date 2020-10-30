package instago

// This file provides methods to create file names of posts, stories, and
// postlives.

import (
	"path"
	"regexp"
	"strconv"
	"strings"
)

func BuildStoryFilename2(url, username, id, timef, times string) string {
	url, err := StripQueryString(url)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(path.Base(url))
	return username + "-" +
		id +
		"-story-" +
		timef + "-" +
		times +
		ext
}

func BuildStoryFilename(url, username, id string, timestamp int64) string {
	return BuildStoryFilename2(url, username, id, FormatTimestamp(timestamp), strconv.FormatInt(timestamp, 10))
}

func BuildFilename(url, username, id, middle, last string, timestamp int64) string {
	url, err := StripQueryString(url)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(path.Base(url))
	return username + "-" +
		id +
		middle +
		FormatTimestamp(timestamp) + "-" +
		last +
		strconv.FormatInt(timestamp, 10) +
		ext
}

func GetPostFilename(username, id, code, url string, timestamp int64, taggedusers []IGTaggedUser) (filename string) {
	filename = BuildFilename(url, username, id, "-post-", code+"-", timestamp)
	filename = AppendUsernameToFilename(username, id, filename, taggedusers)
	return
}

func AppendUsernameToFilename(username, id, filename string, appendIdUsernames []IGTaggedUser) string {
	prefix := username + "-" + id

	usednames := make(map[string]bool)
	usednames[username] = true
	for _, n := range appendIdUsernames {
		taggedname := n.Username
		newprefix := prefix + "-" + taggedname
		newfilename := strings.Replace(filename, prefix, newprefix, 1)

		// cannot use 256 here. will get filename too long error.
		// use 240
		if len(newfilename) > 240 {
			continue
		}

		if _, ok := usednames[taggedname]; ok {
			continue
		} else {
			usednames[taggedname] = true
		}

		prefix = newprefix
		filename = newfilename
	}

	return filename
}

// GetStoryFilename is the same as getStoryFilePath, except adding usernames in reel_mentions
func GetStoryFilename(username, id, code, url string, timestamp int64, rms []ItemReelMention) (filename string) {
	//filename = BuildFilename(url, username, id, "-story-", code+"-", timestamp)
	filename = BuildStoryFilename(url, username, id, timestamp)
	var appendIdUsernames []IGTaggedUser
	for _, rm := range rms {
		pair := IGTaggedUser{Id: rm.GetUserId(), Username: rm.GetUsername()}
		appendIdUsernames = append(appendIdUsernames, pair)
	}
	filename = AppendUsernameToFilename(username, id, filename, appendIdUsernames)
	return
}

// AppendIndexToFilename appends index to the filename. For example, filename
// "abcdef.jpg" with index "1" will return "abcdef-1.jpg", and filename
// "123456.mp4" with index "2" will return "123456-2.mp4".
// Only used for single post/item with several photos/videos of the same TakenAt
// time.
func AppendIndexToFilename(filename string, index int) string {
	ext := path.Ext(filename)
	fne := strings.TrimSuffix(filename, ext)
	return fne + "-" + strconv.Itoa(index) + ext
}

// GetRFC3339String returns RFC3339 time string, given the input string.
func GetRFC3339String(s string) string {
	// Google search: regex rfc3339 golang
	pattern := regexp.MustCompile(`([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`)
	return string(pattern.Find([]byte(s)))
}

func ExtractPostCodeFromFilename(filename string) (code string) {
	// remove ext
	f1 := strings.TrimSuffix(filename, path.Ext(filename))

	rfc3339s := GetRFC3339String(f1)
	pieces := strings.Split(f1, "-"+rfc3339s+"-")
	if len(pieces) != 2 {
		return
	}

	f2 := pieces[1]
	pieces = strings.Split(f2, "-")
	if len(pieces) < 2 {
		return
	}

	utime := pieces[len(pieces)-1]
	if len(utime) < 3 && len(pieces) > 2 {
		utime = pieces[len(pieces)-2]
	}

	pieces = strings.Split(f2, "-"+utime)
	if len(pieces) < 1 {
		return
	}

	return pieces[0]
}

func ExtractUsernameIdFromFilename(filename string) (username, id string) {
	pieces := strings.Split(filename, "-")
	if len(pieces) < 2 {
		return
	}
	username = pieces[0]
	id = pieces[1]
	return
}
