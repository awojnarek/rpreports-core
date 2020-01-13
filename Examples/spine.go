package main

import "github.com/awojnarek/rpreports-core/pkg/reports"

func main() {

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
