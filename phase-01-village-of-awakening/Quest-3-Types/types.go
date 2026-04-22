package main

import "fmt"

func main() {
	hero := "Amon"
	level := 1
	HP := 100
	mana := 85.5
	alive := true

	var marriage bool

	marriage = true

	fmt.Println(hero, "who has level", level, "arrives at the destiny tavern")
	fmt.Println("We are happy that you came with full HP at", HP, "Looking the usage of your mana", mana, "I can see why alive status keeps showing", alive)
	fmt.Println("If you allowed us to ask about this gorgeous litlle teenage elf with you. Whath is her marriage status at moment")
	marriage = true
	fmt.Println("Amon quickly holds one of Lilith hands and the sword in the other, answerin a sonorous", marriage)

	/*
		## Mini Boss Fight

		The Royal Clerk made a terrible mistake.

		He wrote the hero level as text.

		Correct the registry.

		Create:

		- one wrong example as string: "41"
		- one correct example as int: 41

		Print both and identify which one is valid for calculations.

		Example:

		Wrong Level: 41
		Correct Level: 41
		Only one of them can be used in math.

		Bonus challenge:

		Print the result of:

		Correct Level + 1

		Expected:

		42
	*/

	fmt.Println("Suddenly, Amon is summoned to the Castle to provide clarification regarding his level.")
	amon_citizen_registry := 41
	clerk_citizen_registry := "41"

	fmt.Println("If all the problem is about only this, I show you my documents provided when I arrive, says Amon", amon_citizen_registry+1)
	fmt.Println("Doing a simple comparison, everybody could see that Amons's Registry was correct. The Clerk registry was", clerk_citizen_registry)

}
