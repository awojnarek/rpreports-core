package reports

import (
	"errors"
	"github.com/fogleman/gg"
)

type Text struct {
	String   string
	Font     string
	FontSize float64
	X        float64
	Y        float64
}

func Draw(dc *gg.Context, t string, x float64, y float64, font string, size float64) error {

	if err := dc.LoadFontFace(font, size); err != nil {
		return errors.New("Unable to load font: " + font)
	}

	dc.DrawStringAnchored(t, x, y, 0.5, 0.5)

	return nil
}