package instago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var saveRawJsonByte = false

// If set to true, the JSON data returned by API endpoint will be saved. For
// development purpose. default is false.
func SetSaveRawJsonByte(b bool) {
	saveRawJsonByte = b
}

func SaveRawJsonByte(prefix string, b []byte) (err error) {
	filename := prefix + time.Now().Format(time.RFC3339) + ".json"
	return ioutil.WriteFile(filename, b, 0644)
}

func jsonPrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func printTimestamp(timestamp int64) {
	fmt.Println(FormatTimestamp(timestamp))
}

func printPostCount(c int, url string) {
	url = strings.Replace(url, "__a=1&", "", 1)
	fmt.Printf("Getting %d from %s ...\n", c, url)
}
