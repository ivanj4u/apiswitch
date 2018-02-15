/*
 * Copyright (c) 2018.
 */

package services

import (
	"github.com/ivanj4u/apiswitch/dto"
	"log"
	"github.com/ivanj4u/apiswitch/util"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func InquiryMikro(req dto.Dto, str string) (dto.Dto){
	log.Println("Start services Inquiry Mikro : &s", req.JenisTransaksi)

	var res dto.Dto

	url:= util.UrlKonven + "mikro/" + req.RequestType
	log.Println("URL :>", url)

	coreRequest, err := http.NewRequest("POST", url, bytes.NewBufferString(str))
	if err != nil {
		log.Panic(err)
	}

	coreRequest.Header.Set("Authorization", util.TokenOauth)
	coreRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(coreRequest)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	log.Println("resp Status : ", resp.Status)
	log.Println("resp Header : ", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response Core : ", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
	}

	log.Println("End Services Inquiry Mikro : " + req.JenisTransaksi)
	return res
}
