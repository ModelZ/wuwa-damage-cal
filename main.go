package main

import (
	"fmt"

	mzwuwacal "./mz_wuwa_cal"
)

func main() {
	// Create an attacker
	attacker := mzwuwacal.Attacker{
		Name:               "Attacker",
		Level:              90,
		BaseDmg:            1000,
		AbilityDmg:         1.5,
		TotalAtk:           2000,
		ElementDmgBonus:    0.2,
		AttackTypeDmgBonus: 0.3,
		DmgAmplify:         0.4,
		CritDmg:            0.5,
		DefIgnore:          0.6,
		ElementResIgnore:   0.7,
	}

	fmt.Println("Attacker:", attacker)

	// Create a target
	target := mzwuwacal.Target{
		Name:                 "Target",
		Level:                90,
		BaseElementRes:       0.1,
		DmgTaken:             0.2,
		DmgResBase:           0.3,
		DmgResAdditional:     0.4,
		ElementResBase:       0.5,
		ElementResAdditional: 0.6,
		DmgAmplify:           0.7,
	}

	fmt.Println("Target:", target)
}
