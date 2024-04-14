package models

//
// Game Data API structs
//

type ActionCost struct {
	Code                  int     `json:"code"`
	Type                  string  `json:"type"`
	Sp                    int     `json:"sp"`
	Time1                 float64 `json:"time1"`
	Time2                 float64 `json:"time2"`
	ActionWaitTime        float64 `json:"actionWaitTime"`
	CastingAnimTrigger    string  `json:"castingAnimTrigger"`
	EffectCancelCondition string  `json:"effectCancelCondition"`
	CastingBarImgName     string  `json:"castingBarImgName"`
}

type Area struct {
	Code                     int    `json:"code"`
	Name                     string `json:"name"`
	ModeType                 int    `json:"modeType"`
	MaskCode                 int    `json:"maskCode"`
	StartingArea             bool   `json:"startingArea"`
	AreaType                 string `json:"areaType"`
	IsProvideCollectibleItem bool   `json:"isProvideCollectibleItem"`
	RouteCalcBitCode         int    `json:"routeCalcBitCode"`
	IsHyperLoopInstalled     bool   `json:"isHyperLoopInstalled"`
}

type BattleZoneReward struct {
	Code                           int    `json:"code"`
	ModeType                       int    `json:"modeType"`
	AreaAttributesCreateEventCount int    `json:"areaAttributesCreateEventCount"`
	ItemCode                       int    `json:"itemCode"`
	Type                           string `json:"type"`
	Value                          int    `json:"value"`
	Selectable                     bool   `json:"selectable"`
}

type BulletCapacity struct {
	ItemCode  int     `json:"itemCode"`
	Capacity  int     `json:"capacity"`
	LoadType  string  `json:"loadType"`
	Time      float64 `json:"time"`
	InitCount int     `json:"initCount"`
	Count     int     `json:"count"`
}

type Character struct {
	Code                           int     `json:"code"`
	MaxHp                          int     `json:"maxHp"`
	MaxSp                          int     `json:"maxSp"`
	InitExtraPoint                 int     `json:"initExtraPoint"`
	MaxExtraPoint                  int     `json:"maxExtraPoint"`
	AttackPower                    int     `json:"attackPower"`
	Defense                        int     `json:"defense"`
	SkillAmp                       int     `json:"skillAmp"`
	AdaptiveForce                  int     `json:"adaptiveForce"`
	CriticalStrikeChance           float64 `json:"criticalStrikeChance"`
	HpRegen                        float64 `json:"hpRegen"`
	SpRegen                        float64 `json:"spRegen"`
	AttackSpeed                    float64 `json:"attackSpeed"`
	AttackSpeedRatio               float64 `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio float64 `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  float64 `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio float64 `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       float64 `json:"preventSkillDamagedRatio"`
	AttackSpeedLimit               float64 `json:"attackSpeedLimit"`
	AttackSpeedMin                 float64 `json:"attackSpeedMin"`
	MoveSpeed                      float64 `json:"moveSpeed"`
	SightRange                     float64 `json:"sightRange"`
	Radius                         float64 `json:"radius"`
	PathingRadius                  float64 `json:"pathingRadius"`
	UIHeight                       float64 `json:"uiHeight"`
	InitStateDisplayIndex          int     `json:"initStateDisplayIndex"`
	LocalScaleInCutscene           int     `json:"localScaleInCutscene"`
	LocalScaleInVictoryScene       string  `json:"localScaleInVictoryScene"`
	Resource                       string  `json:"resource"`
	LobbySubObject                 string  `json:"lobbySubObject"`
	Name                           string  `json:"name"`
	StrLearnStartSkill             string  `json:"strLearnStartSkill"`
	StrUsePointLearnStartSkill     string  `json:"strUsePointLearnStartSkill"`
}

type CharacterAttributes struct {
	Character         string `json:"character"`
	CharacterCode     int    `json:"characterCode"`
	Mastery           string `json:"mastery"`
	ControlDifficulty int    `json:"controlDifficulty"`
	Attack            int    `json:"attack"`
	Defense           int    `json:"defense"`
	Disruptor         int    `json:"disruptor"`
	Move              int    `json:"move"`
	Assistance        int    `json:"assistance"`
}

type CharacterExp struct {
	Level      int `json:"level"`
	LevelUpExp int `json:"levelUpExp"`
}

type CharacterLevelUpStat struct {
	Code                           int     `json:"code"`
	MaxHp                          int     `json:"maxHp"`
	MaxSp                          int     `json:"maxSp"`
	SkillAmp                       int     `json:"skillAmp"`
	AdaptiveForce                  int     `json:"adaptiveForce"`
	AttackPower                    float64 `json:"attackPower"`
	Defense                        float64 `json:"defense"`
	CriticalChance                 float64 `json:"criticalChance"`
	HpRegen                        float64 `json:"hpRegen"`
	SpRegen                        float64 `json:"spRegen"`
	AttackSpeed                    float64 `json:"attackSpeed"`
	MoveSpeed                      float64 `json:"moveSpeed"`
	AttackSpeedRatio               float64 `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio float64 `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  float64 `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio float64 `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       float64 `json:"preventSkillDamagedRatio"`
	Name                           string  `json:"name"`
}

