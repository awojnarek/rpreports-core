package reports

import (
	"errors"
	"github.com/fogleman/gg"
	"log"
)

type PoliceReportTemplate struct {
	File             string
	X                int
	Y                int
	Case             Text
	Date             Text
	Suspect          Text
	OffenseOne       Text
	OffenseTwo       Text
	OffenseThree     Text
	OffenseFour      Text
	OffenseFive      Text
	ReportingOfficer Text
	Signature        Text
}

type Text struct {
	String string
	Font   string
	X      float64
	Y      float64
}

func Draw(dc *gg.Context, t string, x float64, y float64, font string) error {

	if err := dc.LoadFontFace(font, 45); err != nil {
		return errors.New("Unable to load font: " + font)
	}

	dc.DrawStringAnchored(t, x, y, 0.5, 0.5)

	return nil
}

func (r PoliceReportTemplate) GeneratePolice() {

	im, err := gg.LoadImage(r.File)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(r.X, r.Y)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.DrawImage(im, 0, 0)

	Draw(dc, r.Case.String, r.Case.X, r.Case.Y, r.Case.Font)
	Draw(dc, r.Suspect.String, r.Suspect.X, r.Suspect.Y, r.Suspect.Font)
	Draw(dc, r.Date.String, r.Date.X, r.Date.Y, r.Date.Font)
	Draw(dc, r.OffenseOne.String, r.OffenseOne.X, r.OffenseOne.Y, r.OffenseOne.Font)
	Draw(dc, r.OffenseTwo.String, r.OffenseTwo.X, r.OffenseTwo.Y, r.OffenseTwo.Font)
	Draw(dc, r.OffenseThree.String, r.OffenseThree.X, r.OffenseThree.Y, r.OffenseThree.Font)
	Draw(dc, r.OffenseFour.String, r.OffenseFour.X, r.OffenseFour.Y, r.OffenseFour.Font)
	Draw(dc, r.OffenseFive.String, r.OffenseFive.X, r.OffenseFive.Y, r.OffenseFive.Font)
	Draw(dc, r.ReportingOfficer.String, r.ReportingOfficer.X, r.ReportingOfficer.Y, r.ReportingOfficer.Font)
	Draw(dc, r.Signature.String, r.Signature.X, r.Signature.Y, r.Signature.Font)

	dc.Clip()

	dc.SavePNG("out.png")

}
