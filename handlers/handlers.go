package handlers

import (
	"apiGoSQL/db"
	"apiGoSQL/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	models.SendData(rw, users)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		db.Database.Save(&user)
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil {
		models.SendNotFound(rw)
	} else {
		db.Database.Delete(&user)
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	userAnt, err := getUserById(r)
	if err != nil {
		models.SendNotFound(rw)
	} else {
		userId := userAnt.Id
		user := models.User{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			models.SendUnprocessableEntity(rw)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			models.SendData(rw, user)
		}

	}
}

func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	}

	return user, nil
}
