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
	IncreaseBasicAttackDamageRatioByLv                 float64 `json:"increaseBasicAttackDamageRatioByLv" db:"increase_basic_attack_damage_ratio_by_lv"`
	IsCompletedItem                                    bool    `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool    `json:"alertInSpectator" db:"alert_in_spectator"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	NotDisarm                                          bool    `json:"notDisarm" db:"not_disarm"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	Name                                               string  `json:"name" db:"name"`
	ItemType                                           string  `json:"itemType" db:"item_type"`
	ArmorType                                          string  `json:"armorType" db:"armor_type"`
	ItemGrade                                          string  `json:"itemGrade" db:"item_grade"`
	MarkingType                                        string  `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	ItemUsableType                                     string  `json:"itemUsableType" db:"item_usable_type"`
	MakeCustomAction                                   string  `json:"makeCustomAction" db:"make_custom_action"`
}

type CharacterMastery struct {
	Code      int    `json:"code" db:"code"`
	Weapon1   string `json:"weapon1" db:"weapon_1"`
	Weapon2   string `json:"weapon2" db:"weapon_2"`
	Weapon3   string `json:"weapon3" db:"weapon_3"`
	Weapon4   string `json:"weapon4" db:"weapon_4"`
	Combat1   string `json:"combat1" db:"combat_1"`
	Combat2   string `json:"combat2" db:"combat_2"`
	Survival1 string `json:"survival1" db:"survival_1"`
	Survival2 string `json:"survival2" db:"survival_2"`
	Survival3 string `json:"survival3" db:"survival_3"`
}

type CharacterModeModifier struct {
	CharacterCode                           int    `json:"characterCode" db:"character_code"`
	WeaponType                              string `json:"weaponType" db:"weapon_type"`
	SoloIncreaseModeDamageRatio             int    `json:"soloIncreaseModeDamageRatio" db:"solo_increase_mode_damage_ratio"`
	SoloPreventModeDamageRatio              int    `json:"soloPreventModeDamageRatio" db:"solo_prevent_mode_damage_ratio"`
	SoloIncreaseModeHealRatio               int    `json:"soloIncreaseModeHealRatio" db:"solo_increase_mode_heal_ratio"`
	SoloIncreaseModeShieldRatio             int    `json:"soloIncreaseModeShieldRatio" db:"solo_increase_mode_shield_ratio"`
	DuoIncreaseModeDamageRatio              int    `json:"duoIncreaseModeDamageRatio" db:"duo_increase_mode_damage_ratio"`
	DuoPreventModeDamageRatio               int    `json:"duoPreventModeDamageRatio" db:"duo_prevent_mode_damage_ratio"`
	DuoIncreaseModeHealRatio                int    `json:"duoIncreaseModeHealRatio" db:"duo_increase_mode_heal_ratio"`
	DuoIncreaseModeHealerGiveHealRatio      int    `json:"duoIncreaseModeHealerGiveHealRatio" db:"duo_increase_mode_healer_give_heal_ratio"`
	DuoIncreaseModeShieldRatio              int    `json:"duoIncreaseModeShieldRatio" db:"duo_increase_mode_shield_ratio"`
	DuoIncreaseModeHealerGiveShieldRatio    int    `json:"duoIncreaseModeHealerGiveShieldRatio" db:"duo_increase_mode_healer_give_shield_ratio"`
	SquadIncreaseModeDamageRatio            int    `json:"squadIncreaseModeDamageRatio" db:"squad_increase_mode_damage_ratio"`
	SquadPreventModeDamageRatio             int    `json:"squadPreventModeDamageRatio" db:"squad_prevent_mode_damage_ratio"`
	SquadIncreaseModeHealRatio              int    `json:"squadIncreaseModeHealRatio" db:"squad_increase_mode_heal_ratio"`
	SquadIncreaseModeHealerGiveHealRatio    int    `json:"squadIncreaseModeHealerGiveHealRatio" db:"squad_increase_mode_healer_give_heal_ratio"`
	SquadIncreaseModeShieldRatio            int    `json:"squadIncreaseModeShieldRatio" db:"squad_increase_mode_shield_ratio"`
	SquadIncreaseModeHealerGiveShieldRatio  int    `json:"squadIncreaseModeHealerGiveShieldRatio" db:"squad_increase_mode_healer_give_shield_ratio"`
	CobaltIncreaseModeDamageRatio           int    `json:"cobaltIncreaseModeDamageRatio" db:"cobalt_increase_mode_damage_ratio"`
	CobaltPreventModeDamageRatio            int    `json:"cobaltPreventModeDamageRatio" db:"cobalt_prevent_mode_damage_ratio"`
	CobaltIncreaseModeHealRatio             int    `json:"cobaltIncreaseModeHealRatio" db:"cobalt_increase_mode_heal_ratio"`
	CobaltIncreaseModeHealerGiveHealRatio   int    `json:"cobaltIncreaseModeHealerGiveHealRatio" db:"cobalt_increase_mode_healer_give_shield_ratio"`
	CobaltIncreaseModeShieldRatio           int    `json:"cobaltIncreaseModeShieldRatio" db:"cobalt_increase_mode_shield_ratio"`
	CobaltIncreaseModeHealerGiveShieldRatio int    `json:"cobaltIncreaseModeHealerGiveShieldRatio" db:"cobalt_increase_mode_healer_give_shield_ratio"`
	CobaltIncreaseModeUltCooldownRatio      int    `json:"cobaltIncreaseModeUltCooldownRatio" db:"cobalt_increase_mode_ult_cooldown_ratio"`
	CobaltIncreaseModeMaxSpRatio            int    `json:"cobaltIncreaseModeMaxSpRatio" db:"cobalt_increase_mode-max_sp_ratio"`
	CobaltIncreaseModeSpRegenRatio          int    `json:"cobaltIncreaseModeSpRegenRatio" db:"cobalt_increase_mode_sp_regen_ratio"`
	SoloIncreaseModeDamageToMonsterRatio    int    `json:"soloIncreaseModeDamageToMonsterRatio" db:"solo_increase_mode_damage_to_monster_ratio"`
	DuoIncreaseModeDamageToMonsterRatio     int    `json:"duoIncreaseModeDamageToMonsterRatio" db:"duo_increase_mode_damage_to_monster_ratio"`
	SquadIncreaseModeDamageToMonsterRatio   int    `json:"squadIncreaseModeDamageToMonsterRatio" db:"squad_increase_mode_damage_to_monster_ratio"`
	CobaltIncreaseModeDamageToMonsterRatio  int    `json:"cobaltIncreaseModeDamageToMonsterRatio" db:"cobalt_increase_mode_damage_to_monster_ratio"`
}

