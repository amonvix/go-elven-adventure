package main

import "fmt"

var hero string

func main() {
	hero = "Amon"
	profession := "Apprentice"
	age := 41
	lost := true

	fmt.Println("Our history begins with a hero named:", hero)
	fmt.Println("At the moment, our hero holds the profession:", profession)
	fmt.Println("He has:", age)
	fmt.Println("Is he totally lost?", lost)

	/*
		# 👹 Mini Boss Fight — The Twin Messengers

		The roads of Eldora are under threat.

		Two messengers must cross the kingdom carrying news of hope, but the Royal Registry only recognizes warriors with complete records.

		The Elder Scribe commands:

		> "Create the identity of two heroes and present them to the realm.
		> Only then shall the gates remain open."

		## Objective:

		Use variables to create **two different hero profiles**.
	*/

	messenger1, profession, age, lost := "Amon", "messenger", 41, false
	messenger2, profession2, age2, lost2 := "Lilyth", "messenger", 18, false

	fmt.Println("The new heroes starting a new journey today are!!!")
	fmt.Println(messenger1, profession, "with", age, "years old.", "unknown location:", lost)
	fmt.Println(messenger2, profession2, "with", age2, "years old.", "unknown location:", lost2)
	fmt.Println("Let's cheer the begin of a great journey now. CHEERS!!!")
}
