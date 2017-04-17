package controllers

import (
	"github.com/gorilla/mux"
	"github.com/kohrVid/go-tutorials/pluralsight/creating-web-applications-with-go/viewmodels"
	"net/http"
	"strconv"
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
	vars := mux.Vars(req)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProduct(id)

		w.Header().Add("Content-Type", "text/html")
		this.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}
