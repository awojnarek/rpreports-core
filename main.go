package main

import "github.com/awojnarek/rpreports-core/pkg/reports"

func main() {

	/*
		    police := reports.PoliceReportTemplate{
		   	  File:             "police_report_template.jpeg",
		   	  X: 				1200,
		   	  Y:                1650,
		   	  Case:             reports.Text{String: "123-123", Font: "standard.ttf", X: 400, Y: 455, FontSize: 45},
		   	  Date:             reports.Text{String: "12-3-2019 11:00:00", Font: "standard.ttf", X: 870, Y: 455, FontSize: 34 },
		   	  Suspect:          reports.Text{String: "Sudo Kill", Font: "standard.ttf", X: 550, Y: 605, FontSize: 45 },
		   	  OffenseOne:       reports.Text{String: "Larceny", Font: "standard.ttf", X: 570 , Y: 755, FontSize: 45  },
		   	  OffenseTwo:       reports.Text{String: "Resisting Arrest", Font: "standard.ttf", X: 570, Y: 845, FontSize: 45 },
		   	  OffenseThree:     reports.Text{String: "N/A", Font: "standard.ttf", X: 570, Y: 935, FontSize: 45 },
		   	  OffenseFour:      reports.Text{String: "N/A", Font: "standard.ttf", X: 570, Y: 1025, FontSize: 45 },
		   	  OffenseFive:      reports.Text{String: "N/A", Font: "standard.ttf", X: 570, Y: 1115, FontSize: 45 },
		   	  ReportingOfficer: reports.Text{String: "John Smith", Font: "standard.ttf", X: 700, Y: 1355, FontSize: 45 },
		   	  Signature:        reports.Text{String: "John Smith", Font: "signature2.ttf", X: 700, Y: 1455, FontSize: 45 },
		     }

		     police.GeneratePolice()

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
*/

	tox := reports.SpineMRIReportTemplate{
		File:      "spinemri_1.jpg",
		X:         483,
		Y:         866,
		Patient:      reports.Text{String: "Patient: Parker", Font: "1942.ttf", X: 250, Y: 30, FontSize: 24},
		Date: 		  reports.Text{String: "March 29th 2013 22:41:02", Font: "1942.ttf", X: 250, Y: 50, FontSize: 12},
		Type: 	      reports.Text{String: "MRI Scan: Spine - Gunshot Wound", Font: "1942.ttf", X: 250, Y: 825, FontSize: 24},
	}

	tox.GenerateMRISpine()

}
