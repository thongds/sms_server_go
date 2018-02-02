package v1

import "github.com/julienschmidt/httprouter"

var(
	getDataByCategory GetDataByCategoryRouter
)

type Router struct {

}

func (r Router)RegisterRouter(mux *httprouter.Router){
	getDataByCategory.RegisterHandler(mux)
}