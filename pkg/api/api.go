// API приложения GoNews.
package api

import (
	"ValidateComment/pkg/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type API struct {
	db *db.DB
	r  *mux.Router
}

type News struct {
	Id    int
	Title string
}

// Конструктор API.
func New(db *db.DB) *API {
	a := API{db: db, r: mux.NewRouter()}
	a.endpoints()
	return &a
}

// Router возвращает маршрутизатор для использования
// в качестве аргумента HTTP-сервера.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.HandleFunc("/validate", api.validate).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) validate(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var comment db.PostComment
	err = json.Unmarshal(b, &comment)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(comment)

	lookFor := []string{
		"плохое",
		"слово",
	}

	for _, v := range lookFor {
		if !strings.Contains(comment.Comment, v) {
			return
		}
	}

	err = api.db.Validate(comment.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
