package gui_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/vantaboard/pdfgo/gui"
)

func TestGreeting(t *testing.T) {
	out, in := gui.MakeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

	test.Type(in, "Andy")
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}
}
