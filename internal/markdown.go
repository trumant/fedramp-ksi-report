package internal

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func RenderMarkdown(ksiMetrics *KSIMetrics, templatePath string) error {
	funcMap := template.FuncMap{
		"formatPercentage": FormatPercentage,
	}
	name := path.Base(templatePath)
	template := template.Must(template.New(name).Funcs(funcMap).ParseFiles(templatePath))
	return template.Execute(os.Stdout, ksiMetrics)
}

func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value*100)
}
