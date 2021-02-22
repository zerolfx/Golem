package main

import (
	"fmt"
	"github.com/zerolfx/golem/exif"
	"os"
)

func main() {
	fhnd, err := os.Open("test.jpg")
	if err != nil {
		return
	}

	image, err := ImgMeta.ReadJpeg(fhnd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	basicInfo := ImgMeta.GetBasicInfo(image)
	fmt.Printf("IPTC Description: %v\n", basicInfo.Description)
	fmt.Printf("IPTC Keywords: %v\n", basicInfo.Keywords)
}
