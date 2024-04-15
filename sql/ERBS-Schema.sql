DO $$
    BEGIN
        --
        -- Extensions
        --

        CREATE EXTENSION hstore;
        CREATE EXTENSION intarray;
        CREATE EXTENSION pg_stat_statements;

        --
        -- Type Definitions
        --

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mastery_level') THEN
            CREATE TYPE mastery_level AS
            (
                none INT,
                glove INT,
                tonfa INT,
                bat INT,
                whip INT,
                high_angle_fire INT,
                direct_fire INT,
                bow INT,
                crossbow INT,
                pistol INT,
                assault_rifle INT,
                sniper_rifle INT,
                hammer INT,
                axe INT,
                one_handed_sword INT,
                two_handed_sword INT,
                polearm INT,
                dual_sword INT,
                spear INT,
                nunchaku INT,
                rapier INT,
                guitar INT,
                camera INT,
                arcana INT,
                vf_arm INT,
                craft INT,
                search INT,
                move INT,
                defense INT,
                hunt INT
            );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'equipment') THEN
            CREATE TYPE equipment AS
                (
                    weapon INT,
                    chest INT,
                    head INT,
                    arm INT,
                    legs INT
                );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'monsters_killed') THEN
            CREATE TYPE monsters_killed AS
                (
                    chickens INT,
                    bats INT,
                    boars INT,
                    dogs INT,
                    wolves INT,
                    bears INT,
                    wickeline INT,
                    alpha INT,
                    omega INT,
                    gamma INT,
                    mutant_chickens INT,
                    mutant_bats INT,
                    mutant_boars INT,
                    mutant_dogs INT,
                    mutant_wolves INT,
                    mutant_bears INT
                );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'credit_source') THEN
            CREATE TYPE credit_source AS
                (
                    preliminary_phase INT,
                    tactical_skill_upgrade INT,
                    time_elapsed_compensation_by_ms FLOAT4,
                    time_elapsed_credit_bonus_by_ms FLOAT4,
                    kill_chicken INT,
                    kill_bat INT,
                    kill_boar INT,
                    kill_dog INT,
                    kill_wolf INT,
                    gold_security_console_access INT,
                    kill_mutant_boar INT,
                    kill_bear INT,
                    kill_mutant_chicken INT,
                    kill_player_merge INT,
                    kill_alpha INT,
                    kill_mutant_dog INT,
                    kill_mutant_wolf INT,
                    transfer_console_special_material INT,
                    kill_assist_divide_contribute INT,
                    transfer_console_remote_drone_myself INT,
                    kill_mutant_bear INT,
                    item_bounty INT,
                    kill_omega INT,
                    transfer_console_escape_key INT,
                    kill_mutant_bat INT,
                    transfer_console_resurrection INT,
                    kill_wickeline INT
                );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_equip_log') THEN
            CREATE TYPE item_equip_log AS
                (
                    weapon INT[],
                    chest INT[],
                    head INT[],
                    arm INT[],
                    legs INT[]
                );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'recommended_weapon_route') THEN
            CREATE TYPE recommended_weapon_route AS
                (
                    id INT,
                    title TEXT,
                    user_num INT,
                    user_nickname TEXT,
                    character_code INT,
                    slot_id INT,
                    weapon_type INT,
                    weapon_codes TEXT,
                    tactical_skill_group_code INT,
                    paths TEXT,
                    count INT,
                    version TEXT,
                    team_mode INT,
                    language_code TEXT,
                    route_version INT,
                    share BOOLEAN,
                    update_dtm INT8,
                    v2_like INT,
                    v2_accumulate_like INT,
                    v2_accumulate_win_rate FLOAT8,
                    v2_accumulate_season_id INT
                );
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'recommended_weapon_route_desc') THEN
            CREATE TYPE recommended_weapon_route_desc AS
                (
                    recommended_weapon_route_id INT,
                    skill_path TEXT,
                    desc TEXT
                );
        END IF;

        --
        -- Tables
        --
        CREATE TABLE IF NOT EXISTS action_cost (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            type TEXT,
            sp INT,
            time_1 FLOAT4,
            time_2 FLOAT4,
            action_wait_time FLOAT4,
            casting_anim_trigger TEXT,
            effect_cancel_condition TEXT,
            casting_bar_img_name TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS area (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            mode_type INT,
            mask_code INT,
            starting_area BOOLEAN,
            area_type TEXT,
            is_provide_collectible_item BOOLEAN,
            route_calc_bit_code INT,
            is_hyperloop_installed BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS battle_zone_reward (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode_type INT,
            area_attributes_create_event_count INT,
            item_code INT,
            type TEXT,
            value INT,
            selectable BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS bullet_capacity (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            item_code INT,
            capacity INT,
            load_type TEXT,
            time FLOAT4,
            init_count INT,
            count INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            max_hp INT,
            max_sp INT,
            init_extra_point INT,
            max_extra_point INT,
            attack_power INT,
            defense INT,
            skill_amp INT,
            adaptive_force INT,
            critical_strike_chance FLOAT4,
            hp_regen FLOAT4,
            sp_regen FLOAT4,
            attack_speed FLOAT4,
            attack_speed_ratio FLOAT4,
            increase_basic_attack_damage_ratio FLOAT4,
            skill_amp_ratio FLOAT4,
            prevent_basic_attack_damaged_ratio FLOAT4,
            prevent_skill_damaged_ratio FLOAT4,
            attack_speed_limit FLOAT4,
            attack_speed_min FLOAT4,
            move_speed FLOAT4,
            sight_range FLOAT4,
            radius FLOAT4,
            pathing_radius FLOAT4,
            ui_height FLOAT4,
            init_state_display_index INT,
            local_scale_in_cutscene INT,
            local_scale_in_victory_scene TEXT,
            resource TEXT,
            lobby_sub_object TEXT,
            name TEXT,
            str_learn_start_skill TEXT,
            str_use_point_learn_start_skill TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character_attributes (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            character TEXT,
            character_code INT,
            mastery TEXT,
            control_difficulty INT,
            attack INT,
            defense INT,
            disruptor INT,
            move INT,
            assistance INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character_exp (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            level INT,
            level_up_exp INT
        );

        CREATE TABLE IF NOT EXISTS character_level_up_stat (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            max_hp INT,
            max_sp INT,
            skill_amp INT,
            adaptive_force INT,
            attack_power FLOAT4,
            defense FLOAT4,
            critical_strike_chance FLOAT4,
            hp_regen FLOAT4,
            sp_regen FLOAT4,
            attack_speed FLOAT4,
            move_speed FLOAT4,
            attack_speed_ratio FLOAT4,
            increase_basic_attack_damage_ratio FLOAT4,
            skill_amp_ratio FLOAT4,
            prevent_basic_attack_damaged_ratio FLOAT4,
            prevent_skill_damaged_ratio FLOAT4,
            name TEXT
        );

        CREATE TABLE IF NOT EXISTS item_armor (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            make_material_1 INT,
            make_material_2 INT,
            credit_value_when_converted_to_bounty INT,
            manufacturable_type INT,
            attack_power INT,
            attack_power_by_lv INT,
            defense INT,
            defense_by_lv INT,
            skill_amp INT,
            skilL_amp_by_level INT,
            adaptive_force INT,
            adaptive_force_by_level INT,
            max_hp INT,
            max_hp_by_lv INT,
            max_sp INT,
            prevent_critical_strike_damaged INT,
            attack_range INT,
            increase_basic_attack_damage INT,
            increase_basic_attack_damage_by_lv INT,
            prevent_basic_attack_damaged INT,
            prevent_basic_attack_damaged_by_lv INT,
            increase_basic_attack_damaged_ratio INT,
            prevent_skill_damaged INT,
            prevent_skill_damaged_by_lv INT,
            penetration_defense INT,
            trap_damage_reduce INT,
            unique_attack_range INT,
            unique_penetration_defense INT,
            item_usable_value_list INT,
            exclusive_producer INT,
            mode_type INT,
            stackable INT,
            intial_count INT,
            sight_range INT,
            hp_regen_ratio FLOAT4,
            hp_regen FLOAT4,
            sp_regen_ratio FLOAT4,
            sp_regen FLOAT4,
            attack_speed_ratio FLOAT4,
            attack_speed_ratio_by_lv FLOAT4,
            critical_strike_chance FLOAT4,
            critical_strike_damage FLOAT4,
            skill_amp_ratio FLOAT4,
            skill_amp_ratio_by_level FLOAT4,
            cooldown_reduction FLOAT4,
            cooldown_limit FLOAT4,
            lifesteal FLOAT4,
            normal_lifesteal FLOAT4,
            skill_lifesteal FLOAT4,
            move_speed FLOAT4,
            move_speed_out_of_combat FLOAT4,
            unique_hp_healed_increase_ratio FLOAT4,
            unique_cooldown_limit FLOAT4,
            unique_tenacity FLOAT4,
            unique_move_speed FLOAT4,
            prevent_basic_attack_damaged_ratio FLOAT4,
            prevent_basic_attack_damaged_ratio_by_lv FLOAT4,
            prevent_skill_damaged_ratio FLOAT4,
            prevent_skill_damaged_ratio_by_lv FLOAT4,
            penetration_defense_ratio FLOAT4,
            trap_damage_reduce_ratio FLOAT4,
            hp_healed_increase_ratio FLOAT4,
            healer_give_hp_heal_ratio FLOAT4,
            unique_penetration_defense_ratio FLOAT4,
            unique_lifesteal FLOAT4,
            unique_skill_amp_ratio FLOAT4,
            increase_basic_attack_damage_ratio_by_lv FLOAT4,
            is_completed_item BOOLEAN,
            alert_in_spectator BOOLEAN,
            is_removed_from_player_inventory_on_death BOOLEAN,
            not_disarm BOOLEAN,
            restore_item_when_resurrected BOOLEAN,
            name TEXT,
            item_type TEXT,
            armor_type TEXT,
            item_grade TEXT,
            marking_type TEXT,
            craft_anim_trigger TEXT,
            item_usable_type TEXT,
            make_custom_action TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character_mastery (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            weapon_1 TEXT,
            weapon_2 TEXT,
            weapon_3 TEXT,
            weapon_4 TEXT,
            combat_1 TEXT,
            combat_2 TEXT,
            survival_1 TEXT,
            survival_2 TEXT,
            survival_3 TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character_mode_modifier (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            character_code INT,
            weapon_type TEXT,
            solo_increase_mode_damage_ratio INT,
            solo_prevent_mode_damage_ratio INT,
            solo_increase_mode_heal_ratio INT,
            solo_increase_mode_shield_ratio INT,
            duo_increase_mode_damage_ratio INT,
            duo_prevent_mode_damage_ratio INT,
            duo_increase_mode_heal_ratio INT,
            duo_increase_mode_healer_give_ratio INT,
            duo_increase_mode_shield_ratio INT,
            duo_increase_mode_healer_give_shield_ratio INT,
            squad_increase_mode_damage_ratio INT,
            squad_prevent_mode_damage_ratio INT,
            squad_increase_mode_heal_ratio INT,
            squad_increase_mode_healer_give_ratio INT,
            squad_increase_mode_shield_ratio INT,
            squad_increase_mode_healer_give_shield_ratio INT,
            cobalt_increase_mode_damage_ratio INT,
            cobalt_prevent_mode_damage_ratio INT,
            cobalt_increase_mode_heal_ratio INT,
            cobalt_increase_mode_healer_give_ratio INT,
            cobalt_increase_mode_shield_ratio INT,
            cobalt_increase_mode_healer_give_shield_ratio INT,
            cobalt_increase_mode_ult_cooldown_ratio INT,
            cobalt_increase_mode_max_sp_ratio INT,
            cobalt_increase_mode_sp_regen_ratio INT,
            solo_increase_mode_damage_to_monster_ratio INT,
            duo_increase_mode_damage_to_monster_ratio INT,
            squad_increase_mode_damage_to_monster_ratio INT,
            cobalt_increase_mode_damage_to_monster_ratio INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character_skin (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            name TEXT,
            code INT,
            character_code INT,
            index INT,
            grade INT,
            event_free BOOLEAN,
            purchase_type TEXT,
            effects_path TEXT,
            projectiles_path TEXT,
            object_path TEXT,
            fx_sound_path TEXT,
            voice_path TEXT,
            weapon_mount_path TEXT,
            weapon_mount_common_path TEXT,
            indicator_path TEXT,
            projectiles_deflector_path TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS collectible (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            cooldown INT,
            item_code_1 TEXT,
            item_code_2 TEXT,
            probablity_1 INT,
            probablity_2 INT,
            drop_count INT,
            casting_action_type TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS drop_group (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            group_code INT,
            item_code TEXT,
            min INT,
            max INT,
            probability INT,
            drop_type INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS gain_exp (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            start_time INT,
            end_time INT,
            gain_exp INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS game_tip (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            key INT,
            code INT,
            game_tip_type TEXT,
            game_tip_category TEXT,
            sequence INT,
            title_text_key TEXT,
            context_text_key TEXT,
            image_name TEXT,
            korea_word TEXT,
            link TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS how_to_find_items (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            item_code INT,
            hunt_chicken INT,
            hunt_bat INT,
            hunt_boar INT,
            hunt_dog INT,
            hunt_wolf INT,
            hunt_bear INT,
            hunt_wickeline INT,
            hunt_alpha INT,
            hunt_omega INT,
            collectible_code INT,
            air_supply INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS infusion_product (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            product_type TEXT,
            product_group INT,
            product_code INT,
            store_type TEXT,
            stock_type TEXT,
            stock INT,
            is_restore BOOLEAN,
            price INT,
            special_weight INT,
            weight INT,
            requirement INT,
            icon TEXT,
            simple_icon TEXT,
            alert_in_spectator BOOLEAN,
            character_codes TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_consumable (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            NAME TEXT,
            mode_type INT,
            item_type TEXT,
            consumable_type TEXT,
            consumable_tag TEXT,
            item_grade TEXT,
            is_completed_item BOOLEAN,
            alert_in_spectator BOOLEAN,
            marking_type TEXT,
            craft_anim_trigger TEXT,
            stackable INT,
            initial_count INT,
            item_usable_type TEXT,
            item_usable_value_list INT,
            exclusive_producer INT,
            is_removed_from_player_inventory_on_death BOOLEAN,
            manufacturable_type INT,
            make_material_1 INT,
            make_material_2 INT,
            heal INT,
            hp_recover INT,
            sp_recover INT,
            attack_power_by_buff INT,
            defense_by_buff INT,
            skill_amp_by_buff INT,
            skill_amp_ratio_by_buff INT,
            add_state_code INT,
            is_vpad_quick_slot_item BOOLEAN,
            restore_item_when_resurrected BOOLEAN,
            credit_value_when_converted_to_bounty INT,
            is_reduce_loot_on_death BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_misc (
            id SMALLINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            mode_type INT,
            item_type TEXT,
            misc_item_type TEXT,
            item_grade TEXT,
            grade_bg_override TEXT,
            is_completed_item BOOLEAN,
            alert_in_spectator BOOLEAN,
            marking_type TEXT,
            craft_anim_trigger TEXT,
            stackable INT,
            initial_count INT,
            item_usable_type TEXT,
            item_usable_value_list INT,
            exclusive_producer INT,
            is_removed_from_player_inventory_on_death BOOLEAN,
            manufacturable_type INT,
            make_material_1 INT,
            make_material_2 INT,
            make_custom_action TEXT,
            restore_item_when_resurrected BOOLEAN,
            credit_value_when_converted_to_bounty INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_search_option (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            tag_1 TEXT,
            tag_2 TEXT,
            tag_3 TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_spawn (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            area_code INT,
            area_spawn_group INT,
            item_code INT,
            drop_point TEXT,
            drop_count INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_special (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            mode_type INT,
            item_type TEXT,
            special_item_type TEXT,
            item_grade TEXT,
            is_completed_item BOOLEAN,
            alert_in_spectator BOOLEAN,
            marking_type TEXT,
            craft_anim_trigger TEXT,
            stackable INT,
            initial_count INT,
            cooldown_group_code INT,
            cooldown FLOAT4,
            item_usable_type TEXT,
            item_usable_value_list INT,
            exclusive_producer INT[],
            is_removed_from_player_inventory_on_death BOOLEAN,
            manufacturable_type INT,
            make_material_1 INT,
            make_material_2 INT,
            make_custom_action TEXT,
            consume_count INT,
            summon_code INT,
            ghost_item_state_group INT,
            is_vpad_quick_slot_item BOOLEAN,
            restore_item_when_resurrected BOOLEAN,
            credit_value_when_converted_to_bounty INT,
            is_reduce_loot_on_death BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_weapon (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            mode_type INT,
            item_type TEXT,
            weapon_type TEXT,
            item_grade TEXT,
            grade_bg_override TEXT,
            is_completed_item BOOLEAN,
            alert_in_spectator BOOLEAN,
            marking_type TEXT,
            craft_anim_trigger TEXT,
            stackable INT,
            initial_count INT,
            item_usable_type TEXT,
            item_usable_value_list INT,
            exclusive_producer INT[],
            is_removed_from_player_inventory_on_death BOOLEAN,
            make_material_1 INT,
            make_material_2 INT,
            make_custom_action TEXT,
            not_disarm BOOLEAN,
            consumable BOOLEAN,
            manufacturable_type INT,
            attack_power INT,
            attack_power_by_lv INT,
            defense INT,
            defense_by_lv INT,
            skill_amp INT,
            skill_amp_by_lv INT,
            skill_amp_ratio INT,
            skill_amp_ratio_by_lv INT,
            adaptive_force INT,
            adaptive_force_by_lv INT,
            max_hp INT,
            max_hp_by_lv INT,
            hp_regen_ratio FLOAT4,
            hp_regen INT,
            max_sp INT,
            sp_regen_ratio FLOAT4,
            sp_regen INT,
            attack_speed_ratio FLOAT4,
            attack_speed_ratio_by_lv FLOAT4,
            critical_strike_chance FLOAT4,
            critical_strike_damage FLOAT4,
            cooldown_reduction FLOAT4,
            prevent_critical_strike_damage INT,
            cooldown_limit INT,
            lifesteal FLOAT4,
            normal_lifesteal FLOAT4,
            skill_lifesteal FLOAT4,
            move_speed FLOAT4,
            move_speed_out_of_combat FLOAT4,
            sight_range FLOAT4,
            attack_range FLOAT4,
            increase_basic_attack_damage INT,
            increase_basic_attack_damage_by_lv INT,
            increase_basic_attack_damage_ratio FLOAT4,
            increase_basic_attack_damage_ratio_by_lv FLOAT4,
            prevent_basic_attack_damage INT,
            prevent_basic_attack_damage_by_lv INT,
            prevent_basic_attack_damage_ratio FLOAT4,
            prevent_basic_attack_damage_ratio_by_lv FLOAT4,
            prevent_skill_damage INT,
            prevent_skill_damage_by_lv INT,
            prevent_skill_damage_ratio FLOAT4,
            prevent_skill_damage_ratio_by_lv FLOAT4,
            penetration_defense INT,
            penetration_defense_ratio FLOAT4,
            trap_damage_reduce INT,
            trap_damage_reduce_ratio FLOAT4,
            hp_healed_increase_ratio FLOAT4,
            healer_give_hp_heal_ratio FLOAT4,
            unique_attack_range FLOAT4,
            unique_hp_healed_increase_ratio FLOAT4,
            unique_cooldown_limit FLOAT4,
            unique_tenacity FLOAT4,
            unique_move_speed FLOAT4,
            unique_penetration_defense INT,
            unique_penetration_defense_ratio FLOAT4,
            unique_lifesteal FLOAT4,
            unique_skill_amp_ratio FLOAT4,
            restore_item_when_resurrected BOOLEAN,
            credit_value_when_converted_to_bounty INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS level (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            level INT,
            needed_exp INT,
            accumulate_exp INT,
            reward_acoin INT,
            reward INT
        );

        CREATE TABLE IF NOT EXISTS loading_tip (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            loading_tip_type TEXT,
            min_lv INT,
            max_lv INT,
            text_key TEXT,
            image_name TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS mastery_exp (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode_type INT,
            condition_type TEXT,
            grade TEXT,
            conditional_value INT,
            mastery_type_1 TEXT,
            value_1 INT,
            mastery_type_2 TEXT,
            value_2 INT,
            mastery_type_3 TEXT,
            value_3 INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS mastery_stat (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            type TEXT,
            character_code INT,
            first_option TEXT,
            first_option_section_1_value FLOAT4,
            first_option_section_2_value FLOAT4,
            first_option_section_3_value FLOAT4,
            first_option_section_4_value FLOAT4,
            second_option TEXT,
            second_option_section_1_value FLOAT4,
            second_option_section_2_value FLOAT4,
            second_option_section_3_value FLOAT4,
            second_option_section_4_value FLOAT4,
            third_option TEXT,
            third_option_section_1_value FLOAT4,
            third_option_section_2_value FLOAT4,
            third_option_section_3_value FLOAT4,
            third_option_section_4_value FLOAT4,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS monster (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            monster TEXT,
            is_mutant BOOLEAN,
            grade TEXT,
            mode INT,
            create_day TEXT,
            create_time INT,
            regen_time INT,
            level_up_period INT,
            level_up_amount INT,
            level_up_max INT,
            max_hp INT,
            max_ep INT,
            init_extra_point INT,
            attack_power INT,
            defense INT,
            attack_speed FLOAT4,
            move_speed FLOAT4,
            sight_range INT,
            chasing_range INT,
            attack_range FLOAT4,
            first_attack_range FLOAT4,
            aggressive TEXT,
            detect_invisible BOOLEAN,
            radius FLOAT4,
            pathing_radius FLOAT4,
            ui_height FLOAT4,
            gain_exp INT,
            target_on_range INT,
            random_drop_count INT,
            resource TEXT,
            corpse_resource TEXT,
            appear_time FLOAT4,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS monster_drop_group (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            monster_code INT,
            monster_level INT,
            drop_group INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS monster_level_up_stat (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            monster TEXT,
            mode INT,
            max_hp INT,
            attack_power INT,
            defense FLOAT4,
            move_speed FLOAT4,
            gain_exp INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS monster_spawn_level (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode INT,
            player_level INT,
            monster_code INT,
            spawn_level INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS navi_collect_and_hunt (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            item_code INT,
            area_code_list TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS near_by_area (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            area_code INT,
            near_by_area_code INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS random_equipment (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            "group" TEXT,
            item_code INT,
            weight INT,
            item_grade TEXT,
            tag_multiplier INT,
            character_num INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS recommended_list (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            character TEXT,
            character_code INT,
            mastery TEXT,
            start_weapon INT,
            cobalt_start_weapon INT,
            start_item_group_code INT,
            cobalt_draft INT,
            cobalt_extra_draft TEXT,
            cobalt_can_choose_weapon BOOLEAN,
            favorite_main_tag TEXT,
            opposite_tag TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS season (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            season_id INT,
            season_name DATE,
            season_start DATE,
            season_end DATE,
            is_current BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS summon_object_stat (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            duration FLOAT4,
            create_range FLOAT4,
            pile_range FLOAT4,
            create_visible_time FLOAT4,
            create_stealth_time FLOAT4,
            infiltration_time FLOAT4,
            detection_range FLOAT4,
            max_hp INT,
            max_sp INT,
            init_extra_point INT,
            max_extra_point INT,
            attack_power INT,
            defense INT,
            critical_strike_chance FLOAT4,
            hp_regen INT,
            sp_regen INT,
            range_radius FLOAT4,
            attack_speed FLOAT4,
            attack_range FLOAT4,
            attack_delay FLOAT4,
            move_speed FLOAT4,
            radius FLOAT4,
            ui_height FLOAT4,
            sight_range FLOAT4,
            sight_angle INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS tactical_skill_set (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            next_upgrade_code INT,
            upgrade_credit INT,
            upgrade_material INT,
            skill_code INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS tactical_skill_set_group (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            "group" INT,
            mode_type INT,
            start_code INT,
            equip_with_start BOOLEAN,
            icon TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS trait (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            open_account_lv INT,
            trait_group TEXT,
            trait_type TEXT,
            active BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS transfer_console (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            item_code INT,
            mode INT,
            item_type TEXT,
            transfer_time_safe_area INT,
            subtract_transfer_time_restricted_area INT,
            manufacture_cooldown INT,
            available_time_for_purchase INT,
            consume_vf_credit INT,
            limit_count INT,
            trait_sale BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS vf_credit (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode INT,
            phase INT,
            condition_type TEXT,
            condition_value INT,
            acquire_self INT,
            acquire_team INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS weapon_type_info (
            id INT GENERATED BY DEFAULT AS IDENTITY,
            type TEXT,
            attack_speed FLOAT4,
            attack_range FLOAT4,
            shop_filter INT,
            summon_object_hit_damage INT,

            PRIMARY KEY (id)
        );

        --
        -- Keys and Constraints
        --

        --
        -- Indexes
        --
END $$