type CharacterSkin struct {
	Name                     string `json:"name" db:"name"`
	Code                     int    `json:"code" db:"code"`
	CharacterCode            int    `json:"characterCode" db:"character_code"`
	Index                    int    `json:"index" db:"index"`
	Grade                    int    `json:"grade" db:"grade"`
	EventFree                bool   `json:"eventFree" db:"event_free"`
	PurchaseType             string `json:"purchaseType" db:"purchase_type"`
	EffectsPath              string `json:"effectsPath" db:"effects_path"`
	ProjectilesPath          string `json:"projectilesPath" db:"projectiles_path"`
	ObjectPath               string `json:"objectPath" db:"object_path"`
	FxSoundPath              string `json:"fxSoundPath" db:"fx_sound_path"`
	VoicePath                string `json:"voicePath" db:"voice_path"`
	WeaponMountPath          string `json:"weaponMountPath" db:"weapon_mount_path"`
	WeaponMountCommonPath    string `json:"weaponMountCommonPath" db:"weapon_mount_common_path"`
	IndicatorPath            string `json:"indicatorPath" db:"indicator_path"`
	ProjectilesDeflectorPath string `json:"projectilesDeflectorPath" db:"projectiles_deflector_path"`
}

type Collectible struct {
	Code              int    `json:"code" db:"code"`
	Cooldown          int    `json:"cooldown" db:"cooldown"`
	ItemCode1         string `json:"itemCode1" db:"item_code_1"`
	ItemCode2         string `json:"itemCode2" db:"item_code_2"`
	Probability1      int    `json:"probability1" db:"probability_1"`
	Probability2      int    `json:"probability2" db:"probability_2"`
	DropCount         int    `json:"dropCount" db:"drop_count"`
	CastingActionType string `json:"castingActionType" db:"casting_action_type"`
}

