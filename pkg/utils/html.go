package utils

import (
	"fmt"
	"html/template"
	"strings"
)

func RenderTemplate(pageData interface{}, source string) (string, error) {
	// Read and parse the HTML template file
	//tmpl, err := template.ParseFiles("../../static/add_item_form.html")
	tmpl, err := template.ParseFiles(source)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Create a strings.Builder to store the rendered template
	var renderedTemplate strings.Builder

	err = tmpl.Execute(&renderedTemplate, pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Convert the rendered template to a string
	resultString := renderedTemplate.String()

	return resultString, nil
}
