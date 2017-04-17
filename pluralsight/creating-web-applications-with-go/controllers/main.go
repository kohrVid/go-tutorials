package controllers

import (
	"bufio"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	router := mux.NewRouter()

	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/home", hc.get)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	categoryController := new(categoryController)
	categoryController.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", categoryController.get)

	productController := new(productController)
	productController.template = templates.Lookup("product.html")
	router.HandleFunc("/products/{id}", productController.get)

	http.Handle("/", router)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	switch {
	case strings.HasSuffix(path, ".css"):
		contentType = "text/css"
	case strings.HasSuffix(path, ".html"):
		contentType = "text/html"
	case strings.HasSuffix(path, ".js"):
		contentType = "application/html"
	case strings.HasSuffix(path, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(path, ".mp4"):
		contentType = "video/mp4"
	default:
		contentType = "text/plain"
	}
	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
