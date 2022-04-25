package main

import (
	"fmt"
	"go-unit-testing/prose"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas1(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas1(phrases))
	phrases = []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas1(phrases))
}
