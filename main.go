/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:42:40-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-04T20:03:01-04:30
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
	"github.com/gorilla/handlers"
)

var (
	tomlfile = "config.toml"
	template bool
)

var (
	apiservername string
	apiserverport int
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example config.toml file")
}

func readTomlFile(tomlfile string) (*Tomlconfig, error) {
	var config *Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	flag.Parse()
	if template {
		CreateTemplate()
	}
	config, err := readTomlFile(tomlfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	apiservername = config.APIServer.Name
	apiserverport = config.APIServer.Port

	router := NewRouter() // this func is in router.go
	defer Session.Close() // related to Session in handlers.go

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(apiservername+":"+strconv.Itoa(apiserverport), handlers.CORS()(loggedRouter)))
}
