package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raveger/NuxtTodoApp/backend/api"
	"github.com/raveger/NuxtTodoApp/backend/common"
	"github.com/raveger/NuxtTodoApp/backend/model"
)

var DB *sql.DB

func updateTODO(w http.ResponseWriter, r *http.Request) {
	var todo model.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := api.UpdateTODO(todo, DB); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func deleteTODO(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err = api.DeleteTODO(id, DB); err != nil {
		http.Error(w, err.Error(), 400)
		log.Println(err)
		return
	}
}

func createTODO(w http.ResponseWriter, r *http.Request) {
	var todo model.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), 400)
		log.Println(err)
		return
	}
	if err := api.CreateTODO(todo, DB); err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}
}

func getTODOs(w http.ResponseWriter, r *http.Request) {
	todos := api.GetTODOs(DB)
	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}
}

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// }

func main() {
	DB = common.DbConn()
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		switch r.Method {
		case http.MethodGet:
			getTODOs(w, r)
		case http.MethodPost:
			createTODO(w, r)
		case http.MethodDelete:
			deleteTODO(w, r)
		case http.MethodPut:
			updateTODO(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
