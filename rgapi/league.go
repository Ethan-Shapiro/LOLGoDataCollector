package rgapi

import (
	"LOLDataCollector/constants"
	"fmt"
)

type LeagueEntryDTO struct {
	LeagueId     string        `json:"leagueId"`
	SummonerId   string        `json:"summonerId"`
	SummonerName string        `json:"summonerName"`
	QueueType    string        `json:"queueType"`
	Tier         string        `json:"tier"`
	Rank         string        `json:"rank"`
	LeaguePoints int64         `json:"leaguePoints"`
	Wins         int64         `json:"wins"`
	Losses       int64         `json:"losses"`
	HotStreak    bool          `json:"hotStreak"`
	Veteran      bool          `json:"veteran"`
	FreshBool    bool          `json:"freshBool"`
	Inactive     bool          `json:"inactive"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
}

type MiniSeriesDTO struct {
	Losses   int64 `json:"losses"`
	Progress bool  `json:"progress"`
	Target   int64 `json:"target"`
	Wins     int64 `json:"wins"`
}

var (
	LeagueBaseURL string = "/league/v4"
)

func GetByDivisionTierQueue(division string, tier string, queue string, page int64, region constants.Region) (leagueEntries []*LeagueEntryDTO, err error) {
	// Build the query string
	query := fmt.Sprintf(LeagueBaseURL+"/entries/%s/%s/%s?page=%d", queue, tier, division, page)

	// get the data from the api and return
	err = apiGet("na1", query, &leagueEntries)
	return
}
