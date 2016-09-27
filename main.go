package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	pngpatched "github.com/olt/compressbench/png"
)

func readPNG(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func main() {
	im, _ := os.Create("imagemagick.txt")
	new, _ := os.Create("new.txt")
	old, _ := os.Create("old.txt")

	defer im.Close()
	defer new.Close()
	defer old.Close()

	for _, name := range os.Args[1:] {
		img, err := readPNG(name)
		if err != nil {
			log.Fatal(err)
		}

		var b1, b2 bytes.Buffer
		if err := (&png.Encoder{}).Encode(&b1, img); err != nil {
			log.Fatal(err)
		}

		if err := (&pngpatched.Encoder{}).Encode(&b2, img); err != nil {
			log.Fatal(err)
		}

		stat, err := os.Stat(name)
		if err != nil {
			log.Fatal(err)
		}
		im.WriteString(fmt.Sprintf("Benchmark%s\t1\t%f MB/s\n", name, float64(stat.Size())))
		old.WriteString(fmt.Sprintf("Benchmark%s\t1\t%f MB/s\n", name, float64(b1.Len())))
		new.WriteString(fmt.Sprintf("Benchmark%s\t1\t%f MB/s\n", name, float64(b2.Len())))
	}
}
