/*
 * Copyright (c) 2018.
 */

package helper

import (
	"github.com/ivanj4u/apiswitch/database"
	"log"
)

func ValidateRole(clientId, role, requestType string) (bool) {
	valid := false

	rows, err := database.DBCon.Query("SELECT * FROM tbl_rest_ca_role WHERE username = &u AND role = &r AND request_type = &t",
		clientId, role, requestType)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if rows == nil {
		valid = false
	}
	return valid
}