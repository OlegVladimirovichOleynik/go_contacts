package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Получение идентификатора пользователя, отправившего запрос
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

type ContactID struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint) //Получение идентификатора пользователя, отправившего запрос

	var contact ContactID
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	var data *models.Contact
	data = models.GetContact(contact.ID)
	if userID == data.UserId {
		data.Name = contact.Name
		data.Phone = contact.Phone
		data.UserId = contact.UserId
		data.Update(contact.Name, contact.Phone, contact.UserId)
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	data := models.GetContacts(user)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
