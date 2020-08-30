package main

import (
	"fmt"
	"github.com/idoqo/avi"
	"log"
	"time"
)

func main() {
	config := avi.DefaultConfig()
	config.Width = 500
	config.Height = 500
	config.FontSize = 200
	pic, err := avi.Create("AA", config)
	if err != nil {
		log.Fatal(err.Error())
	}
	pic.Save(fmt.Sprintf("out-%d.png", time.Now().Unix()))
}