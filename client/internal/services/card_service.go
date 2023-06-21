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

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

func (cs *CardService) Save(number, cvs, dateEnd, bank, name, meta string) (int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/card")

	postBody, _ := json.Marshal(map[string]interface{}{
		"meta":      meta,
		"name_pair": name,
		"bank":      bank,
		"number":    number,
		"cvs":       cvs,
		"date_end":  dateEnd,
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

func (cs *CardService) Get(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/get_card")

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

	return string(b), res.StatusCode, nil
}

func (cs *CardService) Delete(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/card")

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
