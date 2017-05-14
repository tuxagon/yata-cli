package yata

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	//"github.com/mattn/go-colorable"`
	"github.com/mgutz/ansi"
)

var (
	out io.Writer = os.Stdout
	in  io.Reader = os.Stdin
)

// Printf will display a string with the configured io.Writer
// using the provided format
func Printf(format, s string) {
	fmt.Printf(format, s)
}

// PrintfColor will display a string styled with mgutz/ansi
// styles with the configured io.Writer using the provided
// format
func PrintfColor(style, format, s string) {
	colorizer := ansi.ColorFunc(style)
	fmt.Fprintf(out, format, colorizer(s))
}

// Println will display a string with the configured io.Writer
func Println(s string) {
	fmt.Fprintln(out, s)
}

// PrintlnColor will display a string styled with mgutz/ansi
// styles with the configured io.Writer
func PrintlnColor(style, s string) {
	colorizer := ansi.ColorFunc(style)
	fmt.Fprintln(out, colorizer(s))
}

// Readln will get input terminated with a newline from
// the configured io.Reader
func Readln() (str string) {
	r := bufio.NewReader(in)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err.Error())
	}
	str = string(line)
	return
}

// ReadInt will read an integer from the configured io.Reader
func ReadInt() (n int) {
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
