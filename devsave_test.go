package instago

// This file is for development purpose

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var jsonKind = flag.String("jsonKind", "", "what kind of json to save?")

func saveJson(b []byte, prefix string) (filepath string, err error) {
	filename := formatTimestamp(time.Now().Unix()) + ".json"
	filepath = "/tmp/" + prefix + filename
	err = ioutil.WriteFile(filepath, b, 0644)
	return
}

func TestSaveJson(t *testing.T) {
	switch *jsonKind {
	case "reels_tray":
		b, err := getHTTPResponse(urlReelsTray,
			os.Getenv("IG_DS_USER_ID"),
			os.Getenv("IG_SESSIONID"),
			os.Getenv("IG_CSRFTOKEN"))
		if err != nil {
			t.Error(err)
			return
		}

		fp, err := saveJson(b, "reels-tray-")
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(fp + " saved!")

	default:
	}
}
