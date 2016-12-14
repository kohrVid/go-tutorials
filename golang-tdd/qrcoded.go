package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello QR Code")

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, "555-2368", Version(1))
}

func GenerateQRCode(w io.Writer, code string, version Version) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}

type Version int8
