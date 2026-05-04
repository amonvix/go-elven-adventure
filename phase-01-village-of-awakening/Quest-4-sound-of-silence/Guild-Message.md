# ⚔️ Phase 01 — Village of Awakening

## Quest 4: The Silent Defaults

The winds of Eldora grow quiet.

Beyond the Hall of Records lies a hidden truth —  
even when nothing is declared, something already exists.

The Guild Master whispers:

> "Not all values are given by choice.  
> Some are granted by the language itself.  
> Learn what is given… before you try to control it."

---

## Objective

Understand how Go initializes variables automatically when no value is assigned.

---

## Concepts Learned

- zero values
- default initialization
- var without assignment
- behavior of different types when uninitialized

---

## Your Task

Create a new file:

nameIt.go

Inside it, declare the following variables using `var`, WITHOUT assigning any value:

- heroName (string)
- heroLevel (int)
- heroMana (float64)
- heroAlive (bool)

Then print all variables.

---

## Expected Discovery

You will notice:

- string → empty
- int → 0
- float64 → 0.0
- bool → false

Example output:

Hero Name:  
Hero Level: 0  
Hero Mana: 0  
Hero Alive: false  

---

## Side Quest

After printing the default values, assign real values to the same variables and print again.

Observe the difference between:

- uninitialized state
- initialized state

---

## Mini Boss Fight

The Royal Oracle has failed to initialize a warrior.

She insists:

> "The hero was never assigned a value… yet he still exists."

Your challenge:

1. Declare a new variable:

   var unknownHero string

2. Print it BEFORE assigning anything.

3. Then assign:

   unknownHero = "Elyra"

4. Print it again.

---

## Bonus Insight

Ask yourself:

- Why does Go not crash when variables are uninitialized?
- What advantage does this bring in production systems?

---

## Reward Unlocked

- You understand default memory behavior in Go
- You reduce risk of undefined values
- You gain safer initialization patterns

---

## Victory Condition

- Program runs successfully
- Default values are printed
- Updated values are printed
- Mini Boss completed

---

## Next Phase →

move.forward()