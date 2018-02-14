/*
 * Copyright (c) 2018.
 */

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ivanj4u/apiswitch/database"
	"github.com/ivanj4u/apiswitch/dto"
	"github.com/ivanj4u/apiswitch/services"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"database/sql"
	"github.com/ivanj4u/apiswitch/helper"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", req.URL.Path[1:])
}

func tabunganHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var r dto.TabunganRequest
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	if helper.ValidateRole(r.ClientId, "tabunganemas", "INQUIRY") {
		services.InquiryTabungan(r)
	} else {
		log.Fatal("CA tidak memiliki akses tabunganemas")
	}
}

func main() {
	log.Println("Starting Application")
	// Open Database Connection
	openConnection()

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/tabungan", tabunganHandler)
	r.Methods("POST")
	http.ListenAndServe(":9090", r)

	log.Println("Application Running on ports : 9090")
}

func openConnection() {
	log.Println("Open Connection to Database")
	var err error
	database.DBCon, err = sql.Open("mysql", "rest:rest@tcp(127.0.0.1:3306)/restswitch")
	if err != nil {
		log.Fatal(err)
	}
}
