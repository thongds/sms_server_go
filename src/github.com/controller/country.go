package controller

import (
	"net/http"
	"github.com/model"
	"html/template"
	"strconv"
)

type CountryController struct {
	countryTemplate *template.Template
}
type CountryDataResponse struct {
	CountryData *[]model.Country
	CountryEdit *model.Country
}
var countryModel model.Country
func (f CountryController) registerFormController(){
	http.HandleFunc("/country-setting",f.handleRequest)
}

func (f CountryController) handleRequest(w http.ResponseWriter, r *http.Request){
	var editData = &model.Country{}
	if r.Method == "POST"{
		r.ParseForm()
		country := r.FormValue("country")
		active := r.FormValue("active")
		idString := r.FormValue("id")
		id,err := strconv.ParseInt(idString,0,4)
		if err == nil && id >0 {
			countryById := *countryModel.GetCountryById(idString)
			countryById.Active = active
			countryById.Name = country
			model.ConnectDB().Save(countryById)
		}else{
			if active == ""{
				active = "off"
			} else {
				active = "on"
			}
			var newCountryData = &model.Country{}
			newCountryData.Name = country
			newCountryData.Active = active
			model.InsertNewData(newCountryData)
		}
	}
	if r.Method == "GET"{
		query := r.URL.Query()
		action := query["action"]
		id := query["id"]
		country := &model.Country{}
		if action != nil && len(action)>0{
			switch action[0] {
			case "on":
				//country := *countryModel.GetCountryById(id[0])
				model.GetDataById(country,id[0])
				country.Active = "on"
				model.ConnectDB().Save(country)
			case "off":
				//country := *countryModel.GetCountryById(id[0])
				model.GetDataById(country,id[0])
				country.Active = "off"
				model.ConnectDB().Save(country)
			case "edit":
				//editData = *countryModel.GetCountryById(id[0])
				model.GetDataById(editData,id[0])
			}
		}
	}
	countryData := &[]model.Country{}
	model.GetAllData(countryData)
	f.countryTemplate.Execute(w,CountryDataResponse{CountryData:countryData,CountryEdit: editData})

}