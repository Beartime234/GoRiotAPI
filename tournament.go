package goriot

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
func (c *Client) PostTournamentProviders(provider ProviderRegistrationParameters) (uint64, error) {
	json, _ := json.Marshal(provider)
	data, networkError := postRequestWithBody(c.EndPoint+"/lol/tournament/v4/providers", c.Key, json)
	if networkError == nil {
		return binary.BigEndian.Uint64(data), nil
	}
	return binary.BigEndian.Uint64(data), networkError
}
