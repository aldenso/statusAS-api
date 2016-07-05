/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:43:02-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-05T13:08:07-04:30
 */
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aldenso/statusAS-api/db"
	"github.com/aldenso/statusAS-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var (
	//Session Establish the main session, this comes from db.go
	Session = db.NewConnection()
	// DBNAME mongodb database name
	DBNAME = "statusAS"
	// SERVICES mongodb collection for services
	SERVICES = "services"
)

//JSONResponse function to help in responses
func JSONResponse(w http.ResponseWriter, r *http.Request, response []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if string(response) != "" {
		w.Write(response)
	}
}

//JSONError function to help in error responses
func JSONError(w http.ResponseWriter, r *http.Request, message string, code int) {
	j := map[string]string{"message": message}
	response, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}

//GetServices handler to route services
func GetServices(w http.ResponseWriter, r *http.Request) {
	var services []models.Service
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(DBNAME).C(SERVICES)
	collection.Find(bson.M{}).All(&services)
	response, err := json.MarshalIndent(services, "", "    ")
	if err != nil {
		panic(err)
	}
	JSONResponse(w, r, response, http.StatusOK)
}

// AddService handler to add new service
func AddService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	json.NewDecoder(r.Body).Decode(&service)
	if service.Name == "" || service.Description == "" {
		JSONError(w, r, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	service.ID = objID
	service.CreatedAt = time.Now()
	service.UpdatedAt = service.CreatedAt
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(DBNAME).C(SERVICES)
	err := collection.Insert(service)
	if err != nil {
		JSONError(w, r, "Failed to insert service", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(service.ID.Hex()))
	JSONResponse(w, r, []byte{}, http.StatusCreated)
}

//UpdateService handler to update a service
func UpdateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["serviceID"]) != true {
		JSONError(w, r, "bad entry for id", http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&service)
	if service.Name == "" || service.Description == "" {
		JSONError(w, r, "Incorrect body", http.StatusBadRequest)
		return
	}
	serviceID := bson.ObjectIdHex(vars["serviceID"])
	session := Session.Copy()
	defer session.Close()
	service.ID = serviceID
	service.UpdatedAt = time.Now()
	collection := session.DB(DBNAME).C(SERVICES)
	err := collection.Update(bson.M{"_id": serviceID}, &service)
	if err != nil {
		JSONError(w, r, "Could not find service "+string(serviceID.Hex())+" to update", http.StatusNotFound)
		return
	}
	JSONResponse(w, r, []byte{}, http.StatusNoContent)
}

//DeleteService handler to delete a todo
func DeleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID := bson.ObjectIdHex(vars["serviceID"])
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(DBNAME).C(SERVICES)
	err := collection.Remove(bson.M{"_id": serviceID})
	if err != nil {
		JSONError(w, r, "Could not find service "+string(serviceID.Hex())+" to delete", http.StatusNotFound)
		return
	}
	JSONResponse(w, r, []byte{}, http.StatusNoContent)
}
