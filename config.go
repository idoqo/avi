package avi

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image/jpeg"
	"io/ioutil"
)

type Config struct {
	Width     int
	Height    int
	HexColors []string
	Font      *truetype.Font
	FontSize  float64
	JpegOptions *jpeg.Options
}

// NewCanvas sets up a new avatar configuration with width and height in pixels
func NewCanvas(width, height int) (*Config, error) {
	config := DefaultConfig()
	config.Width = width

	return config, nil
}

// SetFontFace accepts path to a TTF file and uses it to render the text in the image
func (config *Config) SetFontFace(fontPath string) error {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}
	config.Font = font
	return nil
}

//SetJpegQuality sets the image quality and only needed if you are saving the image as a jpeg file
// 1 <= quality <= 100 and higher is better (albeit larger in size).
func (config *Config) SetJpegQuality(quality int) {
	config.JpegOptions = &jpeg.Options{
		Quality: quality,
	}
}

// DefaultConfig returns a usable configuration for generating images
func DefaultConfig() *Config {
	colors := []string{"#002b36", "#e6194b", "#0c4c98", "#ca8866"}
	regular, _ := freetype.ParseFont(goregular.TTF)

	return &Config{
		Width:     100,
		Height:    100,
		HexColors: colors,
		Font:      regular,
		FontSize:  48.0,
		JpegOptions: &jpeg.Options{
			Quality: 100,
		},
	}
}
