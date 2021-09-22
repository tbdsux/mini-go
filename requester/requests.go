// Package requester implements a request function with auto json decode.
package requester

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Requester struct {
	client *http.Client
	config *RequesterConfig
}

type RequesterConfig struct {
	NotOKError bool
}

// Define new requester object. You are free to define the http.Client yourself.
// Config is needed for other stuff requests.
func NewRequester(client *http.Client, config *RequesterConfig) *Requester {
	return &Requester{
		client: client, config: config,
	}
}

// Get sends a get request and decodes the json response to v. Any other internal error will be returned.
func (r *Requester) Get(url string, v interface{}, e interface{}) error {
	req, err := r.client.Get(url)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	// if defined with parsing notokerrors
	if r.config.NotOKError && req.StatusCode != 200 {
		return json.NewDecoder(req.Body).Decode(&e)
	}

	return json.NewDecoder(req.Body).Decode(v)
}

// Post sends a post request and decods the json response to v. Any other internal error will be returned.
func (r *Requester) Post(url string, data interface{}, v interface{}, e interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := r.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer req.Body.Close()

	// if defined with parsing notokerrors
	if r.config.NotOKError && req.StatusCode != 200 {
		return json.NewDecoder(req.Body).Decode(&e)
	}

	return json.NewDecoder(req.Body).Decode(v)
}

// Request sends a request with a configured *http.Request object. Any other internal error will be returned.
func (r *Requester) Request(request *http.Request, v interface{}, e interface{}) error {
	req, err := r.client.Do(request)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	// if defined with parsing notokerrors
	if r.config.NotOKError && req.StatusCode != 200 {
		return json.NewDecoder(req.Body).Decode(&e)
	}

	return json.NewDecoder(req.Body).Decode(v)
}
