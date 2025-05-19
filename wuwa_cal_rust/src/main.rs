mod mz_wuwa_cal;

use mz_wuwa_cal::*;
use std::f64;

fn main() {
    let attacker = Attacker {
        name: "RoverAero".to_string(),
        level: 90,
        ability_dmg: 0.3531,
        total_atk: 1183,
        crit: false,
        element_dmg_bonus: 0.0,
        attack_type_dmg_bonus: 0.0,
        dmg_amplify: 0.0,
        crit_dmg: 1.5,
        def_ignore: 0.0,
        element_res_ignore: 0.0,
        special_dmg: 0.0,
    };

    let mut target = Target {
        name: "Mourning Aix".to_string(),
        level: 85,
        base_element_res: 0.1,
        dmg_taken: 0.0,
        dmg_res_base: 0.0,
        dmg_res_additional: 0.0,
        element_res_base: 0.0,
        element_res_additional: 0.0,
        dmg_amplify: 0.0,
    };

    println!("True Damage: {}", get_true_damage(&attacker, &mut target).ceil());
    println!("Base Damage: {}", get_base_damage(&attacker));
    println!("All Resistances: {}", get_all_resistances(&attacker, &mut target));
    println!("Base Element Res: {}", get_base_element_res(&attacker, &target));
    println!("Def Multiplier: {}", get_def_multiplier(&attacker, &target));
    println!("Dmg Reduce: {}", get_dmg_reduce(&mut target));
    println!("Element Reduce: {}", get_element_reduce(&mut target));
    println!("Bonus Multipliers: {}", get_bonus_multipliers(&attacker, &mut target));
    println!("Dmg Bonus: {}", get_dmg_bonus(&attacker));
    println!("Dmg Amplify: {}", get_dmg_amplify(&attacker, &mut target));
    println!("Special Dmg: {}", get_special_dmg(&attacker));
}
