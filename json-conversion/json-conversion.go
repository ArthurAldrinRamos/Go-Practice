package main

import (
	"encoding/json"
	"fmt"
)

type Weapon struct {
	Name string `json:"weapon_name"`
	Level int `json:"weapon_level"`
}

type Ninja struct {
	Name string `json:"full_name"`
	Weapon Weapon
	Level int
}

func main(){

	fmt.Println("Hello World")


	sFrom := `{"full_name": "Arthur", "weapon": {"weapon_name": "Katana", "weapon_level": 3}, "level": 69}`

	fmt.Println(sFrom)

	var Arthur Ninja
	err := json.Unmarshal([]byte(sFrom), &Arthur)
	if err != nil {
		fmt.Println("sad la")		
	}
	fmt.Println(Arthur)

	sTo, err := json.Marshal(Arthur)
	if err != nil {
		fmt.Println("sad")
	}
	fmt.Println(string(sTo))

}
