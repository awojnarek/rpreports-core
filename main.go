package main

import "github.com/awojnarek/rpreports-core/pkg/reports"

func main() {

		tox := reports.ToxicologyReportTemplate{
			File:      "toxicology_template.jpeg",
			X:         1200,
			Y:         1650,
			Case:      reports.Text{String: "12-26", Font: "1942.ttf", X: 1000, Y: 315, FontSize: 36},
			Suspect:   reports.Text{String: "Parker Winslow", Font: "1942.ttf", X: 1000, Y: 360, FontSize: 36},
			Addressed: reports.Text{String: "LSPD Command", Font: "1942.ttf", X: 250, Y: 440, FontSize: 36},
			Alcohol:   reports.Text{String: "N/A", Font: "1942.ttf", X: 840, Y: 770, FontSize: 30},
			Drug:      reports.Text{String: "Drugs detected in urine: Oxycodone", Font: "1942.ttf", X: 840, Y: 880, FontSize: 30},
			Date:      reports.Text{String: "12/21/2019 11:00:23", Font: "1942.ttf", X: 340, Y: 1225, FontSize: 24},
			EMS:       reports.Text{String: "Anderson Smith", Font: "1942.ttf", X: 465, Y: 1325, FontSize: 36},
			Signature: reports.Text{String: "Anderson Smith", Font: "signature2.ttf", X: 465, Y: 1395, FontSize: 48},
			Printed:      reports.Text{String: "12/21/2019 11:00:23", Font: "1942.ttf", X: 285, Y: 1570, FontSize: 24},
		}

		tox.GenerateToxicology()

}
