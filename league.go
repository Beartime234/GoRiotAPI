package goriot

import (
	"encoding/json"
)

// LeagueListDTO represents the challenger league.
type LeagueListDTO struct {
	Tier     string `json:"tier"`
	LeagueID string `json:"leagueId"`
	Entries  []struct {
		SummonerName string `json:"summonerName"`
		HotStreak    bool   `json:"hotStreak"`
		Wins         int    `json:"wins"`
		Veteran      bool   `json:"veteran"`
		Losses       int    `json:"losses"`
		Rank         string `json:"rank"`
		Inactive     bool   `json:"inactive"`
		FreshBlood   bool   `json:"freshBlood"`
		SummonerID   string `json:"summonerId"`
		LeaguePoints int    `json:"leaguePoints"`
	} `json:"entries"`
	Queue string `json:"queue"`
	Name  string `json:"name"`
}

// LeagueEntryDTO represents league entry.
type LeagueEntryDTO struct {
	QueueType    string `json:"queueType"`
	SummonerName string `json:"summonerName"`
	HotStreak    bool   `json:"hotStreak"`
	Wins         int    `json:"wins"`
	Veteran      bool   `json:"veteran"`
	Losses       int    `json:"losses"`
	Rank         string `json:"rank"`
	Tier         string `json:"tier"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	LeagueID     string `json:"leagueId"`
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
}

//GetChallengerLeaguesByQueue is a function returns CurrentGameInfo.
func (c *Client) GetChallengerLeaguesByQueue(q QueueType) (LeagueListDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/challengerleagues/by-queue/", c.Key, q.String())
	var leagueListDTO LeagueListDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueListDTO); decodeError != nil {
			return leagueListDTO, decodeError
		}
		return leagueListDTO, nil
	}
	return leagueListDTO, networkError
}

//GetGrandmasterLeaguesByQueue is a function returns CurrentGameInfo.
func (c *Client) GetGrandmasterLeaguesByQueue(q QueueType) (LeagueListDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/grandmasterleagues/by-queue/", c.Key, q.String())
	var leagueListDTO LeagueListDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueListDTO); decodeError != nil {
			return leagueListDTO, decodeError
		}
		return leagueListDTO, nil
	}
	return leagueListDTO, networkError
}

//GetMasterLeaguesByQueue is a function returns CurrentGameInfo.
func (c *Client) GetMasterLeaguesByQueue(q QueueType) (LeagueListDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/masterleagues/by-queue/", c.Key, q.String())
	var leagueListDTO LeagueListDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueListDTO); decodeError != nil {
			return leagueListDTO, decodeError
		}
		return leagueListDTO, nil
	}
	return leagueListDTO, networkError
}

//GetLeaguesByQueue is a function returns CurrentGameInfo.
//TO DO: Optional
func (c *Client) GetLeaguesByQueue(q QueueType, tier Tier, division int) ([]LeagueEntryDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/entries/" + q.String() + "/" + tier.String() + "/" + toRoma(division), c.Key, "")
	var leagueEntryDTO []LeagueEntryDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueEntryDTO); decodeError != nil {
			return leagueEntryDTO, decodeError
		}
		return leagueEntryDTO, nil
	}
	return leagueEntryDTO, networkError
}

func toRoma(division int) string {
	switch division {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	default:
		return "Unknown"
	}
}

//GetLeaguesByLeagueID is a function returns CurrentGameInfo.
func (c *Client) GetLeaguesByLeagueID(leagueID string) (LeagueListDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/leagues/", c.Key, leagueID)
	var leagueListDTO LeagueListDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueListDTO); decodeError != nil {
			return leagueListDTO, decodeError
		}
		return leagueListDTO, nil
	}
	return leagueListDTO, networkError
}

//GetEntriesbySummoner is a function returns league entries in all queues for a given summoner ID.
func (c *Client) GetEntriesbySummoner(encryptedSummonerID string) ([]LeagueEntryDTO, error) {
	data, networkError := getRequest(c.EndPoint+"/lol/league/v4/entries/by-summoner/", c.Key, encryptedSummonerID)
	var leagueEntryDTO []LeagueEntryDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &leagueEntryDTO); decodeError != nil {
			return leagueEntryDTO, decodeError
		}
		return leagueEntryDTO, nil
	}
	return leagueEntryDTO, networkError
}