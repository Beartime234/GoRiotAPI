package riotapi

import (
	"encoding/json"
)

//CurrentGameInfo represents active-game's information
type CurrentGameInfo struct {
	GameID            int    `json:"gameId"`
	GameStartTime     int64  `json:"gameStartTime"`
	PlatformID        string `json:"platformId"`
	GameMode          string `json:"gameMode"`
	MapID             int    `json:"mapId"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	Participants []struct {
		ProfileIconID            int           `json:"profileIconId"`
		ChampionID               int           `json:"championId"`
		SummonerName             string        `json:"summonerName"`
		GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
		Bot                      bool          `json:"bot"`
		Perks                    struct {
			PerkStyle    int   `json:"perkStyle"`
			PerkIds      []int `json:"perkIds"`
			PerkSubStyle int   `json:"perkSubStyle"`
		} `json:"perks"`
		Spell2ID   int    `json:"spell2Id"`
		TeamID     int    `json:"teamId"`
		Spell1ID   int    `json:"spell1Id"`
		SummonerID string `json:"summonerId"`
	} `json:"participants"`
	GameLength      int `json:"gameLength"`
	BannedChampions []struct {
		TeamID     int `json:"teamId"`
		ChampionID int `json:"championId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
}

//FeaturedGames represents matches gotten randomly.
type FeaturedGames struct {
	ClientRefreshInterval int `json:"clientRefreshInterval"`
	GameList              []struct {
		GameID            int    `json:"gameId"`
		GameStartTime     int64  `json:"gameStartTime"`
		PlatformID        string `json:"platformId"`
		GameMode          string `json:"gameMode"`
		MapID             int    `json:"mapId"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int    `json:"gameQueueConfigId"`
		Observers         struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		Participants []struct {
			ProfileIconID int    `json:"profileIconId"`
			ChampionID    int    `json:"championId"`
			SummonerName  string `json:"summonerName"`
			Bot           bool   `json:"bot"`
			Spell2ID      int    `json:"spell2Id"`
			TeamID        int    `json:"teamId"`
			Spell1ID      int    `json:"spell1Id"`
		} `json:"participants"`
		GameLength      int           `json:"gameLength"`
		BannedChampions []interface{} `json:"bannedChampions"`
	} `json:"gameList"`
}

//GetActivegamesBySummoner is a function returns CurrentGameInfo.
func (me *Client) GetActivegamesBySummoner(encryptedSummonerID string) (CurrentGameInfo, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/spectator/v4/active-games/by-summoner/", me.Key, encryptedSummonerID)
	var currentGame CurrentGameInfo
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &currentGame); decodeError != nil {
			return currentGame, decodeError
		}
		return currentGame, nil
	}
	return currentGame, networkError
}

//GetFeaturedGames is a function returns FeaturedGames.
func (me *Client) GetFeaturedGames(encryptedSummonerID string) (FeaturedGames, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/spectator/v4/active-games/by-summoner/", me.Key, encryptedSummonerID)
	var featuredGames FeaturedGames
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &featuredGames); decodeError != nil {
			return featuredGames, decodeError
		}
		return featuredGames, nil
	}
	return featuredGames, networkError
}