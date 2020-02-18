package goriot

import (
	"encoding/json"
)

// ShardStatus represents LoL service status.
type ShardStatus struct {
	Name      string `json:"name"`
	RegionTag string `json:"region_tag"`
	Hostname  string `json:"hostname"`
	Services  []struct {
		Status    string        `json:"status"`
		Incidents []interface{} `json:"incidents"`
		Name      string        `json:"name"`
		Slug      string        `json:"slug"`
	} `json:"services"`
	Slug    string   `json:"slug"`
	Locales []string `json:"locales"`
}

//GetStatusShardData is a function returns ShardStatus.
func (me *Client) GetStatusShardData() (ShardStatus, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/status/v3/shard-data", me.Key, "")
	var shardStatus ShardStatus
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &shardStatus); decodeError != nil {
			return shardStatus, decodeError
		}
		return shardStatus, nil
	}
	return shardStatus, networkError
}
