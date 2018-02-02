package controller

import (
	"html/template"
	"net/http"
	"github.com/model"
	"strconv"
)

type ResponseSongData struct {
	CategoryData *[]model.Category
	SingerData *[]model.Singer
	CountryData *[]model.Country
	SongData *[]model.Songs

}


type SongDataController struct {
	template *template.Template
}
func (f SongDataController) registerFormController(){
	http.HandleFunc("/song-data",f.handlePostForm)
}
func (f SongDataController) handlePostForm(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodPost {
		songData := model.Songs{}
		sourceFile := progressFile(req,"source_file")
		sourceLyric := progressFile(req,"source_lyric")
		image := progressFile(req,"image")
		thumb := progressFile(req,"thumb")
		req.ParseForm()
		name := req.FormValue("song_name")
		metaData := req.FormValue("meta_data")
		active := req.FormValue("active")
		categoryIdString := req.FormValue("category_id")
		categoryId,_ := strconv.ParseInt(categoryIdString,0,4)

		singerIdString := req.FormValue("singer_id")
		singerId,_ := strconv.ParseInt(singerIdString,0,4)
		countryIdString := req.FormValue("country_id")
		countryId,_ := strconv.ParseInt(countryIdString,0,4)
		if active == ""{
			active = "off"
		} else {
			active = "on"
		}
		songData.Name = name
		songData.Meta = metaData
		songData.SourceAudio = sourceFile
		songData.SourceLyric = sourceLyric
		songData.Image = image
		songData.Thumb = thumb
		songData.CategoriesId = uint(categoryId)
		songData.SingersId = uint(singerId)
		songData.CountriesId = uint(countryId)
		songData.Active = active
		model.InsertNewData(&songData)
	}
	responseData := ResponseSongData{}
	categoryData := &[]model.Category{}
	singerData := &[]model.Singer{}
	songData := &[]model.Songs{}
	countryData := &[]model.Country{}
	model.GetAllData(categoryData)
	model.GetAllData(singerData)
	model.GetAllData(songData)
	model.GetAllData(countryData)
	responseData.CategoryData = categoryData
	responseData.SingerData = singerData
	responseData.SongData = songData
	responseData.CountryData = countryData
	f.template.Execute(w,responseData)
}