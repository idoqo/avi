## idoqo/avi
A Go package to generate user avatars based on their initials.

## Installation
```
go get github.com/idoqo/avi
```
## Documentation

## Usage
### Create and save as file
```go
config := avi.DefaultConfig()
initials := "MO"
avatar, err := avi.Create(initials, config)
if err != nil {
    log.Fatal(err.Error())
}
avatar.Save("avatar.png")
``` 
### Create and save as SVG (WIP haha)
```go
initials := "MO"
avatar, _ := avi.Create(initials, config)
svgString := avatar.ToSVG()
```
### Create and save as Base64 string
```go
initials := "MO"
cfg := avi.DefaultConfig()
pic, err := avi.Create(initials, cfg)
if err != nil {
	log.Fatal(err.Error())
}
str, err := pic.ToBase64()
if err != nil {
	log.Fatal(err.Error())
}
fmt.Println(str)
```
### Get the underlying `*image.RGBA` value
Avi exposes the picture it generates for you to further process. You can retrieve it with:
```go
rgba := avatar.Picture()
```

## Configuration
Set up image width and height
```go
width, height := 100, 100
config := avi.NewCanvas(width, height)
```

Set up font style to use
```go
// pass in a TTF file path and avi will take care of parsing.
config := avi.NewCanvas(width, height) // you can also use avi.DefaultConfig() here 
fontPath := "/path/to/your/ttf-file"
config.SetFontFace(fontPath)
```
Set image quality (for JPEG files)
```go
config := avi.NewCanvas(width, height)
config.SetJpegQuality(90)
```
- `Width`     int: Width of the image, in pixel
- `Height`    int: Height of the image, in pixel
- `HexColors` []string: array of hex color codes to use as background
- `Font`      *truetype.Font: The font face to use.
- `FontSize`  float64: Font size, take care, so it doesn't exceed `Width` and `Height`.
Alternatively, you can call `avatar.DefaultConfig()` and override individual configurations
as you deem fit. See `examples/create-image.go` for how to do that.

Thanks to [laravolt/avatar](https://github.com/laravolt/avatar) for inspiring the API design :)