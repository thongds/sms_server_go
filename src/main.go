package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"os"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"github.com/controller"
	"github.com/database"
	"github.com/api/v1"
)

var templates map[string]*template.Template
var db database.Mysql
var router v1.Router
func main() {
	templates = populateTemplates()
	mux := httprouter.New()
	controller.Register(mux,templates)
	router.RegisterRouter(mux)
	http.ListenAndServe(":8080", nil)
}
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "../templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html",basePath+"/_sidebar.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}