package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os/exec"
	"strconv"
)

func main() {

	code := flag.String("code", "", "The language to do")
	fullPtr := flag.Bool("full", false, "Output the full language: verbs, tenses and pronouns")

	flag.Parse()

	if *code == "" {
		fmt.Println("Please specify a language code.")
	} else {
		service := RealLanguageService{}

		var (
			output []byte
			filename string
			err error
		)

		if *fullPtr {
			lang, err := service.GetLang(*code)
			if err == nil {
				output, err = lang.MarshalJSON()
			}
			filename =  strconv.Itoa(lang.Id) + lang.Code + ".json.full"
		} else {
			langID, verbConf, err := service.GetVerbsOnly(*code)
			if err == nil {
				output, err = verbConf.MarshalJSON()
			}
			filename =  strconv.Itoa(langID) + *code + ".json"
		}

		if err == nil {
			err = ioutil.WriteFile(filename, output, 0644)
		}

		if err == nil {
			err = exec.Command("tar", "-cvzf", filename + ".gz ", filename).Run()
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
