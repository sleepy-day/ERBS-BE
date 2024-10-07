package main

import "time"

type Match struct {
	UserNum                                   int64              `json:"userNum"`
	Nickname                                  string             `json:"nickname"`
	GameID                                    int64              `json:"gameId"`
	SeasonID                                  int64              `json:"seasonId"`
	MatchingMode                              int64              `json:"matchingMode"`
	MatchingTeamMode                          int64              `json:"matchingTeamMode"`
	CharacterNum                              int64              `json:"characterNum"`
	SkinCode                                  int64              `json:"skinCode"`
	CharacterLevel                            int64              `json:"characterLevel"`
	GameRank                                  int64              `json:"gameRank"`
	PlayerKill                                int64              `json:"playerKill"`
	PlayerAssistant                           int64              `json:"playerAssistant"`
	MonsterKill                               int64              `json:"monsterKill"`
	BestWeapon                                int64              `json:"bestWeapon"`
	BestWeaponLevel                           int64              `json:"bestWeaponLevel"`
	MasteryLevel                              map[string]int64   `json:"masteryLevel"`
	Equipment                                 map[string]int64   `json:"equipment"`
	VersionMajor                              int64              `json:"versionMajor"`
	VersionMinor                              int64              `json:"versionMinor"`
	Language                                  string             `json:"language"`
	SkillLevelInfo                            map[string]int64   `json:"skillLevelInfo"`
	SkillOrderInfo                            map[string]int64   `json:"skillOrderInfo"`
	ServerName                                string             `json:"serverName"`
	MaxHP                                     int64              `json:"maxHp"`
	MaxSP                                     int64              `json:"maxSp"`
	AttackPower                               int64              `json:"attackPower"`
	Defense                                   int64              `json:"defense"`
	HPRegen                                   float64            `json:"hpRegen"`
	SPRegen                                   float64            `json:"spRegen"`
	AttackSpeed                               float64            `json:"attackSpeed"`
	MoveSpeed                                 float64            `json:"moveSpeed"`
	OutOfCombatMoveSpeed                      float64            `json:"outOfCombatMoveSpeed"`
	SightRange                                float64            `json:"sightRange"`
	AttackRange                               float64            `json:"attackRange"`
	CriticalStrikeChance                      float64            `json:"criticalStrikeChance"`
	CriticalStrikeDamage                      float64            `json:"criticalStrikeDamage"`
	CoolDownReduction                         float64            `json:"coolDownReduction"`
	LifeSteal                                 float64            `json:"lifeSteal"`
	NormalLifeSteal                           float64            `json:"normalLifeSteal"`
	SkillLifeSteal                            float64            `json:"skillLifeSteal"`
	AmplifierToMonster                        float64            `json:"amplifierToMonster"`
	TrapDamage                                float64            `json:"trapDamage"`
	BonusCoin                                 int64              `json:"bonusCoin"`
	GainExp                                   int64              `json:"gainExp"`
	BaseExp                                   int64              `json:"baseExp"`
	BonusExp                                  int64              `json:"bonusExp"`
	StartDtm                                  string             `json:"startDtm"`
	Duration                                  int64              `json:"duration"`
	MmrBefore                                 *int64             `json:"mmrBefore,omitempty"`
	MmrGain                                   *int64             `json:"mmrGain,omitempty"`
	MmrAfter                                  *int64             `json:"mmrAfter,omitempty"`
	PlayTime                                  int64              `json:"playTime"`
	WatchTime                                 int64              `json:"watchTime"`
	TotalTime                                 int64              `json:"totalTime"`
	SurvivableTime                            int64              `json:"survivableTime"`
	BotAdded                                  int64              `json:"botAdded"`
	BotRemain                                 int64              `json:"botRemain"`
	RestrictedAreaAccelerated                 int64              `json:"restrictedAreaAccelerated"`
	SafeAreas                                 int64              `json:"safeAreas"`
	TeamNumber                                int64              `json:"teamNumber"`
	PreMade                                   int64              `json:"preMade"`
	EventMissionResult                        string             `json:"eventMissionResult"`
	GainedNormalMmrKFactor                    float64            `json:"gainedNormalMmrKFactor"`
	Victory                                   int64              `json:"victory"`
	CraftUncommon                             int64              `json:"craftUncommon"`
	CraftRare                                 int64              `json:"craftRare"`
	CraftEpic                                 int64              `json:"craftEpic"`
	CraftLegend                               int64              `json:"craftLegend"`
	DamageToPlayer                            int64              `json:"damageToPlayer"`
	DamageToPlayerTrap                        int64              `json:"damageToPlayer_trap"`
	DamageToPlayerBasic                       int64              `json:"damageToPlayer_basic"`
	DamageToPlayerSkill                       int64              `json:"damageToPlayer_skill"`
	DamageToPlayerItemSkill                   int64              `json:"damageToPlayer_itemSkill"`
	DamageToPlayerDirect                      int64              `json:"damageToPlayer_direct"`
	DamageToPlayerUniqueSkill                 int64              `json:"damageToPlayer_uniqueSkill"`
	DamageFromPlayer                          int64              `json:"damageFromPlayer"`
	DamageFromPlayerTrap                      int64              `json:"damageFromPlayer_trap"`
	DamageFromPlayerBasic                     int64              `json:"damageFromPlayer_basic"`
	DamageFromPlayerSkill                     int64              `json:"damageFromPlayer_skill"`
	DamageFromPlayerItemSkill                 int64              `json:"damageFromPlayer_itemSkill"`
	DamageFromPlayerDirect                    int64              `json:"damageFromPlayer_direct"`
	DamageFromPlayerUniqueSkill               int64              `json:"damageFromPlayer_uniqueSkill"`
	DamageToMonster                           int64              `json:"damageToMonster"`
	DamageToMonsterTrap                       int64              `json:"damageToMonster_trap"`
	DamageToMonsterBasic                      int64              `json:"damageToMonster_basic"`
	DamageToMonsterSkill                      int64              `json:"damageToMonster_skill"`
	DamageToMonsterItemSkill                  int64              `json:"damageToMonster_itemSkill"`
	DamageToMonsterDirect                     int64              `json:"damageToMonster_direct"`
	DamageToMonsterUniqueSkill                int64              `json:"damageToMonster_uniqueSkill"`
	DamageFromMonster                         int64              `json:"damageFromMonster"`
	DamageToPlayerShield                      int64              `json:"damageToPlayer_Shield"`
	DamageOffsetedByShieldPlayer              int64              `json:"damageOffsetedByShield_Player"`
	DamageOffsetedByShieldMonster             int64              `json:"damageOffsetedByShield_Monster"`
	KillMonsters                              map[string]int64   `json:"killMonsters"`
	HealAmount                                int64              `json:"healAmount"`
	TeamRecover                               int64              `json:"teamRecover"`
	ProtectAbsorb                             int64              `json:"protectAbsorb"`
	AddSurveillanceCamera                     int64              `json:"addSurveillanceCamera"`
	AddTelephotoCamera                        int64              `json:"addTelephotoCamera"`
	RemoveSurveillanceCamera                  int64              `json:"removeSurveillanceCamera"`
	RemoveTelephotoCamera                     int64              `json:"removeTelephotoCamera"`
	UseHyperLoop                              int64              `json:"useHyperLoop"`
	UseSecurityConsole                        int64              `json:"useSecurityConsole"`
	GiveUp                                    int64              `json:"giveUp"`
	TeamSpectator                             int64              `json:"teamSpectator"`
	RouteIDOfStart                            int64              `json:"routeIdOfStart"`
	RouteSlotID                               int64              `json:"routeSlotId"`
	PlaceOfStart                              string             `json:"placeOfStart"`
	MmrAvg                                    *int64             `json:"mmrAvg,omitempty"`
	MatchSize                                 int64              `json:"matchSize"`
	TeamKill                                  int64              `json:"teamKill"`
	TotalFieldKill                            int64              `json:"totalFieldKill"`
	AccountLevel                              int64              `json:"accountLevel"`
	KillerUserNum                             int64              `json:"killerUserNum"`
	Killer                                    string             `json:"killer"`
	KillDetail                                string             `json:"killDetail"`
	CauseOfDeath                              string             `json:"causeOfDeath"`
	PlaceOfDeath                              string             `json:"placeOfDeath"`
	KillerCharacter                           string             `json:"killerCharacter"`
	KillerWeapon                              string             `json:"killerWeapon"`
	KillerUserNum2                            int64              `json:"killerUserNum2"`
	Killer2                                   *string            `json:"killer2,omitempty"`
	KillDetail2                               *string            `json:"killDetail2,omitempty"`
	CauseOfDeath2                             *string            `json:"causeOfDeath2,omitempty"`
	PlaceOfDeath2                             *string            `json:"placeOfDeath2,omitempty"`
	KillerCharacter2                          *string            `json:"killerCharacter2,omitempty"`
	KillerWeapon2                             *string            `json:"killerWeapon2,omitempty"`
	KillerUserNum3                            int64              `json:"killerUserNum3"`
	FishingCount                              int64              `json:"fishingCount"`
	UseEmoticonCount                          int64              `json:"useEmoticonCount"`
	ExpireDtm                                 string             `json:"expireDtm"`
	TraitFirstCore                            int64              `json:"traitFirstCore"`
	TraitFirstSub                             []int64            `json:"traitFirstSub"`
	TraitSecondSub                            []int64            `json:"traitSecondSub"`
	AirSupplyOpenCount                        []int64            `json:"airSupplyOpenCount"`
	FoodCraftCount                            []int64            `json:"foodCraftCount"`
	BeverageCraftCount                        []int64            `json:"beverageCraftCount"`
	RankPoint                                 int64              `json:"rankPoint"`
	TotalTurbineTakeOver                      int64              `json:"totalTurbineTakeOver"`
	UsedNormalHealPack                        int64              `json:"usedNormalHealPack"`
	UsedReinforcedHealPack                    int64              `json:"usedReinforcedHealPack"`
	UsedNormalShieldPack                      int64              `json:"usedNormalShieldPack"`
	UsedReinforceShieldPack                   int64              `json:"usedReinforceShieldPack"`
	TotalVFCredits                            []int64            `json:"totalVFCredits"`
	ActivelyGainedCredits                     int64              `json:"activelyGainedCredits"`
	UsedVFCredits                             []int64            `json:"usedVFCredits"`
	SumUsedVFCredits                          int64              `json:"sumUsedVFCredits"`
	CraftMythic                               int64              `json:"craftMythic"`
	PlayerDeaths                              int64              `json:"playerDeaths"`
	KillGamma                                 bool               `json:"killGamma"`
	ScoredPoint                               []int64            `json:"scoredPoint"`
	KillDetails                               string             `json:"killDetails"`
	DeathDetails                              string             `json:"deathDetails"`
	KillsPhaseOne                             int64              `json:"killsPhaseOne"`
	KillsPhaseTwo                             int64              `json:"killsPhaseTwo"`
	KillsPhaseThree                           int64              `json:"killsPhaseThree"`
	DeathsPhaseOne                            int64              `json:"deathsPhaseOne"`
	DeathsPhaseTwo                            int64              `json:"deathsPhaseTwo"`
	DeathsPhaseThree                          int64              `json:"deathsPhaseThree"`
	UsedPairLoop                              int64              `json:"usedPairLoop"`
	CcTimeToPlayer                            float64            `json:"ccTimeToPlayer"`
	CreditSource                              map[string]float64 `json:"creditSource"`
	BoughtInfusion                            string             `json:"boughtInfusion"`
	ItemTransferredConsole                    []int64            `json:"itemTransferredConsole"`
	ItemTransferredDrone                      []int64            `json:"itemTransferredDrone"`
	EscapeState                               int64              `json:"escapeState"`
	TotalDoubleKill                           int64              `json:"totalDoubleKill"`
	TotalTripleKill                           int64              `json:"totalTripleKill"`
	TotalQuadraKill                           int64              `json:"totalQuadraKill"`
	TotalExtraKill                            int64              `json:"totalExtraKill"`
	CollectItemForLog                         []int64            `json:"collectItemForLog"`
	EquipFirstItemForLog                      map[string][]int64 `json:"equipFirstItemForLog"`
	BattleZone1AreaCode                       int64              `json:"battleZone1AreaCode"`
	BattleZone1BattleMark                     int64              `json:"battleZone1BattleMark"`
	BattleZone1ItemCode                       []int64            `json:"battleZone1ItemCode"`
	BattleZone2AreaCode                       int64              `json:"battleZone2AreaCode"`
	BattleZone2BattleMark                     int64              `json:"battleZone2BattleMark"`
	BattleZone2ItemCode                       []int64            `json:"battleZone2ItemCode"`
	BattleZone3AreaCode                       int64              `json:"battleZone3AreaCode"`
	BattleZone3BattleMark                     int64              `json:"battleZone3BattleMark"`
	BattleZone3ItemCode                       []int64            `json:"battleZone3ItemCode"`
	BattleZonePlayerKill                      int64              `json:"battleZonePlayerKill"`
	BattleZoneDeaths                          int64              `json:"battleZoneDeaths"`
	BattleZone1Winner                         int64              `json:"battleZone1Winner"`
	BattleZone2Winner                         int64              `json:"battleZone2Winner"`
	BattleZone3Winner                         int64              `json:"battleZone3Winner"`
	BattleZone1BattleMarkCount                int64              `json:"battleZone1BattleMarkCount"`
	BattleZone2BattleMarkCount                int64              `json:"battleZone2BattleMarkCount"`
	BattleZone3BattleMarkCount                int64              `json:"battleZone3BattleMarkCount"`
	TacticalSkillGroup                        int64              `json:"tacticalSkillGroup"`
	TacticalSkillLevel                        int64              `json:"tacticalSkillLevel"`
	TotalGainVFCredit                         int64              `json:"totalGainVFCredit"`
	KillPlayerGainVFCredit                    int64              `json:"killPlayerGainVFCredit"`
	KillChickenGainVFCredit                   int64              `json:"killChickenGainVFCredit"`
	KillBoarGainVFCredit                      int64              `json:"killBoarGainVFCredit"`
	KillWildDogGainVFCredit                   int64              `json:"killWildDogGainVFCredit"`
	KillWolfGainVFCredit                      int64              `json:"killWolfGainVFCredit"`
	KillBearGainVFCredit                      int64              `json:"killBearGainVFCredit"`
	KillOmegaGainVFCredit                     int64              `json:"killOmegaGainVFCredit"`
	KillBatGainVFCredit                       int64              `json:"killBatGainVFCredit"`
	KillWicklineGainVFCredit                  int64              `json:"killWicklineGainVFCredit"`
	KillAlphaGainVFCredit                     int64              `json:"killAlphaGainVFCredit"`
	KillItemBountyGainVFCredit                int64              `json:"killItemBountyGainVFCredit"`
	KillDroneGainVFCredit                     int64              `json:"killDroneGainVFCredit"`
	KillGammaGainVFCredit                     int64              `json:"killGammaGainVFCredit"`
	KillTurretGainVFCredit                    int64              `json:"killTurretGainVFCredit"`
	ItemShredderGainVFCredit                  int64              `json:"itemShredderGainVFCredit"`
	TotalUseVFCredit                          int64              `json:"totalUseVFCredit"`
	RemoteDroneUseVFCreditMySelf              int64              `json:"remoteDroneUseVFCreditMySelf"`
	RemoteDroneUseVFCreditAlly                int64              `json:"remoteDroneUseVFCreditAlly"`
	TransferConsoleFromMaterialUseVFCredit    int64              `json:"transferConsoleFromMaterialUseVFCredit"`
	TransferConsoleFromEscapeKeyUseVFCredit   int64              `json:"transferConsoleFromEscapeKeyUseVFCredit"`
	TransferConsoleFromRevivalUseVFCredit     int64              `json:"transferConsoleFromRevivalUseVFCredit"`
	TacticalSkillUpgradeUseVFCredit           int64              `json:"tacticalSkillUpgradeUseVFCredit"`
	InfusionReRollUseVFCredit                 int64              `json:"infusionReRollUseVFCredit"`
	InfusionTraitUseVFCredit                  int64              `json:"infusionTraitUseVFCredit"`
	InfusionRelicUseVFCredit                  int64              `json:"infusionRelicUseVFCredit"`
	InfusionStoreUseVFCredit                  int64              `json:"infusionStoreUseVFCredit"`
	TeamElimination                           int64              `json:"teamElimination"`
	TeamDown                                  int64              `json:"teamDown"`
	TeamBattleZoneDown                        int64              `json:"teamBattleZoneDown"`
	TeamRepeatDown                            int64              `json:"teamRepeatDown"`
	AdaptiveForce                             int64              `json:"adaptiveForce"`
	AdaptiveForceAttack                       int64              `json:"adaptiveForceAttack"`
	AdaptiveForceAmplify                      int64              `json:"adaptiveForceAmplify"`
	SkillAmp                                  int64              `json:"skillAmp"`
	CampFireCraftUncommon                     int64              `json:"campFireCraftUncommon"`
	CampFireCraftRare                         int64              `json:"campFireCraftRare"`
	CampFireCraftEpic                         int64              `json:"campFireCraftEpic"`
	CampFireCraftLegendary                    int64              `json:"campFireCraftLegendary"`
	CobaltRandomPickRemoveCharacter           int64              `json:"cobaltRandomPickRemoveCharacter"`
	TacticalSkillUseCount                     int64              `json:"tacticalSkillUseCount"`
	CreditRevivalCount                        int64              `json:"creditRevivalCount"`
	CreditRevivedOthersCount                  int64              `json:"creditRevivedOthersCount"`
	TimeSpentInBriefingRoom                   int64              `json:"timeSpentInBriefingRoom"`
	IsLeavingBeforeCreditRevivalTerminate     bool               `json:"IsLeavingBeforeCreditRevivalTerminate"`
	CRGetAnimal                               int64              `json:"crGetAnimal"`
	CRGetMutant                               int64              `json:"crGetMutant"`
	CRGetPhaseStart                           int64              `json:"crGetPhaseStart"`
	CRGetKill                                 int64              `json:"crGetKill"`
	CRGetAssist                               int64              `json:"crGetAssist"`
	CRGetTimeElapsed                          int64              `json:"crGetTimeElapsed"`
	CRGetCreditBonus                          int64              `json:"crGetCreditBonus"`
	CRUseRemoteDrone                          int64              `json:"crUseRemoteDrone"`
	CRUseUpgradeTacticalSkill                 int64              `json:"crUseUpgradeTacticalSkill"`
	CRUseTreeOfLife                           int64              `json:"crUseTreeOfLife"`
	CRUseMeteorite                            int64              `json:"crUseMeteorite"`
	CRUseMythril                              int64              `json:"crUseMythril"`
	CRUseForceCore                            int64              `json:"crUseForceCore"`
	CRUseVFBloodSample                        int64              `json:"crUseVFBloodSample"`
	CRUseActivationModule                     int64              `json:"crUseActivationModule"`
	CRUseRootkit                              int64              `json:"crUseRootkit"`
	MmrGainInGame                             int64              `json:"mmrGainInGame"`
	MmrLossEntryCost                          int64              `json:"mmrLossEntryCost"`
	PremadeMatchingType                       int64              `json:"premadeMatchingType"`
	ViewContribution                          int64              `json:"viewContribution"`
	UseReconDrone                             int64              `json:"useReconDrone"`
	UseEmpDrone                               int64              `json:"useEmpDrone"`
	ExceptPreMadeTeam                         bool               `json:"exceptPreMadeTeam"`
	TerminateCount                            int64              `json:"terminateCount"`
	ClutchCount                               int64              `json:"clutchCount"`
	UnknownKill                               int64              `json:"unknownKill"`
	MainWeather                               int64              `json:"mainWeather"`
	SubWeather                                int64              `json:"subWeather"`
	ActiveInstallation                        map[string]int64   `json:"activeInstallation"`
	UseGuideRobot                             int64              `json:"useGuideRobot"`
	GuideRobotRadial                          int64              `json:"guideRobotRadial"`
	GuideRobotFlagShip                        int64              `json:"guideRobotFlagShip"`
	GuideRobotSignature                       int64              `json:"guideRobotSignature"`
	CRGetByGuideRobot                         int64              `json:"crGetByGuideRobot"`
	DamageToGuideRobot                        int64              `json:"damageToGuideRobot"`
	GetBuffCubeRed                            int64              `json:"getBuffCubeRed"`
	GetBuffCubePurple                         int64              `json:"getBuffCubePurple"`
	GetBuffCubeGreen                          int64              `json:"getBuffCubeGreen"`
	GetBuffCubeGold                           int64              `json:"getBuffCubeGold"`
	GetBuffCubeSkyBlue                        int64              `json:"getBuffCubeSkyBlue"`
	SumGetBuffCube                            int64              `json:"sumGetBuffCube"`
	GameIsLeavingBeforeCreditRevivalTerminate bool               `json:"isLeavingBeforeCreditRevivalTerminate"`
	Killer3                                   *string            `json:"killer3,omitempty"`
	KillDetail3                               *string            `json:"killDetail3,omitempty"`
	CauseOfDeath3                             *string            `json:"causeOfDeath3,omitempty"`
	PlaceOfDeath3                             *string            `json:"placeOfDeath3,omitempty"`
	KillerCharacter3                          *string            `json:"killerCharacter3,omitempty"`
	KillerWeapon3                             *string            `json:"killerWeapon3,omitempty"`
	SumTotalVFCredits                         *int64             `json:"sumTotalVFCredits,omitempty"`
}

