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

	"github.com/ff2b/gif-maker/config"
)

const (
	DEFAULT_GIF_RATE = 20
	DEFAULT_GIF_LOOP = 0
)

func GIFEncode() string {
	var images []*image.Paletted
	var delays []int
	fname := fmt.Sprintf("%d.gif", time.Now().Unix())
	outPath := path.Join("tmp", fname)
	uriList := GetWorkFolder().GetSelectedURIs()
	delay, loop := loadGIFConfigs()

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
		delays = append(delays, delay)
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
		LoopCount: loop,
	}
	gif.EncodeAll(f, opts)

	return outPath
}

// Load config and return delay, loopFlag
// delay: 20-9999 [ms]
// loopFlag: 0(infinite loop) or 1(only 1time loop)
func loadGIFConfigs() (int, int) {
	delay := DEFAULT_GIF_RATE
	loop := DEFAULT_GIF_LOOP
	conf := config.NewConfig()
	conf.Load()
	if conf.GIFRate != delay {
		delay = conf.GIFRate
	}
	if !conf.GIFLoop {
		loop = -1
	}
	return delay, loop
}
