package goriot

import (
	"encoding/json"
	"strconv"
)

// ChampionMasteryDTO represents mastery each champions.
type ChampionMasteryDTO struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}

//GetChampionMasteriesBySummoner is a function returns ChampionMasteryDTO.
func (c *Client) GetChampionMasteriesBySummoner(encryptedSummonerID string) ([]ChampionMasteryDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/champion-mastery/v4/champion-masteries/by-summoner/", c.Key, encryptedSummonerID)
	var championMasteryDTO []ChampionMasteryDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &championMasteryDTO); decodeError != nil {
			return championMasteryDTO, decodeError
		}
		return championMasteryDTO, nil
	}
	return championMasteryDTO, networkError
}

//GetChampionMasteryScoreBySummoner is a function returns ChampionMasteryDTO.
func (c *Client) GetChampionMasteryScoreBySummoner(encryptedSummonerID string) (int, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/champion-mastery/v4/scores/by-summoner/", c.Key, encryptedSummonerID)
	score, er := strconv.Atoi(string(data))
	if er != nil {
		return score, nil
	}
	if networkError == nil {
		return score, nil
	}
	return score, networkError
}

//GetChampionMasteriesBySummonerByChampion is a function returns ChampionMasteryDTO.
func (c *Client) GetChampionMasteriesBySummonerByChampion(encryptedSummonerID string, championID int) (ChampionMasteryDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID+"/by-champion/", c.Key, strconv.Itoa(championID))
	var championMasteryDTO ChampionMasteryDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &championMasteryDTO); decodeError != nil {
			return championMasteryDTO, decodeError
		}
		return championMasteryDTO, nil
	}
	return championMasteryDTO, networkError
}