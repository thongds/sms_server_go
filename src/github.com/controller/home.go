package controller

import (
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"github.com/database"
)

type home struct {
	homeTemlate *template.Template
	a           database.Mysql
}

func (h home) registerHandler(mux *httprouter.Router){
	//mux.GET("/",h.handlerIndex)
	http.HandleFunc("/",h.handlerIndex)
}

func  (h home)handlerIndex(w http.ResponseWriter,r *http.Request){
	h.homeTemlate.Execute(w,nil)

}
