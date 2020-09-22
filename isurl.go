package instago

import (
	"regexp"
)

func IsWebStoryUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/stories\/[a-zA-Z\d_.]+\/\d+\/$`)
	return re.MatchString(url)
}

func IsWebRootUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/$`)
	return re.MatchString(url)
}

func IsWebUserUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/[a-zA-Z\d_.]+\/$`)
	return re.MatchString(url)
}

func IsWebSavedUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/[a-zA-Z\d_.]+\/saved/$`)
	return re.MatchString(url)
}

func IsWebTaggedUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/[a-zA-Z\d_.]+\/tagged/$`)
	return re.MatchString(url)
}

func IsWebPostUrl(url string) bool {
	re := regexp.MustCompile(`^https:\/\/www\.instagram\.com\/p\/[a-zA-Z\d_-]+\/$`)
	return re.MatchString(url)
}
