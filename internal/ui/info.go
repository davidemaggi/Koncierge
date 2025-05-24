package ui

import (
	"github.com/davidemaggi/koncierge/internal/version"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func PrintInfo() {

	//pterm.DefaultCenter.Println("This text is centered!\nIt centers the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Koncierge")).Srender()

	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s + version.Version() + " - " + version.Name())
	block := "  "
	// Print each line of the text separately centered in the terminal

	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("Proudly made with 🩷 in Italy " + pterm.BgGreen.Sprint(block) + pterm.BgLightWhite.Sprint(block) + pterm.BgRed.Sprint(block))

	var t string = `
       &&&&&&&
      &&(+.+)&&
      ___\=/___
     (|_ ~~~ _|)
        )___(
      /'     '\
     ~~~~~~~~~~~
     '~//~~~\\~'
      /_)   (_\`

	pterm.DefaultCenter.Println(t)

}
