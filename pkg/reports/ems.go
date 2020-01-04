package reports

import (
	"errors"
	"github.com/fogleman/gg"
	"log"
)

type ToxicologyReportTemplate struct {
	File      string
	X         int
	Y         int
	Case      Text
	Suspect	  Text
	Addressed Text
	Alcohol   Text
	Drug      Text
	Date      Text
	EMS       Text
	Signature Text
	Printed   Text
}

type SpineMRIReportTemplate struct {
	File string
	X    int
	Y    int
	Patient Text
	Date    Text
	Type    Text
}


func (r ToxicologyReportTemplate) GenerateToxicology() error {

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

	err = Draw(dc, r.Addressed.String, r.Addressed.X, r.Addressed.Y, r.Addressed.Font, r.Addressed.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Addressed.String)
	}

	err = Draw(dc, r.Alcohol.String, r.Alcohol.X, r.Alcohol.Y, r.Alcohol.Font, r.Alcohol.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Alcohol.String)
	}

	err = Draw(dc, r.Drug.String, r.Drug.X, r.Drug.Y, r.Drug.Font, r.Drug.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Drug.String)
	}

	err = Draw(dc, r.Date.String, r.Date.X, r.Date.Y, r.Date.Font, r.Date.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Date.String)
	}

	err = Draw(dc, r.EMS.String, r.EMS.X, r.EMS.Y, r.EMS.Font, r.EMS.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.EMS.String)
	}

	err = Draw(dc, r.Signature.String, r.Signature.X, r.Signature.Y, r.Signature.Font, r.Signature.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Signature.String)
	}

	err = Draw(dc, r.Printed.String, r.Printed.X, r.Printed.Y, r.Printed.Font, r.Printed.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Printed.String)
	}

	//img, _ := gg.LoadPNG("test.png")
	//dc.DrawImageAnchored(img, 1000, 440, 0.5, 0.5)

	dc.Clip()

	err = dc.SavePNG("ToxicologyReport.png")
	if err != nil {
		return errors.New("unable to save file ToxicologyReport.png")
	}

	return nil
}

func (r SpineMRIReportTemplate) GenerateMRISpine() error {

	im, err := gg.LoadImage(r.File)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(r.X, r.Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawImage(im, 0, 0)

	err = Draw(dc, r.Patient.String, r.Patient.X, r.Patient.Y, r.Patient.Font, r.Patient.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Patient.String)
	}

	err = Draw(dc, r.Date.String, r.Date.X, r.Date.Y, r.Date.Font, r.Date.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Date.String)
	}

	err = Draw(dc, r.Type.String, r.Type.X, r.Type.Y, r.Type.Font, r.Type.FontSize)
	if err != nil {
		return errors.New("Could not draw string: " + r.Type.String)
	}

	dc.Clip()

	err = dc.SavePNG("SpineMRIReport.png")
	if err != nil {
		return errors.New("unable to save file SpineMRIReport.png")
	}

	return nil
}