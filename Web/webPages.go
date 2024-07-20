package Web

import (
	"net/http"
	"text/template"

	"web/Lib"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	var bnStyle, inputStr string
	var tmpl *template.Template

	// Serve form at initial visit of site
	if r.Method == http.MethodGet {
		tmpl = template.Must(template.ParseFiles("static/placeHolder.html"))
		tmpl.Execute(w, nil)

		// Serve form and ascii-art/error after form submission
	} else if r.Method == http.MethodPost {

		// Extract banner style selected and text inputed in form
		bnStyle = r.FormValue("style")
		inputStr = r.FormValue("inputStr")

		// Run AsciiArt function with banner style selected and input string
		output, err := Lib.AsciiArt(inputStr, bnStyle+".txt")

		// Should there occur an error, serve errorPrinter.html with the nature of error
		if err != "" {
			tmpl = template.Must(template.ParseFiles("static/errorPrinter.html"))
			tmpl.Execute(w, struct{ Issue string }{Issue: err})

			// If no error print ascii-art below form on submitForm.html
		} else {
			//Safely load html template from submitForm.html
			tmpl = template.Must(template.ParseFiles("static/submitForm.html"))
			tmpl.Execute(w, struct{ AsciiArt, Input string }{AsciiArt: output, Input: inputStr})
		}

		// Label any requests other than 'GET' and 'POST' requests as 'invalid requests'
	} else {
		http.Error(w, "Invalid Request Method!", http.StatusMethodNotAllowed)
	}
}