type DropGroup struct {
	GroupCode   int    `json:"groupCode" db:"group_code"`
	ItemCode    string `json:"itemCode" db:"item_code"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
	Probability int    `json:"probability" db:"probability"`
	DropType    string `json:"dropType" db:"drop_type"`
}

type GainExp struct {
	StartTime int `json:"startTime" db:"start_time"`
	EndTime   int `json:"endTime" db:"end_time"`
	GainExp   int `json:"gainExp" db:"gain_exp"`
}

type GainScore struct {
	Code           int    `json:"code" db:"code"`
	Phase          int    `json:"phase" db:"phase"`
	ConditionType  string `json:"conditionType" db:"condition_type"`
	ConditionValue int    `json:"conditionValue" db:"condition_value"`
	PointsEnemy    int    `json:"pointsEnemy" db:"points_enemy"`
	PointsAlly     int    `json:"pointsAlly" db:"points_ally"`
}

type GameTip struct {
	Key             int    `json:"key" db:"key"`
	Code            int    `json:"code" db:"code"`
	GameTipType     string `json:"gameTipType" db:"game_tip_type"`
	GameTipCategory string `json:"gameTipCategory" db:"game_tip_category"`
	Sequence        int    `json:"sequence" db:"sequence"`
	TitleTextKey    string `json:"titleTextKey" db:"title_text_key"`
	ContentTextKey  string `json:"contentTextKey" db:"context_text_key"`
	ImageName       string `json:"imageName" db:"image_name"`
	KoreaWord       string `json:"##참고" db:"korea_word"` // idk what this even is
	Link            string `json:"link" db:"link"`
}

type HowToFindItem struct {
	Code            int `json:"code" db:"code"`
	ItemCode        int `json:"itemCode" db:"item_code"`
	HuntChicken     int `json:"huntChicken" db:"hunt_chicken"`
	HuntBat         int `json:"huntBat" db:"hunt_bat"`
	HuntBoar        int `json:"huntBoar" db:"hunt_boar"`
	HuntWildDog     int `json:"huntWildDog" db:"hunt_dog"`
	HuntWolf        int `json:"huntWolf" db:"hunt_wolf"`
	HuntBear        int `json:"huntBear" db:"hunt_bear"`
	HuntWickeline   int `json:"huntWickline" db:"hunt_wickeline"`
	HuntAlpha       int `json:"huntAlpha" db:"hunt_alpha"`
	HuntOmega       int `json:"huntOmega" db:"hunt_omega"`
	CollectibleCode int `json:"collectibleCode" db:"collectible_code"`
	AirSupply       int `json:"airSupply" db:"air_supply"`
}

type InfusionProduct struct {
	Code             int         `json:"code" db:"code"`
	ProductType      string      `json:"productType" db:"product_type"`
	ProductGroup     int         `json:"productGroup" db:"product_group"`
	ProductCode      int         `json:"productCode" db:"product_code"`
	StoreType        string      `json:"storeType" db:"store_type"`
	StockType        string      `json:"stockType" db:"stock_type"`
	Stock            int         `json:"stock" db:"stock"`
	IsRestore        bool        `json:"isRestore" db:"is_restore"`
	Price            int         `json:"price" db:"price"`
	SpecialWeight    int         `json:"specialWeight" db:"special_weight"`
	Weight           int         `json:"weight" db:"weight"`
	Requirement      int         `json:"requirement" db:"requirement"`
	Icon             string      `json:"icon" db:"icon"`
	SimpleIcon       string      `json:"simpleIcon" db:"simple_icon"`
	AlertInSpectator bool        `json:"alertInSpectator" db:"alert_in_spectator"`
	CharacterCodes   interface{} `json:"characterCodes"` // string of comma joined ints or int
}

type ItemConsumable struct {
	Code                                               int    `json:"code" db:"code"`
	Name                                               string `json:"name" db:"name"`
	ModeType                                           int    `json:"modeType" db:"mode_type"`
	ItemType                                           string `json:"itemType" db:"item_type"`
	ConsumableType                                     string `json:"consumableType" db:"consumable_type"`
	ConsumableTag                                      string `json:"consumableTag" db:"consumable_tag"`
	ItemGrade                                          string `json:"itemGrade" db:"item_grade"`
	IsCompletedItem                                    bool   `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool   `json:"alertInSpectator" db:"alert_in_spectator"`
	MarkingType                                        string `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	Stackable                                          int    `json:"stackable" db:"stackable"`
	InitialCount                                       int    `json:"initialCount" db:"initial_count"`
	ItemUsableType                                     string `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int    `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  int    `json:"exclusiveProducer" db:"exclusive_producer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	ManufacturableType                                 int    `json:"manufacturableType" db:"manufacturable_type"`
	MakeMaterial1                                      int    `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int    `json:"makeMaterial2" db:"make_material_2"`
	Heal                                               int    `json:"heal" db:"heal"`
	HpRecover                                          int    `json:"hpRecover" db:"hp_recover"`
	SpRecover                                          int    `json:"spRecover" db:"sp_recover"`
	AttackPowerByBuff                                  int    `json:"attackPowerByBuff" db:"attack_power_by_buff"`
	DefenseByBuff                                      int    `json:"defenseByBuff" db:"defense_by_buff"`
	SkillAmpByBuff                                     int    `json:"skillAmpByBuff" db:"skill_amp_by_buff"`
	SkillAmpRatioByBuff                                int    `json:"skillAmpRatioByBuff" db:"skill_amp_ratio_by_buff"`
	AddStateCode                                       int    `json:"addStateCode" db:"add_state_code"`
	IsVPadQuickSlotItem                                bool   `json:"isVPadQuickSlotItem" db:"is_vpad_quick_slot_item"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int    `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
	IsReduceLootOnDeath                                bool   `json:"isReduceLootOnDeath" db:"is_reduce_loot_on_death"`
}

type ItemMisc struct {
	Code                                               int    `json:"code" db:"code"`
	Name                                               string `json:"name" db:"name"`
	ModeType                                           int    `json:"modeType" db:"mode_type"`
	ItemType                                           string `json:"itemType" db:"item_type"`
	MiscItemType                                       string `json:"miscItemType" db:"misc_item_type"`
	ItemGrade                                          string `json:"itemGrade" db:"item_grade"`
	GradeBgOverride                                    string `json:"gradeBgOverride" db:"grade_bg_override"`
	IsCompletedItem                                    bool   `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool   `json:"alertInSpectator" db:"alert_in_spectator"`
	MarkingType                                        string `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	Stackable                                          int    `json:"stackable" db:"stackable"`
	InitialCount                                       int    `json:"initialCount" db:"initial_count"`
	ItemUsableType                                     string `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int    `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  int    `json:"exclusiveProducer" db:"exclusive_producer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	ManufacturableType                                 int    `json:"manufacturableType" db:"manufacturable_type"`
	MakeMaterial1                                      int    `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int    `json:"makeMaterial2" db:"make_material_2"`
	MakeCustomAction                                   string `json:"makeCustomAction" db:"make_custom_action"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int    `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
}

type ItemSearchOption struct {
	Code int    `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
	Tag1 string `json:"tag1" db:"tag_1"`
	Tag2 string `json:"tag2" db:"tag_2"`
	Tag3 string `json:"tag3" db:"tag_3"`
}

