package riotapi

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