package main

import (
	"fmt"
	"math"
	mzwuwacal "wuwa_cal/mz_wuwa_cal"
)

func main() {
	// Create an attacker
	attacker := mzwuwacal.Attacker{
		Name:               "RoverAero",
		Level:              90,
		AbilityDmg:         0.3531, // 35.31% of base damage (NA1 Rover Aero level 10)
		TotalAtk:           1183,
		Crit:               false,
		ElementDmgBonus:    0.0,
		AttackTypeDmgBonus: 0.0,
		DmgAmplify:         0.0,
		CritDmg:            1.5, // 150% of crit damage
		DefIgnore:          0.0,
		ElementResIgnore:   0.0,
	}

	// Create a target for testing
	target := mzwuwacal.Target{
		Name:           "Mourning Aix",
		Level:          85,
		BaseElementRes: 0.1,
		DmgTaken:       0.0,
	}

	// Print all calculation steps
	fmt.Println("True Damage:", math.Ceil(mzwuwacal.GetTrueDamage(&attacker, &target)))
	fmt.Println("Base Damage:", mzwuwacal.GetBaseDamage(&attacker))
	fmt.Println("All Resistances:", mzwuwacal.GetAllResistances(&attacker, &target))
	fmt.Println("Base Element Res:", mzwuwacal.GetBaseElementRes(&attacker, &target))
	fmt.Println("Def Multiplier:", mzwuwacal.GetDefMultiplier(&attacker, &target))
	fmt.Println("Dmg Reduce:", mzwuwacal.GetDmgReduce(&target))
	fmt.Println("Element Reduce:", mzwuwacal.GetElementReduce(&target))
	fmt.Println("Bonus Multipliers:", mzwuwacal.GetBonusMultipliers(&attacker, &target))
	fmt.Println("Dmg Bonus:", mzwuwacal.GetDmgBonus(&attacker))
	fmt.Println("Dmg Amplify:", mzwuwacal.GetDmgAmplify(&attacker, &target))
	fmt.Println("Special Dmg:", mzwuwacal.GetSpecialDmg(&attacker))
}
