package main

import "fmt"

func main() {
	undeadServants := "Amon's puppet"
	forestVoices := 10

	for i := 1; i <= forestVoices; i++ {
		fmt.Println("GET OUT OF OUR MASTERS LANDS!!! said the voice number", i)

	}
	fmt.Println("this voices can be heard from meters around the once, Ogre's tower")

	fmt.Println("When asked from far who was there, the response was always the same")

	for i := forestVoices; i >= 1; i-- {
		fmt.Println("WE SAID GET OFF. I'M", undeadServants, "AND I GUARD THE", i, "FLOOR!!!")
		if i == 1 {
			fmt.Print("The", undeadServants, "turns to inside the tower and said with a weak voice\nThey have been warned, master!")
		}

	}
}
