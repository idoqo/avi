package avi

import (
	"errors"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

var (
	errInvalidHexCode = errors.New("invalid format")
)

func Create(initials string, config *Config)  (picture *image.RGBA, err error) {
	canvas := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))

	// todo: make this deterministic based on input string
	bg, err := hexToRGBA(config.HexColors[1])
	if err != nil {
		return nil, err
	}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: bg}, image.Point{}, draw.Src)

	fontSize := config.FontSize

	fontDrawer := &font.Drawer{
		Dst: canvas,
		Src: image.White,
		Face: truetype.NewFace(config.Font, &truetype.Options{
			Size: fontSize,
			Hinting: font.HintingFull,
		}),
	}
	ctx := freetype.NewContext()
	ctx.SetClip(canvas.Bounds())

	bounds, _ := fontDrawer.BoundString(initials)
	xIndex := (fixed.I(config.Width) - fontDrawer.MeasureString(initials)) / 2
	textHeight := bounds.Max.Y - bounds.Min.Y
	yIndex := fixed.I((config.Height) - textHeight.Ceil()) / 2 + fixed.I(textHeight.Ceil())
	fontDrawer.Dot = fixed.Point26_6{
		X: xIndex,
		Y: yIndex,
	}
	fontDrawer.DrawString(initials)
	return canvas, nil
}

// hexToRGBA parses a web color given by its hex RGB format.
// lifted from https://github.com/icza/gox/blob/7dc3510ae515f0a6e8479d9a382bc8bb04f3a37d/imagex/colorx/colorx.go#L14
func hexToRGBA(hex string) (c color.RGBA, err error) {
	c.A = 0xff

	if hex[0] != '#' {
		return c, errInvalidHexCode
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidHexCode
		return 0
	}

	switch len(hex) {
	case 7:
		c.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		c.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		c.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		c.R = hexToByte(hex[1]) * 17
		c.G = hexToByte(hex[2]) * 17
		c.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidHexCode
	}
	return
}

