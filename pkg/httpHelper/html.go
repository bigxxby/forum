package httpHelper

import (
	"html/template"
	"log"
	"net/http"
)

func ParseHTMLError(w http.ResponseWriter, statusCode int, message string) {

	temp, err := template.ParseFiles("web/ui/templates/error.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Whoops, something broken :(", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.WriteHeader(statusCode)

	err = temp.Execute(w, map[string]interface{}{
		"Message":   message,
		"ErrorCode": statusCode,
	})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Whoops, something broken :(", http.StatusInternalServerError)
		return
	}
}
func RenderHTMLPage(w http.ResponseWriter, templatePath string, data interface{}) {
	temp, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Whoops, something broken :(", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = temp.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Whoops, something broken :(", http.StatusInternalServerError)
		return
	}
}
