package controllers

import (
	"encoding/json"
	"innohack-backend/models"
	u "innohack-backend/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	res := acc.Create()
	u.Respond(w, res)
}

var Auth = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	res := models.Login(acc.Login, acc.Password)
	u.Respond(w, res)
}
