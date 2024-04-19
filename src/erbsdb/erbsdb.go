package erbsdb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sleepy-day/ERBS-BE/src/models"
)

var db *sqlx.DB

func InitDb(connStr string) {
	var err error
	if db, err = sqlx.Connect("postgres", connStr); err != nil {
		panic(err)
	}

	games := []models.UserGame{}
	err = db.Select(&games, "SELECT * FROM user_game;")
	if err != nil {
		panic(err)
	}
	fmt.Println("user_game queried successfully\n")

	topRanks := []models.TopRank{}
	err = db.Select(&topRanks, "SELECT * FROM top_rank;")
	if err != nil {
		panic(err)
	}
	fmt.Println("top_rank queried successfully\n")

	userRanks := []models.UserRank{}
	err = db.Select(&userRanks, "SELECT * FROM user_rank;")
	if err != nil {
		panic(err)
	}
	fmt.Println("user_rank queried successfully\n")

	userInfo := []models.UserInfo{}
	err = db.Select(&userInfo, "SELECT * FROM user_info;")
	if err != nil {
		panic(err)
	}
	fmt.Println("user_info queried successfully\n")

	userStats := []models.UserStats{}
	err = db.Select(&userStats, "SELECT * FROM user_stats;")
	if err != nil {
		panic(err)
	}
	fmt.Println("user_stats queried successfully\n")

	weaponRoutes := []models.WeaponRoute{}
	err = db.Select(&weaponRoutes, "SELECT * FROM weapon_route;")
	if err != nil {
		panic(err)
	}
	fmt.Println("weapon_route queried successfully\n")

	actionCosts := []models.ActionCost{}
	err = db.Select(&actionCosts, "SELECT * FROM action_cost;")
	if err != nil {
		panic(err)
	}
	fmt.Println("action_cost queried successfully\n")

	areas := []models.Area{}
	err = db.Select(&areas, "SELECT * FROM area;")
	if err != nil {
		panic(err)
	}
	fmt.Println("areas queried successfully\n")

	battleZoneRewards := []models.BattleZoneReward{}
	err = db.Select(&battleZoneRewards, fmt.Sprintf("SELECT * FROM %s;", "battle_zone_reward"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "battle_zone_reward")

	bulletCapacities := []models.BulletCapacity{}
	err = db.Select(&bulletCapacities, fmt.Sprintf("SELECT * FROM %s;", "bullet_capacity"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "bullet_capacity")

	characters := []models.Character{}
	err = db.Select(&characters, fmt.Sprintf("SELECT * FROM %s;", "character"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character")

	characterAttributes := []models.CharacterAttributes{}
	err = db.Select(&characterAttributes, fmt.Sprintf("SELECT * FROM %s;", "character_attributes"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_attributes")

	characterExp := []models.CharacterExp{}
	err = db.Select(&characterExp, fmt.Sprintf("SELECT * FROM %s;", "character_exp"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_exp")

	characterLevelUpStats := []models.CharacterLevelUpStat{}
	err = db.Select(&characterLevelUpStats, fmt.Sprintf("SELECT * FROM %s;", "character_level_up_stat"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_level_up_stat")

	itemArmor := []models.ItemArmor{}
	err = db.Select(&itemArmor, fmt.Sprintf("SELECT * FROM %s;", "item_armor"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_armor")

	characterMasteries := []models.CharacterMastery{}
	err = db.Select(&characterMasteries, fmt.Sprintf("SELECT * FROM %s;", "character_mastery"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_mastery")

	characterModeModifiers := []models.CharacterModeModifier{}
	err = db.Select(&characterModeModifiers, fmt.Sprintf("SELECT * FROM %s;", "character_mode_modifier"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_mode_modifier")

	characterSkins := []models.CharacterSkin{}
	err = db.Select(&characterSkins, fmt.Sprintf("SELECT * FROM %s;", "character_skin"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "character_skin")

	collectibles := []models.Collectible{}
	err = db.Select(&collectibles, fmt.Sprintf("SELECT * FROM %s;", "collectible"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "collectible")

	dropGroups := []models.DropGroup{}
	err = db.Select(&dropGroups, fmt.Sprintf("SELECT * FROM %s;", "drop_group"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "drop_group")

	gainExp := []models.GainExp{}
	err = db.Select(&gainExp, fmt.Sprintf("SELECT * FROM %s;", "gain_exp"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "gain_exp")

	gainScore := []models.GainScore{}
	err = db.Select(&gainScore, fmt.Sprintf("SELECT * FROM %s;", "gain_score"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "gain_score")

	gameTips := []models.GameTip{}
	err = db.Select(&gameTips, fmt.Sprintf("SELECT * FROM %s;", "game_tip"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "game_tip")

	howToFindItems := []models.HowToFindItem{}
	err = db.Select(&howToFindItems, fmt.Sprintf("SELECT * FROM %s;", "how_to_find_items"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "how_to_find_items")

	infusionProducts := []models.InfusionProduct{}
	err = db.Select(&infusionProducts, fmt.Sprintf("SELECT * FROM %s;", "infusion_product"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "infusion_product")

	itemConsumables := []models.ItemConsumable{}
	err = db.Select(&itemConsumables, fmt.Sprintf("SELECT * FROM %s;", "item_consumable"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_consumable")

	itemMisc := []models.ItemMisc{}
	err = db.Select(&itemMisc, fmt.Sprintf("SELECT * FROM %s;", "item_misc"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_misc")

	itemSearchOptions := []models.ItemSearchOption{}
	err = db.Select(&itemSearchOptions, fmt.Sprintf("SELECT * FROM %s;", "item_search_option"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_search_option")

	itemSpawns := []models.ItemSpawn{}
	err = db.Select(&itemSpawns, fmt.Sprintf("SELECT * FROM %s;", "item_spawn"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_spawn")

	itemSpecial := []models.ItemSpecial{}
	err = db.Select(&itemSpecial, fmt.Sprintf("SELECT * FROM %s;", "item_special"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_special")

	itemWeapons := []models.ItemWeapon{}
	err = db.Select(&itemWeapons, fmt.Sprintf("SELECT * FROM %s;", "item_weapon"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "item_weapon")

	levels := []models.Level{}
	err = db.Select(&levels, fmt.Sprintf("SELECT * FROM %s;", "level"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "level")

	loadingTips := []models.LoadingTip{}
	err = db.Select(&loadingTips, fmt.Sprintf("SELECT * FROM %s;", "loading_tip"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "loading_tip")

	masteryExp := []models.MasteryExp{}
	err = db.Select(&masteryExp, fmt.Sprintf("SELECT * FROM %s;", "mastery_exp"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "mastery_exp")

	masteryLevels := []models.MasteryLevel{}
	err = db.Select(&masteryLevels, fmt.Sprintf("SELECT * FROM %s;", "mastery_levels"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "mastery_levels")

	masteryStats := []models.MasteryStat{}
	err = db.Select(&masteryStats, fmt.Sprintf("SELECT * FROM %s;", "mastery_stat"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "mastery_stat")

	monsters := []models.Monster{}
	err = db.Select(&monsters, fmt.Sprintf("SELECT * FROM %s;", "monster"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "monster")

	monsterDropGroups := []models.MonsterDropGroup{}
	err = db.Select(&monsterDropGroups, fmt.Sprintf("SELECT * FROM %s;", "monster_drop_group"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "monster_drop_group")

	monsterLevelUpStats := []models.MonsterLevelUpStat{}
	err = db.Select(&monsterLevelUpStats, fmt.Sprintf("SELECT * FROM %s;", "monster_level_up_stat"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "monster_level_up_stat")

	monsterSpawnLevels := []models.MonsterSpawnLevel{}
	err = db.Select(&monsterSpawnLevels, fmt.Sprintf("SELECT * FROM %s;", "monster_spawn_level"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "monster_spawn_level")

	naviCollectAndHunt := []models.NaviCollectAndHunt{}
	err = db.Select(&naviCollectAndHunt, fmt.Sprintf("SELECT * FROM %s;", "navi_collect_and_hunt"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "navi_collect_and_hunt")

	nearByAreas := []models.NearByArea{}
	err = db.Select(&nearByAreas, fmt.Sprintf("SELECT * FROM %s;", "near_by_area"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "near_by_area")

	randomEquipment := []models.RandomEquipment{}
	err = db.Select(&randomEquipment, fmt.Sprintf("SELECT * FROM %s;", "random_equipment"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "random_equipment")

	recommendedList := []models.RecommendedList{}
	err = db.Select(&recommendedList, fmt.Sprintf("SELECT * FROM %s;", "recommended_list"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "recommended_list")

	seasons := []models.Season{}
	err = db.Select(&seasons, fmt.Sprintf("SELECT * FROM %s;", "season"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "season")

	summonObjectStats := []models.SummonObjectStat{}
	err = db.Select(&summonObjectStats, fmt.Sprintf("SELECT * FROM %s;", "summon_object_stat"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "summon_object_stat")

	tacticalSkillSets := []models.TacticalSkillSet{}
	err = db.Select(&tacticalSkillSets, fmt.Sprintf("SELECT * FROM %s;", "tactical_skill_set"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "tactical_skill_set")

	tacticalSkillSetGroups := []models.TacticalSkillSetGroup{}
	err = db.Select(&tacticalSkillSetGroups, fmt.Sprintf("SELECT * FROM %s;", "tactical_skill_set_group"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "tactical_skill_set_group")

	traits := []models.Trait{}
	err = db.Select(&traits, fmt.Sprintf("SELECT * FROM %s;", "trait"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "trait")

	transferConsole := []models.TransferConsole{}
	err = db.Select(&transferConsole, fmt.Sprintf("SELECT * FROM %s;", "transfer_console"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "transfer_console")

	vfCredit := []models.VfCredit{}
	err = db.Select(&vfCredit, fmt.Sprintf("SELECT * FROM %s;", "vf_credit"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "vf_credit")

	weaponTypeInfo := []models.WeaponTypeInfo{}
	err = db.Select(&weaponTypeInfo, fmt.Sprintf("SELECT * FROM %s;", "weapon_type_info"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s areas queried successfully\n", "weapon_type_info")
}
