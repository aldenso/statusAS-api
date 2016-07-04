/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T19:42:40-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-03T20:35:48-04:30
 */
package main

import (
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
