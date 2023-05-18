package rgapi

import (
	"LOLDataCollector/constants"
	"fmt"
)

type ChampionMasteryDTO struct {
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionId                   int64  `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionLevel                int64  `json:"championLevel"`
	SummonerId                   string `json:"summonerId"`
	ChampionPoints               int64  `json:"championPoints"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"`
	TokensEarned                 int64  `json:"tokensEarned"`
}

var (
	ChampionMasteryBaseURL string = "/champion-mastery/v4/champion-masteries/by-summoner"
)

func GetTopNChampionMastery(summonerId string, count int, region constants.Region) (championMasteries []*ChampionMasteryDTO, err error) {
	// TODO check count doesn't exceed # of champions
	// Format the url to query
	queryURL := fmt.Sprintf(ChampionMasteryBaseURL+"/%s/top?count=%d", summonerId, count)

	// perform the request and return
	err = apiGet(string(region), queryURL, &championMasteries)
	return
}
