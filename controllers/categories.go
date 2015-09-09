package controllers

import (
	"net/http"
	"text/template"
	"gowebapp/viewmodels"
	"cut2it/externals/gorilla/mux"
	"strconv"
	"fmt"
)

type categoriesController struct{
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request){
	vm := viewmodels.GetCategories()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}


type categoryController struct{
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request){
	//using mux to get id from URL
	vars := mux.Vars(req)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)

	fmt.Println("id:" , id)

	//Get info base on the product id
	if err == nil{
		vm := viewmodels.GetProducts(1)
		w.Header().Add("Content Type", "text/html")
		this.template.Execute(w, vm)
	}else{
		w.WriteHeader(404)
	}
}