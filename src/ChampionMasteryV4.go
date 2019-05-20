package riotapi

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
func (me *Client) GetChampionMasteriesBySummoner(encryptedSummonerID string) ([]ChampionMasteryDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/champion-mastery/v4/champion-masteries/by-summoner/", me.Key, encryptedSummonerID)
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
func (me *Client) GetChampionMasteryScoreBySummoner(encryptedSummonerID string) (int, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/champion-mastery/v4/scores/by-summoner/", me.Key, encryptedSummonerID)
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
func (me *Client) GetChampionMasteriesBySummonerByChampion(encryptedSummonerID string, championID int) (ChampionMasteryDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID+"/by-champion/", me.Key, strconv.Itoa(championID))
	var championMasteryDTO ChampionMasteryDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &championMasteryDTO); decodeError != nil {
			return championMasteryDTO, decodeError
		}
		return championMasteryDTO, nil
	}
	return championMasteryDTO, networkError
}