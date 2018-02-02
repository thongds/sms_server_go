package controller

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"strings"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
)

var (
	homeController home
	categoryController CategoryController
	countryController CountryController
	singerSetting SingerSettingController
	songDataController SongDataController
)

func Register(mux *httprouter.Router,templates map[string]*template.Template)  {
	homeController.homeTemlate = templates["index.html"]
	homeController.registerHandler(mux)
	categoryController.formsTemplate = templates["category-setting.html"]
	categoryController.registerFormController(mux)
	countryController.countryTemplate = templates["country_setting.html"]
	countryController.registerFormController()
	singerSetting.singerTemplate = templates["singer_setting.html"]
	singerSetting.registerFormController()
	songDataController.template = templates["song_data.html"]
	songDataController.registerFormController()

	http.Handle("/images/", http.FileServer(http.Dir("../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../public")))
	http.Handle("/js/", http.FileServer(http.Dir("../public")))
	http.Handle("/node_modules/", http.FileServer(http.Dir("../public")))
}
func progressFile(req *http.Request,key string) string {
	var fname string
	mf, fh, err := req.FormFile(key)
	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()
	// create sha for file name
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	fname = fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	// create new file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "files/avatars", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)
	return fname

}
