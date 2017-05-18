package yata

import (
	"bytes"
	"testing"

	"github.com/mgutz/ansi"
)

func TestPrint(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := "Kamehameha"

	Print("Kamehameha")
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestPrintColor(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := ansi.ColorFunc("green+h")("Kamehameha")

	PrintColor("green+h", "Kamehameha")
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestPrintf(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := "42"

	Printf("%d", 42)
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestPrintfColor(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := ansi.ColorFunc("green+h")("42")

	PrintfColor("green+h", "%d", 42)
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestPrintln(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := "You shall not pass!\n"

	Println("You shall not pass!")
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestPrintlnColor(t *testing.T) {
	b := &bytes.Buffer{}
	out = b
	expected := ansi.ColorFunc("green+h")("You shall not pass!") + "\n"

	PrintlnColor("green+h", "You shall not pass!")
	if b.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, b.String())
	}
}

func TestReadln(t *testing.T) {
	b := &bytes.Buffer{}
	in = b
	expected := "Inspector Spacetime"
	b.WriteString(expected + "\n")

	actual := Readln()
	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestReadInt(t *testing.T) {
	b := &bytes.Buffer{}
	in = b
	b.WriteString("42\n")

	actual := ReadInt()
	if actual != 42 {
		t.Errorf("expected %d, got %d", 42, actual)
	}
}
