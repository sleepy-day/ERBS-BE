package models

import (
	"github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

//
// Game Data API structs
//

type ActionCost struct {
	Id                    int64   `db:"id"`
	Code                  int     `json:"code" db:"code"`
	Type                  string  `json:"type" db:"type"`
	Sp                    int     `json:"sp" db:"sp"`
	Time1                 float64 `json:"time1" db:"time_1"`
	Time2                 float64 `json:"time2" db:"time_2"`
	ActionWaitTime        float64 `json:"actionWaitTime" db:"action_wait_time"`
	CastingAnimTrigger    string  `json:"castingAnimTrigger" db:"casting_anim_trigger"`
	EffectCancelCondition string  `json:"effectCancelCondition" db:"effect_cancel_condition"`
	CastingBarImgName     string  `json:"castingBarImgName" db:"casting_bar_img_name"`
}

type Area struct {
	Id                       int64  `db:"id"`
	Code                     int    `json:"code" db:"code"`
	Name                     string `json:"name" db:"name"`
	ModeType                 int    `json:"modeType" db:"mode_type"`
	MaskCode                 int    `json:"maskCode" db:"mask_code"`
	StartingArea             bool   `json:"startingArea" db:"starting_area"`
	AreaType                 string `json:"areaType" db:"area_type"`
	IsProvideCollectibleItem bool   `json:"isProvideCollectibleItem" db:"is_provide_collectible_item"`
	RouteCalcBitCode         int    `json:"routeCalcBitCode" db:"route_calc_bit_code"`
	IsHyperLoopInstalled     bool   `json:"isHyperLoopInstalled" db:"is_hyperloop_installed"`
}

type BattleZoneReward struct {
	Id                             int64  `db:"id"`
	Code                           int    `json:"code" db:"code"`
	ModeType                       int    `json:"modeType" db:"mode_type"`
	AreaAttributesCreateEventCount int    `json:"areaAttributesCreateEventCount" db:"area_attributes_create_event_count"`
	ItemCode                       int    `json:"itemCode" db:"item_code"`
	Type                           string `json:"type" db:"type"`
	Value                          int    `json:"value" db:"value"`
	Selectable                     bool   `json:"selectable" db:"selectable"`
}

type BulletCapacity struct {
	Id        int64   `db:"id"`
	ItemCode  int     `json:"itemCode" db:"item_code"`
	Capacity  int     `json:"capacity" db:"capacity"`
	LoadType  string  `json:"loadType" db:"load_type"`
	Time      float64 `json:"time" db:"time"`
	InitCount int     `json:"initCount" db:"init_count"`
	Count     int     `json:"count" db:"count"`
}

type Character struct {
	Id                             int64   `db:"id"`
	Code                           int     `json:"code" db:"code"`
	MaxHp                          int     `json:"maxHp" db:"max_hp"`
	MaxSp                          int     `json:"maxSp" db:"max_sp"`
	InitExtraPoint                 int     `json:"initExtraPoint" db:"init_extra_point"`
	MaxExtraPoint                  int     `json:"maxExtraPoint" db:"max_extra_point"`
	AttackPower                    int     `json:"attackPower" db:"attack_power"`
	Defense                        int     `json:"defense" db:"defense"`
	SkillAmp                       int     `json:"skillAmp" db:"skill_amp"`
	AdaptiveForce                  int     `json:"adaptiveForce" db:"adaptive_force"`
	CriticalStrikeChance           float64 `json:"criticalStrikeChance" db:"critical_strike_chance"`
	HpRegen                        float64 `json:"hpRegen" db:"hp_regen"`
	SpRegen                        float64 `json:"spRegen" db:"sp_regen"`
	AttackSpeed                    float64 `json:"attackSpeed" db:"attack_speed"`
	AttackSpeedRatio               float64 `json:"attackSpeedRatio" db:"attack_speed_ratio"`
	IncreaseBasicAttackDamageRatio float64 `json:"increaseBasicAttackDamageRatio" db:"increase_basic_attack_damage_ratio"`
	SkillAmpRatio                  float64 `json:"skillAmpRatio" db:"skill_amp_ratio"`
	PreventBasicAttackDamageRatio  float64 `json:"preventBasicAttackDamagedRatio" db:"prevent_basic_attack_damage_ratio"`
	PreventSkillDamageRatio        float64 `json:"preventSkillDamagedRatio" db:"prevent_skill_damage_ratio"`
	AttackSpeedLimit               float64 `json:"attackSpeedLimit" db:"attack_speed_limit"`
	AttackSpeedMin                 float64 `json:"attackSpeedMin" db:"attack_speed_min"`
	MoveSpeed                      float64 `json:"moveSpeed" db:"move_speed"`
	SightRange                     float64 `json:"sightRange" db:"sight_range"`
	Radius                         float64 `json:"radius" db:"radius"`
	PathingRadius                  float64 `json:"pathingRadius" db:"pathing_radius"`
	UIHeight                       float64 `json:"uiHeight" db:"ui_height"`
	InitStateDisplayIndex          int     `json:"initStateDisplayIndex" db:"init_state_display_index"`
	LocalScaleInCutscene           int     `json:"localScaleInCutscene" db:"local_scale_in_cutscene"`
	LocalScaleInVictoryScene       string  `json:"localScaleInVictoryScene" db:"local_scale_in_victory_scene"`
	Resource                       string  `json:"resource" db:"resource"`
	LobbySubObject                 string  `json:"lobbySubObject" db:"lobby_sub_object"`
	Name                           string  `json:"name" db:"name"`
	StrLearnStartSkill             string  `json:"strLearnStartSkill" db:"str_learn_start_skill"`
	StrUsePointLearnStartSkill     string  `json:"strUsePointLearnStartSkill" db:"str_use_point_learn_start_skill"`
}

type CharacterAttributes struct {
	Id                int64  `db:"id"`
	Character         string `json:"character" db:"character"`
	CharacterCode     int    `json:"characterCode" db:"character_code"`
	Mastery           string `json:"mastery" db:"mastery"`
	ControlDifficulty int    `json:"controlDifficulty" db:"control_difficulty"`
	Attack            int    `json:"attack" db:"attack"`
	Defense           int    `json:"defense" db:"defense"`
	Disruptor         int    `json:"disruptor" db:"disruptor"`
	Move              int    `json:"move" db:"move"`
	Assistance        int    `json:"assistance" db:"assistance"`
}

type CharacterExp struct {
	Id         int64 `db:"id"`
	Level      int   `json:"level" db:"level"`
	LevelUpExp int   `json:"levelUpExp" db:"level_up_exp"`
}

type CharacterLevelUpStat struct {
	Id                             int64   `db:"id"`
	Code                           int     `json:"code" db:"code"`
	MaxHp                          int     `json:"maxHp" db:"max_hp"`
	MaxSp                          int     `json:"maxSp" db:"max_sp"`
	SkillAmp                       int     `json:"skillAmp" db:"skill_amp"`
	AdaptiveForce                  int     `json:"adaptiveForce" db:"adaptive_force"`
	AttackPower                    float64 `json:"attackPower" db:"attack_power"`
	Defense                        float64 `json:"defense" db:"defense"`
	CriticalChance                 float64 `json:"criticalChance" db:"critical_strike_chance"`
	HpRegen                        float64 `json:"hpRegen" db:"hp_regen"`
	SpRegen                        float64 `json:"spRegen" db:"sp_regen"`
	AttackSpeed                    float64 `json:"attackSpeed" db:"attack_speed"`
	MoveSpeed                      float64 `json:"moveSpeed" db:"move_speed"`
	AttackSpeedRatio               float64 `json:"attackSpeedRatio" db:"attack_speed_ratio"`
	IncreaseBasicAttackDamageRatio float64 `json:"increaseBasicAttackDamageRatio" db:"increase_basic_attack_damage_ratio"`
	SkillAmpRatio                  float64 `json:"skillAmpRatio" db:"skill_amp_ratio"`
	PreventBasicAttackDamageRatio  float64 `json:"preventBasicAttackDamagedRatio" db:"prevent_basic_attack_damage_ratio"`
	PreventSkillDamageRatio        float64 `json:"preventSkillDamagedRatio" db:"prevent_skill_damage_ratio"`
	Name                           string  `json:"name" db:"name"`
}

type ItemArmor struct {
	Id                                                 int64   `db:"id"`
	Code                                               int     `json:"code" db:"code"`
	MakeMaterial1                                      int     `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int     `json:"makeMaterial2" db:"make_material_2"`
	CreditValueWhenConvertedToBounty                   int     `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
	ManufacturableType                                 int     `json:"manufacturableType" db:"manufacturable_type"`
	AttackPower                                        int     `json:"attackPower" db:"attack_power"`
	AttackPowerByLv                                    int     `json:"attackPowerByLv" db:"attack_power_by_lv"`
	Defense                                            int     `json:"defense" db:"defense"`
	DefenseByLv                                        int     `json:"defenseByLv" db:"defense_by_lv"`
	SkillAmp                                           int     `json:"skillAmp" db:"skill_amp"`
	SkillAmpByLevel                                    int     `json:"skillAmpByLevel" db:"skill_amp_by_level"`
	AdaptiveForce                                      int     `json:"adaptiveForce" db:"adaptive_force"`
	AdaptiveForceByLevel                               int     `json:"adaptiveForceByLevel" db:"adaptive_force_by_level"`
	MaxHp                                              int     `json:"maxHp" db:"max_hp"`
	MaxHpByLv                                          int     `json:"maxHpByLv" db:"max_hp_by_lv"`
	MaxSp                                              int     `json:"maxSp" db:"max_sp"`
	PreventCriticalStrikeDamage                        int     `json:"preventCriticalStrikeDamaged" db:"prevent_critical_strike_damage"`
	AttackRange                                        int     `json:"attackRange" db:"attack_range"`
	IncreaseBasicAttackDamage                          int     `json:"increaseBasicAttackDamage" db:"increase_basic_attack_damage"`
	IncreaseBasicAttackDamageByLv                      int     `json:"increaseBasicAttackDamageByLv" db:"increase_basic_attack_damage_by_lv"`
	PreventBasicAttackDamage                           int     `json:"preventBasicAttackDamaged" db:"prevent_basic_attack_damage"`
	PreventBasicAttackDamageByLv                       int     `json:"preventBasicAttackDamagedByLv" db:"prevent_basic_attack_damage_by_lv"`
	IncreaseBasicAttackDamageRatio                     int     `json:"increaseBasicAttackDamageRatio" db:"increase_basic_attack_damage_ratio"`
	PreventSkillDamage                                 int     `json:"preventSkillDamaged" db:"prevent_skill_damage"`
	PreventSkillDamageByLv                             int     `json:"preventSkillDamagedByLv" db:"prevent_skill_damage_by_lv"`
	PenetrationDefense                                 int     `json:"penetrationDefense" db:"penetration_defense"`
	TrapDamageReduce                                   int     `json:"trapDamageReduce" db:"trap_damage_reduce"`
	UniqueAttackRange                                  int     `json:"uniqueAttackRange" db:"unique_attack_range"`
	UniquePenetrationDefense                           int     `json:"uniquePenetrationDefense" db:"unique_penetration_defense"`
	ItemUsableValueList                                int     `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  int     `json:"exclusiveProducer" db:"exclusive_producer"`
	ModeType                                           int     `json:"modeType" db:"mode_type"`
	Stackable                                          int     `json:"stackable" db:"stackable"`
	InitialCount                                       int     `json:"initialCount" db:"initial_count"`
	SightRange                                         int     `json:"sightRange" db:"sight_range"`
	HpRegenRatio                                       float64 `json:"hpRegenRatio" db:"hp_regen_ratio"`
	HpRegen                                            float64 `json:"hpRegen" db:"hp_regen"`
	SpRegenRatio                                       float64 `json:"spRegenRatio" db:"sp_regen_ratio"`
	SpRegen                                            float64 `json:"spRegen" db:"sp_regen"`
	AttackSpeedRatio                                   float64 `json:"attackSpeedRatio" db:"attack_speed_ratio"`
	AttackSpeedRatioByLv                               float64 `json:"attackSpeedRatioByLv" db:"attack_speed_ratio_by_lv"`
	CriticalStrikeChance                               float64 `json:"criticalStrikeChance" db:"critical_strike_chance"`
	CriticalStrikeDamage                               float64 `json:"criticalStrikeDamage" db:"critical_strike_damage"`
	SkillAmpRatio                                      float64 `json:"skillAmpRatio" db:"skill_amp_ratio"`
	SkillAmpRatioByLevel                               float64 `json:"skillAmpRatioByLevel" db:"skill_amp_ratio_by_level"`
	CooldownReduction                                  float64 `json:"cooldownReduction" db:"cooldown_reduction"`
	CooldownLimit                                      float64 `json:"cooldownLimit" db:"cooldown_limit"`
	LifeSteal                                          float64 `json:"lifeSteal" db:"lifesteal"`
	NormalLifeSteal                                    float64 `json:"normalLifeSteal" db:"normal_lifesteal"`
	SkillLifeSteal                                     float64 `json:"skillLifeSteal" db:"skill_lifesteal"`
	MoveSpeed                                          float64 `json:"moveSpeed" db:"move_speed"`
	MoveSpeedOutOfCombat                               float64 `json:"moveSpeedOutOfCombat" db:"move_speed_out_of_combat"`
	UniqueHpHealedIncreaseRatio                        float64 `json:"uniqueHpHealedIncreaseRatio" db:"unique_hp_healed_increase_ratio"`
	UniqueCooldownLimit                                float64 `json:"uniqueCooldownLimit" db:"unique_cooldown_limit"`
	UniqueTenacity                                     float64 `json:"uniqueTenacity" db:"unique_tenacity"`
	UniqueMoveSpeed                                    float64 `json:"uniqueMoveSpeed" db:"unique_move_speed"`
	PreventBasicAttackDamageRatio                      float64 `json:"preventBasicAttackDamagedRatio" db:"prevent_basic_attack_damage_ratio"`
	PreventBasicAttackDamageRatioByLv                  float64 `json:"preventBasicAttackDamagedRatioByLv" db:"prevent_basic_attack_damage_ratio_by_lv"`
	PreventSkillDamageRatio                            float64 `json:"preventSkillDamagedRatio" db:"prevent_skill_damage_ratio"`
	PreventSkillDamageRatioByLv                        float64 `json:"preventSkillDamagedRatioByLv" db:"prevent_skill_damage_ratio_by_lv"`
	PenetrationDefenseRatio                            float64 `json:"penetrationDefenseRatio" db:"penetration_defense_ratio"`
	TrapDamageReduceRatio                              float64 `json:"trapDamageReduceRatio" db:"trap_damage_reduce_ratio"`
	HpHealedIncreaseRatio                              float64 `json:"hpHealedIncreaseRatio" db:"hp_healed_increase_ratio"`
	HealerGiveHpHealRatio                              float64 `json:"healerGiveHpHealRatio" db:"healer_give_hp_heal_ratio"`
	UniquePenetrationDefenseRatio                      float64 `json:"uniquePenetrationDefenseRatio" db:"unique_penetration_defense_ratio"`
	UniqueLifeSteal                                    float64 `json:"uniqueLifeSteal" db:"unique_lifesteal"`
	UniqueSkillAmpRatio                                float64 `json:"uniqueSkillAmpRatio" db:"unique_skill_amp_ratio"`
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
	Id        int64  `db:"id"`
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
	Id                                      int64  `db:"id"`
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
	CobaltIncreaseModeHealerGiveHealRatio   int    `json:"cobaltIncreaseModeHealerGiveHealRatio" db:"cobalt_increase_mode_healer_give_heal_ratio"`
	CobaltIncreaseModeShieldRatio           int    `json:"cobaltIncreaseModeShieldRatio" db:"cobalt_increase_mode_shield_ratio"`
	CobaltIncreaseModeHealerGiveShieldRatio int    `json:"cobaltIncreaseModeHealerGiveShieldRatio" db:"cobalt_increase_mode_healer_give_shield_ratio"`
	CobaltIncreaseModeUltCooldownRatio      int    `json:"cobaltIncreaseModeUltCooldownRatio" db:"cobalt_increase_mode_ult_cooldown_ratio"`
	CobaltIncreaseModeMaxSpRatio            int    `json:"cobaltIncreaseModeMaxSpRatio" db:"cobalt_increase_mode_max_sp_ratio"`
	CobaltIncreaseModeSpRegenRatio          int    `json:"cobaltIncreaseModeSpRegenRatio" db:"cobalt_increase_mode_sp_regen_ratio"`
	SoloIncreaseModeDamageToMonsterRatio    int    `json:"soloIncreaseModeDamageToMonsterRatio" db:"solo_increase_mode_damage_to_monster_ratio"`
	DuoIncreaseModeDamageToMonsterRatio     int    `json:"duoIncreaseModeDamageToMonsterRatio" db:"duo_increase_mode_damage_to_monster_ratio"`
	SquadIncreaseModeDamageToMonsterRatio   int    `json:"squadIncreaseModeDamageToMonsterRatio" db:"squad_increase_mode_damage_to_monster_ratio"`
	CobaltIncreaseModeDamageToMonsterRatio  int    `json:"cobaltIncreaseModeDamageToMonsterRatio" db:"cobalt_increase_mode_damage_to_monster_ratio"`
}

type CharacterSkin struct {
	Id                       int64  `db:"id"`
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
	Id                int64  `db:"id"`
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
	Id            int64       `db:"id"`
	GroupCode     int         `json:"groupCode" db:"group_code"`
	ItemCodeValue int         `db:"item_code"`
	ItemCode      interface{} `json:"itemCode"`
	Min           int         `json:"min" db:"min"`
	Max           int         `json:"max" db:"max"`
	Probability   int         `json:"probability" db:"probability"`
	DropType      string      `json:"dropType" db:"drop_type"`
}

func (dg *DropGroup) Prepare() {
	switch dg.ItemCode.(type) {
	case int:
		dg.ItemCodeValue = dg.ItemCode.(int)
	case string:
		val, err := strconv.Atoi(dg.ItemCode.(string))
		if err != nil {
			break
		}
		dg.ItemCodeValue = val
	}
}

type GainExp struct {
	Id        int64 `db:"id"`
	StartTime int   `json:"startTime" db:"start_time"`
	EndTime   int   `json:"endTime" db:"end_time"`
	GainExp   int   `json:"gainExp" db:"gain_exp"`
}

type GainScore struct {
	Id             int64  `db:"id"`
	Code           int    `json:"code" db:"code"`
	Phase          int    `json:"phase" db:"phase"`
	ConditionType  string `json:"conditionType" db:"condition_type"`
	ConditionValue int    `json:"conditionValue" db:"condition_value"`
	PointsEnemy    int    `json:"pointsEnemy" db:"points_enemy"`
	PointsAlly     int    `json:"pointsAlly" db:"points_ally"`
}

type GameTip struct {
	Id              int64  `db:"id"`
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
	Id              int64       `db:"id"`
	Code            int         `json:"code" db:"code"`
	ItemCodeValue   int         `db:"item_code"`
	ItemCode        interface{} `json:"itemCode"`
	HuntChicken     int         `json:"huntChicken" db:"hunt_chicken"`
	HuntBat         int         `json:"huntBat" db:"hunt_bat"`
	HuntBoar        int         `json:"huntBoar" db:"hunt_boar"`
	HuntWildDog     int         `json:"huntWildDog" db:"hunt_dog"`
	HuntWolf        int         `json:"huntWolf" db:"hunt_wolf"`
	HuntBear        int         `json:"huntBear" db:"hunt_bear"`
	HuntWickeline   int         `json:"huntWickline" db:"hunt_wickeline"`
	HuntAlpha       int         `json:"huntAlpha" db:"hunt_alpha"`
	HuntOmega       int         `json:"huntOmega" db:"hunt_omega"`
	CollectibleCode int         `json:"collectibleCode" db:"collectible_code"`
	AirSupply       int         `json:"airSupply" db:"air_supply"`
}

func (fi *HowToFindItem) Prepare() {
	switch fi.ItemCode.(type) {
	case int:
		fi.ItemCodeValue = fi.ItemCode.(int)
	case string:
		val, err := strconv.Atoi(fi.ItemCode.(string))
		if err != nil {
			break
		}
		fi.ItemCodeValue = val
	}
}

type InfusionProduct struct {
	Id                 int64       `db:"id"`
	Code               int         `json:"code" db:"code"`
	ProductType        string      `json:"productType" db:"product_type"`
	ProductGroup       int         `json:"productGroup" db:"product_group"`
	ProductCode        int         `json:"productCode" db:"product_code"`
	StoreType          string      `json:"storeType" db:"store_type"`
	StockType          string      `json:"stockType" db:"stock_type"`
	Stock              int         `json:"stock" db:"stock"`
	IsRestore          bool        `json:"isRestore" db:"is_restore"`
	Price              int         `json:"price" db:"price"`
	SpecialWeight      int         `json:"specialWeight" db:"special_weight"`
	Weight             int         `json:"weight" db:"weight"`
	Requirement        int         `json:"requirement" db:"requirement"`
	Icon               string      `json:"icon" db:"icon"`
	SimpleIcon         string      `json:"simpleIcon" db:"simple_icon"`
	AlertInSpectator   bool        `json:"alertInSpectator" db:"alert_in_spectator"`
	CharacterCodeArray interface{} `db:"character_codes"`
	CharacterCodes     interface{} `json:"characterCodes"` // string of comma joined ints or int
}

func (ip *InfusionProduct) Prepare() {
	switch ip.CharacterCodes.(type) {
	case int:
		val := ip.CharacterCodes.(int)
		ip.CharacterCodeArray = pq.Array(val)
	case string:
		values := strings.Split(ip.CharacterCodes.(string), ",")
		var result []int
		for _, e := range values {
			res, err := strconv.Atoi(e)
			if err != nil {
				continue
			} // TODO: log
			result = append(result, res)
		}
		ip.CharacterCodeArray = pq.Array(result)
	}
}

type ItemConsumable struct {
	Id                                                 int64  `db:"id"`
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
	Id                                                 int64  `db:"id"`
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
	Id   int64  `db:"id"`
	Code int    `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
	Tag1 string `json:"tag1" db:"tag_1"`
	Tag2 string `json:"tag2" db:"tag_2"`
	Tag3 string `json:"tag3" db:"tag_3"`
}

type ItemSpawn struct {
	Id             int64  `db:"id"`
	Code           int    `json:"code" db:"code"`
	AreaCode       int    `json:"areaCode" db:"area_code"`
	AreaSpawnGroup int    `json:"areaSpawnGroup" db:"area_spawn_group"`
	ItemCode       int    `json:"itemCode" db:"item_code"`
	DropPoint      string `json:"dropPoint" db:"drop_point"`
	DropCount      int    `json:"dropCount" db:"drop_count"`
}

type ItemSpecial struct {
	Id                                                 int64       `db:"id"`
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
	CooldownGroupCode                                  int         `json:"cooldownGroupCode" db:"cooldown_group_code"`
	Cooldown                                           float64     `json:"cooldown" db:"cooldown"`
	ItemUsableType                                     string      `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int         `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducerArray                             interface{} `db:"exclusive_producer"`
	ExclusiveProducer                                  interface{} `json:"exclusiveProducer"` // number or comma separated numbers in a string
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool        `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	ManufacturableType                                 int         `json:"manufacturableType" db:"manufacturable_type"`
	MakeMaterial1                                      int         `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int         `json:"makeMaterial2" db:"make_material_2"`
	MakeCustomAction                                   string      `json:"makeCustomAction" db:"make_custom_action"`
	ConsumeCount                                       int         `json:"consumeCount" db:"consume_count"`
	SummonCode                                         int         `json:"summonCode" db:"summon_code"`
	GhostItemStateGroup                                int         `json:"ghostItemStateGroup" db:"ghost_item_state_group"`
	IsVPadQuickSlotItem                                bool        `json:"isVPadQuickSlotItem" db:"is_vpad_quick_slot_item"`
	RestoreItemWhenResurrected                         bool        `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int         `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
	IsReduceLootOnDeath                                bool        `json:"isReduceLootOnDeath" db:"is_reduce_loot_on_death"`
}

func (ip *ItemSpecial) Prepare() {
	switch ip.ExclusiveProducer.(type) {
	case int:
		val := ip.ExclusiveProducer.(int)
		ip.ExclusiveProducerArray = pq.Array(val)
	case string:
		values := strings.Split(ip.ExclusiveProducer.(string), ",")
		var result []int
		for _, e := range values {
			res, err := strconv.Atoi(e)
			if err != nil {
				continue
			} // TODO: log
			result = append(result, res)
		}
		ip.ExclusiveProducerArray = pq.Array(result)
	}
}

type ItemWeapon struct {
	Id                                                 int64       `db:"id"`
	Code                                               int         `json:"code" db:"code"`
	Name                                               string      `json:"name" db:"name"`
	ModeType                                           int         `json:"modeType" db:"mode_type"`
	ItemType                                           string      `json:"itemType" db:"item_type"`
	WeaponType                                         string      `json:"weaponType" db:"weapon_type"`
	ItemGrade                                          string      `json:"itemGrade" db:"item_grade"`
	GradeBgOverride                                    string      `json:"gradeBgOverride" db:"grade_bg_override"`
	IsCompletedItem                                    bool        `json:"isCompletedItem" db:"is_completed_item"`
	AlertInSpectator                                   bool        `json:"alertInSpectator" db:"alert_in_spectator"`
	MarkingType                                        string      `json:"markingType" db:"marking_type"`
	CraftAnimTrigger                                   string      `json:"craftAnimTrigger" db:"craft_anim_trigger"`
	Stackable                                          int         `json:"stackable" db:"stackable"`
	InitialCount                                       int         `json:"initialCount" db:"initial_count"`
	ItemUsableType                                     string      `json:"itemUsableType" db:"item_usable_type"`
	ItemUsableValueList                                int         `json:"itemUsableValueList" db:"item_usable_value_list"`
	ExclusiveProducer                                  interface{} `json:"exclusiveProducer"`
	ExclusiveProducerArray                             interface{} `db:"exclusive_producer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool        `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled" db:"is_removed_from_player_inventory_on_death"`
	MakeMaterial1                                      int         `json:"makeMaterial1" db:"make_material_1"`
	MakeMaterial2                                      int         `json:"makeMaterial2" db:"make_material_2"`
	MakeCustomAction                                   string      `json:"makeCustomAction" db:"make_custom_action"`
	NotDisarm                                          bool        `json:"notDisarm" db:"not_disarm"`
	Consumable                                         bool        `json:"consumable" db:"consumable"`
	ManufacturableType                                 int         `json:"manufacturableType" db:"manufacturable_type"`
	AttackPower                                        int         `json:"attackPower" db:"attack_power"`
	AttackPowerByLv                                    int         `json:"attackPowerByLv" db:"attack_power_by_lv"`
	Defense                                            int         `json:"defense" db:"defense"`
	DefenseByLv                                        int         `json:"defenseByLv" db:"defense_by_lv"`
	SkillAmp                                           int         `json:"skillAmp" db:"skill_amp"`
	SkillAmpByLevel                                    int         `json:"skillAmpByLevel" db:"skill_amp_by_lv"`
	SkillAmpRatio                                      int         `json:"skillAmpRatio" db:"skill_amp_ratio"`
	SkillAmpRatioByLevel                               int         `json:"skillAmpRatioByLevel" db:"skill_amp_ratio_by_lv"`
	AdaptiveForce                                      int         `json:"adaptiveForce" db:"adaptive_force"`
	AdaptiveForceByLevel                               int         `json:"adaptiveForceByLevel" db:"adaptive_force_by_lv"`
	MaxHp                                              int         `json:"maxHp" db:"max_hp"`
	MaxHpByLv                                          int         `json:"maxHpByLv" db:"max_hp_by_lv"`
	HpRegenRatio                                       float64     `json:"hpRegenRatio" db:"hp_regen_ratio"`
	HpRegen                                            int         `json:"hpRegen" db:"hp_regen"`
	MaxSP                                              int         `json:"maxSP" db:"max_sp"`
	SpRegenRatio                                       float64     `json:"spRegenRatio" db:"sp_regen_ratio"`
	SpRegen                                            int         `json:"spRegen" db:"sp_regen"`
	AttackSpeedRatio                                   float64     `json:"attackSpeedRatio" db:"attack_speed_ratio"`
	AttackSpeedRatioByLv                               float32     `json:"attackSpeedRatioByLv" db:"attack_speed_ratio_by_lv"`
	CriticalStrikeChance                               float64     `json:"criticalStrikeChance" db:"critical_strike_chance"`
	CriticalStrikeDamage                               float64     `json:"criticalStrikeDamage" db:"critical_strike_damage"`
	CooldownReduction                                  float64     `json:"cooldownReduction" db:"cooldown_reduction"`
	PreventCriticalStrikeDamage                        int         `json:"preventCriticalStrikeDamaged" db:"prevent_critical_strike_damage"`
	CooldownLimit                                      int         `json:"cooldownLimit" db:"cooldown_limit"`
	LifeSteal                                          float64     `json:"lifeSteal" db:"lifesteal"`
	NormalLifeSteal                                    float64     `json:"normalLifeSteal" db:"normal_lifesteal"`
	SkillLifeSteal                                     float64     `json:"skillLifeSteal" db:"skill_lifesteal"`
	MoveSpeed                                          float64     `json:"moveSpeed" db:"move_speed"`
	MoveSpeedOutOfCombat                               float64     `json:"moveSpeedOutOfCombat" db:"move_speed_out_of_combat"`
	SightRange                                         float64     `json:"sightRange" db:"sight_range"`
	AttackRange                                        float64     `json:"attackRange" db:"attack_range"`
	IncreaseBasicAttackDamage                          int         `json:"increaseBasicAttackDamage" db:"increase_basic_attack_damage"`
	IncreaseBasicAttackDamageByLv                      int         `json:"increaseBasicAttackDamageByLv" db:"increase_basic_attack_damage_by_lv"`
	IncreaseBasicAttackDamageRatio                     float64     `json:"increaseBasicAttackDamageRatio" db:"increase_basic_attack_damage_ratio"`
	IncreaseBasicAttackDamageRatioByLv                 float64     `json:"increaseBasicAttackDamageRatioByLv" db:"increase_basic_attack_damage_ratio_by_lv"`
	PreventBasicAttackDamaged                          int         `json:"preventBasicAttackDamaged" db:"prevent_basic_attack_damage"`
	PreventBasicAttackDamagedByLv                      int         `json:"preventBasicAttackDamagedByLv" db:"prevent_basic_attack_damage_by_lv"`
	PreventBasicAttackDamagedRatio                     float64     `json:"preventBasicAttackDamagedRatio" db:"prevent_basic_attack_damage_ratio"`
	PreventBasicAttackDamagedRatioByLv                 float64     `json:"preventBasicAttackDamagedRatioByLv" db:"prevent_basic_attack_damage_ratio_by_lv"`
	PreventSkillDamage                                 int         `json:"preventSkillDamaged" db:"prevent_skill_damage"`
	PreventSkillDamageByLv                             int         `json:"preventSkillDamagedByLv" db:"prevent_skill_damage_by_lv"`
	PreventSkillDamageRatio                            float64     `json:"preventSkillDamagedRatio" db:"prevent_skill_damage_ratio"`
	PreventSkillDamageRatioByLv                        float64     `json:"preventSkillDamagedRatioByLv" db:"prevent_skill_damage_ratio_by_lv"`
	PenetrationDefense                                 int         `json:"penetrationDefense" db:"penetration_defense"`
	PenetrationDefenseRatio                            float64     `json:"penetrationDefenseRatio" db:"penetration_defense_ratio"`
	TrapDamageReduce                                   int         `json:"trapDamageReduce" db:"trap_damage_reduce"`
	TrapDamageReduceRatio                              float64     `json:"trapDamageReduceRatio" db:"trap_damage_reduce_ratio"`
	HpHealedIncreaseRatio                              float64     `json:"hpHealedIncreaseRatio" db:"hp_healed_increase_ratio"`
	HealerGiveHpHealRatio                              float64     `json:"healerGiveHpHealRatio" db:"healer_give_hp_heal_ratio"`
	UniqueAttackRange                                  float64     `json:"uniqueAttackRange" db:"unique_attack_range"`
	UniqueHpHealedIncreaseRatio                        float64     `json:"uniqueHpHealedIncreaseRatio" db:"unique_hp_healed_increase_ratio"`
	UniqueCooldownLimit                                float64     `json:"uniqueCooldownLimit" db:"unique_cooldown_limit"`
	UniqueTenacity                                     float64     `json:"uniqueTenacity" db:"unique_tenacity"`
	UniqueMoveSpeed                                    float64     `json:"uniqueMoveSpeed" db:"unique_move_speed"`
	UniquePenetrationDefense                           int         `json:"uniquePenetrationDefense" db:"unique_penetration_defense"`
	UniquePenetrationDefenseRatio                      float64     `json:"uniquePenetrationDefenseRatio" db:"unique_penetration_defense_ratio"`
	UniqueLifeSteal                                    float64     `json:"uniqueLifeSteal" db:"unique_lifesteal"`
	UniqueSkillAmpRatio                                float64     `json:"uniqueSkillAmpRatio" db:"unique_skill_amp_ratio"`
	RestoreItemWhenResurrected                         bool        `json:"restoreItemWhenResurrected" db:"restore_item_when_resurrected"`
	CreditValueWhenConvertedToBounty                   int         `json:"creditValueWhenConvertedToBounty" db:"credit_value_when_converted_to_bounty"`
}

func (ip *ItemWeapon) Prepare() {
	switch ip.ExclusiveProducer.(type) {
	case int:
		val := ip.ExclusiveProducer.(int)
		ip.ExclusiveProducerArray = pq.Array(val)
	case string:
		values := strings.Split(ip.ExclusiveProducer.(string), ",")
		var result []int
		for _, e := range values {
			res, err := strconv.Atoi(e)
			if err != nil {
				continue
			} // TODO: log
			result = append(result, res)
		}
		ip.ExclusiveProducerArray = pq.Array(result)
	}
}

type Level struct {
	Id            int64 `db:"id"`
	Level         int   `json:"level" db:"level"`
	NeedExp       int   `json:"needExp" db:"needed_exp"`
	AccumulateExp int   `json:"accumulateExp" db:"accumulate_exp"`
	RewardAcoin   int   `json:"rewardAcoin" db:"reward_acoin"`
	Reward        int   `json:"reward" db:"reward"`
}

type LoadingTip struct {
	Id             int64  `db:"id"`
	Code           int    `json:"code" db:"code"`
	LoadingTipType string `json:"loadingTipType" db:"loading_tip_type"`
	MinLv          int    `json:"minLv" db:"min_lv"`
	MaxLv          int    `json:"maxLv" db:"max_lv"`
	TextKey        string `json:"textKey" db:"text_key"`
	ImageName      string `json:"imageName" db:"image_name"`
}

type MasteryExp struct {
	Id             int64  `db:"id"`
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
	Id                int64  `db:"id"`
	Code              int    `json:"code" db:"code"`
	Type              string `json:"type" db:"type"`
	MasteryLevel      int    `json:"masteryLevel" db:"mastery_level"`
	NextMasteryExp    int    `json:"nextMasteryExp" db:"next_mastery_exp"`
	GiveLevelExp      int    `json:"giveLevelExp" db:"give_level_exp"`
	ExpGrowthCapRatio int    `json:"expGrowthCapRatio" db:"exp_growth_cap_ratio"`
}

type MasteryStat struct {
	Id                        int64   `db:"id"`
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
	Id               int64   `db:"id"`
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
	Id           int64 `db:"id"`
	MonsterCode  int   `json:"monsterCode" db:"monster_code"`
	MonsterLevel int   `json:"monsterLevel" db:"monster_level"`
	DropGroup    int   `json:"dropGroup" db:"drop_group"`
}

type MonsterLevelUpStat struct {
	Id          int64   `db:"id"`
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
	Id          int64 `db:"id"`
	Code        int   `json:"code" db:"code"`
	Mode        int   `json:"mode" db:"mode"`
	PlayerLevel int   `json:"playerLevel" db:"player_level"`
	MonsterCode int   `json:"monsterCode" db:"monster_code"`
	SpawnLevel  int   `json:"spawnLevel" db:"spawn_level"`
}

type NaviCollectAndHunt struct {
	Id           int64  `db:"id"`
	Code         int    `json:"code" db:"code"`
	ItemCode     int    `json:"itemCode" db:"item_code"`
	AreaCodeList string `json:"areaCodeList" db:"area_code_list"`
}

type NearByArea struct {
	Id             int64 `db:"id"`
	Code           int   `json:"code" db:"code"`
	AreaCode       int   `json:"areaCode" db:"area_code"`
	NearByAreaCode int   `json:"nearByAreaCode" db:"near_by_area_code"`
}

type RandomEquipment struct {
	Id            int64       `db:"id"`
	Code          int         `json:"code" db:"code"`
	Group         string      `json:"group" db:"group"`
	ItemCode      interface{} `json:"itemcode"` // string or number but always seems to be a number
	ItemCodeValue int         `db:"item_code"`
	Weight        int         `json:"weight" db:"weight"`
	ItemGrade     string      `json:"itemGrade" db:"item_grade"`
	TagMultiplier int         `json:"tagMultiplier" db:"tag_multiplier"`
	CharacterNum  int         `json:"characterNum" db:"character_num"`
}

func (ip *RandomEquipment) Prepare() {
	switch ip.ItemCode.(type) {
	case int:
		ip.ItemCodeValue = ip.ItemCode.(int)
	case string:
		res, err := strconv.Atoi(ip.ItemCode.(string))
		if err != nil {
			break
		}
		ip.ItemCodeValue = res
	}
}

type RecommendedList struct {
	Id                    int64  `db:"id"`
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
	Id              int64     `db:"id"`
	SeasonID        int       `json:"seasonID" db:"season_id"`
	SeasonName      string    `json:"seasonName" db:"season_name"` // no tags as they need to be converted to date
	SeasonStart     string    `json:"seasonStart"`
	SeasonEnd       string    `json:"seasonEnd"`
	SeasonStartDate time.Time `db:"season_start"`
	SeasonEndDate   time.Time `db:"season_end"`
	IsCurrent       int       `json:"isCurrent" db:"is_current"`
}

func (s *Season) Prepare() {
	val, err := time.Parse("2006-01-02 03:04:05", s.SeasonStart)
	if err == nil {
		s.SeasonStartDate = val
	}

	val, err = time.Parse("2006-01-02 03:04:05", s.SeasonEnd)
	if err == nil {
		s.SeasonEndDate = val
	}
}

type SummonObjectStat struct {
	Id                   int64   `db:"id"`
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
	Id              int64 `db:"id"`
	Code            int   `json:"code" db:"code"`
	NextUpgradeCode int   `json:"nextUpgradecode" db:"next_upgrade_code"`
	UpgradeCredit   int   `json:"upgradeCredit" db:"upgrade_credit"`
	UpgradeMaterial int   `json:"upgradeMaterial" db:"upgrade_material"`
	SkillCode       int   `json:"skillCode" db:"skill_code"`
}

type TacticalSkillSetGroup struct {
	Id             int64  `db:"id"`
	Group          int    `json:"group" db:"group"`
	ModeType       int    `json:"modeType" db:"mode_type"`
	StartCode      int    `json:"startCode" db:"start_code"`
	EquipWithStart bool   `json:"equipWithStart" db:"equip_with_start"`
	Icon           string `json:"icon" db:"icon"`
}

type Trait struct {
	Id            int64  `db:"id"`
	Code          int    `json:"code" db:"code"`
	OpenAccountLv int    `json:"openAccountLv" db:"open_account_lv"`
	TraitGroup    string `json:"traitGroup" db:"trait_group"`
	TraitType     string `json:"traitType" db:"trait_type"`
	Active        bool   `json:"active" db:"active"`
}

type TransferConsole struct {
	Id                                 int64  `db:"id"`
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
	Id             int64   `db:"id"`
	Code           int     `json:"code" db:"code"`
	Mode           int     `json:"mode" db:"mode"`
	Phase          int     `json:"phase" db:"phase"`
	ConditionType  string  `json:"conditionType" db:"condition_type"`
	ConditionValue int     `json:"conditionValue" db:"condition_value"`
	AcquireSelf    float64 `json:"acquireSelf" db:"acquire_self"`
	AcquireTeam    int     `json:"acquireTeam" db:"acquire_team"`
}

type WeaponTypeInfo struct {
	Id                    int64   `db:"id"`
	Type                  string  `json:"type" db:"type"`
	AttackSpeed           float64 `json:"attackSpeed" db:"attack_speed"`
	AttackRange           float64 `json:"attackRange" db:"attack_range"`
	ShopFilter            int     `json:"shopFilter" db:"shop_filter"`
	SummonObjectHitDamage int     `json:"summonObjectHitDamage" db:"summon_object_hit_damage"`
}

// Item/Name/{int}
// Item/Help/{int}
// Item/Effect/{int}
type ItemEnglishInfo struct {
	Code        int    `db:"item_code"`
	Name        string `db:"name_english"`
	Description string `db:"description_english"`
	Effect      string `db:"effect_english"`
}

// Trait/Name/{int}
// Trait/StatTooltip/{int}
// Infusion/Trait/Desc/{int}
type TraitEnglishInfo struct {
	Code        int    `db:"code"`
	Name        string `db:"name_english"`
	StatTooltip string `db:"stat_tooltip_english"`
	Description string `db:"description_english"`
}

// Character/Name/{int}
// CharacterStoryName/{int}
// CharacterStoryAge/{int}
// CharacterStoryHeight/{int}
// CharacterStoryGender/{int}
// CharacterStoryDesc/{int}
// CharacterStoryTitle/{int}
type CharacterEnglishInfo struct {
	Code            int    `db:"code"`
	Name            string `db:"name_english"`
	LoreName        string `db:"lore_name_english"`
	LoreHeight      string `db:"lore_height_english"`
	LoreAge         string `db:"lore_age_english"`
	LoreTitle       string `db:"lore_title_english"`
	LoreDescription string `db:"lore_description_english"`
	LoreGender      string `db:"lore_gender_english"`
}

// Monster/Name/{int}
type MonsterEnglishInfo struct {
	Code int    `db:"code"`
	Name string `db:"name_english"`
}

// Area/Name/{int}
type AreaEnglishInfo struct {
	Code int    `db:"code"`
	Name string `db:"name_english"`
}

// ItemType/{string}
type ItemTypeEnglishInfo struct {
	Type string `db:"item_type"`
	Name string `db:"item_type_english"`
}

// WeaponType/{string}
type WeaponTypeEnglishInfo struct {
	Type string `db:"weapon_type"`
	Name string `db:"weapon_type_english"`
}

// ArmorType/{string}
type ArmorTypeEnglishInfo struct {
	Type string `db:"armor_type"`
	Name string `db:"armor_type_english"`
}

// SpecialItemType/{string}
type SpecialItemTypeEnglishInfo struct {
	Type string `db:"special_item_type"`
	Name string `db:"special_item_type_english"`
}

// ItemConsumableType/{string}
type ItemConsumableTypeEnglishInfo struct {
	Type string `db:"consumable_type"`
	Name string `db:"consumable_type_english"`
}

// MasteryType/{string}
type MasteryTypeEnglishInfo struct {
	Type string `db:"type"`
	Name string `db:"type_english"`
}

// ItemGrade/{string}
type ItemGradeEnglishInfo struct {
	Grade string `db:"item_grade"`
	Name  string `db:"item_grade_english"`
}

// MasteryConditionType/{string}
type MasteryConditionTypeEnglishInfo struct {
	Type string `db:"condition_type"`
	Name string `db:"condition_type_english"`
}

// StatType/{string}
// StatType/{string}/Desc
type StatTypeDescriptionEnglish struct {
	Type        string `db:"stat_type"`
	Name        string `db:"name_english"`
	Description string `db:"description_english"`
}

// ToolTipType/{string}
type ToolTipTypeEnglish struct {
	Type        string `db:"tool_tip_type"`
	Description string `db:"description"`
}

// SummonData/Name/{int}
type SummonDataEnglish struct {
	Code int    `db:"code"`
	Name string `db:"name_english"`
}

// Skin/Name/{int}
// Skin/Desc/{int}
type SkinEnglish struct {
	Code        int    `db:"code"`
	Name        string `db:"name_english"`
	Description string `db:"description"`
}

// Skill/Group/Name/{int}
// Skill/Group/Desc/{int}
// Skill/Group/Coef/{int}
// Skill/Group/ExpansionTip/{int}
type SkillGroupEnglish struct {
	Code         int    `db:"code"`
	Name         string `db:"name_english"`
	Description  string `db:"description_english"`
	Coefficient  string `db:"coefficient_english"`
	ExpansionTip string `db:"expansion_tip"`
}

// Skill/LobbyDesc/{int}
// Skill/LobbyName/{int}
type SkillEnglish struct {
	Code        int    `db:"code"`
	Name        string `db:"name_english"`
	Description string `db:"description_english"`
}

// Skill/Code/Name/{int}
// Skill/Code/Desc/{int}
// Skill/Code/Coef/{int}
// Skill/Code/ExpansionTip/{int}
type TacticalSkillEnglish struct {
	Code         int    `db:"code"`
	Name         string `db:"name_english"`
	Description  string `db:"description_english"`
	Coefficient  string `db:"coefficient_english"`
	ExpansionTip string `db:"expansion_tip"`
}

// CharacterState/Group/Name/{int}
// CharacterState/Group/Desc/{int}
type CharacterStateEnglish struct {
	Code        int    `db:"code"`
	Name        string `db:"name_english"`
	Description string `db:"description_english"`
}

// Item/Skills/{int} <-- item_name
// Item/Skills/{int}/Body
// Item/Skills/{int}/Name
// Item/Skills/{int}/Desc
type ItemSkillEnglish struct {
	Code                 int    `db:"code"`
	ItemName             string `db:"item_name"`
	ItemSkillName        string `db:"item_skill_name_english"`
	ItemSkillBody        string `db:"item_skill_body_english"`
	ItemSkillDescription string `db:"item_skill_description_english"`
}
