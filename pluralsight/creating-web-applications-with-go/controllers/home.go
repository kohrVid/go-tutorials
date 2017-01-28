package controllers

import (
	"github.com/kohrVid/go-tutorials/pluralsight/creating-web-applications-with-go/viewmodels"
	"net/http"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
