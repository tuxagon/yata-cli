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

// Print will display a string with the configured io.Writer
func Print(a ...interface{}) {
	fmt.Fprint(out, a...)
}

// PrintColor will display a string styled with mgutz/ansi
// styles with the configured io.Writer
func PrintColor(style string, a ...interface{}) {
	s := fmt.Sprint(a...)
	colorizer := ansi.ColorFunc(style)
	fmt.Fprint(out, colorizer(s))
}

// Printf will display a string with the configured io.Writer
// using the provided format
func Printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}

// PrintfColor will display a string styled with mgutz/ansi
// styles with the configured io.Writer using the provided
// format
func PrintfColor(style, format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	colorizer := ansi.ColorFunc(style)
	fmt.Fprint(out, colorizer(s))
}

// Println will display a string with the configured io.Writer
func Println(a ...interface{}) {
	fmt.Fprintln(out, a...)
}

// PrintlnColor will display a string styled with mgutz/ansi
// styles with the configured io.Writer
func PrintlnColor(style string, a ...interface{}) {
	s := fmt.Sprint(a...)
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