type Nickname struct {
	UserNum  int64  `json:"userNum"`
	Nickname string `json:"nickname"`
}

type Rank struct {
	UserNum     int64         `json:"userNum"`
	Mmr         int64         `json:"mmr"`
	Nickname    string        `json:"nickname"`
	Rank        int64         `json:"rank"`
	UserEmblems []interface{} `json:"userEmblems,omitempty"`
}

type UserStat struct {
	SeasonID          int64           `json:"seasonId"`
	UserNum           int64           `json:"userNum"`
	MatchingMode      int64           `json:"matchingMode"`
	MatchingTeamMode  int64           `json:"matchingTeamMode"`
	Mmr               int64           `json:"mmr"`
	Nickname          string          `json:"nickname"`
	Rank              int64           `json:"rank"`
	RankSize          int64           `json:"rankSize"`
	TotalGames        int64           `json:"totalGames"`
	TotalWINS         int64           `json:"totalWins"`
	TotalTeamKills    int64           `json:"totalTeamKills"`
	TotalDeaths       int64           `json:"totalDeaths"`
	EscapeCount       int64           `json:"escapeCount"`
	RankPercent       float64         `json:"rankPercent"`
	AverageRank       float64         `json:"averageRank"`
	AverageKills      float64         `json:"averageKills"`
	AverageAssistants float64         `json:"averageAssistants"`
	AverageHunts      float64         `json:"averageHunts"`
	Top1              float64         `json:"top1"`
	Top2              float64         `json:"top2"`
	Top3              float64         `json:"top3"`
	Top5              float64         `json:"top5"`
	Top7              float64         `json:"top7"`
	CharacterStats    []CharacterStat `json:"characterStats"`
}

