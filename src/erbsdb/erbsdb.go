package erbsdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sleepy-day/ERBS-BE/src/models"
	pb "github.com/sleepy-day/ERBS-BE/src/protos"
	"github.com/sleepy-day/ERBS-BE/src/util"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const DATA_URL = "https://open-api.bser.io/v2/data/"
const LANGUAGE_URL = "https://open-api.bser.io/v1/l10n/"

var db *sqlx.DB
var client *http.Client
var limiter *rate.Limiter

var API_KEY ApiKey

func InitDb(connStr string, apiKey string) {
	var err error
	if db, err = sqlx.Connect("postgres", connStr); err != nil {
		panic(err)
	}

	if apiKey == "" {
		panic("Empty api key, unable to continue")
	}

	API_KEY = setApiKey(apiKey)

	tr := &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       15 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		DisableKeepAlives:     false,
	}
	client = &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	limiter = rate.NewLimiter(rate.Every(1200*time.Millisecond), 1)
}

func PopulateGameDataTables() error {
	tx := db.MustBegin()
	{
		fmt.Println("Processing Table: " + "ActionCost")
		req, err := createHttpRequest(DATA_URL + "ActionCost")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ActionCost]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "action_cost", cols, vals)
		truncateQuery := "TRUNCATE TABLE action_cost;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Area")
		req, err := createHttpRequest(DATA_URL + "Area")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Area]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "area", cols, vals)
		truncateQuery := "TRUNCATE TABLE area;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "BattleZoneReward")
		req, err := createHttpRequest(DATA_URL + "BattleZoneReward")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.BattleZoneReward]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "battle_zone_reward", cols, vals)
		truncateQuery := "TRUNCATE TABLE battle_zone_reward;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "BulletCapacity")
		req, err := createHttpRequest(DATA_URL + "BulletCapacity")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.BulletCapacity]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "bullet_capacity", cols, vals)
		truncateQuery := "TRUNCATE TABLE bullet_capacity;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Character")
		req, err := createHttpRequest(DATA_URL + "Character")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Character]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character", cols, vals)
		truncateQuery := "TRUNCATE TABLE character;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterAttributes")
		req, err := createHttpRequest(DATA_URL + "CharacterAttributes")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterAttributes]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_attributes", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_attributes;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterExp")
		req, err := createHttpRequest(DATA_URL + "CharacterExp")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterExp]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_exp", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_exp;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterLevelUpStat")
		req, err := createHttpRequest(DATA_URL + "CharacterLevelUpStat")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterLevelUpStat]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_level_up_stat", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_level_up_stat;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterMastery")
		req, err := createHttpRequest(DATA_URL + "CharacterMastery")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterMastery]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_mastery", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_mastery;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterModeModifier")
		req, err := createHttpRequest(DATA_URL + "CharacterModeModifier")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterModeModifier]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_mode_modifier", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_mode_modifier;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "CharacterSkin")
		req, err := createHttpRequest(DATA_URL + "CharacterSkin")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.CharacterSkin]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "character_skin", cols, vals)
		truncateQuery := "TRUNCATE TABLE character_skin;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Collectible")
		req, err := createHttpRequest(DATA_URL + "Collectible")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Collectible]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "collectible", cols, vals)
		truncateQuery := "TRUNCATE TABLE collectible;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "DropGroup")
		req, err := createHttpRequest(DATA_URL + "DropGroup")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.DropGroup]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "drop_group", cols, vals)
		truncateQuery := "TRUNCATE TABLE drop_group;"

		for _, v := range response.Data {
			v.Prepare()
		}

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "GainExp")
		req, err := createHttpRequest(DATA_URL + "GainExp")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.GainExp]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "gain_exp", cols, vals)
		truncateQuery := "TRUNCATE TABLE gain_exp;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "GainScore")
		req, err := createHttpRequest(DATA_URL + "GainScore")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.GainScore]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "gain_score", cols, vals)
		truncateQuery := "TRUNCATE TABLE gain_score;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "GameTip")
		req, err := createHttpRequest(DATA_URL + "GameTip")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.GameTip]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "game_tip", cols, vals)
		truncateQuery := "TRUNCATE TABLE game_tip;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "HowToFindItem")
		req, err := createHttpRequest(DATA_URL + "HowToFindItem")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.HowToFindItem]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "how_to_find_item", cols, vals)
		truncateQuery := "TRUNCATE TABLE how_to_find_item;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "InfusionProduct")
		req, err := createHttpRequest(DATA_URL + "InfusionProduct")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.InfusionProduct]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "infusion_product", cols, vals)
		truncateQuery := "TRUNCATE TABLE infusion_product;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemArmor")
		req, err := createHttpRequest(DATA_URL + "ItemArmor")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemArmor]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "items", cols, vals)
		truncateQuery := "TRUNCATE TABLE items;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}

		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemConsumable")
		req, err := createHttpRequest(DATA_URL + "ItemConsumable")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemConsumable]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "items", cols, vals)
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemMisc")
		req, err := createHttpRequest(DATA_URL + "ItemMisc")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemMisc]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "items", cols, vals)
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemSearchOptionV2")
		req, err := createHttpRequest(DATA_URL + "ItemSearchOptionV2")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemSearchOption]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "item_search_option", cols, vals)
		truncateQuery := "TRUNCATE TABLE item_search_option;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemSpawn")
		req, err := createHttpRequest(DATA_URL + "ItemSpawn")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemSpawn]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "item_spawn", cols, vals)
		truncateQuery := "TRUNCATE TABLE item_spawn;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemSpecial")
		req, err := createHttpRequest(DATA_URL + "ItemSpecial")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemSpecial]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "items", cols, vals)
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "ItemWeapon")
		req, err := createHttpRequest(DATA_URL + "ItemWeapon")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.ItemWeapon]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "items", cols, vals)
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Level")
		req, err := createHttpRequest(DATA_URL + "Level")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Level]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "level", cols, vals)
		truncateQuery := "TRUNCATE TABLE level;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "LoadingTip")
		req, err := createHttpRequest(DATA_URL + "LoadingTip")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.LoadingTip]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "loading_tip", cols, vals)
		truncateQuery := "TRUNCATE TABLE loading_tip;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MasteryExp")
		req, err := createHttpRequest(DATA_URL + "MasteryExp")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MasteryExp]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "mastery_exp", cols, vals)
		truncateQuery := "TRUNCATE TABLE mastery_exp;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MasteryLevel")
		req, err := createHttpRequest(DATA_URL + "MasteryLevel")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MasteryLevel]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "mastery_levels", cols, vals)
		truncateQuery := "TRUNCATE TABLE mastery_levels;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MasteryStat")
		req, err := createHttpRequest(DATA_URL + "MasteryStat")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MasteryStat]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "mastery_stat", cols, vals)
		truncateQuery := "TRUNCATE TABLE mastery_stat;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Monster")
		req, err := createHttpRequest(DATA_URL + "Monster")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Monster]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "monster", cols, vals)
		truncateQuery := "TRUNCATE TABLE monster;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MonsterDropGroup")
		req, err := createHttpRequest(DATA_URL + "MonsterDropGroup")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MonsterDropGroup]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "monster_drop_group", cols, vals)
		truncateQuery := "TRUNCATE TABLE monster_drop_group;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MonsterLevelUpStat")
		req, err := createHttpRequest(DATA_URL + "MonsterLevelUpStat")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MonsterLevelUpStat]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "monster_level_up_stat", cols, vals)
		truncateQuery := "TRUNCATE TABLE monster_level_up_stat;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "MonsterSpawnLevel")
		req, err := createHttpRequest(DATA_URL + "MonsterSpawnLevel")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.MonsterSpawnLevel]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "monster_spawn_level", cols, vals)
		truncateQuery := "TRUNCATE TABLE monster_spawn_level;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "NaviCollectAndHunt")
		req, err := createHttpRequest(DATA_URL + "NaviCollectAndHunt")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.NaviCollectAndHunt]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "navi_collect_and_hunt", cols, vals)
		truncateQuery := "TRUNCATE TABLE navi_collect_and_hunt;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "NearByArea")
		req, err := createHttpRequest(DATA_URL + "NearByArea")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.NearByArea]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "near_by_area", cols, vals)
		truncateQuery := "TRUNCATE TABLE near_by_area;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "RandomEquipment")
		req, err := createHttpRequest(DATA_URL + "RandomEquipment")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.RandomEquipment]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "random_equipment", cols, vals)
		truncateQuery := "TRUNCATE TABLE random_equipment;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "RecommendedList")
		req, err := createHttpRequest(DATA_URL + "RecommendedList")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.RecommendedList]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "recommended_list", cols, vals)
		truncateQuery := "TRUNCATE TABLE recommended_list;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Season")
		req, err := createHttpRequest(DATA_URL + "Season")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Season]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		for _, v := range response.Data {
			v.Prepare()
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "season", cols, vals)
		truncateQuery := "TRUNCATE TABLE season;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "SummonObjectStat")
		req, err := createHttpRequest(DATA_URL + "SummonObjectStat")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.SummonObjectStat]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "summon_object_stat", cols, vals)
		truncateQuery := "TRUNCATE TABLE summon_object_stat;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "TacticalSkillSet")
		req, err := createHttpRequest(DATA_URL + "TacticalSkillSet")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.TacticalSkillSet]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "tactical_skill_set", cols, vals)
		truncateQuery := "TRUNCATE TABLE tactical_skill_set;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "TacticalSkillSetGroup")
		req, err := createHttpRequest(DATA_URL + "TacticalSkillSetGroup")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.TacticalSkillSetGroup]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "tactical_skill_set_group", cols, vals)
		truncateQuery := "TRUNCATE TABLE tactical_skill_set_group;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "Trait")
		req, err := createHttpRequest(DATA_URL + "Trait")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.Trait]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "trait", cols, vals)
		truncateQuery := "TRUNCATE TABLE trait;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "TransferConsole")
		req, err := createHttpRequest(DATA_URL + "TransferConsole")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.TransferConsole]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "transfer_console", cols, vals)
		truncateQuery := "TRUNCATE TABLE transfer_console;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "VFCredit")
		req, err := createHttpRequest(DATA_URL + "VFCredit")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.VfCredit]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "vf_credit", cols, vals)
		truncateQuery := "TRUNCATE TABLE vf_credit;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	{
		fmt.Println("Processing Table: " + "WeaponTypeInfo")
		req, err := createHttpRequest(DATA_URL + "WeaponTypeInfo")
		if err != nil {
			return err
		}
		body, err := makeHttpRequest(req)
		if err != nil {
			return err
		}
		response := &models.Response[[]models.WeaponTypeInfo]{}
		err = json.Unmarshal(*body, &response)
		if err != nil {
			return err
		}

		e := reflect.ValueOf(response.Data[0])
		cols, vals := getDbFields(e)

		query := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", "weapon_type_info", cols, vals)
		truncateQuery := "TRUNCATE TABLE weapon_type_info;"

		_, err = tx.Exec(truncateQuery)
		if err != nil {
			return err
		}
		_, err = tx.NamedExec(query, response.Data)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func createHttpRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", API_KEY.Get())

	return req, nil
}

