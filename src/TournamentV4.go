package riotapi

import (
	"encoding/binary"
	"encoding/json"
)

// ProviderRegistrationParameters represents aaa.
type ProviderRegistrationParameters struct {
	Region string `json:"region"`
	URL    string `json:"url"`
}

// PostTournamentProviders is
func (me *Client) PostTournamentProviders(provider ProviderRegistrationParameters) (uint64, error) {
	json, _ := json.Marshal(provider)
	data, networkError := postRequestWithBody(me.EndPoint+"/lol/tournament/v4/providers", me.Key, json)
	if networkError == nil {
		return binary.BigEndian.Uint64(data), nil
	}
	return binary.BigEndian.Uint64(data), networkError
}
