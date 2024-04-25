package util

import (
	"bufio"
	"github.com/sleepy-day/ERBS-BE/src/models"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type EnglishInfo struct {
	ItemInfo                 map[string]models.ItemEnglishInfo
	TraitInfo                map[string]models.TraitEnglishInfo
	CharacterInfo            map[string]models.CharacterEnglishInfo
	MonsterInfo              map[string]models.MonsterEnglishInfo
	AreaInfo                 map[string]models.AreaEnglishInfo
	ItemTypeInfo             map[string]models.ItemTypeEnglishInfo
	WeaponTypeInfo           map[string]models.WeaponTypeEnglishInfo
	ArmorTypeInfo            map[string]models.ArmorTypeEnglishInfo
	SpecialItemTypeInfo      map[string]models.SpecialItemTypeEnglishInfo
	ItemConsumableInfo       map[string]models.ItemConsumableTypeEnglishInfo
	MasteryTypeInfo          map[string]models.MasteryTypeEnglishInfo
	ItemGradeInfo            map[string]models.ItemGradeEnglishInfo
	MasteryConditionTypeInfo map[string]models.MasteryConditionTypeEnglishInfo
	StatTypeInfo             map[string]models.StatTypeDescriptionEnglish
	ToolTipTypeInfo          map[string]models.ToolTipTypeEnglish
	SkinInfo                 map[string]models.SkinEnglish
	SkillGroupInfo           map[string]models.SkillGroupEnglish
	SkillInfo                map[string]models.SkillEnglish
	TacticalSkillInfo        map[string]models.TacticalSkillEnglish
	CharacterStateInfo       map[string]models.CharacterStateEnglish
	ItemSkillInfo            map[string]models.ItemSkillEnglish
	SummonInfo               map[string]models.SummonDataEnglish
}

func initEnglishInfo() EnglishInfo {
	info := EnglishInfo{}
	info.ItemInfo = make(map[string]models.ItemEnglishInfo)
	info.TraitInfo = make(map[string]models.TraitEnglishInfo)
	info.CharacterInfo = make(map[string]models.CharacterEnglishInfo)
	info.MonsterInfo = make(map[string]models.MonsterEnglishInfo)
	info.AreaInfo = make(map[string]models.AreaEnglishInfo)
	info.ItemTypeInfo = make(map[string]models.ItemTypeEnglishInfo)
	info.WeaponTypeInfo = make(map[string]models.WeaponTypeEnglishInfo)
	info.ArmorTypeInfo = make(map[string]models.ArmorTypeEnglishInfo)
	info.SpecialItemTypeInfo = make(map[string]models.SpecialItemTypeEnglishInfo)
	info.ItemConsumableInfo = make(map[string]models.ItemConsumableTypeEnglishInfo)
	info.MasteryTypeInfo = make(map[string]models.MasteryTypeEnglishInfo)
	info.ItemGradeInfo = make(map[string]models.ItemGradeEnglishInfo)
	info.MasteryConditionTypeInfo = make(map[string]models.MasteryConditionTypeEnglishInfo)
	info.StatTypeInfo = make(map[string]models.StatTypeDescriptionEnglish)
	info.ToolTipTypeInfo = make(map[string]models.ToolTipTypeEnglish)
	info.SkinInfo = make(map[string]models.SkinEnglish)
	info.SkillGroupInfo = make(map[string]models.SkillGroupEnglish)
	info.SkillInfo = make(map[string]models.SkillEnglish)
	info.TacticalSkillInfo = make(map[string]models.TacticalSkillEnglish)
	info.CharacterStateInfo = make(map[string]models.CharacterStateEnglish)
	info.ItemSkillInfo = make(map[string]models.ItemSkillEnglish)
	info.SummonInfo = make(map[string]models.SummonDataEnglish)
	return info
}

func ParseEnglishFile(file *os.File) EnglishInfo {
	scanner := bufio.NewScanner(file)
	info := initEnglishInfo()

	for scanner.Scan() {
		line := scanner.Text()

		segments := strings.Split(line, "┃")
		if len(segments) != 2 {
			continue
		}

		key, value := segments[0], segments[1]

		segments = strings.Split(key, "/")

		if len(segments) == 2 {
			mapKey := segments[1]
			intVal, err := strconv.Atoi(segments[1])
			if err != nil {
				switch segments[0] {
				case "ItemType":
					val := initStruct[models.ItemTypeEnglishInfo](info.ItemTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.ItemTypeInfo[mapKey] = val
				case "WeaponType":
					val := initStruct[models.WeaponTypeEnglishInfo](info.WeaponTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.WeaponTypeInfo[mapKey] = val
				case "ArmorType":
					val := initStruct[models.ArmorTypeEnglishInfo](info.ArmorTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.ArmorTypeInfo[mapKey] = val
				case "SpecialItemType":
					val := initStruct[models.SpecialItemTypeEnglishInfo](info.SpecialItemTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.SpecialItemTypeInfo[mapKey] = val
				case "ItemConsumableType":
					val := initStruct[models.ItemConsumableTypeEnglishInfo](info.ItemConsumableInfo, mapKey, "Type", "Name", mapKey, value)
					info.ItemConsumableInfo[mapKey] = val
				case "MasteryType":
					val := initStruct[models.MasteryTypeEnglishInfo](info.MasteryTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.MasteryTypeInfo[mapKey] = val
				case "ItemGrade":
					val := initStruct[models.ItemGradeEnglishInfo](info.ItemGradeInfo, mapKey, "Grade", "Name", mapKey, value)
					info.ItemGradeInfo[mapKey] = val
				case "MasteryConditionType":
					val := initStruct[models.MasteryConditionTypeEnglishInfo](info.MasteryConditionTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.MasteryConditionTypeInfo[mapKey] = val
				case "StatType":
					val := initStruct[models.StatTypeDescriptionEnglish](info.StatTypeInfo, mapKey, "Type", "Name", mapKey, value)
					info.StatTypeInfo[mapKey] = val
				case "ToolTipType":
					val := initStruct[models.ToolTipTypeEnglish](info.ToolTipTypeInfo, mapKey, "Type", "Description", mapKey, value)
					info.ToolTipTypeInfo[mapKey] = val
				}

				continue
			}

			switch segments[0] {
			case "CharacterStoryName":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreName", mapKey, value)
				info.CharacterInfo[mapKey] = val
			case "CharacterStoryAge":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreAge", mapKey, value)
				info.CharacterInfo[mapKey] = val
			case "CharacterStoryHeight":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreHeight", mapKey, value)
				info.CharacterInfo[mapKey] = val
			case "CharacterStoryGender":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreGender", mapKey, value)
				info.CharacterInfo[mapKey] = val
			case "CharacterStoryDesc":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreDescription", mapKey, value)
				info.CharacterInfo[mapKey] = val
			case "CharacterStoryTitle":
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "LoreTitle", mapKey, value)
				info.CharacterInfo[mapKey] = val
			}

			continue
		}

		if len(segments) == 3 && segments[0] == "StatType" && segments[2] == "Desc" {
			val := initStruct[models.StatTypeDescriptionEnglish](info.StatTypeInfo, segments[1], "Type", "Description", segments[1], value)
			info.StatTypeInfo[segments[1]] = val
			continue
		}

		if len(segments) == 3 {
			mapKey := segments[2]
			intVal, err := strconv.Atoi(mapKey)
			if err != nil {
				continue
			}

			if segments[0] == "Item" {
				switch segments[1] {
				case "Name":
					val := initStruct[models.ItemEnglishInfo](info.ItemInfo, intVal, "Code", "Name", mapKey, value)
					info.ItemInfo[mapKey] = val
				case "Help":
					val := initStruct[models.ItemEnglishInfo](info.ItemInfo, intVal, "Code", "Description", mapKey, value)
					info.ItemInfo[mapKey] = val
				case "Effect":
					val := initStruct[models.ItemEnglishInfo](info.ItemInfo, intVal, "Code", "Effect", mapKey, value)
					info.ItemInfo[mapKey] = val
				case "Skills":
					val := initStruct[models.ItemSkillEnglish](info.ItemSkillInfo, intVal, "Code", "ItemName", mapKey, value)
					info.ItemSkillInfo[mapKey] = val
				}
				continue
			}

			if segments[0] == "Trait" {
				switch segments[1] {
				case "Name":
					val := initStruct[models.TraitEnglishInfo](info.TraitInfo, intVal, "Code", "Name", mapKey, value)
					info.TraitInfo[mapKey] = val
				case "StatToolTip":
					val := initStruct[models.TraitEnglishInfo](info.TraitInfo, intVal, "Code", "StatTooltip", mapKey, value)
					info.TraitInfo[mapKey] = val
				}
				continue
			}

			if segments[0] == "Character" && segments[1] == "Name" {
				val := initStruct[models.CharacterEnglishInfo](info.CharacterInfo, intVal, "Code", "Name", mapKey, value)
				info.CharacterInfo[mapKey] = val
				continue
			}

			if segments[0] == "Monster" && segments[1] == "Name" {
				val := initStruct[models.MonsterEnglishInfo](info.MonsterInfo, intVal, "Code", "Name", mapKey, value)
				info.MonsterInfo[mapKey] = val
				continue
			}

			if segments[0] == "Area" && segments[1] == "Name" {
				val := initStruct[models.AreaEnglishInfo](info.AreaInfo, intVal, "Code", "Name", mapKey, value)
				info.AreaInfo[mapKey] = val
				continue
			}

			if segments[0] == "SummonData" && segments[1] == "Name" {
				val := initStruct[models.SummonDataEnglish](info.SummonInfo, intVal, "Code", "Name", mapKey, value)
				info.SummonInfo[mapKey] = val
				continue
			}

			if segments[0] == "Skin" && segments[1] == "Name" {
				val := initStruct[models.SkinEnglish](info.SkinInfo, intVal, "Code", "Name", mapKey, value)
				info.SkinInfo[mapKey] = val
				continue
			}

			if segments[0] == "Skin" && segments[1] == "Desc" {
				val := initStruct[models.SkinEnglish](info.SkinInfo, intVal, "Code", "Description", mapKey, value)
				info.SkinInfo[mapKey] = val
				continue
			}

			if segments[0] == "Skill" && segments[1] == "LobbyName" {
				val := initStruct[models.SkillEnglish](info.SkillInfo, intVal, "Code", "Name", mapKey, value)
				info.SkillInfo[mapKey] = val
				continue
			}

			if segments[0] == "Skill" && segments[1] == "LobbyDesc" {
				val := initStruct[models.SkillEnglish](info.SkillInfo, intVal, "Code", "Description", mapKey, value)
				info.SkillInfo[mapKey] = val
				continue
			}

			continue
		}

		if len(segments) == 4 && segments[0] == "Item" && segments[1] == "Skills" {
			intVal, err := strconv.Atoi(segments[2])
			if err != nil {
				continue
			}
			switch segments[3] {
			case "Body":
				val := initStruct[models.ItemSkillEnglish](info.ItemSkillInfo, intVal, "Code", "ItemSkillBody", segments[2], value)
				info.ItemSkillInfo[segments[2]] = val
			case "Name":
				val := initStruct[models.ItemSkillEnglish](info.ItemSkillInfo, intVal, "Code", "ItemSkillName", segments[2], value)
				info.ItemSkillInfo[segments[2]] = val
			case "Desc":
				val := initStruct[models.ItemSkillEnglish](info.ItemSkillInfo, intVal, "Code", "ItemSkillDescription", segments[2], value)
				info.ItemSkillInfo[segments[2]] = val
			}
		}

		if len(segments) == 4 {
			mapKey := segments[3]
			intVal, err := strconv.Atoi(mapKey)
			if err != nil {
				continue
			}

			if segments[0] == "Infusion" && segments[1] == "Trait" && segments[2] == "Desc" {
				val := initStruct[models.TraitEnglishInfo](info.TraitInfo, intVal, "Code", "Description", mapKey, value)
				info.TraitInfo[mapKey] = val
				continue
			}

			if segments[0] == "Skill" && segments[1] == "Group" {
				switch segments[3] {
				case "Name":
					val := initStruct[models.SkillGroupEnglish](info.SkillGroupInfo, intVal, "Code", "Name", mapKey, value)
					info.SkillGroupInfo[mapKey] = val
				case "Description":
					val := initStruct[models.SkillGroupEnglish](info.SkillGroupInfo, intVal, "Code", "Description", mapKey, value)
					info.SkillGroupInfo[mapKey] = val
				case "Coef":
					val := initStruct[models.SkillGroupEnglish](info.SkillGroupInfo, intVal, "Code", "Coefficient", mapKey, value)
					info.SkillGroupInfo[mapKey] = val
				case "ExpansionTip":
					val := initStruct[models.SkillGroupEnglish](info.SkillGroupInfo, intVal, "Code", "ExpansionTip", mapKey, value)
					info.SkillGroupInfo[mapKey] = val
				}
				continue
			}

			if segments[0] == "Skill" && segments[1] == "Code" {
				switch segments[3] {
				case "Name":
					val := initStruct[models.TacticalSkillEnglish](info.TacticalSkillInfo, intVal, "Code", "Name", mapKey, value)
					info.TacticalSkillInfo[mapKey] = val
				case "Desc":
					val := initStruct[models.TacticalSkillEnglish](info.TacticalSkillInfo, intVal, "Code", "Description", mapKey, value)
					info.TacticalSkillInfo[mapKey] = val
				case "Coef":
					val := initStruct[models.TacticalSkillEnglish](info.TacticalSkillInfo, intVal, "Code", "Coefficient", mapKey, value)
					info.TacticalSkillInfo[mapKey] = val
				case "ExpansionTip":
					val := initStruct[models.TacticalSkillEnglish](info.TacticalSkillInfo, intVal, "Code", "ExpansionTip", mapKey, value)
					info.TacticalSkillInfo[mapKey] = val
				}
				continue
			}

			if segments[0] == "CharacterState" && segments[1] == "Group" && segments[2] == "Name" {
				val := initStruct[models.CharacterStateEnglish](info.CharacterStateInfo, intVal, "Code", "Name", mapKey, value)
				info.CharacterStateInfo[mapKey] = val
			}

			if segments[0] == "CharacterState" && segments[1] == "Group" && segments[2] == "Desc" {
				val := initStruct[models.CharacterStateEnglish](info.CharacterStateInfo, intVal, "Code", "Description", mapKey, value)
				info.CharacterStateInfo[mapKey] = val
			}
		}
	}

	return info
}

func initStruct[T any](structMap map[string]T, id interface{}, idField, field, key, value string) T {
	val, ok := structMap[key]
	if !ok {
		val = reflect.New(reflect.TypeOf(structMap).Elem()).Elem().Interface().(T)
		switch id.(type) {
		case int:
			reflect.ValueOf(&val).Elem().FieldByName(idField).SetInt(int64(id.(int)))
		case string:
			reflect.ValueOf(&val).Elem().FieldByName(idField).SetString(id.(string))
		}
	}
	reflect.ValueOf(&val).Elem().FieldByName(field).SetString(value)

	return val
}
