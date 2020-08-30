package avi

import (
	"errors"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
)

var (
	errInvalidHexCode = errors.New("invalid format")
)

func Create(initials string, config *Config)  (picture *image.RGBA, err error) {
	canvas := image.NewRGBA(image.Rect(0, 0, config.width, config.height))

	// todo: make this deterministic based on input string
	bg, err := hexToRGBA(config.hexColors[0])
	if err != nil {
		return nil, err
	}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: bg}, image.Point{}, draw.Src)

	fontSize := config.fontSize

	ctx := freetype.NewContext()
	ctx.SetDst(canvas)
	ctx.SetSrc(image.White)
	ctx.SetFont(config.font)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(canvas.Bounds())
	ctx.SetHinting(font.HintingFull)

	pt := freetype.Pt(10, 10+int(ctx.PointToFixed(fontSize)>>6))
	_, err = ctx.DrawString(initials, pt)
	if err != nil {
		return canvas, err
	}

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

