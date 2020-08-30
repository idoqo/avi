package avi

import (
	"image/color"
	"testing"
)

func TestValidHexToRGBA(t *testing.T) {
	hex := "#002b36"
	expected := color.RGBA{R:0x0, G:0x2b, B:0x36, A:0xff}
	c, err := hexToRGBA(hex)
	if err != nil {
		t.Errorf("expected %s to be a valid hex code, got %s", hex, err.Error())
	}
	if &c == nil {
		t.Error("expected c to be a valid RGBA, got nil")
	}
	if c.R != 0x0 || c.G != 0x2b || c.B != 0x36 || c.A != 0xff {
		t.Errorf("expected %#v, got %#v", expected, c)
	}
}

func TestInvalidHexToRGBA(t *testing.T) {
	hex := "00"
	_, err := hexToRGBA(hex)
	if err == nil {
		t.Errorf("expected %s to produce an error.", hex)
	}
}

func TestNumberFromText(t *testing.T) {
	text := "AA"
	expected := 6565
	num, err := numberFromText(text)
	if err != nil {
		t.Errorf("expecting conversion, got %s", err.Error())
	}

	if num != expected {
		t.Errorf("expected %s to produce %d, got %d", text, expected, num)
	}
}
