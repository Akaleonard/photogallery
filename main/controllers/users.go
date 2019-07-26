package controllers

import (
	"../views"
	"net/http"
)

func NewUsers() *Users {
	return &Users {
		NewView: views.NewView("bootstrap", "views/users/new.html"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
