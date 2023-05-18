package rgapi

import (
	"LOLDataCollector/constants"
	"fmt"
)

type MatchDto struct {
	MetaData MetadataDto `json:"metadata"`
	InfoData InfoDto     `json:"info"`
}

type MetadataDto struct {
	DataVersion  string     `json:"dataVersion"`
	MatchId      string     `json:"matchId"`
	Participants [10]string `json:"participants"`
}

type InfoDto struct {
	GameCreation       int64              `json:"gameCreation"`
	GameDuration       int64              `json:"gameDuration"`
	GameEndTimestamp   int64              `json:"gameEndTimestamp"`
	GameId             int64              `json:"gameId"`
	GameMode           string             `json:"gameMode"`
	GameName           string             `json:"gameName"`
	GameStartTimestamp int64              `json:"gameStartTimestamp"`
	GameType           string             `json:"gameType"`
	GameVersion        string             `json:"gameVersion"`
	MapId              int                `json:"mapId"`
	Participants       [10]ParticipantDto `json:"participants"`
	PlatformId         string             `json:"platformId"`
	QueueId            int                `json:"queueId"`
	Teams              [2]TeamDto         `json:"teams"`
	TournamentCode     string             `json:"tournamentCode"`
}

type ParticipantDto struct {
	Assists                        int64    `json:"assists"`
	BaronKills                     int64    `json:"baronKills"`
	BountyLevel                    int64    `json:"bountyLevel"`
	ChampExperience                int64    `json:"champExperience"`
	ChampLevel                     int64    `json:"champLevel"`
	ChampionId                     int64    `json:"championId"`
	ChampionName                   string   `json:"championName"`
	ChampionTransform              int64    `json:"championTransform"`
	ConsumablesPurchased           int64    `json:"consumablesPurchased"`
	DamageDealtToBuildings         int64    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int64    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int64    `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int64    `json:"damageSelfMitigated"`
	Deaths                         int64    `json:"deaths"`
	DetectorWardsPlaced            int64    `json:"detectorWardsPlaced"`
	DoubleKills                    int64    `json:"doubleKills"`
	DragonKills                    int64    `json:"dragonKills"`
	FirstBloodAssist               bool     `json:"firstBloodAssist"`
	FirstBloodKill                 bool     `json:"firstBloodKill"`
	FirstTowerAssist               bool     `json:"firstTowerAssist"`
	FirstTowerKill                 bool     `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool     `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool     `json:"gameEndedInSurrender"`
	GoldEarned                     int64    `json:"goldEarned"`
	GoldSpent                      int64    `json:"goldSpent"`
	IndividualPosition             string   `json:"individualPosition"`
	InhibitorKills                 int64    `json:"inhibitorKills"`
	InhibitorTakedowns             int64    `json:"inhibitorTakedowns"`
	InhibitorsLost                 int64    `json:"inhibitorsLost"`
	Item0                          int64    `json:"item0"`
	Item1                          int64    `json:"item1"`
	Item2                          int64    `json:"item2"`
	Item3                          int64    `json:"item3"`
	Item4                          int64    `json:"item4"`
	Item5                          int64    `json:"item5"`
	Item6                          int64    `json:"item6"`
	ItemsPurchased                 int64    `json:"itemsPurchased"`
	KillingSprees                  int64    `json:"killingSprees"`
	Kills                          int64    `json:"kills"`
	Lane                           string   `json:"lane"`
	LargestCriticalStrike          int64    `json:"largestCriticalStrike"`
	LargestKillingSpree            int64    `json:"largestKillingSpree"`
	LargestMultiKill               int64    `json:"largestMultiKill"`
	LongestTimeSpentLiving         int64    `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int64    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int64    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int64    `json:"magicDamageTaken"`
	NeutralMinionsKilled           int64    `json:"neutralMinionsKilled"`
	NexusKills                     int64    `json:"nexusKills"`
	NexusTakedowns                 int64    `json:"nexusTakedowns"`
	NexusLost                      int64    `json:"nexusLost"`
	ObjectivesStolen               int64    `json:"objectivesStolen"`
	ObjectivesStolenAssists        int64    `json:"objectivesStolenAssists"`
	ParticipantId                  int64    `json:"participantId"`
	PentaKills                     int64    `json:"pentaKills"`
	Perks                          PerksDto `json:"perks"`
	PhysicalDamageDealt            int64    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int64    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int64    `json:"physicalDamageTaken"`
	ProfileIcon                    int64    `json:"profileIcon"`
	Puuid                          string   `json:"puuid"`
	QuadraKills                    int64    `json:"quadraKills"`
	RiotIdTagline                  string   `json:"riotIdTagline"`
	Role                           string   `json:"role"`
	SightWardsBoughtInGame         int64    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int64    `json:"spell1Casts"`
	Spell2Casts                    int64    `json:"spell2Casts"`
	Spell3Casts                    int64    `json:"spell3Casts"`
	Spell4Casts                    int64    `json:"spell4Casts"`
	Summoner1Casts                 int64    `json:"summoner1Casts"`
	Summoner1Id                    int64    `json:"summoner1Id"`
	Summoner2Casts                 int64    `json:"summoner2Casts"`
	Summoner2Id                    int64    `json:"summoner2Id"`
	SummonerId                     string   `json:"summonerId"`
	SummonerLevel                  int64    `json:"summonerLevel"`
	SummonerName                   string   `json:"summonerName"`
	TeamEarlySurrendered           bool     `json:"teamEarlySurrendered"`
	TeamId                         int      `json:"teamId"`
	TeamPosition                   string   `json:"teamPosition"`
	TimeCCingOthers                int64    `json:"timeCCingOthers"`
	TimePlayed                     int64    `json:"timePlayed"`
	TotalDamageDealt               int64    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int64    `json:"totalDamageDealtToChampions"`
	TotalDamageTaken               int64    `json:"totalDamageTaken"`
	TotalHeal                      int64    `json:"totalHeal"`
	TotalHealsOnTeammates          int64    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int64    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int64    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int64    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int64    `json:"totalUnitsHealed"`
	TripleKills                    int64    `json:"tripleKills"`
	TrueDamageDealt                int64    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int64    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int64    `json:"trueDamageTaken"`
	TurretKills                    int64    `json:"turretKills"`
	TurretTakedowns                int64    `json:"turretTakedowns"`
	TurretsLost                    int64    `json:"turretsLost"`
	UnrealKills                    int64    `json:"unrealKills"`
	VisionScore                    int64    `json:"visionScore"`
	VisionWardsBoughtInGame        int64    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int64    `json:"wardsKilled"`
	WardsPlaced                    int64    `json:"wardsPlaced"`
	Win                            bool     `json:"win"`
}

