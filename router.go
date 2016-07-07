/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:42:53-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-06T21:54:09-04:30
 */
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//NotFound responses to routes not defined
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s\t%s\t%d\t",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		http.StatusNotFound,
	)
	w.WriteHeader(http.StatusNotFound)
}

//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/v1/services", GetServices).Methods("GET")
	r.HandleFunc("/api/v1/services/{serviceID}", GetService).Methods("GET")
	r.HandleFunc("/api/v1/services", AddService).Methods("POST")
	r.HandleFunc("/api/v1/services/{serviceID}", UpdateService).Methods("PUT")
	r.HandleFunc("/api/v1/services/{serviceID}", DeleteService).Methods("DELETE")
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}
