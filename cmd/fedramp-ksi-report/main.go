package main

import (
	"flag"
	"fmt"

	"github.com/trumant/fedramp-ksi-report/internal"
)

func main() {
	// Define flags
	name := flag.String("input", "ksi.yaml", "input data YAML file")
	template := flag.String("template", "ksi_report.md.tmpl", "Markdown output template file")

	// Parse command-line flags
	flag.Parse()

	ksiMetrics, err := internal.NewKSIMetrics(*name)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
	err = internal.RenderMarkdown(ksiMetrics, *template)
	if err != nil {
		fmt.Printf("Error templating report: %v\n", err)
	}
}
