package grafanastack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateStackRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Region string `json:"region"`
}

type Client struct {
	client *http.Client
}

type InitialData struct {
	url       string
	accessKey string
	stackName string
	slug      string
	region    string
}

func CreateStack(id InitialData) {
	client := &http.Client{}
	request := CreateStackRequest{Name: id.stackName, Slug: id.slug, Region: id.region}

	data, err := json.Marshal(request)

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", id.url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+id.accessKey)
	req.Header.Add("Content-Type", "application/json")

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

}
