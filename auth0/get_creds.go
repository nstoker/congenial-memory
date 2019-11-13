package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Token what we get back from Auth0
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func main() {

	url := "https://dev-vqn41l5l.eu.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"6mcMYygQleg7uiKFB6tUP5RyiDfz1UpP\",\"client_secret\":\"sqyI8BQjUPFaiaaVD8hr7_339aKvPbx6PT7NAlpTbuYiNU2eIrrFcz4FBy8MBoAW\",\"audience\":\"https://my-congenial-memory-api\",\"grant_type\":\"client_credentials\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatalf("error from request to '%s' with '%v' was: %s", url, payload, err)
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error in response %s", err)
	}
	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	token, httpStatus, err := convertHTTPBodyToToken(res.Body)
	if err != nil {
		log.Fatalf("Error %v in conversion %s", httpStatus, err)
	}
	if httpStatus != http.StatusOK {
		log.Fatalf("didn't expect %v", httpStatus)
	}

	fmt.Printf("ACCESS_TOKEN=%v\n", token.AccessToken)
}

func convertHTTPBodyToToken(httpBody io.ReadCloser) (Token, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return Token{}, http.StatusInternalServerError, err
	}

	defer httpBody.Close()
	return convertJSONBodyToToken(body)
}

func convertJSONBodyToToken(jsonBody []byte) (Token, int, error) {
	var token Token

	err := json.Unmarshal(jsonBody, &token)
	if err != nil {
		return Token{}, http.StatusBadRequest, err
	}

	return token, http.StatusOK, nil
}