type ItemSpawn struct {
	Code           int    `json:"code" db:"code"`
	AreaCode       int    `json:"areaCode" db:"area_code"`
	AreaSpawnGroup int    `json:"areaSpawnGroup" db:"area_spawn_group"`
	ItemCode       int    `json:"itemCode" db:"item_code"`
	DropPoint      string `json:"dropPoint" db:"drop_point"`
	DropCount      int    `json:"dropCount" db:"drop_count"`
}

type ItemSpecial struct {
	Code                                               int         `json:"code" db:"code"`
	Name                                               string      `json:"name" db:"name"`
	ModeType                                           int         `json:"modeType" db:"mode_type"`
	ItemType                                           string      `json:"itemType" db:"item_type"`
	SpecialItemType                                    string      `json:"specialItemType" db:"special_item_type"`
	ItemGrade                                          string      `json:"itemGrade" db:"item_grade"`
	IsCompletedItem                                    bool        `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool        `json:"alertInSpectator" db:"alert_in_spectator"`
	MarkingType                                        string      `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string      `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	Stackable                                          int         `json:"stackable" db:"stackable"`
	InitialCount                                       int         `json:"initialCount" db:"initial_count"`
	CooldownGroupCode                                  int         `json:"cooldownGroupCode" db:"int"`
	Cooldown                                           float64     `json:"cooldown" db:"cooldown"`
	ItemUsableType                                     string      `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int         `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  interface{} `json:"exclusiveProducer"` // number or comma separated numbers in a string
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool        `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	ManufacturableType                                 int         `json:"manufacturableType" db:"manufacturable_type"`
	MakeMaterial1                                      int         `json:"makeMaterial1" db:"make_material_2"`
	MakeMaterial2                                      int         `json:"makeMaterial2" db:"make_material_2"`
	MakeCustomAction                                   string      `json:"makeCustomAction" db:"make_custom_action"`
	ConsumeCount                                       int         `json:"consumeCount" db:"consume_count"`
	SummonCode                                         int         `json:"summonCode" db:"summon_code"`
	GhostItemStateGroup                                int         `json:"ghostItemStateGroup" db:"ghost_item_state_group"`
	IsVPadQuickSlotItem                                bool        `json:"isVPadQuickSlotItem" db:"is_vpad_quick_slot_item"`
	RestoreItemWhenResurrected                         bool        `json:"restoreItemWhenResurrected" db:"remove_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int         `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
	IsReduceLootOnDeath                                bool        `json:"isReduceLootOnDeath" db:"is_reduce_loot_on_death"`
}

