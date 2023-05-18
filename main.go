package main

import (
	"LOLDataCollector/constants"
	"LOLDataCollector/rgapi"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// League-v4 (accountIds) -> Summoner-v4 (puuids) -> Match-v5 (get clash match ids: 700) -> Match-v5 (get clash match data) -> Champion mastery

func main() {
	// Set the api key
	// TODO make it from os secret
	//riotAPIKey := os.Getenv("RIOT_API_KEY")
	riotAPIKey := "RGAPI-aa2f31ad-5f0c-4533-bd48-27694d838478"
	if riotAPIKey == "" {
		log.Fatal("$RIOT_API_KEY must be set")
	}
	rgapi.SetAPIKey(riotAPIKey)

	// Set the long and short API rates
	var shortLimit, longLimit int
	var err error
	shortLimit, err = strconv.Atoi(os.Getenv("SHORT_RATE_LIMIT"))
	shortLimit = 10
	err = nil
	if err != nil {
		log.Fatal("Unable to get the short rate limit")
	} else {
		rgapi.SetShortRateLimit(shortLimit, 10*time.Second)
	}
	longLimit, err = strconv.Atoi(os.Getenv("LONG_RATE_LIMIT"))
	longLimit = 100
	err = nil
	if err != nil {
		log.Fatal("Unable to get the long rate limit")
	} else {
		rgapi.SetLongRateLimit(longLimit, 10*time.Minute)
	}

	queue, err := rgapi.GetByDivisionTierQueue("II", "DIAMOND", "RANKED_SOLO_5x5", 1, constants.RegionNorthAmerica)
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get Summoner info")
	} else {
		for i, v := range queue {
			fmt.Println(fmt.Sprintf("%d: %s", i, v.SummonerName))
		}
	}

	// We get the PUUID from this endpoint
	summonerInfo, err := rgapi.GetSummonerInfo("SEGW2YJaxJW0Ic86xDVs05ikhAAApGl96jqXb5Wb1Xmcx3k", constants.RegionNorthAmerica)

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get Summoner info")
	} else {
		fmt.Println(summonerInfo.Puuid)
	}

	// Here we can pass the encrypted PUUID to get match id's for clash games
	matchIds, err := rgapi.GetMatchIdsByPUUID("53lwcMQaW3mWp8p6pFjiUsUsE0Xd2YRCi6RUXlMtA153q9ZPEm2r6gsYmd53rZCLkPvvrNjyQPfgHg", 700, constants.RouteAmericas)

	// Next we'll have to check if the player has played any clash games
	// If they have, we'll save the player
	// If not, we will disregard the player
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get Summoner info")
	} else {
		for i, v := range matchIds {
			fmt.Println(fmt.Sprintf("%d: %s", i, *v))
		}
	}

	matchData, err := rgapi.GetMatchData("NA1_4637118754", constants.RouteAmericas)

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get Summoner info")
	} else {
		fmt.Println(matchData.InfoData.Participants[0].SummonerId)
	}

	champonMasteries, err := rgapi.GetTopNChampionMastery("bO-y9WRgiEjLSa-L_QthzMoiP8xkwdBXY736uyF3Wqp4T8s", 5, constants.RegionNorthAmerica)

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get Summoner info")
	} else {
		for i, v := range champonMasteries {
			fmt.Println(fmt.Sprintf("%d: %d", i, (*v).ChampionId))
		}
	}

}
