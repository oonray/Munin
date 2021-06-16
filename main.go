package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)


func main() {
	path := "http://icanhazip.com/"

	resp, err := http.Get(path)
	if(err != nil){
		logrus.Errorf("Could not fetch url %s",err)
		return
	}

  data, err := ioutil.ReadAll(resp.Body)
	if(err != nil){
		logrus.Errorf("Could not read response %s",err)
		return
	}

	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("https://api.dnsimple.com/v2/%s/zones/%s/record/%s",
										os.Getenv("ACCOUNT_ID"),
										os.Getenv("ZONE_ID"),
										os.Getenv("RECORD_ID"),
		),
		bytes.NewBuffer([]byte(fmt.Sprintf("{\"content\":\"%s\"}",data))))

	if(err != nil){
		logrus.Errorf("Could not create response %s",err)
		return
	}

	req.Header.Add("Authorization",fmt.Sprintf("Bearer %s",os.Getenv("TOKEN")))
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Accept","application/json")
	
	client := &http.Client{}
	err = client.Do(req)
	if(err != nil){
		logrus.Errorf("Could not send response %s",err)
		return
	}
}
