package main

import (
	"fmt"
	"os"
	"text/template"
)

type pussyCat struct {
	Name string
	Color string
	CorrGram string
}

func main() {
	
	alyanna := pussyCat {"Alyanna", "Black and White", "a"}
	bunso := pussyCat {"Bunso", "Orenj Kalbo", "an"}
	strok := pussyCat {"Strok", "Orenj Baldado", "an"}
	tonky := pussyCat {"Tonky", "Orenj Iyakin", "an"}

	theCats := [4]pussyCat{alyanna, bunso, strok, tonky}

	templatePath := "C:/Users/user/OneDrive/Desktop/Go Practice/text-template-practice/template.txt"

	for _, cat := range theCats {
	
	t, err := template.New("template.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, cat)
	
	if err != nil {
		fmt.Println(err)
	}
  }
}