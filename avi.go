package avi

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	errInvalidHexCode = errors.New("invalid format")
)

type Avi struct {
	picture *image.RGBA
	config  *Config
}

//Create generates an avatar for initials given a valid configuration
func Create(initials string, config *Config) (avi *Avi, err error) {
	avi = &Avi{config: config}
	canvas := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))

	bg, err := colorByText(initials, config.HexColors)
	if err != nil {
		return nil, err
	}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: bg}, image.Point{}, draw.Src)

	fontSize := config.FontSize

	fontDrawer := &font.Drawer{
		Dst: canvas,
		Src: image.White,
		Face: truetype.NewFace(config.Font, &truetype.Options{
			Size:    fontSize,
			Hinting: font.HintingFull,
		}),
	}

	bounds, _ := fontDrawer.BoundString(initials)
	xIndex := (fixed.I(config.Width) - fontDrawer.MeasureString(initials)) / 2
	textHeight := bounds.Max.Y - bounds.Min.Y
	yIndex := fixed.I((config.Height)-textHeight.Ceil())/2 + fixed.I(textHeight.Ceil())
	fontDrawer.Dot = fixed.Point26_6{
		X: xIndex,
		Y: yIndex,
	}
	fontDrawer.DrawString(initials)
	avi.picture = canvas
	return avi, nil
}

// Save saves a generated avatar as `filename`. The file-type is
// guessed based on the filename extension
func (avi *Avi) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
	case ".jpg", "jpeg":
		err = jpeg.Encode(f, avi.picture, avi.config.JpegOptions)
		break
	case ".png":
		err = png.Encode(f, avi.picture)
		break
	default:
		return fmt.Errorf("unsupported file format")
	}
	return err
}

//ToSVG returns the svg string representation of the generated picture
func (avi *Avi) ToSVG() (string, error) {
	//todo
	return "", fmt.Errorf("not implemented")
}

//ToBase64 returns the base64 string equivalent of the generated picture (or any error that occurs
//in the process)
func (avi *Avi) ToBase64() (string, error) {
	var (
		err error
		str string
	)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, avi.picture)
	str = base64.StdEncoding.EncodeToString(buf.Bytes())
	return str, err
}

// Picture returns the underlying image instance
func (avi *Avi) Picture() *image.RGBA {
	return avi.picture
}

func colorByText(text string, colorBucket []string) (c color.RGBA, err error) {
	numValue, err := numberFromText(text)
	if err != nil {
		c := color.RGBA{}
		return c, err
	}
	hexCode := colorBucket[numValue%len(colorBucket)]
	return hexToRGBA(hexCode)
}

func numberFromText(text string) (int, error) {
	charCodes := ""
	for _, ch := range text {
		charCodes += strconv.Itoa(int(ch))
	}
	return strconv.Atoi(charCodes)
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
