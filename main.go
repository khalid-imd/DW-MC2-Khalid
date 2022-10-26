package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(Name string) Profile {
	var Detail Profile
	Detail.Name = Name
	if Name == "Sasuke" {
		Detail.Health = 200
		Detail.Power = 100
		Detail.Exp = 0
	} else if Name == "Goku" {
		Detail.Health = 400
		Detail.Power = 300
		Detail.Exp = 100
	} else if Name == "Naruto" {
		Detail.Health = 150
		Detail.Power = 200
		Detail.Exp = 50
	}

	return Detail
}

func PowerUp(profile Profile, multipier int) Profile {

	profile.Health = profile.Health + (profile.Health * multipier)
	profile.Power = profile.Power + (profile.Power * multipier)
	profile.Exp = profile.Exp + (profile.Exp * multipier)

	return profile
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("=====HEROES POWER UP======")
	PowerUp := PowerUp(profile, 3)
	fmt.Println("Name : ", PowerUp.Name)
	fmt.Println("Health : ", PowerUp.Health)
	fmt.Println("Power : ", PowerUp.Power)
	fmt.Println("Exp : ", PowerUp.Exp)
}
