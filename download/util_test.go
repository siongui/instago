package igdl

import (
	"testing"
)

func TestIsCommandAvailable(t *testing.T) {
	if IsCommandAvailable("ls") == false {
		t.Error("ls command does not exist!")
	}
	if IsCommandAvailable("ls111") == true {
		t.Error("ls111 command should not exist!")
	}
}

func ExampleReadNonCommentLines(t *testing.T) {
	lines, err := ReadNonCommentLines("path/to/your/file")

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		println(line)
	}
}
