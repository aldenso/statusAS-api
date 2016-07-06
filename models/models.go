//Package models for mongo struct and toml file
package models

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Service struct
type Service struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Link        string        `bson:"link" json:"link"`
	Status      int           `bson:"status" json:"status"`
	GroupID     int           `bson:"group_id" json:"group_id"`
	Messages    []string      `bson:"messages" json:"messages"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}

// Tomlconfig struct to read toml file components.
type Tomlconfig struct {
	APIServer APIServerinfo
}

// APIServerinfo struct to configure statusAS-api
type APIServerinfo struct {
	Name        string
	Port        int
	Apitls      bool
	MongoServer string
	MongoPort   int
}

// CreateTemplate function to create a base config.toml file
func CreateTemplate() {
	template := `# Example of config Configuration
[apiserver]
name = "server1.mydom.local"
port = 8080
apitls = true
mongoserver = "serverdb.mydom.local"
mongoport = 27017
`
	tomlfile := "config.toml"
	if _, err := os.Stat(tomlfile); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(tomlfile)
			if err != nil {
				fmt.Println("Error creating config.toml file", err)
				os.Exit(1)
			}
			defer file.Close()
			if _, err := file.Write([]byte(template)); err != nil {
				fmt.Printf("Can't write message\n%v\n", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("config.toml already exist in directory.")
		os.Exit(1)
	}
	fmt.Println("config.toml created.")
	os.Exit(0)
}