func makeHttpRequest(req *http.Request) (*[]byte, error) {
	var err error
	var resp *http.Response

	for i := 0; i < 3; i++ {
		err = limiter.Wait(context.Background())
		if err != nil {
			return nil, err
		}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusTooManyRequests {
			break
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(rune(resp.StatusCode)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	return &body, nil
}

func getDbFields(v reflect.Value) (string, string) {
	columns := make([]string, 0)
	values := make([]string, 0)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			innerType := v.Field(i).Type()
			structTag := field.Tag.Get("db")
			for ii := 0; ii < v.Field(i).NumField(); ii++ {
				innerField := innerType.Field(ii)
				innerTag := innerField.Tag.Get("db")
				if innerTag == "" || innerTag == "id" {
					continue
				}
				values = append(values, ":"+innerTag)
				if innerTag == "group" {
					innerTag = "\"group\""
				}
				columns = append(columns, structTag+innerTag)
			}
		}

		tag := field.Tag.Get("db")
		if tag == "" || tag == "id" {
			continue
		}
		values = append(values, ":"+tag)
		if tag == "group" {
			tag = "\"group\""
		}
		columns = append(columns, tag)
	}
	return "(" + strings.Join(columns, ", ") + ")", "(" + strings.Join(values, ", ") + ")"
}

func PopulateTablesFromLangFile(info util.EnglishInfo) error {
	tx := db.MustBegin()
	{
		fmt.Printf("Processing struct %s\n", "ItemEnglishInfo")
		query :=
			`UPDATE 
				items 
			SET 
				name_english = :name_english, 
				description_english = :description_english, 
				effect_english = :effect_english 
			WHERE 
				code = :code;`
		for _, v := range info.ItemInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "TraitEnglishInfo")
		query :=
			`UPDATE 
				trait 
			SET 
				name_english = :name_english, 
				description_english = :description_english, 
				stat_tooltip_english = :stat_tooltip_english
			WHERE 
				code = :code;`
		for _, v := range info.TraitInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "CharacterEnglishInfo")
		query :=
			`UPDATE 
			character 
		SET 
			lore_name_english = :lore_name_english, 
			lore_description_english = :lore_description_english, 
			lore_height_english = :lore_height_english,
			lore_age_english = :lore_age_english,
			lore_gender_english = :lore_gender_english,
			lore_title_english = :lore_title_english
		WHERE
		    code = :code;`
		for _, v := range info.CharacterInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "MonsterEnglishInfo")
		query :=
			`UPDATE 
			monster 
		SET 
		    name_english = :name_english
		WHERE
		    code = :code;`
		for _, v := range info.MonsterInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ItemTypeEnglishInfo")
		query :=
			`UPDATE 
				items
			SET 
		    	item_type_english = :item_type_english
			WHERE
		    	item_type = :item_type;`
		for _, v := range info.ItemTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "WeaponTypeEnglishInfo")
		query :=
			`UPDATE 
				items 
			SET 
		    	weapon_type_english = :weapon_type_english
			WHERE
		    	weapon_type = :weapon_type;`
		for _, v := range info.WeaponTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ArmorTypeEnglishInfo")
		query :=
			`UPDATE 
				items
			SET 
		    	armor_type_english = :armor_type_english
			WHERE
		    	armor_type = :armor_type;`
		for _, v := range info.ArmorTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "SpecialItemTypeEnglishInfo")
		query :=
			`UPDATE 
				items 
			SET 
		    	special_item_type_english = :special_item_type_english
			WHERE
		    	special_item_type = :special_item_type;`
		for _, v := range info.SpecialItemTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ItemConsumableTypeEnglishInfo")
		query :=
			`UPDATE 
				items 
			SET 
		    	consumable_type_english = :consumable_type_english
			WHERE
		    	consumable_type = :consumable_type;`
		for _, v := range info.ItemConsumableInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "MasteryTypeEnglishInfo")
		query :=
			`UPDATE 
				mastery_stat 
			SET 
		    	type_english = :type_english
			WHERE
		    	type = :type;`
		for _, v := range info.MasteryTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ItemGradeEnglishInfo")
		query :=
			`UPDATE 
				items 
			SET 
		    	item_grade_english = :item_grade_english
			WHERE
			    item_grade = :item_grade;`
		for _, v := range info.ItemGradeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "MasteryConditionTypeEnglishInfo")
		query :=
			`UPDATE 
				 mastery_exp
			SET 
			    condition_type_english = :condition_type_english
			WHERE
		    	condition_type = :condition_type;`
		for _, v := range info.MasteryConditionTypeInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "StatTypeDescriptionEnglish")
		query :=
			`INSERT INTO 
				stat_type (stat_type, name_english, description_english) 
			VALUES  
				(:stat_type, :name_english, :description_english);`

		slc := make([]models.StatTypeDescriptionEnglish, len(info.StatTypeInfo))
		for _, v := range info.StatTypeInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ToolTipTypeEnglish")
		query :=
			`INSERT INTO 
				tooltip_type (tool_tip_type, description_english)  
			VALUES 
				(:tool_tip_type, :description_english);`
		slc := make([]models.ToolTipTypeEnglish, len(info.ToolTipTypeInfo))
		for _, v := range info.ToolTipTypeInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}
	{
		fmt.Printf("Processing struct %s\n", "SummonDataEnglish")
		query :=
			`UPDATE 
				summon_object_stat 
			SET 
		    	name_english = :name_english
			WHERE
		    	code = :code;`
		for _, v := range info.SummonInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "SkinEnglish")
		query :=
			`UPDATE 
				character_skin 
			SET 
		    	name_english = :name_english,
				description_english = :description_english
			WHERE
		    	code = :code;`
		for _, v := range info.SkinInfo {
			_, err := tx.NamedExec(query, v)
			if err != nil {
				return err
			}
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "SkillGroupInfo")
		query :=
			`INSERT INTO 
				skill_group (code, name_english, description_english, coefficient_english, expansion_tip_english) 
			VALUES 
				(:code, :name_english, :description_english, :coefficient_english, :expansion_tip_english);`
		slc := make([]models.SkillGroupEnglish, len(info.SkillGroupInfo))
		for _, v := range info.SkillGroupInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}
	{
		fmt.Printf("Processing struct %s\n", "SkillEnglish")
		query :=
			`INSERT INTO 
				skill (code, name_english, description_english) 
			VALUES 
				(:code, :name_english, :description_english);`
		slc := make([]models.SkillEnglish, len(info.SkillInfo))
		for _, v := range info.SkillInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "TacticalSkillEnglish")
		query :=
			`INSERT INTO 
				tactical_skill (code, name_english, description_english, coefficient_english, expansion_tip_english) 
			VALUES 
				(:code, :name_english, :description_english, :coefficient_english, :expansion_tip_english);`
		slc := make([]models.TacticalSkillEnglish, len(info.TacticalSkillInfo))
		for _, v := range info.TacticalSkillInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "CharacterStateEnglish")
		query :=
			`INSERT INTO 
				character_state (code, name_english, description_english) 
			VALUES 
				(:code, :name_english, :description_english);`
		slc := make([]models.CharacterStateEnglish, len(info.CharacterStateInfo))
		for _, v := range info.CharacterStateInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}

	{
		fmt.Printf("Processing struct %s\n", "ItemSkillEnglish")
		query :=
			`INSERT INTO 
				item_skill (code, item_name_english, item_skill_name_english, item_skill_body_english, item_skill_description_english) 
			VALUES 
				(:code, :item_name_english, :item_skill_name_english, :item_skill_body_english, :item_skill_description_english);`
		slc := make([]models.ItemSkillEnglish, len(info.ItemSkillInfo))
		for _, v := range info.ItemSkillInfo {
			slc = append(slc, v)
		}
		_, err := tx.NamedExec(query, slc)

		if err != nil {
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateLanguageFiles() {
	req, err := createHttpRequest(LANGUAGE_URL + "English")
	if err != nil {
		panic(err)
	}
	body, err := makeHttpRequest(req)
	if err != nil {
		panic(err)
	}

	response := &models.Response[models.Language]{}
	err = json.Unmarshal(*body, &response)
	if err != nil {
		panic(err)
	}

	req, err = createHttpRequest(response.Data.L10Path)
	if err != nil {
		panic(err)
	}
	body, err = makeHttpRequest(req)
	if err != nil {
		panic(err)
	}

	exists := directoryExists("../lang")
	if exists {
		err = os.WriteFile("../lang/English.txt", *body, 0644)
		if err != nil {
			panic(err)
		}
		return
	}

	exists = directoryExists("./lang")
	if exists {
		err = os.WriteFile("./lang/English.txt", *body, 0644)
		if err != nil {
			panic(err)
		}
		return
	}

	panic("Unable to find lang directory")
}

func directoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func GetCharacters() ([]*pb.CharacterEntry, error) {
	query := `
	SELECT
		name,
		code
	FROM 
		character;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	characters := make([]*pb.CharacterEntry, 0)
	for rows.Next() {
		var name string
		var code int32
		err = rows.Scan(&name, &code)
		if err != nil {
			return nil, err
		}
		characters = append(characters, &pb.CharacterEntry{Name: name, Code: code})
	}

	return characters, nil
}

func GetCharacterProfile(code int32) (*pb.CharacterProfile, error) {
	query := `
	SELECT 
		code,
		max_hp,
		max_sp,
		init_extra_point,
		max_extra_point,
		attack_power,
		defense,
		hp_regen,
		sp_regen,
		attack_speed,
		attack_speed_limit,
		attack_speed_min,
		move_speed,
		sight_range,
		radius,
		pathing_radius,
		resource,
		name,
		str_learn_start_skill,
		str_use_point_learn_start_skill,
		lore_name_english,
		lore_height_english,
		lore_age_english,
		lore_title_english,
		lore_description_english,
		lore_gender_english
	FROM
		character
	WHERE
		code = $1;`

	row, err := db.Queryx(query, code)
	if err != nil {
		return nil, err
	}
	row.Next()
	c := &pb.CharacterProfile{}
	err = row.Scan(&c.Code, &c.MaxHp, &c.MaxSp, &c.UniqueResourceInitValue, &c.UniqueResourceMaxValue, &c.AttackPower, &c.Defense, &c.HpRegen,
		&c.SpRegen, &c.AttackSpeed, &c.AttackSpeedLimit, &c.AttackSpeedMin, &c.MoveSpeed, &c.SightRange, &c.SizeRadius, &c.PathingRadius,
		&c.ResourceName, &c.CharacterName, &c.AvailableStartingSkills, &c.StartsWithSkill, &c.FullName, &c.Height, &c.Age, &c.Title, &c.Description,
		&c.Gender)
	if err != nil {
		return nil, err
	}

	return c, nil
}
