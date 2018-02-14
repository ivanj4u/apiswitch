/*
 * Copyright (c) 2018.
 */

package services

import (
	"log"
	"github.com/ivanj4u/apiswitch/dto"
)

func InquiryTabungan(req dto.TabunganRequest) (dto.ResponseServices){
	log.Println("Starting services Inquiry Tabungan : " + req.JenisTransaksi)
	var res dto.ResponseServices

	if req.JenisTransaksi == "SL" {
		res.ResponseCode = "00"
		res.ResponseDesc = "Approved"
		res.Data = req.JenisTransaksi
	}
	return res
}