/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:42:40-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-05T13:07:17-04:30
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/aldenso/statusAS-api/models"
	"github.com/gorilla/handlers"
)

var (
	tomlfile      = "config.toml"
	template      bool
	apiservername string
	apiserverport int
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example config.toml file")
}

func readTomlFile(tomlfile string) (*models.Tomlconfig, error) {
	var config *models.Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	flag.Parse()
	if template {
		models.CreateTemplate()
	}
	config, err := readTomlFile(tomlfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	apiservername = config.APIServer.Name
	apiserverport = config.APIServer.Port

	router := NewRouter() // this func is in router.go
	defer Session.Close()

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(apiservername+":"+strconv.Itoa(apiserverport), handlers.CORS()(loggedRouter)))
}
