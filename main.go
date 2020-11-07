package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"go-contacts/app"
	"go-contacts/controllers"
)

func main() {

	router := mux.NewRouter() // функция для определения маршрутов от библиотеки gorilla
	// метод HandleFunc сопостовляет маршрут с определенным обработчикоом
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")        // отправляет на обработчик создания аккаунта
	router.HandleFunc("/api/user/edit", controllers.EditAccount).Methods("PUT")          // отправляет на обработчик создания аккаунта
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")       // отправляет на обработчик авторизации
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")    // отправляет на обработчик создания контакта
	router.HandleFunc("/api/contacts/update", controllers.UpdateContact).Methods("POST") // отправляет на обработчик обновления контакта
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")     // выводит всех контактов авторизованного юзера

	router.Use(app.JwtAuthentication) // attach JWT auth middleware

	// router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT") // выбор порта из файла .env
	// если порта в файле .env не  указано присваиваем порт 8000
	if port == "" {
		port = "8000" // localhost
	}

	err := http.ListenAndServe(":"+port, router) // Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