type CharacterStat struct {
	CharacterCode int64 `json:"characterCode"`
	TotalGames    int64 `json:"totalGames"`
	Usages        int64 `json:"usages"`
	MaxKillings   int64 `json:"maxKillings"`
	Top3          int64 `json:"top3"`
	WINS          int64 `json:"wins"`
	Top3Rate      int64 `json:"top3Rate"`
	AverageRank   int64 `json:"averageRank"`
}

type WeaponRoute struct {
	RecommendWeaponRoute     RecommendWeaponRoute     `json:"recommendWeaponRoute"`
	RecommendWeaponRouteDesc RecommendWeaponRouteDesc `json:"recommendWeaponRouteDesc"`
}

type RecommendWeaponRoute struct {
	ID                     int64   `json:"id"`
	Title                  string  `json:"title"`
	UserNum                int64   `json:"userNum"`
	UserNickname           string  `json:"userNickname"`
	CharacterCode          int64   `json:"characterCode"`
	SlotID                 int64   `json:"slotId"`
	WeaponType             int64   `json:"weaponType"`
	WeaponCodes            string  `json:"weaponCodes"`
	TacticalSkillGroupCode int64   `json:"tacticalSkillGroupCode"`
	Paths                  string  `json:"paths"`
	Count                  int64   `json:"count"`
	Version                string  `json:"version"`
	TeamMode               int64   `json:"teamMode"`
	LanguageCode           string  `json:"languageCode"`
	RouteVersion           int64   `json:"routeVersion"`
	Share                  bool    `json:"share"`
	UpdateDtm              int64   `json:"updateDtm"`
	V2Like                 int64   `json:"v2Like"`
	V2WinRate              float64 `json:"v2WinRate"`
	V2SeasonID             int64   `json:"v2SeasonId"`
	V2AccumulateLike       int64   `json:"v2AccumulateLike"`
	V2AccumulateWinRate    float64 `json:"v2AccumulateWinRate"`
	V2AccumulateSeasonID   int64   `json:"v2AccumulateSeasonId"`
}

