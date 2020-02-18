package goriot

import (
	"encoding/json"
)

//SummonerDTO represents a summoner
type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

//GetSummonersByName is a function returns SummonerDTO.
func (c *Client) GetSummonersByName(SN string) (SummonerDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/summoner/v4/summoners/by-name/", c.Key, SN)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError
}

//GetSummonersByAccount is a function returns SummonerDTO
func (c *Client) GetSummonersByAccount(encryptedAccountID string) (SummonerDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/summoner/v4/summoners/by-account/", c.Key, encryptedAccountID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError
}

//GetSummonersByPUUID is a function returns SummonerDTO
func (c *Client) GetSummonersByPUUID(encryptedPUUID string) (SummonerDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/summoner/v4/summoners/by-puuid/", c.Key, encryptedPUUID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError
}

//GetSummoners is a function returns SummonerDTO
func (c *Client) GetSummoners(encryptedSummonerID string) (SummonerDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/summoner/v4/summoners/", c.Key, encryptedSummonerID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError
}