package test

import (
	"fmt"
	"github.com/wpf1118/toolbox/toolbox"
	"testing"
)

func Test_color(t *testing.T) {
	s := toolbox.Green("this is a test")

	fmt.Println(s)
}
