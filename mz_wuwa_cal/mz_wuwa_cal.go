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
	bonus_multiplier float64
	dmg_bonus        float64
	dmg_amplify      float64
	special_dmg      float64
	crit_dmg         float64
}

type Target struct {
	// Basic info of the target
	name    string
	def     int64
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
}

// Damage Calculation
func GetTrueDamage(a *Attacker, t *Target) float64 {
	return GetBaseDamage(a) * GetAllResistances(t) * a.bonus_multiplier
}

func GetBaseDamage(a *Attacker) float64 {
	return a.ability_dmg * GetTotalAttack(a)
}

func GetTotalAttack(a *Attacker) float64 {
	return float64(a.base_atk)*(1+a.atk_buff) + float64(a.flat_atk) + float64(a.flat_atk_buff)
}

// Resistances Calculation
func GetAllResistances(t *Target) float64 {
	return GetBaseElementRes(t) * t.def_multiplier * t.dmg_reduce * t.element_dmg_reduce
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
