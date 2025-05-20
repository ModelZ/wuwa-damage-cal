package main

import (
	"fmt"
	"math"
	"os"

	mzwuwacal "wuwa_cal/mz_wuwa_cal"

	"github.com/pelletier/go-toml/v2"
)

type InputData struct {
	Attacker mzwuwacal.Attacker `toml:"attacker"`
	Target   mzwuwacal.Target   `toml:"target"`
}

func main() {
	// Read TOML input from file
	file, err := os.Open("input.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input InputData
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&input); err != nil {
		panic(err)
	}

	attacker := input.Attacker
	target := input.Target

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

	fmt.Println("Press Any Key to Exit...")
	var inputStr string
	fmt.Scanln(&inputStr)
}
