package controllers

import (
	"crud/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request){
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(requestBody, &user)
	result,err:=user.CreateUser();
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	result, err := models.GetUserById(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(requestBody, &user)

	objectId, err := primitive.ObjectIDFromHex(key)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user.ID  =objectId
	err = models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

func Delete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	err := models.DeleteUserById(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}