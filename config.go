package avi

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"io/ioutil"
)

type Config struct {
	width     int
	height    int
	hexColors []string
	font *truetype.Font
	fontSize float64
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
		width:     width,
		height:    height,
		hexColors: colors,
		font: font,
		fontSize: fontSize,
	}
	return c, nil
}

func DefaultConfig() *Config {
	colors := []string{"#002b36"}
	regular, _ := freetype.ParseFont(goregular.TTF)

	return &Config{
		width:     100,
		height:    100,
		hexColors: colors,
		font: regular,
		fontSize: 48.0,
	}
}
