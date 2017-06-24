package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify a language code.")
	} else {
		service := RealLanguageService{}
		lang, err := service.GetLang(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		output, err := lang.MarshalJSON()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(output))
		}
	}
}