type RecommendWeaponRouteDesc struct {
	RecommendWeaponRouteID int64   `json:"recommendWeaponRouteId"`
	SkillPath              *string `json:"skillPath,omitempty"`
	Desc                   *string `json:"desc,omitempty"`
}

type ActionCostElement struct {
	Code                  int64   `json:"code"`
	Type                  string  `json:"type"`
	SP                    int64   `json:"sp"`
	Time1                 float64 `json:"time1"`
	Time2                 float64 `json:"time2"`
	ActionWaitTime        float64 `json:"actionWaitTime"`
	CastingAnimTrigger    string  `json:"castingAnimTrigger"`
	EffectCancelCondition string  `json:"effectCancelCondition"`
	CastingBarImgName     string  `json:"castingBarImgName"`
}

type AreaElement struct {
	Code                     int64  `json:"code"`
	Name                     string `json:"name"`
	ModeType                 int64  `json:"modeType"`
	MaskCode                 int64  `json:"maskCode"`
	StartingArea             bool   `json:"startingArea"`
	AreaType                 string `json:"areaType"`
	IsProvideCollectibleItem bool   `json:"isProvideCollectibleItem"`
	RouteCalcBitCode         int64  `json:"routeCalcBitCode"`
	IsHyperLoopInstalled     bool   `json:"isHyperLoopInstalled"`
	IsNearByOcean            bool   `json:"isNearByOcean"`
}

type BattleZoneRewardElement struct {
	Code                           int64  `json:"code"`
	ModeType                       int64  `json:"modeType"`
	AreaAttributesCreateEventCount int64  `json:"areaAttributesCreateEventCount"`
	ItemCode                       int64  `json:"itemCode"`
	Type                           string `json:"type"`
	Value                          int64  `json:"value"`
	Selectable                     bool   `json:"selectable"`
}

type BulletCapacityElement struct {
	ItemCode  int64  `json:"itemCode"`
	LoadType  string `json:"loadType"`
	Time      int64  `json:"time"`
	InitCount int64  `json:"initCount"`
	Count     int64  `json:"count"`
}

type CharacterElement struct {
	Code                           int64   `json:"code"`
	Name                           string  `json:"name"`
	MaxHP                          int64   `json:"maxHp"`
	MaxSP                          int64   `json:"maxSp"`
	StrLearnStartSkill             string  `json:"strLearnStartSkill"`
	StrUsePointLearnStartSkill     string  `json:"strUsePointLearnStartSkill"`
	InitExtraPoint                 int64   `json:"initExtraPoint"`
	MaxExtraPoint                  int64   `json:"maxExtraPoint"`
	AttackPower                    int64   `json:"attackPower"`
	Defense                        int64   `json:"defense"`
	SkillAmp                       int64   `json:"skillAmp"`
	AdaptiveForce                  int64   `json:"adaptiveForce"`
	CriticalStrikeChance           int64   `json:"criticalStrikeChance"`
	HPRegen                        float64 `json:"hpRegen"`
	SPRegen                        float64 `json:"spRegen"`
	AttackSpeed                    float64 `json:"attackSpeed"`
	AttackSpeedRatio               int64   `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio int64   `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  int64   `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio int64   `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       int64   `json:"preventSkillDamagedRatio"`
	AttackSpeedLimit               float64 `json:"attackSpeedLimit"`
	AttackSpeedMin                 float64 `json:"attackSpeedMin"`
	MoveSpeed                      float64 `json:"moveSpeed"`
	SightRange                     float64 `json:"sightRange"`
	Radius                         float64 `json:"radius"`
	PathingRadius                  float64 `json:"pathingRadius"`
	UIHeight                       float64 `json:"uiHeight"`
	InitStateDisplayIndex          int64   `json:"initStateDisplayIndex"`
	LocalScaleInCutscene           int64   `json:"localScaleInCutscene"`
	LocalScaleInVictoryScene       string  `json:"localScaleInVictoryScene"`
	Resource                       string  `json:"resource"`
	LobbySubObject                 string  `json:"lobbySubObject"`
}

type CharacterAttribute struct {
	Character         string `json:"character"`
	CharacterCode     int64  `json:"characterCode"`
	Mastery           string `json:"mastery"`
	ControlDifficulty int64  `json:"controlDifficulty"`
	Attack            int64  `json:"attack"`
	Defense           int64  `json:"defense"`
	Disruptor         int64  `json:"disruptor"`
	Move              int64  `json:"move"`
	Assistance        int64  `json:"assistance"`
}

type CharacterExpElement struct {
	Level      int64 `json:"level"`
	LevelUpExp int64 `json:"levelUpExp"`
}

type CharacterLevelUpStatElement struct {
	Code                           int64   `json:"code"`
	Name                           string  `json:"name"`
	MaxHP                          int64   `json:"maxHp"`
	MaxSP                          int64   `json:"maxSp"`
	AttackPower                    float64 `json:"attackPower"`
	Defense                        float64 `json:"defense"`
	SkillAmp                       int64   `json:"skillAmp"`
	AdaptiveForce                  int64   `json:"adaptiveForce"`
	CriticalChance                 int64   `json:"criticalChance"`
	HPRegen                        float64 `json:"hpRegen"`
	SPRegen                        float64 `json:"spRegen"`
	AttackSpeed                    int64   `json:"attackSpeed"`
	MoveSpeed                      int64   `json:"moveSpeed"`
	AttackSpeedRatio               int64   `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio int64   `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  int64   `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio int64   `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       int64   `json:"preventSkillDamagedRatio"`
}

