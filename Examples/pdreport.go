package main

import "github.com/awojnarek/rpreports-core/pkg/reports"

func main() {

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
}
