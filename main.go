package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os/exec"
	"strconv"
	"os"
)

func main() {

	code := flag.String("code", "", "The language to do")
	fullPtr := flag.Bool("full", false, "Output the full language: verbs, tenses and pronouns")
	outDir := flag.String("out", ".", "Directory to write output to")

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

		if *outDir != "." {
			if _, err := os.Stat(*outDir); os.IsNotExist(err) {
				os.Mkdir(*outDir, 0755)
			}
		}

		filename = *outDir + "/"

		if *fullPtr {
			lang, err := service.GetLang(*code)
			if err == nil {
				output, err = lang.MarshalJSON()
			}
			filename += strconv.Itoa(lang.Id) + lang.Code + ".json.full"
		} else {
			langID, verbConf, err := service.GetVerbsOnly(*code)
			if err == nil {
				output, err = verbConf.MarshalJSON()
			}
			filename += strconv.Itoa(langID) + *code + ".json"
		}

		if err == nil {
			err = ioutil.WriteFile(filename, output, 0644)
		}

		if err == nil {
			err = exec.Command("bash", "-c", "gzip " + filename + " -c > " + filename + ".gz").Run()
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("SUCCESS: Written " + filename + " and gz'd")
		}
	}
}
