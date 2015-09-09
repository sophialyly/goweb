package controllers

import (
	"net/http"
	"os"
	"text/template"
	"bufio"
	"strings"
	"github.com/gorilla/mux"
)

func Register(templates *template.Template){
	router := mux.NewRouter()

	//register template for home
	hc :=new(homeController)
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)

	//register template for categories
	cc :=new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	//register template for category
	categoryController := new(categoryController)
	categoryController.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", categoryController.get)

	http.Handle("/", router)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}


