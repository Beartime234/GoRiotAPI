//Package riotapi is wrapper for RiotAPI
package riotapi

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var httpclient http.Client
var client *Client
var endPoint string

//Client is for RiotAPI
type Client struct {
	EndPoint string
	Key      string
}

//New is constructor
func New(key string) *Client {
	client = &Client{EndPoint: "https://jp1.api.riotgames.com/", Key: key}
	return client
}

func getRequest(uri, key, param string) ([]byte, error) {
	req, _ := http.NewRequest("GET", uri+url.PathEscape(param), nil)
	req.Header.Set("X-Riot-Token", key)
	data, err := httpclient.Do(req)
	if err == nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, data.Body)
		if data.StatusCode != 200 {
			return buf.Bytes(), errors.New(string(data.StatusCode) + data.Status)
		}
		return buf.Bytes(), err
	}
	return nil, err
}
