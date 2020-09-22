package instago

import (
	"errors"
	"regexp"
)

// https://www.google.com/search?q=instagram+get+query_hash
// https://github.com/mineur/instagram-parser/issues/7
// https://git.kaki87.net/KaKi87/ig-scraper/src/branch/master/index.js#L190

func GetCrxWebQueryHash() (story, unknown1, unknown2 string, err error) {
	url := "https://www.instagram.com/"
	b, err := GetHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	// find JavaScript file which contains the query hash
	patternJs := regexp.MustCompile(`\/static\/bundles\/es6\/Consumer\.js\/[a-zA-Z0-9]+?\.js`)
	jsPath := string(patternJs.Find(b))
	jsUrl := "https://www.instagram.com" + jsPath
	bJs, err := GetHTTPResponseNoLogin(jsUrl)
	if err != nil {
		return
	}

	patternStoryQH := regexp.MustCompile(`50,[a-zA-Z]="([a-zA-Z0-9]{32})",[a-zA-Z]="([a-zA-Z0-9]{32})",[a-zA-Z]="([a-zA-Z0-9]{32})",`)
	qhs := patternStoryQH.FindStringSubmatch(string(bJs))
	if len(qhs) == 4 {
		// qhs[0] is the whole string in regexp
		// qhs[1] is query_hash for story query
		story = qhs[1]
		unknown1 = qhs[2]
		unknown2 = qhs[3]
	} else {
		err = errors.New("fail to get hash string in Consumer.js")
	}
	return
}
