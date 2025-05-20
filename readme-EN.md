# WuWa Damage Calculator

This project implements a basic damage calculation model for a game, based on formulas described in [dmg_formula.txt](dmg_formula.txt).

## Project Structure

- `main.go`: Entry point for the application.
- `mz_wuwa_cal/mz_wuwa_cal.go`: Core structs and functions for damage calculation.
- `dmg_formula.txt`: Reference formulas and explanations for damage calculation.
- `input.toml`: Example input file for attacker and target configuration.
- `readme.md`: Project documentation.

## Prerequisite
- Download programming langauge [Go](https://go.dev/doc/install)

## Usage

1. **Edit `input.toml`**  
   Configure your attacker and target stats in `input.toml`.  
   Example:
   ```toml
    # Attacker name
    Name = "RoverAero"

    # Level
    Level = 90

    # %Skill damage (35.31%)
    AbilityDmg = 0.3531

    # Total attack (as shown in the character stat)
    TotalAtk = 1183

    # Should it crit? false = crits, true = does not crit
    Crit = false

    # %Elemental damage bonus (Spectro/Aero/Fusion/Havoc/Glacio)
    ElementDmgBonus = 0.0

    # %Attack type damage bonus (basic/heavy/resonance skill/resonance liberation)
    AttackTypeDmgBonus = 0.0

    # %Final damage (similar to Verina’s Outro)
    DmgAmplify = 0.0

    # %Critical damage (150%)
    CritDmg = 1.5

    # Ignore defense %
    DefIgnore = 0.0

    # Ignore elemental resistance %
    ElementResIgnore = 0.0

    [target]
    # Monster name
    Name = "Mourning Aix"

    # Monster level
    Level = 85

    # Elemental resistance %
    BaseElementRes = 0.1

    # Takes additional damage %
    DmgTaken = 0.0

   ```

2. **Build the project:**
   ```sh
   go mod tidy
   go build main.go
   ```

## Features

- Calculate true damage based on attacker and target stats.
- Modular design for easy extension.
- User-friendly TOML input with comments support.

## References

See [dmg_formula.txt](dmg_formula.txt) for detailed calculation formulas.