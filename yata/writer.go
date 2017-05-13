package yata

import (
	"fmt"
	"io"
	"os"

	//"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
)

const (
	Red     = "\x1b[31m"
	NoColor = "\x1b[0m"
)

var out io.Writer = os.Stdout

func Println(a ...interface{}) {
	phosphorize := ansi.ColorFunc("green+bh")
	phosphorize("Look, I'm a CRT!")
	fmt.Fprintln(out, phosphorize("Bring back the 80s!"))
}
