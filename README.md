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
### Create and save as Base64 string (WIP haha)
```go
initials := "MO"
avatar, _ := avi.Create(initials, config)
base64 := avatar.ToBase64()
```
### Get the underlying `*image.RGBA` value
Avi exposes the picture it generates for you to further process. You can retrieve it with:
```go
rgba := avatar.Picture()
```

### Configuration
- `Width`     int: Width of the image, in pixel
- `Height`    int: Height of the image, in pixel
- `HexColors` []string: array of hex color codes to use as background
- `Font`      *truetype.Font: The font face to use, if you have a `TTF` file lying around,
you can simply pass the ttf file path to `avi.NewConfig()` when setting up your configuration
and the package will handle the parsing. 
- `FontSize`  float64: Font size, take care, so it doesn't exceed `Width` and `Height`.
Alternatively, you can call `avatar.DefaultConfig()` and override individual configurations
as you deem fit. See `examples/create-image.go` for how to do that.

Thanks to [laravolt/avatar] for inspiring the API design :)