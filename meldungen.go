package meldungen
// Autor: St. Schmidt
// Datum: 18.01.2022
// Zweck: ein Modul mit einem globalen Import
import "fmt"
import "github.com/lwb-informatik/strings"

func Hallo () {
	fmt.Println ("*****************************")
	strings.HalloWelt()
	fmt.Println ("*****************************")
}
