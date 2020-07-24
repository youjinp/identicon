package main

import (
	"image/png"
	"os"

	"github.com/youjinp/identicon"
)

func main() {
	ic := identicon.New("abc", &identicon.Options{
		BackgroundColor: identicon.RGB(220, 220, 220),
		Margin:          20,
	})

	fi, _ := os.Create("my-file.png")
	png.Encode(fi, ic)
}
