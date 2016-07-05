//Package db for mongodb
package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/aldenso/statusAS-api/models"
	"gopkg.in/mgo.v2"
)

var (
	tomlfile    = "config.toml"
	mongoserver string
	mongoport   int
)

func readTomlFile(tomlfile string) (*models.Tomlconfig, error) {
	var config *models.Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
}

// NewConnection create connection to DB
func NewConnection() *mgo.Session {
	config, err := readTomlFile(tomlfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mongoserver = config.APIServer.MongoServer
	mongoport = config.APIServer.MongoPort
	session, err := mgo.Dial(mongoserver + ":" + strconv.Itoa(mongoport))
	//session, err := mgo.Dial("172.17.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
