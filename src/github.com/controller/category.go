package controller

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/model"
)


type CategoryController struct {
	formsTemplate *template.Template
}
type CategoryResponseData struct {
	CategoryData *[]model.Category
	CategoryEdit *model.Category
}
func (f CategoryController) registerFormController(mux *httprouter.Router){
	http.HandleFunc("/category-setting",f.handlePostForm)
}

func (f CategoryController) handleIndexForm(w http.ResponseWriter, r *http.Request){
	f.formsTemplate.Execute(w,nil)
}
func (f CategoryController) handlePostForm(w http.ResponseWriter, r *http.Request){
	categoryEdit := &model.Category{}
	if r.Method == "POST"{
		r.ParseForm()
		category := r.FormValue("category")
		active := r.FormValue("active")
		if active == ""{
			active = "off"
		} else {
			active = "on"
		}
		db := &model.Category{Active:active,Name:category}
		model.InsertNewData(db)
	}
	if r.Method == "GET"{
		query := r.URL.Query()
		action := query["action"]
		id := query["id"]
		category := &model.Category{}
		if action != nil && len(action)>0{
			switch action[0] {
			case "on":
				model.GetDataById(category,id[0])
				category.Active = "on"
				model.ConnectDB().Save(category)
			case "off":
				model.GetDataById(category,id[0])
				category.Active = "off"
				model.ConnectDB().Save(category)
			case "edit":
				model.GetDataById(categoryEdit,id[0])
			}
		}
	}
	category := &[]model.Category{}
	model.GetAllData(category)
	categoryResponse := CategoryResponseData{CategoryData:category,CategoryEdit:categoryEdit}
	f.formsTemplate.Execute(w,categoryResponse)

}
// add func to get cookie
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

