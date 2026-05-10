package main

import "fmt"

func main() {
	hero := "Amon"
	level := 45
	role := "Necromancer"
	master := true
	mana := 15000
	hp := 10000

	fmt.Println("To initiate the ceremony", hero, "present your credentials")

	if master && level > 30 {
		fmt.Println(hero, "are allowed to receive the 'Crow of Fate'")
		for i := 1; i < 6; i++ {
			fmt.Println("Necromancer, Warrior, Druid ,Archer or Cunning, asks the Crown")
		}
		fmt.Println("Let's find out who will be THE ONE!!!")
	} else {
		fmt.Println("The Oracle show us the prophecy and now all the presents nows who is the real Master")
	}

	if role == "Necromancer" && master && level > 30 {
		fmt.Println("Valiant Necromancer", hero, "according to the prophecies, you are the one.\nAnd for this reason, from now to ever, you shall be knowed as a Necropolis Master")
		fmt.Println("Here are you recompenses for the effort in protect all of us\nYour mana will be rised to", mana+10000, ", and your HP", hp+5000)
	} else if role == "Warrior" && master && level > 30 {
		fmt.Println("Valiant Warrior", hero, "according to the prophecies, you are the one.\nAnd for this reason, from now to ever, you shall be knowed as a Crusade Master")
		fmt.Println("Here are you recompenses for the effort in protect all of us\nYour mana will be rised to", mana+100, ", and your HP", hp+25000)
	} else if role == "Druid" && master && level > 30 {
		fmt.Println("Valiant Druid", hero, "according to the prophecies, you are the one.\nAnd for this reason, from now to ever, you shall be knowed as a Yggdrasil Master")
		fmt.Println("Here are you recompenses for the effort in protect all of us\nYour mana will be rised to", mana+50000, ", and your HP", hp+30000)
	} else if role == "Archer" && master && level > 30 {
		fmt.Println("Valiant Archer", hero, "according to the prophecies, you are the one.\nAnd for this reason, from now to ever, you shall be knowed as a Paladin Master")
		fmt.Println("Here are you recompenses for the effort in protect all of us\nYour mana will be rised to", mana+2000, ", and your HP", hp+10000)
	} else if role == "Cunning" && master && level > 30 {
		fmt.Println("Cunning", hero, "according to the prophecies, your role is not reliable enough to wield such power.\nFind a better way to use your skills and maybe the Oracle can recognize you as a real Hero")
		fmt.Println("Since our Kingdom recognize the value of every citizen, as a Guild Master, I offer you a place in our table, if you accept to change your role")
	} else {
		fmt.Println(hero, " was not the name once wrote in the prophecies")
	}

	var newrole string

	switch role {
	case "Necromancer":
		newrole = "Necropolis Master"
	case "Warrior":
		newrole = "Crusade Master"
	case "Druid":
		newrole = "Yggdrasil Master"
	case "Archer":
		newrole = "Paladin Master"
	case "Cunning":
	}

	fmt.Println("At the end of the day, all villagers was celebrating the new Kingdom Protector")
	fmt.Println("And Lilith was more proud than every one, feeling that she made the correct choise, when left the Elven Lands to travel with Amon")
	fmt.Println("After the celebration, the information was spread to the four directions")
	fmt.Println("The great", hero, "was honored as a", newrole, "with master status", master, "at level", level, "and increase mana to", mana+10000, "and HP to", hp+5000)
}
