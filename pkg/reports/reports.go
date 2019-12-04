package reports

import (
	"fmt"
)


type PoliceReportTemplate struct {
	File             string
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
	Text string
	Font string
	X    int
	Y    int
}


func (r PoliceReportTemplate) GeneratePolice() {
	fmt.Println(r)

	/*
		const X = 1200
		const Y = 1650

		t := time.Now()

		im, err := gg.LoadImage("police_report_template.jpeg")
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(X, Y)

		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)

		if err := dc.LoadFontFace("standard.ttf", 30); err != nil {
			panic(err)
		}

		dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
		dc.DrawImage(im, 0, 0)

	// Case number
	dc.DrawStringAnchored("214324-4-234", 400, 460, 0.5, 0.5)
	// Date
	dc.DrawStringAnchored(t.Format("2006-01-02 15:04:05"), 850, 460, 0.5, 0.5)
	// Suspect
	dc.DrawStringAnchored("Anderson Smith", 520, 610, 0.5, 0.5)
	// Offense #1
	dc.DrawStringAnchored("Lewd Behavior", 520, 760, 0.5, 0.5)

	// Offense #2
	dc.DrawStringAnchored("Public Intoxication", 520, 850, 0.5, 0.5)

	// Offense #3
	dc.DrawStringAnchored("N/A", 520, 940, 0.5, 0.5)

	// Offense #4
	dc.DrawStringAnchored("N/A", 520, 1030, 0.5, 0.5)

	// Offense #5
	dc.DrawStringAnchored("N/A", 520, 1120, 0.5, 0.5)

	// Reporting Officer
	dc.DrawStringAnchored("Lt. Farva", 650, 1360, 0.5, 0.5)

	// Signature
	if err := dc.LoadFontFace("signature.ttf", 45); err != nil {
		panic(err)
	}

	dc.DrawStringAnchored("Lt. Farva", 550, 1460, 0.5, 0.5)

	dc.Clip()

	dc.SavePNG("out.png")

	*/
}
