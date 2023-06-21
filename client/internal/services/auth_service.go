package services

import (
	"bytes"
	config "client/init"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Registration(login string, password string) (int, error) {
	postBody, _ := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)

	requestURL := fmt.Sprintf(config.Env.ApiURL + "/user/register")

	res, err := http.Post(requestURL, "application/json", responseBody)
	if err != nil {
		return 500, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func (as *AuthService) Login(login string, password string) (string, int, error) {
	postBody, _ := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)

	requestURL := fmt.Sprintf(config.Env.ApiURL + "/user/login")

	res, err := http.Post(requestURL, "application/json", responseBody)
	if err != nil {
		return "", 500, err
	}
	defer res.Body.Close()

	return res.Header.Get("Token"), res.StatusCode, nil
}

func (as *AuthService) IsAuthenticated() (int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/user/is-authenticated")

	req, _ := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)
	if err != nil {
		return 500, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