type ItemWeapon struct {
	Code                                               int     `json:"code" db:"code"`
	Name                                               string  `json:"name" db:"name"`
	ModeType                                           int     `json:"modeType" db:"mode_type"`
	ItemType                                           string  `json:"itemType" db:"item_type"`
	WeaponType                                         string  `json:"weaponType" db:"weapon_type"`
	ItemGrade                                          string  `json:"itemGrade" db:"item_grade"`
	GradeBgOverride                                    string  `json:"gradeBgOverride" db:"grade_bg_override"`
	IsCompletedItem                                    bool    `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool    `json:"alertInSpectator" db:"alert_in_spectator"`
	MarkingType                                        string  `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	Stackable                                          int     `json:"stackable" db:"stackable"`
	InitialCount                                       int     `json:"initialCount" db:"initial_count"`
	ItemUsableType                                     string  `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int     `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  int     `json:"exclusiveProducer" db:"exclusive_producer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	MakeMaterial1                                      int     `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int     `json:"makeMaterial2" db:"make_material_2"`
	MakeCustomAction                                   string  `json:"makeCustomAction" db:"make_custom_action"`
	NotDisarm                                          bool    `json:"notDisarm" db:"not_disarm"`
	Consumable                                         bool    `json:"consumable" db:"consumable"`
	ManufacturableType                                 int     `json:"manufacturableType" db:"manufacturable_type"`
	AttackPower                                        int     `json:"attackPower" db:"attack_power"`
	AttackPowerByLv                                    int     `json:"attackPowerByLv" db:"attack_power_by_lv"`
	Defense                                            int     `json:"defense" db:"defense"`
	DefenseByLv                                        int     `json:"defenseByLv" db:"defense_by_lv"`
	SkillAmp                                           int     `json:"skillAmp" db:"skill_amp"`
	SkillAmpByLevel                                    int     `json:"skillAmpByLevel" db:"skill_amp_by_lv"`
	SkillAmpRatio                                      int     `json:"skillAmpRatio" db:"skill_amp_ratio"`
	SkillAmpRatioByLevel                               int     `json:"skillAmpRatioByLevel" db:"skill_amp_ratio_by_lv"`
	AdaptiveForce                                      int     `json:"adaptiveForce" db:"adaptive_force"`
	AdaptiveForceByLevel                               int     `json:"adaptiveForceByLevel" db:"adaptive_force_by_lv"`
	MaxHp                                              int     `json:"maxHp" db:"max_hp"`
	MaxHpByLv                                          int     `json:"maxHpByLv" db:"max_hp_by_lv"`
	HpRegenRatio                                       float64 `json:"hpRegenRatio" db:"hp_regen_ratio"`
	HpRegen                                            int     `json:"hpRegen" db:"hp_regen"`
	MaxSP                                              int     `json:"maxSP" db:"max_sp"`
	SpRegenRatio                                       float64 `json:"spRegenRatio" db:"sp_regen_ratio"`
	SpRegen                                            int     `json:"spRegen" db:"sp_regen"`
	AttackSpeedRatio                                   float64 `json:"attackSpeedRatio" db:"attack_speed_ratio"`
	AttackSpeedRatioByLv                               float32 `json:"attackSpeedRatioByLv" db:"attack_speed_ratio_by_lv"`
	CriticalStrikeChance                               float64 `json:"criticalStrikeChance" db:"critical_strike_chance"`
	CriticalStrikeDamage                               float64 `json:"criticalStrikeDamage" db:"critical_strike_damage"`
	CooldownReduction                                  float64 `json:"cooldownReduction" db:"cooldown_reduction"`
	PreventCriticalStrikeDamage                        int     `json:"preventCriticalStrikeDamaged" db:"prevent_critical_strike_damage"`
	CooldownLimit                                      int     `json:"cooldownLimit" db:"cooldown_limit"`
	LifeSteal                                          float64 `json:"lifeSteal" db:"lifesteal"`
	NormalLifeSteal                                    float64 `json:"normalLifeSteal" db:"normal_lifesteal"`
	SkillLifeSteal                                     float64 `json:"skillLifeSteal" db:"skill_lifesteal"`
	MoveSpeed                                          float64 `json:"moveSpeed" db:"move_speed"`
	MoveSpeedOutOfCombat                               float64 `json:"moveSpeedOutOfCombat" db:"move_speed_out_of_combat"`
	SightRange                                         float64 `json:"sightRange" db:"sight_range"`
	AttackRange                                        float64 `json:"attackRange" db:"attack_range"`
	IncreaseBasicAttackDamage                          int     `json:"increaseBasicAttackDamage" db:"increase_basic_attack_damage"`
	IncreaseBasicAttackDamageByLv                      int     `json:"increaseBasicAttackDamageByLv" db:"increase_basic_attack_damage_by_lv"`
	IncreaseBasicAttackDamageRatio                     float64 `json:"increaseBasicAttackDamageRatio" db:"increase_basic_attack_damage_ratio"`
	IncreaseBasicAttackDamageRatioByLv                 float64 `json:"increaseBasicAttackDamageRatioByLv" db:"increase_basic_attack_damage_ratio_by_lv"`
	PreventBasicAttackDamaged                          int     `json:"preventBasicAttackDamaged" db:"prevent_basic_attack_damage"`
	PreventBasicAttackDamagedByLv                      int     `json:"preventBasicAttackDamagedByLv" db:"prevent_basic_attack_damage_by_lv"`
	PreventBasicAttackDamagedRatio                     float64 `json:"preventBasicAttackDamagedRatio" db:"prevent_basic_attack_damage_ratio"`
	PreventBasicAttackDamagedRatioByLv                 float64 `json:"preventBasicAttackDamagedRatioByLv" db:"prevent_basic_attack_damage_ratio_by_lv"`
	PreventSkillDamage                                 int     `json:"preventSkillDamaged" db:"prevent_skill_damage"`
	PreventSkillDamageByLv                             int     `json:"preventSkillDamagedByLv" db:"prevent_skill_damage_by_lv"`
	PreventSkillDamageRatio                            float64 `json:"preventSkillDamagedRatio" db:"prevent_skill_damage_ratio"`
	PreventSkillDamageRatioByLv                        float64 `json:"preventSkillDamagedRatioByLv" db:"prevent_skill_damage_ratio_by_lv"`
	PenetrationDefense                                 int     `json:"penetrationDefense" db:"penetration_defense"`
	PenetrationDefenseRatio                            float64 `json:"penetrationDefenseRatio" db:"penetration_defense_ratio"`
	TrapDamageReduce                                   int     `json:"trapDamageReduce" db:"trap_damage_reduce"`
	TrapDamageReduceRatio                              float64 `json:"trapDamageReduceRatio" db:"trap_damage_reduce_ratio"`
	HpHealedIncreaseRatio                              float64 `json:"hpHealedIncreaseRatio" db:"hp_healed_increase_ratio"`
	HealerGiveHpHealRatio                              float64 `json:"healerGiveHpHealRatio" db:"healer_give_hp_heal_ratio"`
	UniqueAttackRange                                  float64 `json:"uniqueAttackRange" db:"unique_attack_range"`
	UniqueHpHealedIncreaseRatio                        float64 `json:"uniqueHpHealedIncreaseRatio" db:"unique_hp_healed_increase_ratio"`
	UniqueCooldownLimit                                float64 `json:"uniqueCooldownLimit" db:"unique_cooldown_limit"`
	UniqueTenacity                                     float64 `json:"uniqueTenacity" db:"unique_tenacity"`
	UniqueMoveSpeed                                    float64 `json:"uniqueMoveSpeed" db:"unique_move_speed"`
	UniquePenetrationDefense                           int     `json:"uniquePenetrationDefense" db:"unique_penetration_defense"`
	UniquePenetrationDefenseRatio                      float64 `json:"uniquePenetrationDefenseRatio" db:"unique_penetration_defense_ratio"`
	UniqueLifeSteal                                    float64 `json:"uniqueLifeSteal" db:"unique_lifesteal"`
	UniqueSkillAmpRatio                                float64 `json:"uniqueSkillAmpRatio" db:"unique_skill_amp_ratio"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int     `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
}