type CharacterMasteryElement struct {
	Code      int64  `json:"code"`
	Weapon1   string `json:"weapon1"`
	Weapon2   string `json:"weapon2"`
	Weapon3   string `json:"weapon3"`
	Weapon4   string `json:"weapon4"`
	Combat1   string `json:"combat1"`
	Combat2   string `json:"combat2"`
	Survival1 string `json:"survival1"`
	Survival2 string `json:"survival2"`
	Survival3 string `json:"survival3"`
}

type CharacterModeModifierElement struct {
	CharacterCode                           int64  `json:"characterCode"`
	WeaponType                              string `json:"weaponType"`
	SoloIncreaseModeDamageRatio             int64  `json:"soloIncreaseModeDamageRatio"`
	SoloPreventModeDamageRatio              int64  `json:"soloPreventModeDamageRatio"`
	SoloIncreaseModeHealRatio               int64  `json:"soloIncreaseModeHealRatio"`
	SoloIncreaseModeShieldRatio             int64  `json:"soloIncreaseModeShieldRatio"`
	DuoIncreaseModeDamageRatio              int64  `json:"duoIncreaseModeDamageRatio"`
	DuoPreventModeDamageRatio               int64  `json:"duoPreventModeDamageRatio"`
	DuoIncreaseModeHealRatio                int64  `json:"duoIncreaseModeHealRatio"`
	DuoIncreaseModeHealerGiveHealRatio      int64  `json:"duoIncreaseModeHealerGiveHealRatio"`
	DuoIncreaseModeShieldRatio              int64  `json:"duoIncreaseModeShieldRatio"`
	DuoIncreaseModeHealerGiveShieldRatio    int64  `json:"duoIncreaseModeHealerGiveShieldRatio"`
	SquadIncreaseModeDamageRatio            int64  `json:"squadIncreaseModeDamageRatio"`
	SquadPreventModeDamageRatio             int64  `json:"squadPreventModeDamageRatio"`
	SquadIncreaseModeHealRatio              int64  `json:"squadIncreaseModeHealRatio"`
	SquadIncreaseModeHealerGiveHealRatio    int64  `json:"squadIncreaseModeHealerGiveHealRatio"`
	SquadIncreaseModeShieldRatio            int64  `json:"squadIncreaseModeShieldRatio"`
	SquadIncreaseModeHealerGiveShieldRatio  int64  `json:"squadIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeDamageRatio           int64  `json:"cobaltIncreaseModeDamageRatio"`
	CobaltPreventModeDamageRatio            int64  `json:"cobaltPreventModeDamageRatio"`
	CobaltIncreaseModeHealRatio             int64  `json:"cobaltIncreaseModeHealRatio"`
	CobaltIncreaseModeHealerGiveHealRatio   int64  `json:"cobaltIncreaseModeHealerGiveHealRatio"`
	CobaltIncreaseModeShieldRatio           int64  `json:"cobaltIncreaseModeShieldRatio"`
	CobaltIncreaseModeHealerGiveShieldRatio int64  `json:"cobaltIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeUltCooldownRatio      int64  `json:"cobaltIncreaseModeUltCooldownRatio"`
	CobaltIncreaseModeMaxSPRatio            int64  `json:"cobaltIncreaseModeMaxSpRatio"`
	CobaltIncreaseModeSPRegenRatio          int64  `json:"cobaltIncreaseModeSpRegenRatio"`
	SoloIncreaseModeDamageToMonsterRatio    int64  `json:"soloIncreaseModeDamageToMonsterRatio"`
	DuoIncreaseModeDamageToMonsterRatio     int64  `json:"duoIncreaseModeDamageToMonsterRatio"`
	SquadIncreaseModeDamageToMonsterRatio   int64  `json:"squadIncreaseModeDamageToMonsterRatio"`
	CobaltIncreaseModeDamageToMonsterRatio  int64  `json:"cobaltIncreaseModeDamageToMonsterRatio"`
}

type CharacterSkinElement struct {
	Name                     string `json:"name"`
	Code                     int64  `json:"code"`
	CharacterCode            int64  `json:"characterCode"`
	Index                    int64  `json:"index"`
	Grade                    int64  `json:"grade"`
	EventFree                bool   `json:"eventFree"`
	PurchaseType             string `json:"purchaseType"`
	EffectsPath              string `json:"effectsPath"`
	ProjectilesPath          string `json:"projectilesPath"`
	ObjectPath               string `json:"objectPath"`
	FxSoundPath              string `json:"fxSoundPath"`
	VoicePath                string `json:"voicePath"`
	WeaponMountPath          string `json:"weaponMountPath"`
	WeaponMountCommonPath    string `json:"weaponMountCommonPath"`
	IndicatorPath            string `json:"indicatorPath"`
	ProjectilesDeflectorPath string `json:"projectilesDeflectorPath"`
}

type CollectibleElement struct {
	Code              int64  `json:"code"`
	Cooldown          int64  `json:"cooldown"`
	ItemCode1         string `json:"itemCode1"`
	ItemCode2         string `json:"itemCode2"`
	Probability1      int64  `json:"probability1"`
	Probability2      int64  `json:"probability2"`
	DropCount         int64  `json:"dropCount"`
	CastingActionType string `json:"castingActionType"`
}

type DropGroupElement struct {
	GroupCode   int64   `json:"groupCode"`
	ItemCode    *string `json:"itemCode"`
	Min         int64   `json:"min"`
	Max         int64   `json:"max"`
	Probability int64   `json:"probability"`
	DropType    string  `json:"dropType"`
}

type GainExpElement struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	GainExp   int64 `json:"gainExp"`
}

type GainScoreElement struct {
	Code           int64  `json:"code"`
	Phase          int64  `json:"phase"`
	ConditionType  string `json:"conditionType"`
	ConditionValue int64  `json:"conditionValue"`
	PointsEnemy    int64  `json:"pointsEnemy"`
	PointsAlly     int64  `json:"pointsAlly"`
}

type GameTipElement struct {
	Key             int64  `json:"key"`
	Code            int64  `json:"code"`
	GameTipType     string `json:"gameTipType"`
	GameTipCategory string `json:"gameTipCategory"`
	Sequence        int64  `json:"sequence"`
	TitleTextKey    string `json:"titleTextKey"`
	ContentTextKey  string `json:"contentTextKey"`
	ImageName       string `json:"imageName"`
	KR              string `json:"##참고"`
	Link            string `json:"link"`
}

type HowToFindItemElement struct {
	Code            int64   `json:"code"`
	ItemCode        *string `json:"itemCode"`
	HuntChicken     int64   `json:"huntChicken"`
	HuntBat         int64   `json:"huntBat"`
	HuntBoar        int64   `json:"huntBoar"`
	HuntWildDog     int64   `json:"huntWildDog"`
	HuntWolf        int64   `json:"huntWolf"`
	HuntBear        int64   `json:"huntBear"`
	HuntWickline    int64   `json:"huntWickline"`
	HuntAlpha       int64   `json:"huntAlpha"`
	HuntOmega       int64   `json:"huntOmega"`
	CollectibleCode int64   `json:"collectibleCode"`
	AirSupply       int64   `json:"airSupply"`
}

type InfusionProductElement struct {
	Code             int64   `json:"code"`
	ProductType      string  `json:"productType"`
	ProductGroup     int64   `json:"productGroup"`
	ProductCode      int64   `json:"productCode"`
	StoreType        string  `json:"storeType"`
	StockType        string  `json:"stockType"`
	Stock            int64   `json:"stock"`
	IsRestore        bool    `json:"isRestore"`
	Price            int64   `json:"price"`
	SpecialWeight    int64   `json:"specialWeight"`
	Weight           int64   `json:"weight"`
	Requirement      int64   `json:"requirement"`
	Icon             string  `json:"icon"`
	SimpleIcon       string  `json:"simpleIcon"`
	AlertInSpectator bool    `json:"alertInSpectator"`
	CharacterCodes   *string `json:"characterCodes"`
}

