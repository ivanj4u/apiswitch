package services

import (
	"log"
	"github.com/ivanj4u/apiswitch/dto"
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/ivanj4u/apiswitch/util"
	"encoding/json"
)

func InquiryGadai(req dto.Dto, str string) (dto.Dto) {
	log.Println("Start services Inquiry Gadai : " + req.JenisTransaksi)
	var res dto.Dto

	url := util.UrlKonven + "/gadai/" + req.RequestType
	log.Println("URL:>", url)

	r := []byte(str)
	coreRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(r))
	coreRequest.Header.Set("Authorization", util.TokenOauth)
	coreRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(coreRequest)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Core :", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}

	log.Println("End services Inquiry Gadai : " + req.JenisTransaksi)
	return res
}