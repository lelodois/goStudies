package main

import (
	"html/template"
	"os"
)

func main() {
	templateHtml := template.Must(template.ParseFiles("tpl.html"))
	_ = templateHtml.ExecuteTemplate(os.Stdout, "tpl.gohtml", "This is parameter value.")
}