type ItemArmor struct {
	Code                                               int     `json:"code"`
	MakeMaterial1                                      int     `json:"makeMaterial1"`
	MakeMaterial2                                      int     `json:"makeMaterial2"`
	CreditValueWhenConvertedToBounty                   int     `json:"creditValueWhenConvertedToBounty"`
	ManufacturableType                                 int     `json:"manufacturableType"`
	AttackPower                                        int     `json:"attackPower"`
	AttackPowerByLv                                    int     `json:"attackPowerByLv"`
	Defense                                            int     `json:"defense"`
	DefenseByLv                                        int     `json:"defenseByLv"`
	SkillAmp                                           int     `json:"skillAmp"`
	SkillAmpByLevel                                    int     `json:"skillAmpByLevel"`
	AdaptiveForce                                      int     `json:"adaptiveForce"`
	AdaptiveForceByLevel                               int     `json:"adaptiveForceByLevel"`
	MaxHp                                              int     `json:"maxHp"`
	MaxHpByLv                                          int     `json:"maxHpByLv"`
	MaxSp                                              int     `json:"maxSp"`
	PreventCriticalStrikeDamaged                       int     `json:"preventCriticalStrikeDamaged"`
	AttackRange                                        int     `json:"attackRange"`
	IncreaseBasicAttackDamage                          int     `json:"increaseBasicAttackDamage"`
	IncreaseBasicAttackDamageByLv                      int     `json:"increaseBasicAttackDamageByLv"`
	PreventBasicAttackDamaged                          int     `json:"preventBasicAttackDamaged"`
	PreventBasicAttackDamagedByLv                      int     `json:"preventBasicAttackDamagedByLv"`
	IncreaseBasicAttackDamageRatio                     int     `json:"increaseBasicAttackDamageRatio"`
	PreventSkillDamaged                                int     `json:"preventSkillDamaged"`
	PreventSkillDamagedByLv                            int     `json:"preventSkillDamagedByLv"`
	PenetrationDefense                                 int     `json:"penetrationDefense"`
	TrapDamageReduce                                   int     `json:"trapDamageReduce"`
	UniqueAttackRange                                  int     `json:"uniqueAttackRange"`
	UniquePenetrationDefense                           int     `json:"uniquePenetrationDefense"`
	ItemUsableValueList                                int     `json:"itemUsableValueList"`
	ExclusiveProducer                                  int     `json:"exclusiveProducer"`
	ModeType                                           int     `json:"modeType"`
	Stackable                                          int     `json:"stackable"`
	InitialCount                                       int     `json:"initialCount"`
	SightRange                                         int     `json:"sightRange"`
	HpRegenRatio                                       float64 `json:"hpRegenRatio"`
	HpRegen                                            float64 `json:"hpRegen"`
	SpRegenRatio                                       float64 `json:"spRegenRatio"`
	SpRegen                                            float64 `json:"spRegen"`
	AttackSpeedRatio                                   float64 `json:"attackSpeedRatio"`
	AttackSpeedRatioByLv                               float64 `json:"attackSpeedRatioByLv"`
	CriticalStrikeChance                               float64 `json:"criticalStrikeChance"`
	CriticalStrikeDamage                               float64 `json:"criticalStrikeDamage"`
	SkillAmpRatio                                      float64 `json:"skillAmpRatio"`
	SkillAmpRatioByLevel                               float64 `json:"skillAmpRatioByLevel"`
	CooldownReduction                                  float64 `json:"cooldownReduction"`
	CooldownLimit                                      float64 `json:"cooldownLimit"`
	LifeSteal                                          float64 `json:"lifeSteal"`
	NormalLifeSteal                                    float64 `json:"normalLifeSteal"`
	SkillLifeSteal                                     float64 `json:"skillLifeSteal"`
	MoveSpeed                                          float64 `json:"moveSpeed"`
	MoveSpeedOutOfCombat                               float64 `json:"moveSpeedOutOfCombat"`
	UniqueHpHealedIncreaseRatio                        float64 `json:"uniqueHpHealedIncreaseRatio"`
	UniqueCooldownLimit                                float64 `json:"uniqueCooldownLimit"`
	UniqueTenacity                                     float64 `json:"uniqueTenacity"`
	UniqueMoveSpeed                                    float64 `json:"uniqueMoveSpeed"`
	PreventBasicAttackDamagedRatio                     float64 `json:"preventBasicAttackDamagedRatio"`
	PreventBasicAttackDamagedRatioByLv                 float64 `json:"preventBasicAttackDamagedRatioByLv"`
	PreventSkillDamagedRatio                           float64 `json:"preventSkillDamagedRatio"`
	PreventSkillDamagedRatioByLv                       float64 `json:"preventSkillDamagedRatioByLv"`
	PenetrationDefenseRatio                            float64 `json:"penetrationDefenseRatio"`
	TrapDamageReduceRatio                              float64 `json:"trapDamageReduceRatio"`
	HpHealedIncreaseRatio                              float64 `json:"hpHealedIncreaseRatio"`
	HealerGiveHpHealRatio                              float64 `json:"healerGiveHpHealRatio"`
	UniquePenetrationDefenseRatio                      float64 `json:"uniquePenetrationDefenseRatio"`
	UniqueLifeSteal                                    float64 `json:"uniqueLifeSteal"`
	UniqueSkillAmpRatio                                float64 `json:"uniqueSkillAmpRatio"`
	IncreaseBasicAttackDamageRatioByLv                 float64 `json:"increaseBasicAttackDamageRatioByLv"`
	IsCompletedItem                                    bool    `json:"isCompletedItem"`
	AlertInSpectator                                   bool    `json:"alertInSpectator"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	NotDisarm                                          bool    `json:"notDisarm"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected"`
	Name                                               string  `json:"name"`
	ItemType                                           string  `json:"itemType"`
	ArmorType                                          string  `json:"armorType"`
	ItemGrade                                          string  `json:"itemGrade"`
	MarkingType                                        string  `json:"markingType"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger"`
	ItemUsableType                                     string  `json:"itemUsableType"`
	MakeCustomAction                                   string  `json:"makeCustomAction"`
}

