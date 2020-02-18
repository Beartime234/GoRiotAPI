package goriot

import (
	"encoding/json"
)

// ChampionInfo represents freechampions.
type ChampionInfo struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

//GetChampionRotations is a function returns ChampionMasteryDTO.
func (me *Client) GetChampionRotations() (ChampionInfo, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/platform/v3/champion-rotations", me.Key, "")
	var championMasteryDTO ChampionInfo
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &championMasteryDTO); decodeError != nil {
			return championMasteryDTO, decodeError
		}
		return championMasteryDTO, nil
	}
	return championMasteryDTO, networkError
}