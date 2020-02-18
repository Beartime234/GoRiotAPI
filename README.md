# RiotAPI4G
RiotAPI wrapper for golang initially forked from https://github.com/themilkey/RiotAPI4G

This supports the 4th Generation commands. 

## How to use
```go
client := riotapi.New("YOUR_API_KEY")
summoner, err := client.GetSummonersByName("summonerName")
if err == nil {
  Println(summoner.ID)
} else {
  Println(err)
}
```
