// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Profile struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

// @Title getProfileById
// @Description retrieves the profile information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile/{id} [get]
func getProfileById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Profile{}
	arg := Profile{Id: int64(id)}
	nstmt, err := db.PrepareNamed(`select * from profile where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getProfiles
// @Description retrieves the profile information for a certain id
// @Accept  application/json
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile [get]
func getProfiles(db *sqlx.DB) (interface{}, error) {
	ret := []Profile{}
	queryStr := "select * from profile"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postProfile
// @Description enter a new profile
// @Accept  application/json
// @Param                 Body body     Profile   true "Profile object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile [post]
func postProfile(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Profile
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO profile("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putProfile
// @Description modify an existing profileentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Profile   true "Profile object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile/{id}  [put]
func putProfile(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Profile
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE profile SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delProfileById
// @Description deletes profile information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile/{id} [delete]
func delProfile(id int, db *sqlx.DB) (interface{}, error) {
	arg := Profile{Id: int64(id)}
	result, err := db.NamedExec("DELETE FROM profile WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
