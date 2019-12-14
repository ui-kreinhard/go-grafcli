package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	username string
	password string
	baseUrl  string
}

func NewHttpClient(baseUrl, username, password string) *HttpClient {
	return &HttpClient{
		username,
		password,
		baseUrl,
	}
}

func (h *HttpClient) HttpGet(url string) ([]byte, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, h.baseUrl+  url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(h.username, h.password)
	req.Header.Set("Content-Type", "application/json")

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

func (h *HttpClient) HttpPut(url, body string) (string, *http.Response, error){
	req, err := http.NewRequest("PUT", h.baseUrl+  url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", nil, err
	}
	req.SetBasicAuth(h.username, h.password)
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

func (h *HttpClient) HttpPost(url, body string) (string, *http.Response, error) {

	req, err := http.NewRequest("POST", h.baseUrl+  url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", nil, err
	}
	req.SetBasicAuth(h.username, h.password)
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