type Level struct {
	Level         int `json:"level" db:"level"`
	NeedExp       int `json:"needExp" db:"need_exp"`
	AccumulateExp int `json:"accumulateExp" db:"accumulate_exp"`
	RewardAcoin   int `json:"rewardAcoin" db:"reward_acoin"`
	Reward        int `json:"reward" db:"reward"`
}

type LoadingTip struct {
	Code           int    `json:"code" db:"code"`
	LoadingTipType string `json:"loadingTipType" db:"loading_tip_type"`
	MinLv          int    `json:"minLv" db:"min_lv"`
	MaxLv          int    `json:"maxLv" db:"max_lv"`
	TextKey        string `json:"textKey" db:"text_key"`
	ImageName      string `json:"imageName" db:"image_name"`
}

type MasteryExp struct {
	Code           int    `json:"code" db:"code"`
	ModeType       int    `json:"modeType" db:"mode_type"`
	ConditionType  string `json:"conditionType" db:"condition_type"`
	Grade          string `json:"grade" db:"grade"`
	ConditionValue int    `json:"conditionValue" db:"condition_value"`
	MasteryType1   string `json:"masteryType1" db:"mastery_type_1"`
	Value1         int    `json:"value1" db:"value_1"`
	MasteryType2   string `json:"masteryType2" db:"mastery_type_2"`
	Value2         int    `json:"value2" db:"value_2"`
	MasteryType3   string `json:"masteryType3" db:"mastery_type_3"`
	Value3         int    `json:"value3" db:"value_3"`
}

type MasteryLevel struct {
	Code              int    `json:"code" db:"code"`
	Type              string `json:"type" db:"type"`
	MasteryLevel      int    `json:"masteryLevel" db:"mastery_level"`
	NextMasteryExp    int    `json:"nextMasteryExp" db:"next_mastery"`
	GiveLevelExp      int    `json:"giveLevelExp" db:"give_level_exp"`
	ExpGrowthCapRatio int    `json:"expGrowthCapRatio" db:"exp_growth_cap_ratio"`
}

type MasteryStat struct {
	Code                      int     `json:"code" db:"code"`
	Type                      string  `json:"type" db:"type"`
	CharacterCode             int     `json:"characterCode" db:"character_code"`
	FirstOption               string  `json:"firstOption" db:"first_option"`
	FirstOptionSection1Value  float64 `json:"firstOptionSection1Value" db:"first_option_section_1_value"`
	FirstOptionSection2Value  float64 `json:"firstOptionSection2Value" db:"first_option_section_2_value"`
	FirstOptionSection3Value  float64 `json:"firstOptionSection3Value" db:"first_option_section_3_value"`
	FirstOptionSection4Value  float64 `json:"firstOptionSection4Value" db:"first_option_section_4_value"`
	SecondOption              string  `json:"secondOption" db:"second_option"`
	SecondOptionSection1Value float64 `json:"secondOptionSection1Value" db:"second_option_section_1_value"`
	SecondOptionSection2Value float64 `json:"secondOptionSection2Value" db:"second_option_section_2_value"`
	SecondOptionSection3Value float64 `json:"secondOptionSection3Value" db:"second_option_section_3_value"`
	SecondOptionSection4Value float64 `json:"secondOptionSection4Value" db:"second_option_section_4_value"`
	ThirdOption               string  `json:"thirdOption" db:"third_option"`
	ThirdOptionSection1Value  int     `json:"thirdOptionSection1Value" db:"third_option_section_1_value"`
	ThirdOptionSection2Value  int     `json:"thirdOptionSection2Value" db:"third_option_section_2_value"`
	ThirdOptionSection3Value  int     `json:"thirdOptionSection3Value" db:"third_option_section_3_value"`
	ThirdOptionSection4Value  int     `json:"thirdOptionSection4Value" db:"third_option_section_4_value"`
}