type TeamDto struct {
	Bans       [5]BanDto     `json:"bans"`
	Objectives ObjectivesDto `json:"objectives"`
	TeamId     int64         `json:"teamId"`
	Win        bool          `json:"win"`
}

type BanDto struct {
	ChampionId int64 `json:"championId"`
	PickTurn   int64 `json:"pickTurn"`
}

type ObjectivesDto struct {
	Baron      ObjectiveDto `json:"baron"`
	Champion   ObjectiveDto `json:"champion"`
	Dragon     ObjectiveDto `json:"dragon"`
	Inhibitor  ObjectiveDto `json:"inhibitor"`
	RiftHerald ObjectiveDto `json:"riftHerald"`
	Tower      ObjectiveDto `json:"tower"`
}

type ObjectiveDto struct {
	First bool  `json:"first"`
	Kills int64 `json:"kills"`
}

type PerksDto struct {
	StatPerks PerkStatsDto   `json:"statPerks"`
	Styles    []PerkStyleDto `json:"styles"`
}

type PerkStatsDto struct {
	Defense int64 `json:"defense"`
	Flex    int64 `json:"flex"`
	Offense int64 `json:"offense"`
}

type PerkStyleDto struct {
	Description string                  `json:"description"`
	Selections  []PerkStyleSelectionDto `json:"selections"`
	Style       int64                   `json:"style"`
}

type PerkStyleSelectionDto struct {
	Perk int64 `json:"perk"`
	Var1 int64 `json:"var1"`
	Var2 int64 `json:"var2"`
	Var3 int64 `json:"var3"`
}

var (
	MatchBaseURL string = "/match/v5/matches"
)

func GetMatchIdsByPUUID(encryptedPUUID string, queue int, route constants.Route) (matchIds []*string, err error) {
	// Format the url to query
	queryURL := fmt.Sprintf(MatchBaseURL+"/by-puuid/%s/ids?queue=%d", encryptedPUUID, queue)

	// perform the request and return
	err = apiGet(string(route), queryURL, &matchIds)
	return
}

func GetMatchData(matchId string, route constants.Route) (match *MatchDto, err error) {
	// Format the url to query
	queryURL := fmt.Sprintf(MatchBaseURL+"/%s", matchId)

	// perform the request and return
	err = apiGet(string(route), queryURL, &match)
	return
}