type ItemArmorElement struct {
	Code                                               int64   `json:"code"`
	Name                                               string  `json:"name"`
	ModeType                                           int64   `json:"modeType"`
	ItemType                                           string  `json:"itemType"`
	ArmorType                                          *string `json:"armorType,omitempty"`
	ItemGrade                                          string  `json:"itemGrade"`
	GradeBgOverride                                    string  `json:"gradeBgOverride"`
	IsCompletedItem                                    bool    `json:"isCompletedItem"`
	AlertInSpectator                                   bool    `json:"alertInSpectator"`
	MarkingType                                        string  `json:"markingType"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger"`
	Stackable                                          int64   `json:"stackable"`
	InitialCount                                       int64   `json:"initialCount"`
	ItemUsableType                                     string  `json:"itemUsableType"`
	ItemUsableValueList                                int64   `json:"itemUsableValueList"`
	ExclusiveProducer                                  int64   `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	MakeMaterial1                                      int64   `json:"makeMaterial1"`
	MakeMaterial2                                      int64   `json:"makeMaterial2"`
	MakeCustomAction                                   string  `json:"makeCustomAction"`
	CorpseEffectEnable                                 bool    `json:"corpseEffectEnable"`
	NotDisarm                                          bool    `json:"notDisarm"`
	ManufacturableType                                 int64   `json:"manufacturableType"`
	AttackPower                                        int64   `json:"attackPower"`
	AttackPowerByLV                                    int64   `json:"attackPowerByLv"`
	Defense                                            int64   `json:"defense"`
	DefenseByLV                                        int64   `json:"defenseByLv"`
	SkillAmp                                           int64   `json:"skillAmp"`
	SkillAmpByLevel                                    int64   `json:"skillAmpByLevel"`
	SkillAmpRatio                                      int64   `json:"skillAmpRatio"`
	SkillAmpRatioByLevel                               int64   `json:"skillAmpRatioByLevel"`
	AdaptiveForce                                      int64   `json:"adaptiveForce"`
	AdaptiveForceByLevel                               int64   `json:"adaptiveForceByLevel"`
	MaxHP                                              int64   `json:"maxHp"`
	MaxHPByLV                                          int64   `json:"maxHpByLv"`
	ItemMaxSP                                          *int64  `json:"maxSp,omitempty"`
	HPRegenRatio                                       float64 `json:"hpRegenRatio"`
	HPRegen                                            int64   `json:"hpRegen"`
	SPRegenRatio                                       float64 `json:"spRegenRatio"`
	SPRegen                                            int64   `json:"spRegen"`
	AttackSpeedRatio                                   float64 `json:"attackSpeedRatio"`
	AttackSpeedRatioByLV                               int64   `json:"attackSpeedRatioByLv"`
	CriticalStrikeChance                               float64 `json:"criticalStrikeChance"`
	CriticalStrikeDamage                               float64 `json:"criticalStrikeDamage"`
	PreventCriticalStrikeDamaged                       int64   `json:"preventCriticalStrikeDamaged"`
	CooldownReduction                                  float64 `json:"cooldownReduction"`
	CooldownLimit                                      int64   `json:"cooldownLimit"`
	LifeSteal                                          float64 `json:"lifeSteal"`
	NormalLifeSteal                                    float64 `json:"normalLifeSteal"`
	SkillLifeSteal                                     int64   `json:"skillLifeSteal"`
	MoveSpeed                                          float64 `json:"moveSpeed"`
	MoveSpeedOutOfCombat                               int64   `json:"moveSpeedOutOfCombat"`
	SightRange                                         float64 `json:"sightRange"`
	AttackRange                                        int64   `json:"attackRange"`
	IncreaseBasicAttackDamage                          int64   `json:"increaseBasicAttackDamage"`
	IncreaseBasicAttackDamageByLV                      int64   `json:"increaseBasicAttackDamageByLv"`
	PreventBasicAttackDamaged                          int64   `json:"preventBasicAttackDamaged"`
	PreventBasicAttackDamagedByLV                      int64   `json:"preventBasicAttackDamagedByLv"`
	PreventBasicAttackDamagedRatio                     int64   `json:"preventBasicAttackDamagedRatio"`
	PreventBasicAttackDamagedRatioByLV                 int64   `json:"preventBasicAttackDamagedRatioByLv"`
	IncreaseBasicAttackDamageRatio                     int64   `json:"increaseBasicAttackDamageRatio"`
	IncreaseBasicAttackDamageRatioByLV                 float64 `json:"increaseBasicAttackDamageRatioByLv"`
	PreventSkillDamaged                                int64   `json:"preventSkillDamaged"`
	PreventSkillDamagedByLV                            int64   `json:"preventSkillDamagedByLv"`
	PreventSkillDamagedRatio                           float64 `json:"preventSkillDamagedRatio"`
	PreventSkillDamagedRatioByLV                       int64   `json:"preventSkillDamagedRatioByLv"`
	PenetrationDefense                                 int64   `json:"penetrationDefense"`
	PenetrationDefenseRatio                            float64 `json:"penetrationDefenseRatio"`
	TrapDamageReduce                                   int64   `json:"trapDamageReduce"`
	TrapDamageReduceRatio                              int64   `json:"trapDamageReduceRatio"`
	HPHealedIncreaseRatio                              int64   `json:"hpHealedIncreaseRatio"`
	HealerGiveHPHealRatio                              float64 `json:"healerGiveHpHealRatio"`
	UniqueAttackRange                                  float64 `json:"uniqueAttackRange"`
	UniqueHPHealedIncreaseRatio                        int64   `json:"uniqueHpHealedIncreaseRatio"`
	UniqueCooldownLimit                                float64 `json:"uniqueCooldownLimit"`
	UniqueTenacity                                     float64 `json:"uniqueTenacity"`
	UniqueMoveSpeed                                    float64 `json:"uniqueMoveSpeed"`
	UniquePenetrationDefense                           int64   `json:"uniquePenetrationDefense"`
	UniquePenetrationDefenseRatio                      int64   `json:"uniquePenetrationDefenseRatio"`
	UniqueLifeSteal                                    int64   `json:"uniqueLifeSteal"`
	UniqueSkillAmpRatio                                float64 `json:"uniqueSkillAmpRatio"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int64   `json:"creditValueWhenConvertedToBounty"`
	WeaponType                                         *string `json:"weaponType,omitempty"`
	Consumable                                         *bool   `json:"consumable,omitempty"`
	MaxSP                                              *int64  `json:"maxSP,omitempty"`
}

type ItemConsumableElement struct {
	Code                                               int64  `json:"code"`
	Name                                               string `json:"name"`
	ModeType                                           int64  `json:"modeType"`
	ItemType                                           string `json:"itemType"`
	ConsumableType                                     string `json:"consumableType"`
	ConsumableTag                                      string `json:"consumableTag"`
	ItemGrade                                          string `json:"itemGrade"`
	IsCompletedItem                                    bool   `json:"isCompletedItem"`
	AlertInSpectator                                   bool   `json:"alertInSpectator"`
	MarkingType                                        string `json:"markingType"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger"`
	Stackable                                          int64  `json:"stackable"`
	InitialCount                                       int64  `json:"initialCount"`
	ItemUsableType                                     string `json:"itemUsableType"`
	ItemUsableValueList                                int64  `json:"itemUsableValueList"`
	ExclusiveProducer                                  int64  `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int64  `json:"manufacturableType"`
	MakeMaterial1                                      int64  `json:"makeMaterial1"`
	MakeMaterial2                                      int64  `json:"makeMaterial2"`
	CorpseEffectEnable                                 bool   `json:"corpseEffectEnable"`
	Heal                                               int64  `json:"heal"`
	HPRecover                                          int64  `json:"hpRecover"`
	SPRecover                                          int64  `json:"spRecover"`
	AttackPowerByBuff                                  int64  `json:"attackPowerByBuff"`
	DefenseByBuff                                      int64  `json:"defenseByBuff"`
	SkillAmpByBuff                                     int64  `json:"skillAmpByBuff"`
	SkillAmpRatioByBuff                                int64  `json:"skillAmpRatioByBuff"`
	AddStateCode                                       int64  `json:"addStateCode"`
	IsVPadQuickSlotItem                                bool   `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int64  `json:"creditValueWhenConvertedToBounty"`
	IsReduceLootOnDeath                                bool   `json:"isReduceLootOnDeath"`
}

