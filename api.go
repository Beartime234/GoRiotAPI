//Package riotapi is wrapper for RiotAPI
package goriot

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var httpclient http.Client
var client *Client
var endPoint string

// QueueType has consts RANKEDSOLO5x5, RANKEDFLEXSR and RANKEDFLEXTT.
type QueueType int

const (
	_ QueueType = iota
	// RANKEDSOLO5x5 represents solo queue
	RANKEDSOLO5x5
	// RANKEDFLEXSR represents flex queue
	RANKEDFLEXSR
	// RANKEDFLEXTT represents flex queue
	RANKEDFLEXTT
)

func (q QueueType) String() string {
	switch q {
	case RANKEDSOLO5x5:
		return "RANKED_SOLO_5x5"
	case RANKEDFLEXSR:
		return "RANKED_FLEX_SR"
	case RANKEDFLEXTT:
		return "RANKED_FLEX_TT"
	default:
		return "Unknown"
	}
}

// Tier has consts IRON to DIAMOND.
type Tier int

const (
	_ Tier = iota
	// DIAMOND tier
	DIAMOND
	// PLATINUM tier
	PLATINUM
	// GOLD tier
	GOLD
	// SILVER tier
	SILVER
	// BRONZE tier
	BRONZE
	// IRON tier
	IRON
)

func (q Tier) String() string {
	switch q {
	case DIAMOND:
		return "DIAMOND"
	case PLATINUM:
		return "PLATINUM"
	case GOLD:
		return "GOLD"
	case SILVER:
		return "SILVER"
	case BRONZE:
		return "BRONZE"
	case IRON:
		return "IRON"
	default:
		return "Unknown"
	}
}

//Client is for RiotAPI
type Client struct {
	EndPoint string
	Key      string
	ApiRegion string
}

//Constructs the client
func New(key string, region string) *Client {
	client = &Client{EndPoint: fmt.Sprintf("https://%s.api.riotgames.com", region), Key: key, ApiRegion:region}
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
			fmt.Println(data.Status)
			return buf.Bytes(), errors.New(data.Status)
		}
		return buf.Bytes(), err
	}
	return nil, err
}

func postRequestWithBody(uri, key string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	req.Header.Set("X-Riot-Token", key)
	data, err := httpclient.Do(req)
	if err == nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, data.Body)
		if data.StatusCode != 200 {
			fmt.Println(string(data.StatusCode) + data.Status)
			return buf.Bytes(), errors.New(string(data.StatusCode) + data.Status)
		}
		return buf.Bytes(), err
	}
	return nil, err
}
