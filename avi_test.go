package avi

import (
	"image/color"
	"testing"
)

func TestValidHexToRGBA(t *testing.T) {
	hex := "#002b36"
	expected := color.RGBA{R: 0x0, G: 0x2b, B: 0x36, A: 0xff}
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

func TestToBase64(t *testing.T) {
	colors := []string{"#0c4c98"}
	cfg := DefaultConfig()
	cfg.HexColors = colors
	cfg.Width = 100
	cfg.Height = 100
	avatar, err := Create("JJ", cfg)
	if err != nil {
		t.Errorf("could not create avatar: %s", err.Error())
	}

	result, err := avatar.ToBase64()
	if err != nil {
		t.Errorf("could not get base64 string: %s", err.Error())
	}
	expected := "iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAIAAAD/gAIDAAAClklEQVR4nOzcPWgTYRzHcZtcrDGhiSXU2BQpSoWjUNChdtPBIoiDg6CTddDNyUHn7r7g1sFF3UTRpSI6OOgSUaJYsUshlBoJNVRJUojkxa0JGtL7xSf3tPL9TL3m4X/JlzuSwEOc6Om5HfAmYPsJbCfEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBLQCwBsQTEEhBL4PhzmnPHx1KJaIcFtx9nGg3/5nTHp1iXTo1PuXs7LLjz5GOtVvdtTne4DQXEEvh0G05fe/rHf5YeXBiKh23N6Q5XloBYAmIJiCUgloBYAmIJiCUgloBYAmIJiCUgloBYAmIJiCUgloBYAmIJiCUgloBYAmIJiCUglmCrxGoY2vtiak5b1mI5weap1yvVer3LF2lqjhfWYg3sDm38vVasWJ/jhZ1Y0XCo9YrIFcp253hkJ9aBfbHWw/Ri3u4cj+zEOjYx3Hr45lPO7hyPDMcaiof7+jZZ4wQDF0+6G4crq6Xn75Z7NMcsw7GunJnIzJ2/evbwWCredkFyMHL/+vShkeajNx99+HsXqKk5Zpnf+XdwODY7Mzk7M/n1e3khW1jK/fxR+hVyArHIzvHRwSk3GQw0r5n5dPbus889nWNQD7dJphKRVCLSYcF8Onv51ivf5vw7w7FWVkvVWr317bytbL5442Hm3osvvZ5jVp/xH6KORfpPHBk56ibd/XtGkwPxaH90l7Neqa4VK7lC+e1i/vXCt5fvlzf9qG1qjkHmY/3HtsoX6W2BWAJiCYglIJaAWAJiCYglIJaAWAJiCYglIJaAWAJiCYglIJaAWAJiCYglIJbgdwAAAP//xYrfF8X1PV4AAAAASUVORK5CYII="
	if result != expected {
		t.Error("did not get expected base64 string.")
	}
}
