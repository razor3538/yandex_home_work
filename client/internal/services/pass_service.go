package services

import (
	"bytes"
	config "client/init"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type PassService struct{}

func NewPassService() *PassService {
	return &PassService{}
}

var authService = NewAuthService()

func (ps *PassService) Save(login string, password string, name string, meta string) (int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/passwords")

	postBody, _ := json.Marshal(map[string]interface{}{
		"meta":      meta,
		"login":     login,
		"name_pair": name,
		"password":  password,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("POST", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return 500, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func (ps *PassService) Get(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/get_password")

	postBody, _ := json.Marshal(map[string]string{
		"name": name,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("POST", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return "", 500, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	println(string(b))
	return string(b), res.StatusCode, nil
}

func (ps *PassService) Delete(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/password")

	postBody, _ := json.Marshal(map[string]string{
		"name": name,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("DELETE", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return "", 500, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b), res.StatusCode, nil
}
