package instago

import (
	"encoding/json"
	"fmt"
	"strings"
)

func jsonPrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func printTimestamp(timestamp int64) {
	fmt.Println(formatTimestamp(timestamp))
}

func printPostCount(c int, url string) {
	url = strings.Replace(url, "__a=1&", "", 1)
	fmt.Printf("Getting %d from %s ...\n", c, url)
}
