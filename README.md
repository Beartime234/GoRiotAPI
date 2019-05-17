# RiotAPI4G
RiotAPI wrapper for golang

## Installation 
In your shell  
`$go get -u https://github.com/themilkey/RiotAPI4G/src` 

## How to import
`import riotapi "github.com/themilkey/RiotAPI4G/src"`    

## How to use

```go:main
client := riotapi.New("YOUR_API_KEY")
summoner, err := client.GetSummonersByName("summonerName")
if err == nil {
  fmt.Println(summoner.ID)
} else {
  fmt.Println(err)
}

currentGame, err := client.GetActivegamesBySummoner(summoner.ID)
if err == nil {
  fmt.Println(currentGame.GameID)
} else {
  fmt.Println(err)
}

//This always occcurs an error because currentGame has been played NOW.
match, err := client.GetMatches(currentGame.GameID)
if err == nil {
  fmt.Println(match.ParticipantIdentities)
} else {
  fmt.Println(err)
}
```
