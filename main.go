package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) *Profile {
	anime := Profile{Name: name}

	if anime.Name == "Sasuke" {
		anime.Name = name
		anime.Health = 200
		anime.Power = 100
		anime.Exp = 0
	}

	if anime.Name == "Goku" {
		anime.Name = name
		anime.Health = 400
		anime.Power = 300
		anime.Exp = 100
	}
	if anime.Name == "Naruto" {
		anime.Name = name
		anime.Health = 150
		anime.Power = 200
		anime.Exp = 50
	}

	return &anime
}
func PowerUp(anime *Profile, power int) *Profile {
	hero := anime

	if hero.Name == anime.Name {
		hero.Health = hero.Health + (hero.Health * power)
		hero.Power = hero.Power + (hero.Power * power)
		hero.Exp = hero.Exp + (hero.Exp * power)
	}

	return hero
}

// func PowerU
func main() {
	profile := MakeProfile("Naruto")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("----- HEROES POWER UP ------")
	PowerUp := PowerUp(profile, 4)
	fmt.Println("Name: ", PowerUp.Name)
	fmt.Println("Health: ", PowerUp.Health)
	fmt.Println("Power:", PowerUp.Power)
	fmt.Println("Exp:", PowerUp.Exp)
}
