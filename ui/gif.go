package ui

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"log"
	"os"
	"path"
	"time"
)

func GIFEncode() string {
	var images []*image.Paletted
	var delays []int
	fname := fmt.Sprintf("%d.gif", time.Now().Unix())
	outPath := path.Join("tmp", fname)
	uriList := GetWorkFolder().GetSelectedURIs()

	for _, uri := range uriList {
		reader, err := os.Open(uri.Path())
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		img, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		// Get paletted colors from Image
		paletted := image.NewPaletted(img.Bounds(), palette.WebSafe)
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				paletted.Set(x, y, img.At(x, y))
			}
		}

		images = append(images, paletted)
		delays = append(delays, 20)
	}

	// Generate GIF
	f, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	opts := &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: 0,
	}
	gif.EncodeAll(f, opts)

	return outPath
}
