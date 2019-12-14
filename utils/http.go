package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url, user, password string) ([]byte, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth(user, password)
	if err != nil {
		return nil, err
	}

	res, err := spaceClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpPut(url, body, user, password string) (string, *http.Response, error){
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(body)))
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	return string(responseBody), resp, nil
}

func HttpPost(url, body, user, password string) (string, *http.Response, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	return string(responseBody), resp, nil
}