type ItemMiscElement struct {
	Code                                               int64  `json:"code"`
	Name                                               string `json:"name"`
	ModeType                                           int64  `json:"modeType"`
	ItemType                                           string `json:"itemType"`
	MiscItemType                                       string `json:"miscItemType"`
	ItemGrade                                          string `json:"itemGrade"`
	GradeBgOverride                                    string `json:"gradeBgOverride"`
	IsCompletedItem                                    bool   `json:"isCompletedItem"`
	AlertInSpectator                                   bool   `json:"alertInSpectator"`
	MarkingType                                        string `json:"markingType"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger"`
	Stackable                                          int64  `json:"stackable"`
	InitialCount                                       int64  `json:"initialCount"`
	ItemUsableType                                     string `json:"itemUsableType"`
	ItemUsableValueList                                int64  `json:"itemUsableValueList"`
	ExclusiveProducer                                  int64  `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int64  `json:"manufacturableType"`
	MakeMaterial1                                      int64  `json:"makeMaterial1"`
	MakeMaterial2                                      int64  `json:"makeMaterial2"`
	MakeCustomAction                                   string `json:"makeCustomAction"`
	CorpseEffectEnable                                 bool   `json:"corpseEffectEnable"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int64  `json:"creditValueWhenConvertedToBounty"`
}

type ItemSearchOptionV2Element struct {
	Code int64  `json:"code"`
	Name string `json:"name"`
	Tag1 string `json:"tag1"`
	Tag2 string `json:"tag2"`
	Tag3 string `json:"tag3"`
}

type ItemSpawnElement struct {
	Code           int64  `json:"code"`
	AreaCode       int64  `json:"areaCode"`
	AreaSpawnGroup int64  `json:"areaSpawnGroup"`
	ItemCode       int64  `json:"itemCode"`
	DropPoint      string `json:"dropPoint"`
	DropCount      int64  `json:"dropCount"`
}

type ItemSpecialElement struct {
	Code                                               int64   `json:"code"`
	Name                                               string  `json:"name"`
	ModeType                                           int64   `json:"modeType"`
	ItemType                                           string  `json:"itemType"`
	SpecialItemType                                    string  `json:"specialItemType"`
	ItemGrade                                          string  `json:"itemGrade"`
	IsCompletedItem                                    bool    `json:"isCompletedItem"`
	AlertInSpectator                                   bool    `json:"alertInSpectator"`
	MarkingType                                        string  `json:"markingType"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger"`
	Stackable                                          int64   `json:"stackable"`
	InitialCount                                       int64   `json:"initialCount"`
	CooldownGroupCode                                  int64   `json:"cooldownGroupCode"`
	Cooldown                                           float64 `json:"cooldown"`
	ItemUsableType                                     string  `json:"itemUsableType"`
	ItemUsableValueList                                int64   `json:"itemUsableValueList"`
	ExclusiveProducer                                  *string `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int64   `json:"manufacturableType"`
	MakeMaterial1                                      int64   `json:"makeMaterial1"`
	MakeMaterial2                                      int64   `json:"makeMaterial2"`
	MakeCustomAction                                   string  `json:"makeCustomAction"`
	CorpseEffectEnable                                 bool    `json:"corpseEffectEnable"`
	ConsumeCount                                       int64   `json:"consumeCount"`
	SummonCode                                         int64   `json:"summonCode"`
	GhostItemStateGroup                                int64   `json:"ghostItemStateGroup"`
	IsVPadQuickSlotItem                                bool    `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int64   `json:"creditValueWhenConvertedToBounty"`
	IsReduceLootOnDeath                                bool    `json:"isReduceLootOnDeath"`
}

type LevelElement struct {
	Level         int64 `json:"level"`
	NeedExp       int64 `json:"needExp"`
	AccumulateExp int64 `json:"accumulateExp"`
	RewardAcoin   int64 `json:"rewardAcoin"`
	Reward        int64 `json:"reward"`
}

type LoadingTipElement struct {
	Code           int64  `json:"code"`
	LoadingTipType string `json:"loadingTipType"`
	MinLV          int64  `json:"minLv"`
	MaxLV          int64  `json:"maxLv"`
	TextKey        string `json:"textKey"`
	ImageName      string `json:"imageName"`
}

type MasteryExpElement struct {
	Code           int64  `json:"code"`
	ModeType       int64  `json:"modeType"`
	ConditionType  string `json:"conditionType"`
	Grade          string `json:"grade"`
	ConditionValue int64  `json:"conditionValue"`
	MasteryType1   string `json:"masteryType1"`
	Value1         int64  `json:"value1"`
	MasteryType2   string `json:"masteryType2"`
	Value2         int64  `json:"value2"`
	MasteryType3   string `json:"masteryType3"`
	Value3         int64  `json:"value3"`
}

type MasteryLevelElement struct {
	Code              int64  `json:"code"`
	Type              string `json:"type"`
	MasteryLevel      int64  `json:"masteryLevel"`
	NextMasteryExp    int64  `json:"nextMasteryExp"`
	GiveLevelExp      int64  `json:"giveLevelExp"`
	ExpGrowthCapRatio int64  `json:"expGrowthCapRatio"`
}

type MasteryStatElement struct {
	Code                      int64   `json:"code"`
	Type                      string  `json:"type"`
	CharacterCode             int64   `json:"characterCode"`
	FirstOption               string  `json:"firstOption"`
	FirstOptionSection1Value  float64 `json:"firstOptionSection1Value"`
	FirstOptionSection2Value  float64 `json:"firstOptionSection2Value"`
	FirstOptionSection3Value  float64 `json:"firstOptionSection3Value"`
	FirstOptionSection4Value  float64 `json:"firstOptionSection4Value"`
	SecondOption              string  `json:"secondOption"`
	SecondOptionSection1Value float64 `json:"secondOptionSection1Value"`
	SecondOptionSection2Value float64 `json:"secondOptionSection2Value"`
	SecondOptionSection3Value float64 `json:"secondOptionSection3Value"`
	SecondOptionSection4Value float64 `json:"secondOptionSection4Value"`
	ThirdOption               string  `json:"thirdOption"`
	ThirdOptionSection1Value  int64   `json:"thirdOptionSection1Value"`
	ThirdOptionSection2Value  int64   `json:"thirdOptionSection2Value"`
	ThirdOptionSection3Value  int64   `json:"thirdOptionSection3Value"`
	ThirdOptionSection4Value  int64   `json:"thirdOptionSection4Value"`
}

