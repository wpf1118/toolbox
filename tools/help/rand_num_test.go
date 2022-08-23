package help

import (
	"fmt"
	"testing"
)

func TestRandLt(t *testing.T) {
	max := 255

	for i := 0; i < 100; i++ {
		fmt.Println(RandLt(max))

	}
}
