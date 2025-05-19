package mzwuwacal

type Attacker struct {
	// Basic info of the attacker
	Name  string
	Level int

	// Base damage of the attacker
	AbilityDmg float64

	// Total attack of the attacker
	TotalAtk int64

	// DMG Bonus of the attacker
	ElementDmgBonus    float64
	AttackTypeDmgBonus float64
	DmgAmplify         float64
	CritDmg            float64

	// penetration multiplier
	DefIgnore        float64 // <-> def
	ElementResIgnore float64 // <-> base_element_res

	// Unknown Source and not use in json
	SpecialDmg float64
}

type Target struct {
	// Basic info of the target
	Name  string
	Level int
	// Element string

	// all resistances of the target
	BaseElementRes float64

	// negative resistances of the target
	DmgTaken float64 // <-> dmg_reduce

	// Unknown Source and not use in json
	DmgResBase           float64
	DmgResAdditional     float64
	ElementResBase       float64
	ElementResAdditional float64
	DmgAmplify           float64
}

// Damage Calculation
func GetTrueDamage(a *Attacker, t *Target) float64 {
	return GetBaseDamage(a) * GetAllResistances(a, t) * GetBonusMultipliers(a, t)
}

func GetBaseDamage(a *Attacker) float64 {
	return a.AbilityDmg * float64(a.TotalAtk)
}

// No need to use this function TOO DETAILS
// func GetTotalAttack(a *Attacker) float64 {
// 	return float64(a.BaseAtk)*(1+a.AtkBuff) + float64(a.FlatAtk) + float64(a.FlatAtkBuff)
// }

// Resistances Calculation
func GetAllResistances(a *Attacker, t *Target) float64 {
	return GetBaseElementRes(a, t) * GetDefMultiplier(a, t) * GetDmgReduce(t) * GetElementReduce(t)
}

func GetBaseElementRes(a *Attacker, t *Target) float64 {

	resTotal := t.BaseElementRes - a.ElementResIgnore

	if resTotal < 0 {
		return 1 - (resTotal / 2)
	} else if resTotal < 0.8 {
		return 1 - resTotal
	} else { // resTotal => 0.8
		return 1 / (1 + (5 * resTotal))
	}

}

func GetDefMultiplier(a *Attacker, t *Target) float64 {
	targetDef := (8 * t.Level) + 792
	attackerMultiplier := float64(800 + (8 * a.Level))
	return attackerMultiplier / (attackerMultiplier + float64(targetDef)*(1-a.DefIgnore))
}

func GetDmgReduce(t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.DmgResBase = 0.0
	t.DmgResAdditional = 0.0
	// I think damage taken is here
	return 1 - (t.DmgResBase + t.DmgResAdditional) + t.DmgTaken
}

func GetElementReduce(t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.DmgResBase = 0.0
	t.DmgResAdditional = 0.0
	return 1 - (t.ElementResBase + t.ElementResAdditional)
}

// Bonus Multipler Calculation
func GetBonusMultipliers(a *Attacker, t *Target) float64 {
	return GetDmgBonus(a) * GetDmgAmplify(a, t) * GetSpecialDmg(a) * a.CritDmg
}

// element_dmg_bonus% -> Spectro/Aero/Glacio/Fusion/Electro/Havoc DMG bonus
// atk_type_dmg_bonus% -> basic/heavy/reson_skill/reson_riberation DMG Bonus
func GetDmgBonus(a *Attacker) float64 {
	return 1 + (a.ElementDmgBonus + a.AttackTypeDmgBonus)
}

func GetDmgAmplify(a *Attacker, t *Target) float64 {
	// Maybe Unused I just set it as 0
	t.DmgAmplify = 0.0
	return 1 + a.DmgAmplify + t.DmgAmplify
}

func GetSpecialDmg(a *Attacker) float64 {
	// Maybe Unused I just set it as 0
	a.SpecialDmg = 0.0
	return 1 + a.SpecialDmg
}
