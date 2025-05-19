package mzwuwacal

type Attacker struct {
	// Basic info of the attacker
	name        string
	level       int
	element     string
	attack_Type int

	// true damage of the attacker
	true_dmg float64

	// Base damage of the attacker
	base_dmg    int64
	ability_dmg float64
	// total attack of the attacker
	total_atk     int64
	base_atk      int64
	atk_buff      float64
	flat_atk_buff int64
	flat_atk      int64

	// DMG Bonus of the attacker
	bonus_multiplier      float64
	dmg_bonus             float64
	element_dmg_bonus     float64
	attack_type_dmg_bonus float64
	dmg_amplify           float64
	special_dmg           float64
	crit_dmg              float64
}

type Target struct {
	// Basic info of the target
	name    string
	level   int
	element string

	// all resistances of the target
	all_resistances    float64
	base_element_res   float64
	def_multiplier     float64
	dmg_reduce         float64
	element_dmg_reduce float64

	// negative resistances of the target
	def_ignore         float64 // <-> def
	dmg_taken          float64 // <-> dmg_reduce
	element_res_ignore float64 // <-> base_element_res

	// Unknown Source
	dmg_res_base           float64
	dmg_res_additional     float64
	element_res_base       float64
	element_res_additional float64
	dmg_amplify            float64
}

// Damage Calculation
func GetTrueDamage(a *Attacker, t *Target) float64 {
	return GetBaseDamage(a) * GetAllResistances(a, t) * a.bonus_multiplier
}

func GetBaseDamage(a *Attacker) float64 {
	return a.ability_dmg * GetTotalAttack(a)
}

func GetTotalAttack(a *Attacker) float64 {
	return float64(a.base_atk)*(1+a.atk_buff) + float64(a.flat_atk) + float64(a.flat_atk_buff)
}

// Resistances Calculation
func GetAllResistances(a *Attacker, t *Target) float64 {
	return GetBaseElementRes(t) * GetDefMultiplier(a, t) * GetDmgReduce(t) * GetElementReduce(t)
}

func GetBaseElementRes(t *Target) float64 {

	res_total := t.base_element_res - t.element_res_ignore

	if res_total < 0 {
		return 1 - (res_total / 2)
	} else if res_total < 0.8 {
		return 1 - res_total
	} else { // res_total => 0.8
		return 1 / (1 + (5 * res_total))
	}

}

func GetDefMultiplier(a *Attacker, t *Target) float64 {
	target_def := (8 * t.level) + 792
	attacker_multiplier := float64(800 + (8 * a.level))
	return attacker_multiplier / (attacker_multiplier + float64(target_def)*(1-t.def_ignore))
}

func GetDmgReduce(t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.dmg_res_base = 0.0
	t.dmg_res_additional = 0.0
	// I think damage taken is here
	return 1 - (t.dmg_res_base + t.dmg_res_additional) + t.dmg_taken
}

func GetElementReduce(t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.dmg_res_base = 0.0
	t.dmg_res_additional = 0.0
	return 1 - (t.element_res_base + t.element_res_additional)
}

// Bonus Multipler Calculation
func GetBonusMultipliers(a *Attacker) float64 {
	return a.dmg_bonus * a.dmg_amplify * a.special_dmg * a.crit_dmg
}

// element_dmg_bonus% -> Spectro/Aero/Glacio/Fusion/Electro/Havoc DMG bonus
// atk_type_dmg_bonus% -> basic/heavy/reson_skill/reson_riberation DMG Bonus
func GetDmgBonus(a *Attacker) float64 {
	return 1 + (a.element_dmg_bonus + a.attack_type_dmg_bonus)
}

func GetDmgAmplify(a *Attacker, t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.dmg_amplify = 0.0
	return 1 + a.dmg_amplify + t.dmg_amplify
}

func GetSpecialDmg(a *Attacker) float64 {
	// Maybe Unused I just set it as 0
	a.special_dmg = 0.0
	return 1 + a.special_dmg
}
