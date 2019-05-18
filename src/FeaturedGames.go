package riotapi

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