package controller

import (
	"html/template"
	"net/http"
	"github.com/model"
)

type SingerSettingController struct {
	singerTemplate *template.Template
}

func (f SingerSettingController) registerFormController(){
	http.HandleFunc("/singer-setting",f.handlePostForm)
}
func (f SingerSettingController) handlePostForm(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodPost {
		var singerData = &model.Singer{}
		avatartLink := progressFile(req,"avatar_file")
		req.ParseForm()
		singerName := req.FormValue("singer_name")
		metaData := req.FormValue("meta_data")
		active := req.FormValue("active")
		if active == ""{
			active = "off"
		} else {
			active = "on"
		}
		singerData.Active = active
		singerData.Avatar = avatartLink
		singerData.Name = singerName
		singerData.Meta = metaData

		model.InsertNewData(singerData)

	}
	f.singerTemplate.Execute(w,nil)
}

