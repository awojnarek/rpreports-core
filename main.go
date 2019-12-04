package main

import "github.com/awojnarek/rpreports-core/pkg/reports"

func main() {

  x := reports.PoliceReportTemplate{
	  File:             "police_report_template.jpeg",
	  X: 				1200,
	  Y:                1650,
	  Case:             reports.Text{String: "123-123", Font: "standard.ttf", X: 400, Y: 455},
	  Date:             reports.Text{String: "12-3-2019 11:00:00", Font: "standard.ttf", X: 850, Y: 455 },
	  Suspect:          reports.Text{String: "Sudo Kill", Font: "standard.ttf", X: 520, Y: 610 },
	  OffenseOne:       reports.Text{String: "Larceny", Font: "standard.ttf", X: 520 , Y: 760  },
	  OffenseTwo:       reports.Text{String: "Resisting Arrest", Font: "standard.ttf", X: 520, Y: 850 },
	  OffenseThree:     reports.Text{String: "N/A", Font: "standard.ttf", X: 520, Y: 940 },
	  OffenseFour:      reports.Text{String: "N/A", Font: "standard.ttf", X: 520, Y: 1030 },
	  OffenseFive:      reports.Text{String: "N/A", Font: "standard.ttf", X: 520, Y: 1120 },
	  ReportingOfficer: reports.Text{String: "John Smith", Font: "standard.ttf", X: 650, Y: 1360 },
	  Signature:        reports.Text{String: "John Smith", Font: "signature.ttf", X: 550, Y: 1460 },
  }

  x.GeneratePolice()
}
