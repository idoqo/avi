package main

import (
	"fmt"
	"github.com/idoqo/avi"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	config := avi.DefaultConfig()
	config.Width = 500
	config.Height = 500
	config.FontSize = 200
	pic, err := avi.Create("UM", config)
	if err != nil {
		log.Fatal(err.Error())
	}
	out, err := os.Create(fmt.Sprintf("out-%d.png", time.Now().Unix()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer out.Close()
	png.Encode(out, pic)
}