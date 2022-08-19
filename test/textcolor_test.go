package test

import (
	"fmt"
	"github.com/wpf1118/toolbox/tools"
	"testing"
)

func Test_color(t *testing.T) {
	s := tools.Green("this is a test")

	fmt.Println(s)
}
