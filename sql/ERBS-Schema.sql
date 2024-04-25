DO $$
    BEGIN
        --
        -- Extensions
        --

        IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'hstore') THEN
            CREATE EXTENSION hstore;
        END IF;
        IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'intarray') THEN
            CREATE EXTENSION intarray;
        END IF;
        IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pg_stat_statements') THEN
            CREATE EXTENSION pg_stat_statements;
        END IF;
        IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pg_trgm') THEN
            CREATE EXTENSION pg_trgm WITH SCHEMA public;
        END IF;

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

        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_character_stats') THEN
            CREATE TYPE user_character_stats AS
                (
                    character_code INT,
                    total_games INT,
                    usages INT,
                    max_killings INT,
                    top_3 INT,
                    wins INT,
                    top_3_rate INT,
                    average_rank INT
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
                    route_id BIGINT,
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
                    update_dtm INT,
                    v2_like INT,
                    v2_winrate FLOAT4,
                    v2_season_id INT,
                    v2_accumulate_like INT,
                    v2_accumulate_win_rate FLOAT4,
                    v2_accumulate_season_id INT
                );
        END IF;


        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'recommended_weapon_route_desc') THEN
            CREATE TYPE recommended_weapon_route_desc AS
                (
                    recommended_weapon_route_id INT,
                    skill_path TEXT,
                    "desc" TEXT
                );
        END IF;

        --
        -- Tables
        --
        CREATE TABLE IF NOT EXISTS mastery_levels (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            type TEXT,
            mastery_level INT,
            next_mastery_exp INT,
            give_level_exp INT,
            exp_growth_cap_ratio INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS user_info (
            user_num BIGINT,
            nickname TEXT,

            PRIMARY KEY (user_num)
        );

        CREATE TABLE IF NOT EXISTS user_rank (
            user_num INT,
            mmr INT,
            nickname TEXT,
            rank INT
        );

        CREATE TABLE IF NOT EXISTS top_rank (
            user_num INT,
            nickname TEXT,
            rank INT,
            mmr INT,
            user_emblems INT[]
        );

        CREATE TABLE IF NOT EXISTS user_stats (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            season_id INT,
            user_num INT,
            matching_mode INT,
            matching_team_mode INT,
            mmr INT,
            nickname TEXT,
            rank INT,
            rank_size INT,
            total_games INT,
            total_wins INT,
            total_team_kills INT,
            total_deaths INT,
            escape_count INT,
            rank_percent FLOAT4,
            average_rank FLOAT4,
            average_kills FLOAT4,
            average_assistants FLOAT4,
            average_hunts FLOAT4,
            top_1 FLOAT4,
            top_2 FLOAT4,
            top_3 FLOAT4,
            top_5 FLOAT4,
            top_7 FLOAT4,
            character_stats user_character_stats,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS gain_score (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            phase INT,
            condition_type TEXT,
            condition_value INT,
            points_enemy INT,
            points_ally INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS user_game (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            user_num INT,
            nickname TEXT,
            game_id INT,
            season_id INT,
            matching_mode INT,
            matching_team_mode INT,
            character_num INT,
            skin_code INT,
            character_level INT,
            game_rank INT,
            player_kill INT,
            player_assistant INT,
            monster_kill INT,
            best_weapon INT,
            best_weapon_level INT,
            mastery_level mastery_level,
            equipment equipment,
            version_major INT,
            version_minor INT,
            language TEXT,
            skill_level_info hstore,
            skill_order_info hstore,
            server_name TEXT,
            max_hp INT,
            max_sp INT,
            attack_power INT,
            defense INT,
            hp_regen FLOAT4,
            sp_regen FLOAT4,
            attack_speed FLOAT4,
            move_speed FLOAT4,
            out_of_combat_move_speed FLOAT4,
            sight_range FLOAT4,
            attack_range FLOAT4,
            critical_strike_chance FLOAT4,
            critical_strike_damage FLOAT4,
            cooldown_reduction FLOAT4,
            lifesteal FLOAT4,
            normal_lifesteal FLOAT4,
            skill_lifesteal FLOAT4,
            amplifier_to_monster FLOAT4,
            trap_damage INT,
            bonus_coin INT,
            gain_exp INT,
            base_exp INT,
            bonus_exp INT,
            start_dtm TEXT,
            duration INT,
            mmr_before INT,
            mmr_gain INT,
            mmr_after INT,
            play_time INT,
            watch_time INT,
            total_time INT,
            survivable_time INT,
            bot_added INT,
            bot_remain INT,
            restricted_area_accelerated INT,
            safe_areas INT,
            team_number INT,
            pre_made INT,
            gained_normal_mmr_k_factor INT,
            victory INT,
            craft_uncommon INT,
            craft_rare INT,
            craft_epic INT,
            craft_legend INT,
            damage_to_player INT,
            damage_to_player_basic INT,
            damage_to_player_skill INT,
            damage_to_player_item_skill INT,
            damage_to_player_direct INT,
            damage_to_player_unique_skill INT,
            damage_from_player INT,
            damage_from_player_trap INT,
            damage_from_player_basic INT,
            damage_from_player_item_skill INT,
            damage_from_player_direct INT,
            damage_from_player_unique_skill INT,
            damage_to_monster INT,
            damage_to_monster_trap INT,
            damage_to_monster_basic INT,
            damage_to_monster_skill INT,
            damage_to_monster_item_skill INT,
            damage_to_monster_direct INT,
            damage_to_monster_unique_skill INT,
            damage_from_monster INT,
            damage_to_player_shield INT,
            damage_offset_by_shield_player INT,
            damage_offset_by_shield_monster INT,
            kill_monsters monsters_killed,
            heal_amount INT,
            team_recover INT,
            protect_absorb INT,
            add_surveillance_camera INT,
            add_telephoto_camera INT,
            remove_surveillance_camera INT,
            remove_telephoto_camera INT,
            use_hyperloop INT,
            use_security_console INT,
            give_up INT,
            team_spectator INT,
            route_id_of_start INT,
            route_slot_id INT,
            place_of_start TEXT,
            mmr_avg INT,
            match_size INT,
            team_kill INT,
            total_field_kill INT,
            account_level INT,
            killer_user_num INT,
            killer TEXT,
            kill_detail TEXT,
            cause_of_death TEXT,
            place_of_death TEXT,
            killer_character TEXT,
            killer_weapon TEXT,
            killer_user_num_2 INT,
            kill_detail_2 TEXT,
            cause_of_death_2 TEXT,
            place_of_death_2 TEXT,
            killer_character_2 TEXT,
            killer_weapon_2 TEXT,
            killer_user_num_3 INT,
            kill_detail_3 TEXT,
            cause_of_death_3 TEXT,
            place_of_death_3 TEXT,
            killer_character_3 TEXT,
            killer_weapon_3 TEXT,
            fishing_count INT,
            use_emoticon_count INT,
            expire_dtm TEXT,
            trait_first_core INT,
            trait_first_sub INT[],
            trait_second_sub INT[],
            air_supply_open_count INT[],
            food_craft_count INT[],
            beverage_craft_count INT[],
            rank_point INT,
            total_turbine_take_over INT,
            used_normal_heal_pack INT,
            used_reinforced_heal_pack INT,
            used_normal_shield_pack INT,
            used_reinforced_shield_pack INT,
            total_vf_credits INT[],
            sum_total_vf_credits INT,
            sum_used_vf_credits INT,
            craft_mythic INT,
            player_deaths INT,
            kill_gamma BOOLEAN,
            scored_point INT[],
            kill_details TEXT,
            death_details TEXT,
            kills_phase_one INT,
            kills_phase_two INT,
            kills_phase_three INT,
            deaths_phase_one INT,
            deaths_phase_two INT,
            deaths_phase_three INT,
            used_pair_loop INT,
            cc_time_to_player FLOAT4,
            credit_source credit_source,
            bought_infusion TEXT,
            item_transferred_console INT[],
            item_transferred_drone INT[],
            escape_state INT,
            total_double_kill INT,
            total_triple_kill INT,
            total_quadra_kill INT,
            total_extra_kill INT,
            collect_item_for_log INT[],
            equip_first_item_for_log item_equip_log,
            battle_zone_1_area_code INT,
            battle_zone_1_battle_mark INT,
            battle_zone_1_item_code INT[],
            battle_zone_2_area_code INT,
            battle_zone_2_battle_mark INT,
            battle_zone_2_item_code INT[],
            battle_zone_3_area_code INT,
            battle_zone_3_battle_mark INT,
            battle_zone_3_item_code INT[],
            battle_zone_player_kill INT,
            battle_zone_deaths INT,
            battle_zone_1_winner INT,
            battle_zone_2_winner INT,
            battle_zone_3_winner INT,
            battle_zone_1_battle_mark_count INT,
            battle_zone_2_battle_mark_count INT,
            battle_zone_3_battle_mark_count INT,
            tactical_skill_group INT,
            tactical_skill_level INT,
            total_gain_vf_credit INT,
            kill_player_gain_vf_credit INT,
            kill_chicken_gain_vf_credit INT,
            kill_boar_gain_vf_credit INT,
            kill_dog_gain_vf_credit INT,
            kill_wolf_gain_vf_credit INT,
            kill_bear_gain_vf_credit INT,
            kill_omega_gain_vf_credit INT,
            kill_bat_gain_vf_credit INT,
            kill_wickeline_gain_vf_credit INT,
            kill_alpha_gain_vf_credit INT,
            kill_item_bounty_gain_vf_credit INT,
            kill_drone_gain_vf_credit INT,
            kill_gamma_gain_vf_credit INT,
            kill_turret_gain_vf_credit INT,
            item_shredder_gain_vf_credit INT,
            total_use_vf_credit INT,
            remote_drone_use_vf_credit_my_self INT,
            remote_drone_use_vf_credit_ally INT,
            transfer_console_from_material_use_vf_credit INT,
            transfer_console_from_escape_key_use_vf_credit INT,
            transfer_console_from_revival_use_vf_credit INT,
            tactical_skill_upgrade_use_vf_credit INT,
            infusion_reroll_use_vf_credit INT,
            infusion_trait_use_vf_credit INT,
            infusion_relic_use_vf_credit INT,
            infusion_store_use_vf_credit INT,
            team_elimination INT,
            team_down INT,
            team_battle_zone_down INT,
            team_repeat_down INT,
            adaptive_force INT,
            adaptive_force_attack INT,
            adaptive_force_amplify INT,
            skill_amp INT,
            camp_fire_craft_uncommon INT,
            camp_fire_craft_rare INT,
            camp_fire_craft_epic INT,
            camp_fire_craft_legendary INT,
            cobalt_random_pick_remove_character INT,
            tactical_skill_use_count INT,
            credit_revival_count INT,
            time_spent_in_briefing_room INT,
            is_leaving_before_credit_revival_terminate BOOLEAN,
            cr_get_animal INT,
            cr_get_mutant INT,
            cr_get_phase_start INT,
            cr_get_kill INT,
            cr_get_assist INT,
            cr_get_time_elapsed INT,
            cr_get_credit_bonus INT,
            cr_use_remote_drone INT,
            cr_use_upgrade_tactical_skill INT,
            cr_use_tree_of_life INT,
            cr_use_meteorite INT,
            cr_use_mythril INT,
            cr_use_force_core INT,
            cr_use_vf_blood_sample INT,
            cr_use_activation_module INT,
            cr_use_rootkit INT,
            mmr_gain_in_game INT,
            mmr_loss_entry_cost INT,
            premade_matching_type INT,
            view_contribution INT,
            use_recon_drone INT,
            use_emp_drone INT,
            except_premade_team BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS action_cost (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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

        CREATE TABLE IF NOT EXISTS weapon_route (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            weapon_route recommended_weapon_route,
            weapon_route_desc recommended_weapon_route_desc,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS area (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            item_code INT,
            capacity INT,
            load_type TEXT,
            time FLOAT4,
            init_count INT,
            count INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS character (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            prevent_basic_attack_damage_ratio FLOAT4,
            prevent_skill_damage_ratio FLOAT4,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            level INT,
            level_up_exp INT
        );

        CREATE TABLE IF NOT EXISTS character_level_up_stat (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            prevent_basic_attack_damage_ratio FLOAT4,
            prevent_skill_damage_ratio FLOAT4,
            name TEXT
        );

        CREATE TABLE IF NOT EXISTS character_mastery (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            character_code INT,
            weapon_type TEXT,
            solo_increase_mode_damage_ratio INT,
            solo_prevent_mode_damage_ratio INT,
            solo_increase_mode_heal_ratio INT,
            solo_increase_mode_shield_ratio INT,
            duo_increase_mode_damage_ratio INT,
            duo_prevent_mode_damage_ratio INT,
            duo_increase_mode_heal_ratio INT,
            duo_increase_mode_healer_give_heal_ratio INT,
            duo_increase_mode_shield_ratio INT,
            duo_increase_mode_healer_give_shield_ratio INT,
            squad_increase_mode_damage_ratio INT,
            squad_prevent_mode_damage_ratio INT,
            squad_increase_mode_heal_ratio INT,
            squad_increase_mode_healer_give_heal_ratio INT,
            squad_increase_mode_shield_ratio INT,
            squad_increase_mode_healer_give_shield_ratio INT,
            cobalt_increase_mode_damage_ratio INT,
            cobalt_prevent_mode_damage_ratio INT,
            cobalt_increase_mode_heal_ratio INT,
            cobalt_increase_mode_healer_give_heal_ratio INT,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            cooldown INT,
            item_code_1 TEXT,
            item_code_2 TEXT,
            probability_1 INT,
            probability_2 INT,
            drop_count INT,
            casting_action_type TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS drop_group (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            group_code INT,
            item_code INT,
            min INT,
            max INT,
            probability INT,
            drop_type TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS gain_exp (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            start_time INT,
            end_time INT,
            gain_exp INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS game_tip (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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

        CREATE TABLE IF NOT EXISTS how_to_find_item (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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

        CREATE TABLE IF NOT EXISTS item_search_option (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            name TEXT,
            tag_1 TEXT,
            tag_2 TEXT,
            tag_3 TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS item_spawn (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            area_code INT,
            area_spawn_group INT,
            item_code INT,
            drop_point TEXT,
            drop_count INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS items (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            exclusive_producer INT,
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
            is_reduce_loot_on_death BOOLEAN,
            is_vpad_quick_slot_item BOOLEAN,
            consumable_type TEXT,
            consumable_tag TEXT,
            heal INT,
            hp_recover INT,
            sp_recover INT,
            attack_power_by_buff INT,
            defense_by_buff INT,
            skill_amp_by_buff INT,
            skill_amp_ratio_by_buff INT,
            add_state_code INT,
            misc_item_type TEXT,
            special_item_type TEXT,
            cooldown_group_code INT,
            cooldown FLOAT4,
            consume_count INT,
            summon_code INT,
            ghost_item_state_group INT,
            skill_amp_by_level INT,
            adaptive_force_by_level INT,
            skill_amp_ratio_by_level FLOAT4,
            armor_type TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS level (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            level INT,
            needed_exp INT,
            accumulate_exp INT,
            reward_acoin INT,
            reward INT
        );

        CREATE TABLE IF NOT EXISTS loading_tip (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            loading_tip_type TEXT,
            min_lv INT,
            max_lv INT,
            text_key TEXT,
            image_name TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS mastery_exp (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode_type INT,
            condition_type TEXT,
            grade TEXT,
            condition_value INT,
            mastery_type_1 TEXT,
            value_1 INT,
            mastery_type_2 TEXT,
            value_2 INT,
            mastery_type_3 TEXT,
            value_3 INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS mastery_stat (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            monster_code INT,
            monster_level INT,
            drop_group INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS monster_level_up_stat (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode INT,
            player_level INT,
            monster_code INT,
            spawn_level INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS navi_collect_and_hunt (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            item_code INT,
            area_code_list TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS near_by_area (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            area_code INT,
            near_by_area_code INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS random_equipment (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            season_id INT,
            season_name TEXT,
            season_start DATE,
            season_end DATE,
            is_current BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS summon_object_stat (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            next_upgrade_code INT,
            upgrade_credit INT,
            upgrade_material INT,
            skill_code INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS tactical_skill_set_group (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            "group" INT,
            mode_type INT,
            start_code INT,
            equip_with_start BOOLEAN,
            icon TEXT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS trait (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            open_account_lv INT,
            trait_group TEXT,
            trait_type TEXT,
            active BOOLEAN,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS transfer_console (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
            code INT,
            mode INT,
            phase INT,
            condition_type TEXT,
            condition_value INT,
            acquire_self FLOAT4,
            acquire_team INT,

            PRIMARY KEY (id)
        );

        CREATE TABLE IF NOT EXISTS weapon_type_info (
            id BIGINT GENERATED BY DEFAULT AS IDENTITY,
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

        -- user_info
        CREATE INDEX IF NOT EXISTS idx_user_info_user_num ON user_info (user_num);
        CREATE INDEX IF NOT EXISTS idx_user_info_nickname ON user_info USING GIN (nickname gin_trgm_ops);

        -- user_rank
        CREATE INDEX IF NOT EXISTS idx_user_rank_user_num ON user_rank (user_num);

        -- top_rank
        CREATE INDEX IF NOT EXISTS idx_top_rank_user_num ON top_rank (user_num);

        -- user_stats
        CREATE INDEX IF NOT EXISTS idx_user_stats_user_num_season_id ON user_stats (user_num, season_id);
        CREATE INDEX IF NOT EXISTS idx_user_stats_season_id_mmr ON user_stats (season_id, mmr);

        -- user_game
        CREATE INDEX IF NOT EXISTS idx_user_game_game_id_user_num ON user_game (game_id, user_num);
        CREATE INDEX IF NOT EXISTS idx_user_game_character_num_game_id ON user_game (character_num, game_id);
        CREATE INDEX IF NOT EXISTS idx_user_game_season_id ON user_game (season_id);
        CREATE INDEX IF NOT EXISTS idx_user_game_version_major_version_minor ON user_game (version_major, version_minor);
        CREATE INDEX IF NOT EXISTS idx_user_game_server_name_season_id ON user_game (server_name, season_id);

        -- action_cost
        CREATE INDEX IF NOT EXISTS idx_action_cost_code ON action_cost (code);

        -- weapon_route
        CREATE INDEX IF NOT EXISTS idx_weapon_route_id ON weapon_route (((weapon_route).route_id));

        -- area
        CREATE INDEX IF NOT EXISTS idx_area_code ON area (code);
        CREATE INDEX IF NOT EXISTS idx_area_name ON area (name);

        -- battle_zone_reward
        CREATE INDEX IF NOT EXISTS idx_battle_zone_reward_code_mode_type_selectable ON battle_zone_reward (code, mode_type, selectable);

        -- bullet_capacity
        CREATE INDEX IF NOT EXISTS idx_bullet_capacity_item_code ON bullet_capacity (item_code);

        -- character
        CREATE INDEX IF NOT EXISTS idx_character_code ON character (code);
        CREATE INDEX IF NOT EXISTS gidx_character_name ON character USING GIN (name gin_trgm_ops);

        -- character_attributes
        CREATE INDEX IF NOT EXISTS idx_character_attributes_character_code ON character_attributes (character_code);

        -- character_level_up_stat
        CREATE INDEX IF NOT EXISTS idx_character_level_up_stat_code ON character_level_up_stat (code);

        -- item_armor
        CREATE INDEX IF NOT EXISTS idx_item_code ON items (code);
        CREATE INDEX IF NOT EXISTS idx_item_make_material_1 ON items (make_material_1);
        CREATE INDEX IF NOT EXISTS idx_item_make_material_2 ON items (make_material_2);
        CREATE INDEX IF NOT EXISTS gidx_item_name ON items USING GIN (name gin_trgm_ops);

        -- character_mastery
        CREATE INDEX IF NOT EXISTS idx_character_mastery_code ON character_mastery (code);

        -- character_mode_modifier
        CREATE INDEX IF NOT EXISTS idx_character_mode_modifier_character_code_weapon_type ON character_mode_modifier (character_code, weapon_type);

        -- character_skin
        CREATE INDEX IF NOT EXISTS idx_character_skin_character_code ON character_skin (character_code);
        CREATE INDEX IF NOT EXISTS gidx_character_skin_name ON character_skin USING GIN (name gin_trgm_ops);

        -- collectible
        CREATE INDEX IF NOT EXISTS idx_collectible_code ON collectible (code);

        -- drop_group

        -- how_to_find_item
        CREATE INDEX IF NOT EXISTS idx_how_to_find_items_item_code ON how_to_find_item (item_code) INCLUDE (hunt_chicken, hunt_bat, hunt_boar, hunt_dog, hunt_wolf, hunt_bear, hunt_wickeline, hunt_alpha, hunt_omega, air_supply);


        -- item_search_option
        CREATE INDEX IF NOT EXISTS idx_item_search_option ON item_search_option (code);

        -- item_spawn
        CREATE INDEX IF NOT EXISTS idx_item_spawn_item_code ON item_spawn (item_code);
        CREATE INDEX IF NOT EXISTS idx_item_spawn_area_code ON item_spawn (area_code);



END $$