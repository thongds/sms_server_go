package v1

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/model"
)

type GetDataByCategoryRouter struct {

}

func (f GetDataByCategoryRouter )RegisterHandler(mux *httprouter.Router)  {
	http.HandleFunc("/get-data-by-category",getDataHandler)
	http.HandleFunc("/get-song-by-id",getSongIdHandler)
	http.HandleFunc("/get-menu",getMenu)
}
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	categoryData := &model.Category{}
	countryData := &model.Country{}
	songData := &[]model.Songs{}
	if r.Method == "GET"{
		query := r.URL.Query()
		category := query.Get("category")
		country := query.Get("country")
		model.GetDataByField(categoryData,"name",category)
		model.GetDataByField(countryData,"name",country)
		model.ConnectDB().Where("categories_id = ? AND countries_id = ? AND active = ?",categoryData.ID,countryData.ID,"on").Find(songData)
	}
	json.NewEncoder(w).Encode(songData)
}
func getSongIdHandler(w http.ResponseWriter, r *http.Request) {
	songData := &model.Songs{}
	if r.Method == "GET"{
		query := r.URL.Query()
		songId := query.Get("id")
		model.GetDataById(songData,songId)
	}
	json.NewEncoder(w).Encode(songData)
}
func getMenu(w http.ResponseWriter, r *http.Request)  {
	menuData := &[]model.Country{}
	if r.Method == "GET"{
		model.GetAllData(menuData)
	}
	json.NewEncoder(w).Encode(menuData)
}