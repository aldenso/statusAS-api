/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:43:11-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-03T20:36:08-04:30
 */
package main

import "gopkg.in/mgo.v2"

var (
	// DBNAME mongodb database name
	DBNAME = "statusAS"
	// SERVICES mongodb collection for services
	SERVICES = "services"
)

// NewConnection create connection to DB
func NewConnection() *mgo.Session {
	session, err := mgo.Dial("localhost")
	//session, err := mgo.Dial("172.17.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
