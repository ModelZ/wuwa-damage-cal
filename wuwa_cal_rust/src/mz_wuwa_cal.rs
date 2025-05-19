#[derive(Debug)]
pub struct Attacker {
    pub name: String,
    pub level: i32,
    pub ability_dmg: f64,
    pub total_atk: i64,
    pub crit: bool,
    pub element_dmg_bonus: f64,
    pub attack_type_dmg_bonus: f64,
    pub dmg_amplify: f64,
    pub crit_dmg: f64,
    pub def_ignore: f64,
    pub element_res_ignore: f64,
    pub special_dmg: f64,
}

#[derive(Debug)]
pub struct Target {
    pub name: String,
    pub level: i32,
    pub base_element_res: f64,
    pub dmg_taken: f64,
    pub dmg_res_base: f64,
    pub dmg_res_additional: f64,
    pub element_res_base: f64,
    pub element_res_additional: f64,
    pub dmg_amplify: f64,
}

pub fn get_true_damage(a: &Attacker, t: &mut Target) -> f64 {
    get_base_damage(a) * get_all_resistances(a, t) * get_bonus_multipliers(a, t)
}

pub fn get_base_damage(a: &Attacker) -> f64 {
    a.ability_dmg * a.total_atk as f64
}

pub fn get_all_resistances(a: &Attacker, t: &mut Target) -> f64 {
    get_base_element_res(a, t)
        * get_def_multiplier(a, t)
        * get_dmg_reduce(t)
        * get_element_reduce(t)
}

pub fn get_base_element_res(a: &Attacker, t: &Target) -> f64 {
    let res_total = t.base_element_res - a.element_res_ignore;

    if res_total < 0.0 {
        1.0 - (res_total / 2.0)
    } else if res_total < 0.8 {
        1.0 - res_total
    } else {
        1.0 / (1.0 + (5.0 * res_total))
    }
}

pub fn get_def_multiplier(a: &Attacker, t: &Target) -> f64 {
    let target_def = (8 * t.level + 792) as f64;
    let attacker_multiplier = (800 + 8 * a.level) as f64;
    attacker_multiplier / (attacker_multiplier + target_def * (1.0 - a.def_ignore))
}

pub fn get_dmg_reduce(t: &mut Target) -> f64 {
    t.dmg_res_base = 0.0;
    t.dmg_res_additional = 0.0;
    1.0 - (t.dmg_res_base + t.dmg_res_additional) + t.dmg_taken
}

pub fn get_element_reduce(t: &mut Target) -> f64 {
    t.dmg_res_base = 0.0;
    t.dmg_res_additional = 0.0;
    1.0 - (t.element_res_base + t.element_res_additional)
}

pub fn get_bonus_multipliers(a: &Attacker, t: &mut Target) -> f64 {
    if a.crit {
        get_dmg_bonus(a) * get_dmg_amplify(a, t) * get_special_dmg(a) * a.crit_dmg
    } else {
        get_dmg_bonus(a) * get_dmg_amplify(a, t) * get_special_dmg(a)
    }
}

pub fn get_dmg_bonus(a: &Attacker) -> f64 {
    1.0 + a.element_dmg_bonus + a.attack_type_dmg_bonus
}

pub fn get_dmg_amplify(a: &Attacker, t: &mut Target) -> f64 {
    t.dmg_amplify = 0.0;
    1.0 + a.dmg_amplify + t.dmg_amplify
}

pub fn get_special_dmg(a: &Attacker) -> f64 {
    1.0 + a.special_dmg
}