type Monster struct {
	Code             int     `json:"Code" db:"code"`
	Monster          string  `json:"monster" db:"monster"`
	IsMutant         bool    `json:"isMutant" db:"is_mutant"`
	Grade            string  `json:"grade" db:"grade"`
	Mode             int     `json:"mode" db:"mode"`
	CreateDay        string  `json:"createDay" db:"create_day"`
	CreateTime       int     `json:"createTime" db:"create_time"`
	RegenTime        int     `json:"regenTime" db:"regen_time"`
	LevelUpPeriod    int     `json:"levelUpPeriod" db:"level_up_period"`
	LevelUpAmount    int     `json:"levelUpAmount" db:"level_up_amount"`
	LevelUpMax       int     `json:"levelUpMax" db:"level_up_max"`
	MaxHp            int     `json:"maxHp" db:"max_hp"`
	MaxEp            int     `json:"maxEp" db:"max_ep"`
	InitExtraPoint   int     `json:"initExtraPoint" db:"init_extra_point"`
	AttackPower      int     `json:"attackPower" db:"attack_power"`
	Defense          int     `json:"defense" db:"defense"`
	AttackSpeed      float64 `json:"attackSpeed" db:"attack_speed"`
	MoveSpeed        float64 `json:"moveSpeed" db:"move_speed"`
	SightRange       int     `json:"sightRange" db:"sight_range"`
	ChasingRange     int     `json:"chasingRange" db:"chasing_range"`
	AttackRange      float64 `json:"attackRange" db:"attack_range"`
	FirstAttackRange float64 `json:"firstAttackRange" db:"first_attack_range"`
	Aggressive       string  `json:"aggressive" db:"aggressive"`
	DetectInvisible  bool    `json:"detectInvisible" db:"detect_invisible"`
	Radius           float64 `json:"radius" db:"radius"`
	PathingRadius    float64 `json:"pathingRadius" db:"pathing_radius"`
	UiHeight         float64 `json:"uiHeight" db:"ui_height"`
	GainExp          int     `json:"gainExp" db:"gain_exp"`
	TargetOnRange    int     `json:"targetOnRange" db:"target_on_range"`
	RandomDropCount  int     `json:"randomDropCount" db:"random_drop_count"`
	Resource         string  `json:"resource" db:"resource"`
	CorpseResource   string  `json:"corpseResource" db:"corpse_resource"`
	AppearTime       float64 `json:"appearTime" db:"appear_time"`
}

type MonsterDropGroup struct {
	MonsterCode  int `json:"monsterCode" db:"monster_code"`
	MonsterLevel int `json:"monsterLevel" db:"monster_level"`
	DropGroup    int `json:"dropGroup" db:"drop_group"`
}

type MonsterLevelUpStat struct {
	Code        int     `json:"code" db:"code"`
	Monster     string  `json:"monster" db:"monster"`
	Mode        int     `json:"mode" db:"mode"`
	MaxHp       int     `json:"maxHp" db:"max_hp"`
	AttackPower int     `json:"attackPower" db:"attack_power"`
	Defense     float64 `json:"defense" db:"defense"`
	MoveSpeed   float64 `json:"moveSpeed" db:"move_speed"`
	GainExp     int     `json:"gainExp" db:"gain_exp"`
}

type MonsterSpawnLevel struct {
	Code        int `json:"code" db:"code"`
	Mode        int `json:"mode" db:"mode"`
	PlayerLevel int `json:"playerLevel" db:"player_level"`
	MonsterCode int `json:"monsterCode" db:"monster_code"`
	SpawnLevel  int `json:"spawnLevel" db:"spawn_level"`
}

type NaviCollectAndHunt struct {
	Code         int    `json:"code" db:"code"`
	ItemCode     int    `json:"itemCode" db:"item_code"`
	AreaCodeList string `json:"areaCodeList" db:"area_code_list"`
}

type NearByArea struct {
	Code           int `json:"code" db:"code"`
	AreaCode       int `json:"areaCode" db:"area_code"`
	NearByAreaCode int `json:"nearByAreaCode" db:"near_by_area_code"`
}

type RandomEquipment struct {
	Code          int         `json:"code" db:"code"`
	Group         string      `json:"group" db:"group"`
	ItemCode      interface{} `json:"itemcode"` // string or number but always seems to be a number
	Weight        int         `json:"weight" db:"weight"`
	ItemGrade     string      `json:"itemGrade" db:"item_grade"`
	TagMultiplier int         `json:"tagMultiplier" db:"tag_multiplier"`
	CharacterNum  int         `json:"characterNum" db:"character_num"`
}

type RecommendedList struct {
	Code                  int    `json:"code" db:"code"`
	Character             string `json:"character" db:"character"`
	CharacterCode         int    `json:"characterCode" db:"character_code"`
	Mastery               string `json:"mastery" db:"mastery"`
	StartWeapon           int    `json:"startWeapon" db:"start_weapon"`
	CobaltStartWeapon     int    `json:"cobaltStartWeapon" db:"cobalt_start_weapon"`
	StartItemGroupCode    int    `json:"startItemGroupCode" db:"start_item_group_code"`
	CobaltDraft           int    `json:"cobaltDraft" db:"cobalt_draft"`
	CobaltExtraDraft      string `json:"cobaltExtraDraft" db:"cobalt_extra_draft"`
	CobaltCanChooseWeapon bool   `json:"cobaltCanChooseWeapon" db:"cobalt_can_choose_weapon"`
	FavoriteMainTag       string `json:"favoriteMainTag" db:"favorite_main_tag"`
	OppositeTag           string `json:"oppositeTag" db:"opposite_tag"`
}

type Season struct {
	SeasonID    int    `json:"seasonID" db:"season_id"`
	SeasonName  string `json:"seasonName"` // no tags as they need to be converted to date
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
	IsCurrent   int    `json:"isCurrent" db:"is_current"`
}

