package main

import "fmt"

func main() {
	fmt.Println("In this part of the journey, Amon is called to the sacred chamber to pass through the strength ritual")

	hero := "Amon"
	role := "Necromancer"
	mana := 2000
	magicLevel := 300
	sorcererLevel := 200

	maskGolden := "Golden"
	maskSilver := "Silver"
	maskBlack := "Black"

	switch role {
	case "Necromancer":
		fmt.Println("GREAT!!! You choose", maskBlack)
		fmt.Println("This is the mask corret for you", hero, "! From now on, you can carry this black mask to add", mana, "into your skills, letting you with", mana*2, "Mana")

	case "Sorcerer":
		fmt.Println("Sorry. You choose", maskSilver)
		fmt.Println("This is the mask corret for Sorcerers! The possessor of this artifact increase", magicLevel, "into their skills, letting it rise", magicLevel*2, " in the magic level")
	case "Druid":
		fmt.Println("Bad choise, You choose", maskGolden)
		fmt.Println("This is the mask corret for Druids! You as a Necromancer, cannot even touch in one of this.\n The power that holds in this artifact is greatfull. It's almost imortality. The Health Points rises 5000")

	default:
		fmt.Println("If you cannot decide witch one collect now, we can make the ritual in the next year!")
	}
	fmt.Println("At the middle of the cerimony, Amon remembered about the other village, where he mets Lilith and pulls a scroll with the Kingdom Stamp")
	fmt.Println("With this new information, some never saw happened... all 3 masks were collected by only one hero\n Respecting the rules, the Scroll said: Amonvix, Sorcerer, MagicLevel 200 ")

	switch hero {
	case "Necromancer":
		fmt.Print("mask", maskBlack, "Collect the artifact and now his mana is", mana*2)
	case "Sorcerer":
		fmt.Print("mask", maskSilver, "Collect the artifact and her magic level now is", sorcererLevel*2)
	case "Druid":
		fmt.Print(hero, "Cannot collect or touch the artifact")
	default:
		fmt.Println("The real revelation happens.\n", hero, "collect Sorcerer and Necromancer")
		fmt.Println("Lilith shows up flying at the center of the table and the brilliant mask", maskGolden, "reveals to all there, that she was an Elfic Angel, receiving more 5000HP")
	}
	Elven := true
	fmt.Println("Now, everyone in the village nows about the lineage of Lilith, she is a", Elven, "for sure!")

}
