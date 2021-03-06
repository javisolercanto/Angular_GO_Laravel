package routers

import (
	"github.com/gorilla/mux"
	"products/controllers"
)

func setProductRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	return router
}