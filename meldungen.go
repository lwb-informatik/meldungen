package meldungen

// Autor: St. Schmidt
// Datum: 20.02.2024
// Zweck: ein Modul mit globalem Import

import "fmt"
import "github.com/lwb-informatik/strings"

func Hallo () {
	fmt.Println ("*****************************")
	strings.HalloWelt()
	fmt.Println ("*****************************")
}
