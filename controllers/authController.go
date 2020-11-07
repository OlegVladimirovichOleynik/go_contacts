package controllers

import (
	"encoding/json"
	"net/http"

	"go-contacts/models"
	u "go-contacts/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}                   // присваиваем переменной account данные из модели account
	err := json.NewDecoder(r.Body).Decode(account) // декодирует тело запроса в struct и завершается неудачно в случае ошибки
	// если ошибка не пустая то выдаем ее
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() // Создать аккаунт
	u.Respond(w, resp)
}

type UserEditRequest struct {
	Email string `json:"email"`
}

var EditAccount = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint)
	var req UserEditRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	var user *models.Account
	user = models.GetUser(userID)
	user.Email = req.Email
	user.Update(user.Email) // Обновить аккаунт
	resp := u.Message(true, "success")
	resp["user"] = user
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password) // вызывает функцию с входными параметрами login и password
	u.Respond(w, resp)                                    // выводит результат в json
}
