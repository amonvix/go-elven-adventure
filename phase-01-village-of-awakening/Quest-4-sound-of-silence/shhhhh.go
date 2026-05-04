package main

import "fmt"

var unknownHero = "Elyra"

func main() {
	fmt.Println("This chapter of Amon's passage in the village was marked by an unknown hero, and it gets more strange as it goes. A name", unknownHero, "can be smelled in the air, but no one who she is! Maybe Amon...")
	var heroName string
	var heroLevel int
	var heroMana float64
	var heroAlive bool

	fmt.Println("Some guards was interrogating the people in the village regards a body found at the North area of the church")
	fmt.Println("They made the same four questions for everybody, but the responses was allways the same ")
	fmt.Println(heroName, "Was the name of that pour soul. He was found holding a Level badge that says, ", heroLevel)
	fmt.Println("the guards realized that this young man tried to survive at most, since his mana was", heroMana, "while the alive status was totally", heroAlive)

	fmt.Println("Fast-forward two weeks")
	fmt.Println("the guards closed the investigations with a clear explanation of the case")

	heroName = "Gunther"
	heroLevel = 2
	heroMana = 300.0
	heroAlive = true

	fmt.Println("The body found, once was", heroName, "and he was chopping trees, using his magic abilities as a forest dwarf")
	fmt.Println("Checking with more villagers, the guards had the information that Gunther alive status was", heroAlive, "but having", heroLevel, "he did not see the black bear coming very slowly on his back")
	fmt.Println("unfortunately, that was the end of the Gunther lineage, and the case too")

	// Mini Boss Fight

	fmt.Println("for some reason, some people in village was directing the guards to talk to an unknownHero, named", unknownHero, "but she was never found in the village or anywhere else!")

}
