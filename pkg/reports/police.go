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

func (r PoliceReportTemplate) GeneratePolice() error {

	im, err := gg.LoadImage(r.File)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(r.X, r.Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawImage(im, 0, 0)

	err = Draw(dc, r.Case.String, r.Case.X, r.Case.Y, r.Case.Font, r.Case.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Case.String)
	}

	err = Draw(dc, r.Suspect.String, r.Suspect.X, r.Suspect.Y, r.Suspect.Font, r.Suspect.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Suspect.String)
	}

	err = Draw(dc, r.Date.String, r.Date.X, r.Date.Y, r.Date.Font, r.Date.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Date.String)
	}

	err = Draw(dc, r.OffenseOne.String, r.OffenseOne.X, r.OffenseOne.Y, r.OffenseOne.Font, r.OffenseOne.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.OffenseOne.String)
	}

	err = Draw(dc, r.OffenseTwo.String, r.OffenseTwo.X, r.OffenseTwo.Y, r.OffenseTwo.Font, r.OffenseTwo.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.OffenseTwo.String)
	}

	err = Draw(dc, r.OffenseThree.String, r.OffenseThree.X, r.OffenseThree.Y, r.OffenseThree.Font, r.OffenseThree.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.OffenseThree.String)
	}

	err = Draw(dc, r.OffenseFour.String, r.OffenseFour.X, r.OffenseFour.Y, r.OffenseFour.Font, r.OffenseFour.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.OffenseFour.String)
	}

	err = Draw(dc, r.OffenseFive.String, r.OffenseFive.X, r.OffenseFive.Y, r.OffenseFive.Font, r.OffenseFive.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.OffenseFive.String)
	}

	err = Draw(dc, r.ReportingOfficer.String, r.ReportingOfficer.X, r.ReportingOfficer.Y, r.ReportingOfficer.Font, r.ReportingOfficer.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.ReportingOfficer.String)
	}

	err = Draw(dc, r.Signature.String, r.Signature.X, r.Signature.Y, r.Signature.Font, r.Signature.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Signature.String)
	}

	dc.Clip()

	err = dc.SavePNG("PoliceReport.png")
	if err != nil {
		return errors.New("unable to save file PoliceReport.png")
	}

	return nil
}