type CharacterMastery struct {
	Code      int    `json:"code"`
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

type CharacterModeModifier struct {
	CharacterCode                           int    `json:"characterCode"`
	WeaponType                              string `json:"weaponType"`
	SoloIncreaseModeDamageRatio             int    `json:"soloIncreaseModeDamageRatio"`
	SoloPreventModeDamageRatio              int    `json:"soloPreventModeDamageRatio"`
	SoloIncreaseModeHealRatio               int    `json:"soloIncreaseModeHealRatio"`
	SoloIncreaseModeShieldRatio             int    `json:"soloIncreaseModeShieldRatio"`
	DuoIncreaseModeDamageRatio              int    `json:"duoIncreaseModeDamageRatio"`
	DuoPreventModeDamageRatio               int    `json:"duoPreventModeDamageRatio"`
	DuoIncreaseModeHealRatio                int    `json:"duoIncreaseModeHealRatio"`
	DuoIncreaseModeHealerGiveHealRatio      int    `json:"duoIncreaseModeHealerGiveHealRatio"`
	DuoIncreaseModeShieldRatio              int    `json:"duoIncreaseModeShieldRatio"`
	DuoIncreaseModeHealerGiveShieldRatio    int    `json:"duoIncreaseModeHealerGiveShieldRatio"`
	SquadIncreaseModeDamageRatio            int    `json:"squadIncreaseModeDamageRatio"`
	SquadPreventModeDamageRatio             int    `json:"squadPreventModeDamageRatio"`
	SquadIncreaseModeHealRatio              int    `json:"squadIncreaseModeHealRatio"`
	SquadIncreaseModeHealerGiveHealRatio    int    `json:"squadIncreaseModeHealerGiveHealRatio"`
	SquadIncreaseModeShieldRatio            int    `json:"squadIncreaseModeShieldRatio"`
	SquadIncreaseModeHealerGiveShieldRatio  int    `json:"squadIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeDamageRatio           int    `json:"cobaltIncreaseModeDamageRatio"`
	CobaltPreventModeDamageRatio            int    `json:"cobaltPreventModeDamageRatio"`
	CobaltIncreaseModeHealRatio             int    `json:"cobaltIncreaseModeHealRatio"`
	CobaltIncreaseModeHealerGiveHealRatio   int    `json:"cobaltIncreaseModeHealerGiveHealRatio"`
	CobaltIncreaseModeShieldRatio           int    `json:"cobaltIncreaseModeShieldRatio"`
	CobaltIncreaseModeHealerGiveShieldRatio int    `json:"cobaltIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeUltCooldownRatio      int    `json:"cobaltIncreaseModeUltCooldownRatio"`
	CobaltIncreaseModeMaxSpRatio            int    `json:"cobaltIncreaseModeMaxSpRatio"`
	CobaltIncreaseModeSpRegenRatio          int    `json:"cobaltIncreaseModeSpRegenRatio"`
	SoloIncreaseModeDamageToMonsterRatio    int    `json:"soloIncreaseModeDamageToMonsterRatio"`
	DuoIncreaseModeDamageToMonsterRatio     int    `json:"duoIncreaseModeDamageToMonsterRatio"`
	SquadIncreaseModeDamageToMonsterRatio   int    `json:"squadIncreaseModeDamageToMonsterRatio"`
	CobaltIncreaseModeDamageToMonsterRatio  int    `json:"cobaltIncreaseModeDamageToMonsterRatio"`
}

type CharacterSkin struct {
	Name                     string `json:"name"`
	Code                     int    `json:"code"`
	CharacterCode            int    `json:"characterCode"`
	Index                    int    `json:"index"`
	Grade                    int    `json:"grade"`
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

type Collectible struct {
	Code              int    `json:"code"`
	Cooldown          int    `json:"cooldown"`
	ItemCode1         string `json:"itemCode1"`
	ItemCode2         string `json:"itemCode2"`
	Probability1      int    `json:"probability1"`
	Probability2      int    `json:"probability2"`
	DropCount         int    `json:"dropCount"`
	CastingActionType string `json:"castingActionType"`
}

type DropGroup struct {
	GroupCode   int    `json:"groupCode"`
	ItemCode    string `json:"itemCode"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Probability int    `json:"probability"`
	DropType    string `json:"dropType"`
}

type GainExp struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	GainExp   int `json:"gainExp"`
}

type GainScore struct {
	Code           int    `json:"code"`
	Phase          int    `json:"phase"`
	ConditionType  string `json:"conditionType"`
	ConditionValue int    `json:"conditionValue"`
	PointsEnemy    int    `json:"pointsEnemy"`
	PointsAlly     int    `json:"pointsAlly"`
}

type GameTip struct {
	Key             int    `json:"key"`
	Code            int    `json:"code"`
	GameTipType     string `json:"gameTipType"`
	GameTipCategory string `json:"gameTipCategory"`
	Sequence        int    `json:"sequence"`
	TitleTextKey    string `json:"titleTextKey"`
	ContentTextKey  string `json:"contentTextKey"`
	ImageName       string `json:"imageName"`
	KoreaWord       string `json:"##참고"` // idk what this even is
	Link            string `json:"link"`
}

type HowToFindItem struct {
	Code            int `json:"code"`
	ItemCode        int `json:"itemCode"`
	HuntChicken     int `json:"huntChicken"`
	HuntBat         int `json:"huntBat"`
	HuntBoar        int `json:"huntBoar"`
	HuntWildDog     int `json:"huntWildDog"`
	HuntWolf        int `json:"huntWolf"`
	HuntBear        int `json:"huntBear"`
	HuntWickline    int `json:"huntWickline"`
	HuntAlpha       int `json:"huntAlpha"`
	HuntOmega       int `json:"huntOmega"`
	CollectibleCode int `json:"collectibleCode"`
	AirSupply       int `json:"airSupply"`
}

type InfusionProduct struct {
	Code             int         `json:"code"`
	ProductType      string      `json:"productType"`
	ProductGroup     int         `json:"productGroup"`
	ProductCode      int         `json:"productCode"`
	StoreType        string      `json:"storeType"`
	StockType        string      `json:"stockType"`
	Stock            int         `json:"stock"`
	IsRestore        bool        `json:"isRestore"`
	Price            int         `json:"price"`
	SpecialWeight    int         `json:"specialWeight"`
	Weight           int         `json:"weight"`
	Requirement      int         `json:"requirement"`
	Icon             string      `json:"icon"`
	SimpleIcon       string      `json:"simpleIcon"`
	AlertInSpectator bool        `json:"alertInSpectator"`
	CharacterCodes   interface{} `json:"characterCodes"`
}

type ItemConsumable struct {
	Code                                               int    `json:"code"`
	Name                                               string `json:"name"`
	ModeType                                           int    `json:"modeType"`
	ItemType                                           string `json:"itemType"`
	ConsumableType                                     string `json:"consumableType"`
	ConsumableTag                                      string `json:"consumableTag"`
	ItemGrade                                          string `json:"itemGrade"`
	IsCompletedItem                                    bool   `json:"isCompletedItem"`
	AlertInSpectator                                   bool   `json:"alertInSpectator"`
	MarkingType                                        string `json:"markingType"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger"`
	Stackable                                          int    `json:"stackable"`
	InitialCount                                       int    `json:"initialCount"`
	ItemUsableType                                     string `json:"itemUsableType"`
	ItemUsableValueList                                int    `json:"itemUsableValueList"`
	ExclusiveProducer                                  int    `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int    `json:"manufacturableType"`
	MakeMaterial1                                      int    `json:"makeMaterial1"`
	MakeMaterial2                                      int    `json:"makeMaterial2"`
	Heal                                               int    `json:"heal"`
	HpRecover                                          int    `json:"hpRecover"`
	SpRecover                                          int    `json:"spRecover"`
	AttackPowerByBuff                                  int    `json:"attackPowerByBuff"`
	DefenseByBuff                                      int    `json:"defenseByBuff"`
	SkillAmpByBuff                                     int    `json:"skillAmpByBuff"`
	SkillAmpRatioByBuff                                int    `json:"skillAmpRatioByBuff"`
	AddStateCode                                       int    `json:"addStateCode"`
	IsVPadQuickSlotItem                                bool   `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int    `json:"creditValueWhenConvertedToBounty"`
	IsReduceLootOnDeath                                bool   `json:"isReduceLootOnDeath"`
}

type ItemMisc struct {
	Code                                               int    `json:"code"`
	Name                                               string `json:"name"`
	ModeType                                           int    `json:"modeType"`
	ItemType                                           string `json:"itemType"`
	MiscItemType                                       string `json:"miscItemType"`
	ItemGrade                                          string `json:"itemGrade"`
	GradeBgOverride                                    string `json:"gradeBgOverride"`
	IsCompletedItem                                    bool   `json:"isCompletedItem"`
	AlertInSpectator                                   bool   `json:"alertInSpectator"`
	MarkingType                                        string `json:"markingType"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger"`
	Stackable                                          int    `json:"stackable"`
	InitialCount                                       int    `json:"initialCount"`
	ItemUsableType                                     string `json:"itemUsableType"`
	ItemUsableValueList                                int    `json:"itemUsableValueList"`
	ExclusiveProducer                                  int    `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int    `json:"manufacturableType"`
	MakeMaterial1                                      int    `json:"makeMaterial1"`
	MakeMaterial2                                      int    `json:"makeMaterial2"`
	MakeCustomAction                                   string `json:"makeCustomAction"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int    `json:"creditValueWhenConvertedToBounty"`
}

type ItemSearchOption struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Tag1 string `json:"tag1"`
	Tag2 string `json:"tag2"`
	Tag3 string `json:"tag3"`
}

type ItemSpawn struct {
	Code           int    `json:"code"`
	AreaCode       int    `json:"areaCode"`
	AreaSpawnGroup int    `json:"areaSpawnGroup"`
	ItemCode       int    `json:"itemCode"`
	DropPoint      string `json:"dropPoint"`
	DropCount      int    `json:"dropCount"`
}

type ItemSpecial struct {
	Code                                               int         `json:"code"`
	Name                                               string      `json:"name"`
	ModeType                                           int         `json:"modeType"`
	ItemType                                           string      `json:"itemType"`
	SpecialItemType                                    string      `json:"specialItemType"`
	ItemGrade                                          string      `json:"itemGrade"`
	IsCompletedItem                                    bool        `json:"isCompletedItem"`
	AlertInSpectator                                   bool        `json:"alertInSpectator"`
	MarkingType                                        string      `json:"markingType"`
	CraftAnimTrigger                                   string      `json:"craftAnimTrigger"`
	Stackable                                          int         `json:"stackable"`
	InitialCount                                       int         `json:"initialCount"`
	CooldownGroupCode                                  int         `json:"cooldownGroupCode"`
	Cooldown                                           float64     `json:"cooldown"`
	ItemUsableType                                     string      `json:"itemUsableType"`
	ItemUsableValueList                                int         `json:"itemUsableValueList"`
	ExclusiveProducer                                  interface{} `json:"exclusiveProducer"` // number or comma separated numbers in a string
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool        `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	ManufacturableType                                 int         `json:"manufacturableType"`
	MakeMaterial1                                      int         `json:"makeMaterial1"`
	MakeMaterial2                                      int         `json:"makeMaterial2"`
	MakeCustomAction                                   string      `json:"makeCustomAction"`
	ConsumeCount                                       int         `json:"consumeCount"`
	SummonCode                                         int         `json:"summonCode"`
	GhostItemStateGroup                                int         `json:"ghostItemStateGroup"`
	IsVPadQuickSlotItem                                bool        `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected                         bool        `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int         `json:"creditValueWhenConvertedToBounty"`
	IsReduceLootOnDeath                                bool        `json:"isReduceLootOnDeath"`
}

type ItemWeapon struct {
	Code                                               int     `json:"code"`
	Name                                               string  `json:"name"`
	ModeType                                           int     `json:"modeType"`
	ItemType                                           string  `json:"itemType"`
	WeaponType                                         string  `json:"weaponType"`
	ItemGrade                                          string  `json:"itemGrade"`
	GradeBgOverride                                    string  `json:"gradeBgOverride"`
	IsCompletedItem                                    bool    `json:"isCompletedItem"`
	AlertInSpectator                                   bool    `json:"alertInSpectator"`
	MarkingType                                        string  `json:"markingType"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger"`
	Stackable                                          int     `json:"stackable"`
	InitialCount                                       int     `json:"initialCount"`
	ItemUsableType                                     string  `json:"itemUsableType"`
	ItemUsableValueList                                int     `json:"itemUsableValueList"`
	ExclusiveProducer                                  int     `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	MakeMaterial1                                      int     `json:"makeMaterial1"`
	MakeMaterial2                                      int     `json:"makeMaterial2"`
	MakeCustomAction                                   string  `json:"makeCustomAction"`
	NotDisarm                                          bool    `json:"notDisarm"`
	Consumable                                         bool    `json:"consumable"`
	ManufacturableType                                 int     `json:"manufacturableType"`
	AttackPower                                        int     `json:"attackPower"`
	AttackPowerByLv                                    int     `json:"attackPowerByLv"`
	Defense                                            int     `json:"defense"`
	DefenseByLv                                        int     `json:"defenseByLv"`
	SkillAmp                                           int     `json:"skillAmp"`
	SkillAmpByLevel                                    int     `json:"skillAmpByLevel"`
	SkillAmpRatio                                      int     `json:"skillAmpRatio"`
	SkillAmpRatioByLevel                               int     `json:"skillAmpRatioByLevel"`
	AdaptiveForce                                      int     `json:"adaptiveForce"`
	AdaptiveForceByLevel                               int     `json:"adaptiveForceByLevel"`
	MaxHp                                              int     `json:"maxHp"`
	MaxHpByLv                                          int     `json:"maxHpByLv"`
	HpRegenRatio                                       float64 `json:"hpRegenRatio"`
	HpRegen                                            int     `json:"hpRegen"`
	MaxSP                                              int     `json:"maxSP"`
	SpRegenRatio                                       float64 `json:"spRegenRatio"`
	SpRegen                                            int     `json:"spRegen"`
	AttackSpeedRatio                                   float64 `json:"attackSpeedRatio"`
	AttackSpeedRatioByLv                               int     `json:"attackSpeedRatioByLv"`
	CriticalStrikeChance                               float64 `json:"criticalStrikeChance"`
	CriticalStrikeDamage                               float64 `json:"criticalStrikeDamage"`
	CooldownReduction                                  float64 `json:"cooldownReduction"`
	PreventCriticalStrikeDamaged                       int     `json:"preventCriticalStrikeDamaged"`
	CooldownLimit                                      int     `json:"cooldownLimit"`
	LifeSteal                                          float64 `json:"lifeSteal"`
	NormalLifeSteal                                    float64 `json:"normalLifeSteal"`
	SkillLifeSteal                                     int     `json:"skillLifeSteal"`
	MoveSpeed                                          float64 `json:"moveSpeed"`
	MoveSpeedOutOfCombat                               int     `json:"moveSpeedOutOfCombat"`
	SightRange                                         float64 `json:"sightRange"`
	AttackRange                                        int     `json:"attackRange"`
	IncreaseBasicAttackDamage                          int     `json:"increaseBasicAttackDamage"`
	IncreaseBasicAttackDamageByLv                      int     `json:"increaseBasicAttackDamageByLv"`
	IncreaseBasicAttackDamageRatio                     int     `json:"increaseBasicAttackDamageRatio"`
	IncreaseBasicAttackDamageRatioByLv                 float64 `json:"increaseBasicAttackDamageRatioByLv"`
	PreventBasicAttackDamaged                          int     `json:"preventBasicAttackDamaged"`
	PreventBasicAttackDamagedByLv                      int     `json:"preventBasicAttackDamagedByLv"`
	PreventBasicAttackDamagedRatio                     int     `json:"preventBasicAttackDamagedRatio"`
	PreventBasicAttackDamagedRatioByLv                 int     `json:"preventBasicAttackDamagedRatioByLv"`
	PreventSkillDamaged                                int     `json:"preventSkillDamaged"`
	PreventSkillDamagedByLv                            int     `json:"preventSkillDamagedByLv"`
	PreventSkillDamagedRatio                           int     `json:"preventSkillDamagedRatio"`
	PreventSkillDamagedRatioByLv                       int     `json:"preventSkillDamagedRatioByLv"`
	PenetrationDefense                                 int     `json:"penetrationDefense"`
	PenetrationDefenseRatio                            float64 `json:"penetrationDefenseRatio"`
	TrapDamageReduce                                   int     `json:"trapDamageReduce"`
	TrapDamageReduceRatio                              int     `json:"trapDamageReduceRatio"`
	HpHealedIncreaseRatio                              int     `json:"hpHealedIncreaseRatio"`
	HealerGiveHpHealRatio                              float64 `json:"healerGiveHpHealRatio"`
	UniqueAttackRange                                  float64 `json:"uniqueAttackRange"`
	UniqueHpHealedIncreaseRatio                        int     `json:"uniqueHpHealedIncreaseRatio"`
	UniqueCooldownLimit                                float64 `json:"uniqueCooldownLimit"`
	UniqueTenacity                                     float64 `json:"uniqueTenacity"`
	UniqueMoveSpeed                                    int     `json:"uniqueMoveSpeed"`
	UniquePenetrationDefense                           int     `json:"uniquePenetrationDefense"`
	UniquePenetrationDefenseRatio                      int     `json:"uniquePenetrationDefenseRatio"`
	UniqueLifeSteal                                    int     `json:"uniqueLifeSteal"`
	UniqueSkillAmpRatio                                int     `json:"uniqueSkillAmpRatio"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int     `json:"creditValueWhenConvertedToBounty"`
}

type Level struct {
	Level         int `json:"level"`
	NeedExp       int `json:"needExp"`
	AccumulateExp int `json:"accumulateExp"`
	RewardAcoin   int `json:"rewardAcoin"`
	Reward        int `json:"reward"`
}

type LoadingTip struct {
	Code           int    `json:"code"`
	LoadingTipType string `json:"loadingTipType"`
	MinLv          int    `json:"minLv"`
	MaxLv          int    `json:"maxLv"`
	TextKey        string `json:"textKey"`
	ImageName      string `json:"imageName"`
}

type MasteryExp struct {
	Code           int    `json:"code"`
	ModeType       int    `json:"modeType"`
	ConditionType  string `json:"conditionType"`
	Grade          string `json:"grade"`
	ConditionValue int    `json:"conditionValue"`
	MasteryType1   string `json:"masteryType1"`
	Value1         int    `json:"value1"`
	MasteryType2   string `json:"masteryType2"`
	Value2         int    `json:"value2"`
	MasteryType3   string `json:"masteryType3"`
	Value3         int    `json:"value3"`
}

type MasteryLevel struct {
	Code              int    `json:"code"`
	Type              string `json:"type"`
	MasteryLevel      int    `json:"masteryLevel"`
	NextMasteryExp    int    `json:"nextMasteryExp"`
	GiveLevelExp      int    `json:"giveLevelExp"`
	ExpGrowthCapRatio int    `json:"expGrowthCapRatio"`
}

type MasteryStat struct {
	Code                      int     `json:"code"`
	Type                      string  `json:"type"`
	CharacterCode             int     `json:"characterCode"`
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
	ThirdOptionSection1Value  int     `json:"thirdOptionSection1Value"`
	ThirdOptionSection2Value  int     `json:"thirdOptionSection2Value"`
	ThirdOptionSection3Value  int     `json:"thirdOptionSection3Value"`
	ThirdOptionSection4Value  int     `json:"thirdOptionSection4Value"`
}

type Monster struct {
	Code             int     `json:"Code"`
	Monster          string  `json:"monster"`
	IsMutant         bool    `json:"isMutant"`
	Grade            string  `json:"grade"`
	Mode             int     `json:"mode"`
	CreateDay        string  `json:"createDay"`
	CreateTime       int     `json:"createTime"`
	RegenTime        int     `json:"regenTime"`
	LevelUpPeriod    int     `json:"levelUpPeriod"`
	LevelUpAmount    int     `json:"levelUpAmount"`
	LevelUpMax       int     `json:"levelUpMax"`
	MaxHp            int     `json:"maxHp"`
	MaxEp            int     `json:"maxEp"`
	InitExtraPoint   int     `json:"initExtraPoint"`
	AttackPower      int     `json:"attackPower"`
	Defense          int     `json:"defense"`
	AttackSpeed      float64 `json:"attackSpeed"`
	MoveSpeed        float64 `json:"moveSpeed"`
	SightRange       int     `json:"sightRange"`
	ChasingRange     int     `json:"chasingRange"`
	AttackRange      float64 `json:"attackRange"`
	FirstAttackRange int     `json:"firstAttackRange"`
	Aggressive       string  `json:"aggressive"`
	DetectInvisible  bool    `json:"detectInvisible"`
	Radius           float64 `json:"radius"`
	PathingRadius    float64 `json:"pathingRadius"`
	UiHeight         float64 `json:"uiHeight"`
	GainExp          int     `json:"gainExp"`
	TargetOnRange    int     `json:"targetOnRange"`
	RandomDropCount  int     `json:"randomDropCount"`
	Resource         string  `json:"resource"`
	CorpseResource   string  `json:"corpseResource"`
	AppearTime       float64 `json:"appearTime"`
}

type MonsterDropGroup struct {
	MonsterCode  int `json:"monsterCode"`
	MonsterLevel int `json:"monsterLevel"`
	DropGroup    int `json:"dropGroup"`
}

type MonsterLevelUpStat struct {
	Code        int     `json:"code"`
	Monster     string  `json:"monster"`
	Mode        int     `json:"mode"`
	MaxHp       int     `json:"maxHp"`
	AttackPower int     `json:"attackPower"`
	Defense     float64 `json:"defense"`
	MoveSpeed   float64 `json:"moveSpeed"`
	GainExp     int     `json:"gainExp"`
}

type MonsterSpawnLevel struct {
	Code        int `json:"code"`
	Mode        int `json:"mode"`
	PlayerLevel int `json:"playerLevel"`
	MonsterCode int `json:"monsterCode"`
	SpawnLevel  int `json:"spawnLevel"`
}

type NaviCollectAndHunt struct {
	Code         int    `json:"code"`
	ItemCode     int    `json:"itemCode"`
	AreaCodeList string `json:"areaCodeList"`
}

type NearByArea struct {
	Code           int `json:"code"`
	AreaCode       int `json:"areaCode"`
	NearByAreaCode int `json:"nearByAreaCode"`
}

type RandomEquipment struct {
	Code          int         `json:"code"`
	Group         string      `json:"group"`
	ItemCode      interface{} `json:"itemcode"`
	Weight        int         `json:"weight"`
	ItemGrade     string      `json:"itemGrade"`
	TagMultiplier int         `json:"tagMultiplier"`
	CharacterNum  int         `json:"characterNum"`
}

type RecommendedList struct {
	Code                  int    `json:"code"`
	Character             string `json:"character"`
	CharacterCode         int    `json:"characterCode"`
	Mastery               string `json:"mastery"`
	StartWeapon           int    `json:"startWeapon"`
	CobaltStartWeapon     int    `json:"cobaltStartWeapon"`
	StartItemGroupCode    int    `json:"startItemGroupCode"`
	CobaltDraft           int    `json:"cobaltDraft"`
	CobaltExtraDraft      string `json:"cobaltExtraDraft"`
	CobaltCanChooseWeapon bool   `json:"cobaltCanChooseWeapon"`
	FavoriteMainTag       string `json:"favoriteMainTag"`
	OppositeTag           string `json:"oppositeTag"`
}

type Season struct {
	SeasonID    int    `json:"seasonID"`
	SeasonName  string `json:"seasonName"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
	IsCurrent   int    `json:"isCurrent"`
}

type SummonObjectStat struct {
	Code                 int     `json:"code"`
	Name                 string  `json:"name"`
	Duration             float64 `json:"duration"`
	CreateRange          float64 `json:"createRange"`
	PileRange            float64 `json:"pileRange"`
	CreateVisibleTime    float64 `json:"createVisibleTime"`
	CreateStealthTime    float64 `json:"createStealthTime"`
	InfiltrationTime     int     `json:"infiltrationTime"`
	DetectionRange       float64 `json:"detectionRange"`
	MaxHp                int     `json:"maxHp"`
	MaxSp                int     `json:"maxSp"`
	InitExtraPoint       int     `json:"initExtraPoint"`
	MaxExtraPoint        int     `json:"maxExtraPoint"`
	AttackPower          int     `json:"attackPower"`
	Defense              int     `json:"defense"`
	CriticalStrikeChance int     `json:"criticalStrikeChance"`
	HpRegen              int     `json:"hpRegen"`
	SpRegen              int     `json:"spRegen"`
	RangeRadius          float64 `json:"rangeRadius"`
	AttackSpeed          float64 `json:"attackSpeed"`
	AttackRange          float64 `json:"attackRange"`
	AttackDelay          float64 `json:"attackDelay"`
	MoveSpeed            int     `json:"moveSpeed"`
	Radius               float64 `json:"radius"`
	UiHeight             float64 `json:"uiHeight"`
	SightRange           float64 `json:"sightRange"`
	SightAngle           int     `json:"sightAngle"`
}

type TacticalSkillSet struct {
	Code            int `json:"code"`
	NextUpgradeCode int `json:"nextUpgradecode"`
	UpgradeCredit   int `json:"upgradeCredit"`
	UpgradeMaterial int `json:"upgradeMaterial"`
	SkillCode       int `json:"skillCode"`
}

type TacticalSkillSetGroup struct {
	Group          int    `json:"group"`
	ModeType       int    `json:"modeType"`
	StartCode      int    `json:"startCode"`
	EquipWithStart bool   `json:"equipWithStart"`
	Icon           string `json:"icon"`
}

type Trait struct {
	Code          int    `json:"code"`
	OpenAccountLv int    `json:"openAccountLv"`
	TraitGroup    string `json:"traitGroup"`
	TraitType     string `json:"traitType"`
	Active        bool   `json:"active"`
}

type TransferConsole struct {
	ItemCode                           int    `json:"itemCode"`
	Mode                               int    `json:"mode"`
	ItemType                           string `json:"itemType"`
	TransferTimeSafeArea               int    `json:"transferTimeSafeArea"`
	SubtractTransferTimeRestrictedArea int    `json:"subtractTrasferTimeRestrictedArea"`
	ManufactureCooldown                int    `json:"manufactureCooldown"`
	AvailableTimeForPurchase           int    `json:"availableTimeForPurchase"`
	ConsumeVFCredit                    int    `json:"consumeVFCredit"`
	LimitCount                         int    `json:"limitCount"`
	TraitSale                          bool   `json:"traitSale"`
}

type VfCredit struct {
	Code           int     `json:"code"`
	Mode           int     `json:"mode"`
	Phase          int     `json:"phase"`
	ConditionType  string  `json:"conditionType"`
	ConditionValue int     `json:"conditionValue"`
	AcquireSelf    float64 `json:"acquireSelf"`
	AcquireTeam    int     `json:"acquireTeam"`
}

type WeaponTypeInfo struct {
	Type                  string  `json:"type"`
	AttackSpeed           float64 `json:"attackSpeed"`
	AttackRange           float64 `json:"attackRange"`
	ShopFilter            int     `json:"shopFilter"`
	SummonObjectHitDamage int     `json:"summonObjectHitDamage"`
}
