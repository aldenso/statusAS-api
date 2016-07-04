/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:42:40-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-03T21:13:49-04:30
 */
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	router := NewRouter() // this func is in router.go
	defer Session.Close() // related to Session in handlers.go

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(loggedRouter)))
}
