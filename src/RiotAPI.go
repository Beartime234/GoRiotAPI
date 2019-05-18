//Package riotapi is wrapper for RiotAPI
package riotapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"errors"
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

//GetSummonersByName is a function returns SummonerDTO.
func (me *Client) GetSummonersByName(SN string) (SummonerDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/summoner/v4/summoners/by-name/", me.Key, SN)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError	
}

//GetSummonersByAccount is a funtion returns SummonerDTO
func (me *Client) GetSummonersByAccount(encryptedAccountID string) (SummonerDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/summoner/v4/summoners/by-account/", me.Key, encryptedAccountID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError	
}

//GetSummonersByPUUID is a funtion returns SummonerDTO
func (me *Client) GetSummonersByPUUID(encryptedPUUID string) (SummonerDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/summoner/v4/summoners/by-puuid/", me.Key, encryptedPUUID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError	
}

//GetSummoners is a funtion returns SummonerDTO
func (me *Client) GetSummoners(encryptedSummonerID string) (SummonerDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/summoner/v4/summoners/", me.Key, encryptedSummonerID)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, decodeError
		}
		return summoner, nil
	}
	return summoner, networkError	
}

//GetActivegamesBySummoner is a function returns CurrentGameInfo.
func (me *Client) GetActivegamesBySummoner(encryptedSummonerID string) (CurrentGameInfo, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/spectator/v4/active-games/by-summoner/", me.Key, encryptedSummonerID)
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
	data, networkError := getRequest(me.EndPoint + "/lol/spectator/v4/active-games/by-summoner/", me.Key, encryptedSummonerID)
	var featuredGames FeaturedGames
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &featuredGames); decodeError != nil {
			return featuredGames, decodeError
		}
		return featuredGames, nil
	}
	return featuredGames, networkError
}

//GetMatches is a function returns MatchDto.
func (me *Client) GetMatches(gameID int) (MatchDto, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/match/v4/matches/", me.Key, strconv.Itoa(gameID))
	var matchDto MatchDto
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &matchDto); decodeError != nil {
			return matchDto, decodeError
		}
		return matchDto, nil
	}
	return matchDto, networkError
}

//GetTimelineByMatch is a function returns MatchDto.
func (me *Client) GetTimelineByMatch(gameID int) (MatchTimelineDto, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/match/v4/matches/", me.Key, strconv.Itoa(gameID))
	var timeline MatchTimelineDto
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &timeline); decodeError != nil {
			return timeline, decodeError
		}
		return timeline, nil
	}
	return timeline, networkError
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
