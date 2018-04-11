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
