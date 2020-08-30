package avi

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"io/ioutil"
)

type Config struct {
	Width     int
	Height    int
	HexColors []string
	Font      *truetype.Font
	FontSize  float64
}

func NewConfig(width, height int, fontFile string, fontSize float64, colors []string) (*Config, error) {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	c := &Config{
		Width:     width,
		Height:    height,
		HexColors: colors,
		Font:      font,
		FontSize:  fontSize,
	}
	return c, nil
}

func DefaultConfig() *Config {
	colors := []string{"#002b36"}
	regular, _ := freetype.ParseFont(goregular.TTF)

	return &Config{
		Width:     100,
		Height:    100,
		HexColors: colors,
		Font:      regular,
		FontSize:  48.0,
	}
}