type SummonObjectStat struct {
	Code                 int     `json:"code" db:"code"`
	Name                 string  `json:"name" db:"name"`
	Duration             float64 `json:"duration" db:"duration"`
	CreateRange          float64 `json:"createRange" db:"create_range"`
	PileRange            float64 `json:"pileRange" db:"pile_range"`
	CreateVisibleTime    float64 `json:"createVisibleTime" db:"create_visible_time"`
	CreateStealthTime    float64 `json:"createStealthTime" db:"create_stealth_time"`
	InfiltrationTime     float64 `json:"infiltrationTime" db:"infiltration_time"`
	DetectionRange       float64 `json:"detectionRange" db:"detection_range"`
	MaxHp                int     `json:"maxHp" db:"max_hp"`
	MaxSp                int     `json:"maxSp" db:"max_sp"`
	InitExtraPoint       int     `json:"initExtraPoint" db:"init_extra_point"`
	MaxExtraPoint        int     `json:"maxExtraPoint" db:"max_extra_point"`
	AttackPower          int     `json:"attackPower" db:"attack_power"`
	Defense              int     `json:"defense" db:"defense"`
	CriticalStrikeChance float64 `json:"criticalStrikeChance" db:"critical_strike_chance"`
	HpRegen              int     `json:"hpRegen" db:"hp_regen"`
	SpRegen              int     `json:"spRegen" db:"sp_regen"`
	RangeRadius          float64 `json:"rangeRadius" db:"range_radius"`
	AttackSpeed          float64 `json:"attackSpeed" db:"attack_speed"`
	AttackRange          float64 `json:"attackRange" db:"attack_range"`
	AttackDelay          float64 `json:"attackDelay" db:"attack_delay"`
	MoveSpeed            float64 `json:"moveSpeed" db:"move_speed"`
	Radius               float64 `json:"radius" db:"radius"`
	UiHeight             float64 `json:"uiHeight" db:"ui_height"`
	SightRange           float64 `json:"sightRange" db:"sight_range"`
	SightAngle           int     `json:"sightAngle" db:"sight_angle"`
}

type TacticalSkillSet struct {
	Code            int `json:"code" db:"code"`
	NextUpgradeCode int `json:"nextUpgradecode" db:"next_upgrade_code"`
	UpgradeCredit   int `json:"upgradeCredit" db:"upgrade_credit"`
	UpgradeMaterial int `json:"upgradeMaterial" db:"upgrade_material"`
	SkillCode       int `json:"skillCode" db:"skill_code"`
}

type TacticalSkillSetGroup struct {
	Group          int    `json:"group" db:"group"`
	ModeType       int    `json:"modeType" db:"mode_type"`
	StartCode      int    `json:"startCode" db:"start_code"`
	EquipWithStart bool   `json:"equipWithStart" db:"equip_with_start"`
	Icon           string `json:"icon" db:"icon"`
}

type Trait struct {
	Code          int    `json:"code" db:"code"`
	OpenAccountLv int    `json:"openAccountLv" db:"open_account_lv"`
	TraitGroup    string `json:"traitGroup" db:"trait_group"`
	TraitType     string `json:"traitType" db:"trait_type"`
	Active        bool   `json:"active" db:"active"`
}

type TransferConsole struct {
	ItemCode                           int    `json:"itemCode" db:"item_code"`
	Mode                               int    `json:"mode" db:"mode"`
	ItemType                           string `json:"itemType" db:"item_type"`
	TransferTimeSafeArea               int    `json:"transferTimeSafeArea" db:"transfer_time_safe_area"`
	SubtractTransferTimeRestrictedArea int    `json:"subtractTrasferTimeRestrictedArea" db:"subtract_transfer_time_restricted_area"`
	ManufactureCooldown                int    `json:"manufactureCooldown" db:"manufacture_cooldown"`
	AvailableTimeForPurchase           int    `json:"availableTimeForPurchase" db:"available_time_for_purchase"`
	ConsumeVFCredit                    int    `json:"consumeVFCredit" db:"consume_vf_credit"`
	LimitCount                         int    `json:"limitCount" db:"limit_count"`
	TraitSale                          bool   `json:"traitSale" db:"trait_sale"`
}

type VfCredit struct {
	Code           int     `json:"code" db:"code"`
	Mode           int     `json:"mode" db:"mode"`
	Phase          int     `json:"phase" db:"phase"`
	ConditionType  string  `json:"conditionType" db:"condition_type"`
	ConditionValue int     `json:"conditionValue" db:"condition_value"`
	AcquireSelf    float64 `json:"acquireSelf" db:"acquire_self"`
	AcquireTeam    int     `json:"acquireTeam" db:"acquire_team"`
}

type WeaponTypeInfo struct {
	Type                  string  `json:"type" db:"type"`
	AttackSpeed           float64 `json:"attackSpeed" db:"attack_speed"`
	AttackRange           float64 `json:"attackRange" db:"attack_range"`
	ShopFilter            int     `json:"shopFilter" db:"shop_filter"`
	SummonObjectHitDamage int     `json:"summonObjectHitDamage" db:"summon_object_hit_damage"`
}
