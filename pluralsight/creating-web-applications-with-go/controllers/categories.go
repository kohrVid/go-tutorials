package controllers

import (
	"github.com/kohrVid/go-tutorials/pluralsight/creating-web-applications-with-go/viewmodels"
	"net/http"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetProducts()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
