package controllers

import "../views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "views/static/home.html"),
		Contact: views.NewView(
			"bootstrap", "views/static/contact.html"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
}