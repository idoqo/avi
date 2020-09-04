package main

import (
	"fmt"
	"log"

	"github.com/idoqo/avi"
)

func main() {
	cfg := avi.DefaultConfig()
	pic, err := avi.Create("JJ", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	str, err := pic.ToBase64()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(str)
}
