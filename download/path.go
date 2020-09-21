package igdl

import (
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/siongui/instago"
)

var outputDir = "Instagram"
var dataDir = "Data"

func SetOutputDir(s string) {
	outputDir = s
}

func SetDataDir(s string) {
	dataDir = s
}

func GetPostFilePath(username, id, code, url string, timestamp int64, taggedusers []instago.IGTaggedUser) string {
	userDir := path.Join(outputDir, username)
	userPostsDir := path.Join(userDir, "posts")

	return path.Join(userPostsDir, instago.GetPostFilename(username, id, code, url, timestamp, taggedusers))
}

func GetUserDir(username string) (dir string) {
	return path.Join(outputDir, username)
}

func GetUserStoryDir(username string) (dir string) {
	return path.Join(GetUserDir(username), "stories")
}

func GetStoryFilePath(username, id, code, url string, timestamp int64, rms []instago.ItemReelMention) string {
	return path.Join(GetUserStoryDir(username), instago.GetStoryFilename(username, id, code, url, timestamp, rms))
}

func getPostLiveFilePath(username, id, url, typ string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userPostLiveDir := path.Join(userDir, "postlives")
	return path.Join(userPostLiveDir, instago.BuildFilename(url, username, id, "-postlive-"+typ+"-", "", timestamp))
}

func getPostLiveMergedFilePath(vpath, apath string) string {
	filename := path.Base(vpath)
	filename = strings.Replace(filename, "video-", "", 1)
	return path.Join(path.Dir(vpath), filename)
}

// only for post/item with several photos/videos of the same TakenAt time
func appendIndexToFilename(filename string, index int) string {
	ext := path.Ext(filename)
	fne := strings.TrimSuffix(filename, ext)
	return fne + "-" + strconv.Itoa(index) + ext
}

func CreateFilepathDirIfNotExist(filepath string) {
	dir := path.Dir(filepath)
	err := CreateDirIfNotExist(dir)
	if err != nil {
		panic(err)
	}
}

func getUserProfilPicFilePath(username, id, url string, timestamp int64) string {
	url, err := instago.StripQueryString(url)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(url)
	userDir := path.Join(outputDir, username)
	filename := username + "-" + id + "-profile_pic-" + strconv.FormatInt(timestamp, 10) + ext
	return path.Join(userDir, filename)
}

func GetIdUsernameDir() string {
	return path.Join(dataDir, "ID-USERNAME")
}

func GetIdUsernamePath(id, username string) string {
	filename := id + "-" + username
	return path.Join(GetIdUsernameDir(), filename)
}

func GetReelMentionsPath(id, username string) string {
	filename := id + "-" + username
	return path.Join(dataDir, "Reel-Mentions", filename)
}

func GetScreenshotPath(id, username string) string {
	filename := username + "-" + id + "-screenshot.png"
	return path.Join(outputDir, "Auto-Screenshot", filename)
}

func GetFollowDir() string {
	return path.Join(dataDir, "Follow")
}

func getFollowingPath(id string) string {
	filename := id + "-following-" + time.Now().Format(time.RFC3339) + ".json"
	return path.Join(GetFollowDir(), filename)
}

func getFollowersPath(id string) string {
	filename := id + "-followers-" + time.Now().Format(time.RFC3339) + ".json"
	return path.Join(GetFollowDir(), filename)
}

func GetReelMediaUnixTimesInUserStoryDir(username string) (utimes []string, err error) {
	infos, err := ioutil.ReadDir(GetUserStoryDir(username))
	if err != nil {
		return
	}

	for _, info := range infos {
		if info.Mode().IsRegular() {
			filename := info.Name()
			if strings.Contains(filename, username+"-") &&
				strings.Contains(filename, "-story-") {

				// remove ext
				f1 := strings.TrimSuffix(filename, path.Ext(filename))

				pieces := strings.Split(f1, "-")
				if len(pieces) > 0 {
					utime := pieces[len(pieces)-1]
					utimes = append(utimes, utime)
				}
			}
		}
	}
	return
}
