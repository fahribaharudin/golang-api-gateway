package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIGateway is controller to the http proxy handler
type APIGateway struct {
}

// WelcomeHandler to the landing endpoint
func (c *APIGateway) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the KUDO CMS - API Gateway."))
}

// UniversalHandler to handle all request for the API Gateway
func (c *APIGateway) UniversalHandler(w http.ResponseWriter, r *http.Request) {
	var reqData interface{}
	var reqURL = "http://localhost:3000" + r.RequestURI

	if r.Method == "POST" {
		if r.Header.Get("content-type") == "application/json" {
			// decode the request body
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqData)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			// send a request as application/json
			makeJSONPostRequest(w, r, reqURL, reqData)
			return
		}

		output, _ := json.Marshal(map[string]interface{}{
			"code":     http.StatusBadRequest,
			"messagea": "The API Gateway only support 'application/json' content-type",
		})
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)

	}
}

func makeJSONPostRequest(w http.ResponseWriter, r *http.Request, url string, data interface{}) {
	reqDataJSON, err := json.Marshal(&data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqDataJSON))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-type", resp.Header.Get("Content-type"))
	w.Write(body)
}
