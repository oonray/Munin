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
	}

  data, err := ioutil.ReadAll(resp.Body)
	if(err != nil){
		logrus.Errorf("Could not read response %s",err)
	}

	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("https://api.dnsimple.com/v2/%s/zones/%s/record/%s",
										os.Getenv("ACCOUNT_ID"),
										os.Getenv("ZONE_ID"),
										os.Getenv("RECORD_ID"),
	),
		bytes.NewBuffer([]byte(fmt.Sprintf("{\"content\":\"%s\"}",data))),
	)

	req.Header.Add("Authorization",fmt.Sprintf("Bearer %s",os.Getenv("TOKEN")))
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Accept","application/json")
	
	client := &http.Client{}
	client.Do(req)
}
