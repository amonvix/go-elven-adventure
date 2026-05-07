package main

import "fmt"

func main() {
	heroName := "Amon"
	heroLevel := 15
	heroAlive := true

	fmt.Println("As the days goes by, our Hero,", heroName, "spend the days doing little jobs aroud the village")
	fmt.Println("Now, carring a bagde that says 'Hero Level'", heroLevel, ". Amon suddenly receives a Guilds letters.")
	fmt.Println("The letter says... Death hunter Amon, now that we are fortunate enough to have a necromancer among our citizens, could you please lend your attention to this urgent mission?")
	fmt.Println("The only requirement to assume this quest is have the status Alive", heroAlive, ". Now, as a necromancer, you shall decide whether you have the ability to proceed or prefer to refuse.")

	fmt.Println("The letter was just sent to Amon for one reason. Once inside the forest,you have to obey the Alive rule. And the rule is:")

	if heroAlive {
		fmt.Println(heroName, "Will be allowed to enter safelly in the forest if status is marked as Alive")
	} else {
		fmt.Println("The journey ends here. With the forest claims their souls. Amon and Lilith will be traped for ever in the darkness")
	}

	fmt.Println("The Guild mission is enter in the tower in the middle of the forrest, but since this is a forrest of decisions, the door will open only if a condition is respected:")

	if heroLevel > 10 {
		fmt.Println("Only heroes greater than level 10 shall enter")
	} else {
		fmt.Println("The door will not open and a huge noise will be spread to call every single hostile creature to atack the invader")
	}

	// Mini Boss Fight
	fmt.Println("Inside the tower, in the last floor, relays the biggest ogre ever saw in the forrest, and the Gild request was just to bring back a proof of the ogre's death")
	fmt.Println("And cannot be different, we faced a condition here too, that was whrote in the Guild's letter!\nTo defeat the ogre you both have to be Alive and have a level grater than five")
	fmt.Println("If both are true, with the help of the gods, you can defeat that beast!")

	partyHero := "Lilith"
	partyLevel := 35
	partyAlive := true

	if (partyAlive && heroAlive) && (partyLevel > 5 && heroLevel > 5) {
		fmt.Println(heroName, "and", partyHero, "worked together and easily cut the ogre's head off")
	} else {
		fmt.Println(heroName, "and", partyHero, "were not capable to execute the task and their body will be as part of the tower for ever")
	}

	fmt.Println("Returning to the Guild master's at the head quarter, Amon's just drop a dirty bag in the table...\nInside only one thing! The ogre's head, for the honor of the Guild")

}
