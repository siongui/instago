package igdl

import (
	"testing"
)

func ExampleMoveExpiredStory(t *testing.T) {
	MoveExpiredStory("path/to/downloaded/files", "dir/to/move/files/to")
}
