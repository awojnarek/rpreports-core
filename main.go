package main

import (
	"github.com/fogleman/gg"
	"log"
)

func main() {

	const X = 1200
	const Y = 1650
	im, err := gg.LoadImage("police_report_template.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetFontFace()
	dc.DrawStringAnchored("Hello, world!", X/2, Y/2, 0.5, 0.5)

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("Hello, world!", X/2, Y/2, 0.5, 0.5)
	dc.Clip()
	dc.FontHeight()
	dc.SavePNG("out.png")

}
