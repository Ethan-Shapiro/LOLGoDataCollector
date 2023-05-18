package rgapi

import (
	"LOLDataCollector/constants"
	"fmt"
)

type SummonerDTO struct {
	AccountId     string  `json:"accountId"`
	ProfileIconId int64   `json:"profileIconId"`
	RevisionDate  float64 `json:"revisionDate"`
	Name          string  `json:"name"`
	Id            string  `json:"id"`
	Puuid         string  `json:"puuid"`
	SummonerLevel float64 `json:"summonerLevel"`
}

var (
	SummonerBaseURL string = "/summoner/v4/summoners"
)

func GetSummonerInfo(summonerId string, region constants.Region) (summonerDTO *SummonerDTO, err error) {

	// Format the url to query
	queryURL := fmt.Sprintf(SummonerBaseURL+"/%s", summonerId)

	// perform the request and return
	err = apiGet(string(region), queryURL, &summonerDTO)
	return
}
