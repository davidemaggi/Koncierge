package ui

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

var (
	Version = "v1.0.0" // Set during build or manually

)

func PrintInfo() {

	//pterm.DefaultCenter.Println("This text is centered!\nIt centers the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Koncierge")).Srender()

	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s + Version)
	block := "  "
	// Print each line of the text separately centered in the terminal

	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("Proudly made with 🩷 in Italy " + pterm.BgGreen.Sprint(block) + pterm.BgLightWhite.Sprint(block) + pterm.BgRed.Sprint(block))

}