type MonsterElement struct {
	Code             int64   `json:"Code"`
	Monster          string  `json:"monster"`
	IsMutant         bool    `json:"isMutant"`
	Grade            string  `json:"grade"`
	Mode             int64   `json:"mode"`
	CreateDay        string  `json:"createDay"`
	CreateTime       int64   `json:"createTime"`
	RegenTime        int64   `json:"regenTime"`
	RestoreTime      int64   `json:"restoreTime"`
	LevelUpPeriod    int64   `json:"levelUpPeriod"`
	LevelUpAmount    int64   `json:"levelUpAmount"`
	LevelUpMax       int64   `json:"levelUpMax"`
	MaxHP            int64   `json:"maxHp"`
	MaxEp            int64   `json:"maxEp"`
	InitExtraPoint   int64   `json:"initExtraPoint"`
	AttackPower      int64   `json:"attackPower"`
	Defense          int64   `json:"defense"`
	AttackSpeed      float64 `json:"attackSpeed"`
	MoveSpeed        float64 `json:"moveSpeed"`
	SightRange       int64   `json:"sightRange"`
	ChasingRange     int64   `json:"chasingRange"`
	AttackRange      float64 `json:"attackRange"`
	FirstAttackRange int64   `json:"firstAttackRange"`
	Aggressive       string  `json:"aggressive"`
	DetectInvisible  bool    `json:"detectInvisible"`
	Radius           float64 `json:"radius"`
	PathingRadius    float64 `json:"pathingRadius"`
	UIHeight         float64 `json:"uiHeight"`
	GainExp          int64   `json:"gainExp"`
	TargetOnRange    int64   `json:"targetOnRange"`
	RandomDropCount  int64   `json:"randomDropCount"`
	Resource         string  `json:"resource"`
	CorpseResource   string  `json:"corpseResource"`
	AppearTime       float64 `json:"appearTime"`
}

type MonsterDropGroupElement struct {
	MonsterCode  int64 `json:"monsterCode"`
	MonsterLevel int64 `json:"monsterLevel"`
	DropGroup    int64 `json:"dropGroup"`
}

type MonsterLevelUpStatElement struct {
	Code        int64   `json:"code"`
	Monster     string  `json:"monster"`
	Mode        int64   `json:"mode"`
	MaxHP       int64   `json:"maxHp"`
	AttackPower int64   `json:"attackPower"`
	Defense     float64 `json:"defense"`
	MoveSpeed   float64 `json:"moveSpeed"`
	GainExp     int64   `json:"gainExp"`
}

type MonsterSpawnLevelElement struct {
	Code        int64 `json:"code"`
	Mode        int64 `json:"mode"`
	PlayerLevel int64 `json:"playerLevel"`
	MonsterCode int64 `json:"monsterCode"`
	SpawnLevel  int64 `json:"spawnLevel"`
}

type NaviCollectAndHuntElement struct {
	Code         int64  `json:"code"`
	ItemCode     int64  `json:"itemCode"`
	AreaCodeList string `json:"areaCodeList"`
}

type NearByAreaElement struct {
	Code           int64 `json:"code"`
	AreaCode       int64 `json:"areaCode"`
	NearByAreaCode int64 `json:"nearByAreaCode"`
}

type RandomEquipmentElement struct {
	Code          int64   `json:"code"`
	Group         string  `json:"group"`
	Itemcode      *string `json:"itemcode"`
	Weight        int64   `json:"weight"`
	ItemGrade     string  `json:"itemGrade"`
	TagMultiplier int64   `json:"tagMultiplier"`
	CharacterNum  int64   `json:"characterNum"`
}

type RecommendedListElement struct {
	Code                  int64  `json:"code"`
	Character             string `json:"character"`
	CharacterCode         int64  `json:"characterCode"`
	Mastery               string `json:"mastery"`
	StartWeapon           int64  `json:"startWeapon"`
	CobaltStartWeapon     int64  `json:"cobaltStartWeapon"`
	StartItemGroupCode    int64  `json:"startItemGroupCode"`
	CobaltDraft           int64  `json:"cobaltDraft"`
	CobaltExtraDraft      string `json:"cobaltExtraDraft"`
	CobaltCanChooseWeapon bool   `json:"cobaltCanChooseWeapon"`
	FavoriteMainTag       string `json:"favoriteMainTag"`
	OppositeTag           string `json:"oppositeTag"`
}

type SeasonElement struct {
	SeasonID    int64     `json:"seasonID"`
	SeasonName  string    `json:"seasonName"`
	SeasonStart time.Time `json:"seasonStart"`
	SeasonEnd   time.Time `json:"seasonEnd"`
	IsCurrent   int64     `json:"isCurrent"`
}

type SummonObjectStatElement struct {
	Code                 int64   `json:"code"`
	Name                 string  `json:"name"`
	Duration             float64 `json:"duration"`
	CreateRange          float64 `json:"createRange"`
	PileRange            float64 `json:"pileRange"`
	CreateVisibleTime    float64 `json:"createVisibleTime"`
	CreateStealthTime    float64 `json:"createStealthTime"`
	InfiltrationTime     int64   `json:"infiltrationTime"`
	DetectionRange       float64 `json:"detectionRange"`
	MaxHP                int64   `json:"maxHp"`
	MaxSP                int64   `json:"maxSp"`
	InitExtraPoint       int64   `json:"initExtraPoint"`
	MaxExtraPoint        int64   `json:"maxExtraPoint"`
	AttackPower          int64   `json:"attackPower"`
	Defense              int64   `json:"defense"`
	CriticalStrikeChance int64   `json:"criticalStrikeChance"`
	HPRegen              int64   `json:"hpRegen"`
	SPRegen              int64   `json:"spRegen"`
	RangeRadius          float64 `json:"rangeRadius"`
	AttackSpeed          float64 `json:"attackSpeed"`
	AttackRange          float64 `json:"attackRange"`
	AttackDelay          float64 `json:"attackDelay"`
	MoveSpeed            int64   `json:"moveSpeed"`
	Radius               float64 `json:"radius"`
	UIHeight             float64 `json:"uiHeight"`
	SightRange           float64 `json:"sightRange"`
	SightAngle           int64   `json:"sightAngle"`
}

type TacticalSkillSetElement struct {
	Code            int64 `json:"code"`
	NextUpgradecode int64 `json:"nextUpgradecode"`
	UpgradeCredit   int64 `json:"upgradeCredit"`
	UpgradeMaterial int64 `json:"upgradeMaterial"`
	SkillCode       int64 `json:"skillCode"`
}

type TacticalSkillSetGroupElement struct {
	Group          int64  `json:"group"`
	ModeType       int64  `json:"modeType"`
	StartCode      int64  `json:"startCode"`
	EquipWithStart bool   `json:"equipWithStart"`
	Icon           string `json:"icon"`
}

type TraitElement struct {
	Code           int64  `json:"code"`
	OpenAccountLV  int64  `json:"openAccountLv"`
	TraitGroup     string `json:"traitGroup"`
	TraitType      string `json:"traitType"`
	Active         bool   `json:"active"`
	TraitSortOrder int64  `json:"traitSortOrder"`
}

type TransferConsoleElement struct {
	ItemCode                          int64  `json:"itemCode"`
	Mode                              int64  `json:"mode"`
	ItemType                          string `json:"itemType"`
	TransferTimeSafeArea              int64  `json:"transferTimeSafeArea"`
	SubtractTrasferTimeRestrictedArea int64  `json:"subtractTrasferTimeRestrictedArea"`
	ManufactureCooldown               int64  `json:"manufactureCooldown"`
	AvailableTimeForPurchase          int64  `json:"availableTimeForPurchase"`
	ConsumeVFCredit                   int64  `json:"consumeVFCredit"`
	LimitCount                        int64  `json:"limitCount"`
	TraitSale                         bool   `json:"traitSale"`
	KioskType                         string `json:"kioskType"`
}

type VFCreditElement struct {
	Code           int64   `json:"code"`
	Mode           int64   `json:"mode"`
	Phase          int64   `json:"phase"`
	ConditionType  string  `json:"conditionType"`
	ConditionValue int64   `json:"conditionValue"`
	AcquireSelf    float64 `json:"acquireSelf"`
	AcquireTeam    int64   `json:"acquireTeam"`
}

type WeaponTypeInfoElement struct {
	Type                  string  `json:"type"`
	AttackSpeed           float64 `json:"attackSpeed"`
	AttackRange           float64 `json:"attackRange"`
	ShopFilter            int64   `json:"shopFilter"`
	SummonObjectHitDamage int64   `json:"summonObjectHitDamage"`
}
