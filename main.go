package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/cssFiles/", http.FileServer(http.Dir("/HTML_Sites/css")))
	http.HandleFunc("/", basicfunchandler)
	http.ListenAndServe(":9050", nil)
}

func basicfunchandler(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("HTML_Sites/LoginForm.html"))

	err := tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
